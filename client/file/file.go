package file

import "github.com/rs/xid"

type File struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

func newFile(name string, path string) File {
	return File{
		ID:   xid.New().String(),
		Path: path,
		Name: name,
	}
}
