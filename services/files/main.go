package files

import "github.com/goatcms/goat-core/filesystem"

// Files is global file access provider
type Files struct {
	fs filesystem.Filespace
}

// NewFiles create a files instance
func NewFiles(fs filesystem.Filespace) (*Files, error) {
	return &Files{
		fs: fs,
	}, nil
}

// NewImage add image to filespace
func (f *Files) NewImage() filesystem.Filespace {
	return f.fs
}
