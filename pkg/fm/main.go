package fm

import (
	"io"
	"os"
	"path/filepath"
)

func Copy(src, dir string) error {
	out, err := os.Open(src)
	if err != nil {
		return err
	}
	defer out.Close()

	in, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer in.Close()

	if _, err := io.Copy(in, out); err != nil {
		return err
	}

	return nil
}

func CopyWithCreate(src, dir string) error {
	out, err := os.OpenFile(src, os.O_RDONLY, 0777)
	if err != nil {
		return err
	}
	defer out.Close()

	in, err := os.OpenFile(dir, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer in.Close()

	if _, err := io.Copy(in, out); err != nil {
		return err
	}

	return nil
}

func CopyDir(src, dir string) error {
	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, v := range files {
		if v.IsDir() {
			if err := os.Mkdir(filepath.Join(dir, v.Name()), 0777); err != nil {
				return err
			}
			if err := CopyDir(filepath.Join(src, v.Name()), filepath.Join(dir, v.Name())); err != nil {
				return err
			}

		} else {
			if err := CopyWithCreate(filepath.Join(src, v.Name()), filepath.Join(dir, v.Name())); err != nil {
				return err
			}
		}
	}

	return nil
}
