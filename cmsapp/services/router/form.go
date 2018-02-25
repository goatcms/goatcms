package router

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"reflect"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/fsstandalone"
	"github.com/goatcms/goatcore/varutil/r"
)

type FormInjector struct {
	req            *http.Request
	filespace      filesystem.Filespace
	maxMemFileSize int64
	tagname        string
	eventScope     app.EventScope
}

func (fi FormInjector) InjectTo(obj interface{}) error {
	if err := fi.req.ParseForm(); err != nil {
		return err
	}
	if fi.IsMultipart() {
		if err := fi.req.ParseMultipartForm(fi.maxMemFileSize); err != nil {
			return err
		}
	}
	structValue := reflect.ValueOf(obj).Elem()
	return fi.injectToL2(structValue, "")
}

func (fi FormInjector) injectToL2(structValue reflect.Value, base string) error {
	for i := 0; i < structValue.NumField(); i++ {
		var isRequired = true
		valueField := structValue.Field(i)
		structField := structValue.Type().Field(i)
		if valueField.Kind() == reflect.Struct {
			if alias, ok := structField.Tag.Lookup(fi.tagname); ok == true {
				if err := fi.injectToL2(valueField, base+alias); err != nil {
					return err
				}
			}
			continue
		}
		if valueField.Kind() == reflect.Ptr && valueField.Elem().Kind() == reflect.Struct {
			if alias, ok := structField.Tag.Lookup(fi.tagname); ok == true {
				subValueField := valueField.Elem()
				if err := fi.injectToL2(subValueField, base+alias); err != nil {
					return err
				}
			}
			continue
		}
		key := structField.Tag.Get(fi.tagname)
		if key == "" {
			continue
		}
		if strings.HasPrefix(key, "?") {
			isRequired = false
			key = key[1:]
		}
		key = base + key
		switch valueField.Interface().(type) {
		case filesystem.File:
			multipartFile, multipartHeader, err := fi.req.FormFile(key)
			if err != nil {
				if !isRequired {
					continue
				}
				return err
			}
			defer multipartFile.Close()
			tmpFile, err := fsstandalone.NewTMPStandaloneFile(fi.filespace, multipartHeader.Filename, multipartHeader.Header.Get("Content-Type"))
			if err != nil {
				return err
			}
			tmpWriter, err := tmpFile.Writer()
			if err != nil {
				return err
			}
			defer tmpWriter.Close()
			if _, err = io.Copy(tmpWriter, multipartFile); err != nil {
				return err
			}
			valueField.Set(reflect.ValueOf(tmpFile))
		case []filesystem.File:
			multipartFiles, ok := fi.req.MultipartForm.File[key]
			if !ok {
				if !isRequired {
					continue
				}
				return fmt.Errorf("Dont contains files for %v", key)
			}
			accumulator := make([]filesystem.File, len(multipartFiles))
			for i, fileRow := range multipartFiles {
				multipartFile, err := fileRow.Open()
				defer multipartFile.Close()
				if err != nil {
					return err
				}
				tmpFile, err := fsstandalone.NewTMPStandaloneFile(fi.filespace, fileRow.Filename, fileRow.Header.Get("Content-Type"))
				if err != nil {
					return err
				}
				tmpWriter, err := tmpFile.Writer()
				if err != nil {
					return err
				}
				defer tmpWriter.Close()
				if _, err = io.Copy(tmpWriter, multipartFile); err != nil {
					return err
				}
				accumulator[i] = tmpFile
			}
			valueField.Set(reflect.ValueOf(accumulator))
		case []string:
			collection, ok := fi.req.Form[key]
			if !ok {
				if !isRequired {
					continue
				}
				return fmt.Errorf("Dont contains string collection for %v", key)
			}
			valueField.Set(reflect.ValueOf(collection))
		default:
			collection, ok := fi.req.Form[key]
			if !ok || len(collection) == 0 {
				if !isRequired {
					continue
				}
				return fmt.Errorf("Don't contains a string for %v", key)
			}
			if len(collection) > 1 {
				return fmt.Errorf("%v is a collection", key)
			}
			r.SetValueFromString(valueField, collection[0])
		}
	}
	return nil
}

func (fi FormInjector) IsMultipart() bool {
	v := fi.req.Header.Get("Content-Type")
	if v == "" {
		return false
	}
	d, _, err := mime.ParseMediaType(v)
	if err != nil || d != "multipart/form-data" {
		return false
	}
	return true
}
