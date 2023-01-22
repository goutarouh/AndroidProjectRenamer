package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	root := "/mnt/c/Users/81704/Documents/dev/android/SimpleRssReader"
	target := "androidsampleapp"
	newTarget := "simplerssreader"

	// Replace file contents
	renameFilesContent(root, target, newTarget)

	// Rename Directory
	dirs, err := getRenameDirs(root, target)
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(dirs, func(i, j int) bool {
		return len(dirs[i]) > len(dirs[j])
	})
	fmt.Println("------------------------")
	for _, dir := range dirs {
		pdir, _ := filepath.Split(dir)
		new := filepath.Join(pdir, newTarget)
		fmt.Println(new)
		os.Rename(dir, new)
	}

}

func getRenameDirs(root, target string) ([]string, error) {
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
			dirs = append(dirs, path)
			return nil
		}
		return nil
	})
	if err != nil {
		log.Fatal("RenameDirs", err)
	}
	return dirs, err
}

func renameFilesContent(root, target, newTarget string) {
	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {

		if err != nil {
			log.Fatal("WalkDir", err)
		}

		ex := filepath.Ext(path)
		if info.IsDir() {
			return nil
		}
		if ex == ".kt" || ex == ".kts" || ex == ".gradle" {
			fmt.Println(path)
			Replace(path, target, newTarget)
		}

		return nil
	})
	if err != nil {
		log.Fatal("RenameDirs", err)
	}
}

func Replace(path, old, new string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Replace", err)
		os.Exit(1)
	}
	contents := string(bytes)
	contents = strings.ReplaceAll(contents, old, new)
	err = os.WriteFile(path, []byte(contents), os.ModePerm)
	if err != nil {
		log.Fatal("Replace", err)
		os.Exit(1)
	}
	return err
}
