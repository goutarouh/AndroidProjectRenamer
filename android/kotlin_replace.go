package android

import (
	"os"
	"strings"
)

func ReplaceKotlin(path, oldPackageName, newPackageName string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		os.Exit(1)
	}
	contents := string(bytes)
	contents = strings.ReplaceAll(contents, oldPackageName, newPackageName)

	err = os.WriteFile(path, []byte(contents), os.ModePerm)
	if err != nil {
		os.Exit(1)
	}
	return err
}
