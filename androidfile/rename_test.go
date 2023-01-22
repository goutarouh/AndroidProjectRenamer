package androidfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetRenameDirs(t *testing.T) {
	root := "/mnt/c/Users/81704/Documents/dev/android/SimpleRssReader"
	target := "androidsampleapp"
	_, err := GetRenameDirs(root, target)
	if err != nil {
		t.Fatal("TestGetRenameDirs", err)
	}
}

func TestRename(t *testing.T) {
	os.Mkdir("test1", 0777)
	err := Rename("test1", "test2")
	if err != nil {
		t.Fatal("TestRename error", err)
	}
	os.Remove("test2")
}

func TestWalkKt(t *testing.T) {
	os.Mkdir("./kt", 0777)
	os.Create("./kt/test1.kt")
	os.Create("./kt/test2.kt")
	os.Create("./kt/test3.go")
	err := WalkKt(".", func(path string) {
		ex := filepath.Ext(path)
		if ex != ".kt" {
			t.Fatal("TestWalkKt error", path)
		}
	})
	if err != nil {
		t.Fatal("TestWalkKt error", err)
		t.Fatal("Test ")
	}
	os.RemoveAll("./kt")
}
