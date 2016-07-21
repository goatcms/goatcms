package images

import (
	"io"
	"log"
	mp "mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/forms"
	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/models/image"
	"github.com/goatcms/goatcms/services"
	"github.com/gorilla/mux"
)

// ImageController is image controller endpoint
type ImageController struct {
	tmpl       services.Template
	randomID   services.RandomID
	imageDAO   models.ImageDAO
	articleDAO models.ArticleDAO
}

// NewImageController create instance of a image controller
func NewImageController(dp dependency.Provider) (*ImageController, error) {
	ctrl := &ImageController{}
	// load template service from dependency provider
	tmplIns, err := dp.Get(services.TemplateID)
	if err != nil {
		return nil, err
	}
	ctrl.tmpl = tmplIns.(services.Template)
	// load template service from dependency provider
	randidIns, err := dp.Get(services.RandomidID)
	if err != nil {
		return nil, err
	}
	ctrl.randomID = randidIns.(services.RandomID)
	// load imageDAO service from dependency provider
	daoIns, err := dp.Get(models.ImageDAOID)
	if err != nil {
		return nil, err
	}
	ctrl.imageDAO = daoIns.(models.ImageDAO)
	// load articleDAO service from dependency provider
	dao2Ins, err := dp.Get(models.ArticleDAOID)
	if err != nil {
		return nil, err
	}
	ctrl.articleDAO = dao2Ins.(models.ArticleDAO)
	return ctrl, nil
}

const (
	imagesFilePath = "./assets/images/"
	imageIDlength  = 12
)

// newImage create new ImageDTO instance
func (c *ImageController) newImage(articleID int) imagemodel.ImageDTO {
	return imagemodel.ImageDTO{
		// ID:        rand.Intn(100000), // TODO better ID handling
		ArticleID: articleID,
		CreatedAt: time.Now(),
	}
}

// createFromFile persist image from form given file
func (c *ImageController) createFromFile(
	f mp.File, h *mp.FileHeader, d string, articleID int,
) (*imagemodel.ImageDTO, error) {
	image := c.newImage(articleID)
	image.Name = h.Filename
	image.Description = d
	// Move file to an appropriate place, with and appropriate name
	// image.Location = "art" + strconv.Itoa(image.GetArticleID()) + "_" + image.GetName()
	image.Location, _ = c.randomID.GenerateID(
		"art"+strconv.Itoa(image.GetArticleID()),
		imageIDlength,
	)
	image.Location = image.Location + filepath.Ext(h.Filename)
	// Create dir for article's photos if not exists
	artImgsPath := imagesFilePath + "article" + strconv.Itoa(image.GetArticleID()) + "/"
	if _, err := os.Stat(artImgsPath); os.IsNotExist(err) {
		log.Println(artImgsPath, "does not exist")
		os.Mkdir(artImgsPath, 0755) // so we create it
	}
	// if directory exists
	if _, err := os.Stat(artImgsPath); err == nil {
		log.Println(artImgsPath, "exists")
	}
	// Open a file at target location
	savedFile, err := os.Create(artImgsPath + image.Location)
	if err != nil {
		return nil, err
	}
	defer savedFile.Close()
	// Copy the uploaded file to target location
	size, err := io.Copy(savedFile, f)
	if err != nil {
		return nil, err
	}
	image.Size = size
	// Save the image to the database
	return &image, c.imageDAO.PersistOne(models.ImageDTO(&image))
}

// TemplateAddImage is handler to serve template where one can add new image
func (c *ImageController) TemplateAddImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	arguments := map[string]interface{}{
		"articleID": vars["id"],
	}
	if err := c.tmpl.ExecuteTemplate(w, "images/new", arguments); err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// TrySaveImage is handler to save image from form given source
func (c *ImageController) TrySaveImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID, _ := strconv.Atoi(vars["id"])
	// 2 ways of uploading - by URL or from local file
	if r.FormValue("url") != "" {
		// upload file from given URL
		return
	}
	c.handlerImageCreateFromFile(w, r, articleID)
}

func (c *ImageController) handlerImageCreateFromFile(
	w http.ResponseWriter, r *http.Request, articleID int,
) {
	description := r.FormValue("description")
	file, headers, err := r.FormFile("file")
	if file == nil { // if no file uploaded
		c.tmpl.ExecuteTemplate(w, "images/new", map[string]interface{}{
			"Error": forms.ErrNoImage,
			// TODO add description here to persist
		})
		return
	}
	ext := filepath.Ext(headers.Filename)
	if ext != ".png" && ext != ".jpg" && ext != ".gif" {
		c.tmpl.ExecuteTemplate(w, "images/new", map[string]interface{}{
			"Error": forms.ErrInvalidImageType,
			// TODO add description here to persist
		})
		return
	}
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Create file and try to save, if errors reload site with them
	image, err := c.createFromFile(file, headers, description, articleID)
	if err != nil {
		c.tmpl.ExecuteTemplate(w, "images/new", map[string]interface{}{
			"Error": err,
			"Image": image,
		})
		return
	}
	http.Redirect(w, r, "/article/"+strconv.Itoa(articleID), http.StatusFound)
}
