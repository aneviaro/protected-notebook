package file

import (
	"errors"
	"io/ioutil"
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
	root := "resources"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			f := File{
				ID:   xid.New().String(),
				Name: info.Name(),
				Path: path,
			}
			mtx.Lock()
			list = append(list, f)
			mtx.Unlock()
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

//File struct to store file id name and path
type File struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

//Get list of files
func Get() []File {
	return list
}

//Add new file
func Add(name string, content string) string {
	f := newFile(name, content)
	mtx.Lock()
	list = append(list, f)
	mtx.Unlock()
	return f.ID
}

//GetContentByID and return it
func GetContentByID(id string) (string, error) {
	fileItem, err := findFileByID(id)
	if err != nil {
		return "", err
	}
	return readFile(fileItem)
}

//GetContentByName and return it
func GetContentByName(name string) (string, error) {
	fileItem, err := findFileByName(name)
	if err != nil {
		return "", err
	}
	return readFile(fileItem)
}

//Delete file by id
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

func readFile(file File) (string, error) {
	data, err := ioutil.ReadFile(file.Path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func removeElementByLocation(i int) {
	mtx.Lock()
	list = append(list[:i], list[i+1:]...)
	mtx.Unlock()
}

func findFileByID(id string) (File, error) {
	for _, f := range list {
		if isMatchingID(f.ID, id) {
			return f, nil
		}
	}
	return File{}, errors.New("could not find file by id")
}

func findFileByName(name string) (File, error) {
	for _, f := range list {
		if isMatchingID(f.Name, name) {
			return f, nil
		}
	}
	return File{}, errors.New("could not find file by id")
}

func isMatchingID(a string, b string) bool {
	return a == b
}
