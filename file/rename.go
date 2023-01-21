package file

import (
	"io/fs"
	"os"
	"path/filepath"
)

func Rename(oldDirPath, newDirPath string) error {
	error := os.Rename(oldDirPath, newDirPath)
	return error
}

func WalkKt(root string, action func(string)) error {
	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			os.Exit(1)
		}
		if info.IsDir() {
			return nil
		}
		ex := filepath.Ext(info.Name())
		if ex != ".kt" {
			return nil
		}
		action(path)
		return nil
	})
	return err
}
