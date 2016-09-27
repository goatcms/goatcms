package forms

// ImageForm is structure with image form values
type ImageForm struct {
}

const (
	// ErrInvalidImageType temporarily export, until proper Validation will work
	ErrInvalidImageType = "Please upload only jpeg, gif or png images"
	// ErrNoImage temporarily export, until proper Validation will work
	ErrNoImage         = "Please select an image to upload"
	errImageURLInvalid = "Couldn't download image from the URL you provided"
)

var mimeExtensions = map[string]string{
	"image/png":  ".png",
	"image/jpeg": ".jpg",
	"image/gif":  ".gif",
}

// Validate validate form and return bool how validation passed
func (f ImageForm) Validate() (bool, map[string][]string) {
	// TODO to validate here, not in controller
	return false, nil
}
