package repository_os

import (
	"fmt"
	"io"
	"os"
	"patreon/internal/app"
	repFiles "patreon/internal/app/repository/files"
	"time"
)

const createChmod = 744

type FileRepository struct {
	staticDir string
}

func NewFileRepository(staticDir string) *FileRepository {
	return &FileRepository{
		staticDir: staticDir,
	}
}

// SaveFile Errors:
//		app.GeneralError Errors:
//			repository_os.ErrorCreate
//			repository_os.ErrorCopyFile
func (fr *FileRepository) SaveFile(file io.Reader, name repFiles.FileName, typeF repFiles.TypeFiles) (string, error) {
	date := time.Now()
	path := fmt.Sprintf("%s/%s/%d/%s/%d", fr.staticDir, typeF, date.Year(), date.Month(), date.Day())
	if err := os.MkdirAll(path, createChmod); err != nil {
		return "", &app.GeneralError{Err: ErrorCreate, ExternalErr: err}
	}

	name = repFiles.FileName(fmt.Sprintf("%d%d%d%s", date.Hour(), date.Minute(), date.Second(), name))
	out, err := os.Create(fmt.Sprintf("%s/%s", path, name))
	if err != nil {
		return "", &app.GeneralError{Err: ErrorCreate, ExternalErr: err}
	}

	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	_, err = io.Copy(out, file)
	if err != nil {
		return "", &app.GeneralError{Err: ErrorCopyFile, ExternalErr: err}
	}

	return path + "/" + string(name), nil
}

func (fr *FileRepository) LoadFile(path string) (io.Reader, error) {
	file, err := os.Open(fr.staticDir + "/" + path)
	if err != nil {
		return nil, &app.GeneralError{Err: ErrorOpenFile, ExternalErr: err}
	}
	return file, nil
}
