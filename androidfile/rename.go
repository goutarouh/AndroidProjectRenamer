package androidfile

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetRenameDirs(root, target string) ([]string, error) {
	dirs := []string{}
	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {

		if err != nil {
			log.Fatal("WalkDir", err)
		}
		if strings.Contains(path, "build") {
			return nil
		}
		if strings.Contains(path, ".") {
			return nil
		}
		if info.IsDir() && info.Name() == target {
			dirs = append(dirs, target)
			return nil
		}
		return nil
	})
	if err != nil {
		log.Fatal("RenameDirs", err)
	}
	return dirs, err
}

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
