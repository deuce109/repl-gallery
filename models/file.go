package model

import (
	"github.com/google/uuid"
	"os"
  "path/filepath"
)

type File struct {
	UUID     string
	Data     []byte
	Name     string
	FileInfo os.FileInfo
}

func (f *File) WriteToPath(path string) error {
	return os.WriteFile(filepath.Clean(path), f.Data, 0777)
}

func (f *File) WriteToDir(basePath string) error {
	return os.WriteFile(filepath.Clean(basePath) + filepath.Clean(f.Name), f.Data, 0777)
}

func FileFromPath(path string) (file *File, err error) {
	info, err := os.Stat(path)
  
	if err != nil {
		return
	}

  data, err := os.ReadFile(path)

  if err != nil {
    return
  }
  file = &File{
    uuid.New().String(),
    data,
    info.Name(),
    info,
  }
  return

}
