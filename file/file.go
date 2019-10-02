package file

import (
	"errors"
	"os"
	"sync"

	"path/filepath"

	"github.com/rs/xid"
)

var (
	list []File
	mtx  sync.RWMutex
	once sync.Once
)

func init() {
	once.Do(initialiseList)
}

func initialiseList() {
	root := "/Users/xrustalik/Documents/go/protected-notebook/resources"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		println(path)
		f := File{
			ID:   xid.New().String(),
			Name: info.Name(),
			Path: path,
		}
		mtx.Lock()
		list = append(list, f)
		mtx.Unlock()
		return nil
	})
	if err != nil {
		panic(err)
	}
}

type File struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

func Get() []File {
	return list
}

func Add(name string, content string) string {
	f := newFile(name, content)
	mtx.Lock()
	list = append(list, f)
	mtx.Unlock()
	return f.ID
}

func Delete(id string) error {
	location, err := findFileLocation(id)
	if err != nil {
		return err
	}
	removeElementByLocation(location)
	return nil
}

func newFile(name string, path string) File {
	return File{
		ID:   xid.New().String(),
		Path: path,
		Name: name,
	}
}

func findFileLocation(id string) (int, error) {
	mtx.RLock()
	defer mtx.RUnlock()
	for i, f := range list {
		if isMatchingID(f.ID, id) {
			return i, nil
		}
	}
	return 0, errors.New("could not find file by id")
}

func removeElementByLocation(i int) {
	mtx.Lock()
	list = append(list[:i], list[i+1:]...)
	mtx.Unlock()
}

func isMatchingID(a string, b string) bool {
	return a == b
}
