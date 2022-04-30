package files

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

// Local is an implementation of the Storage interface which works with the
// local disk on the current machine
type Local struct {
	maxFileSize int // maximum number of bytes for files
	basePath    string
}

func NewLocal(basePath string, maxSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}
	return &Local{basePath: p}, nil
}

func (l *Local) Save(path string, r io.Reader) error {
	// get the full path for the file
	fp := l.fullPath(path)

	d := filepath.Dir(fp)
	err := os.MkdirAll(d, os.ModePerm)
	if err != nil {
		return xerrors.Errorf("Unable to delete file: %w", err)
	}

	// if the file exists delete it
	_, err = os.Stat(fp)

	if err == nil {
		err = os.Remove(fp)
		if err != nil {
			return xerrors.Errorf("Unable to delete file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		// If the error is anything other than not exists error
		return xerrors.Errorf("Unable to get the file info: %w", err)
	}

	// create a new file at the path
	f, err := os.Create(fp)
	if err != nil {
		return xerrors.Errorf("Unable to create file: %w", err)
	}
	defer f.Close()

	_, err = io.Copy(f, r)
	if err != nil {
		return xerrors.Errorf("Unable to write to file: %w", err)
	}

	return nil

}

// returns the absolute path
func (l *Local) fullPath(path string) string {
	// append the given path with the basepath
	return filepath.Join(l.basePath, path)
}
