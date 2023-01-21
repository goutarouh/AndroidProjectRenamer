package android

import (
	"os"
	"testing"
)

func TestReplaceKotlin(t *testing.T) {
	os.WriteFile("test.kt", []byte("com.github.test1"), os.ModePerm)
	err := ReplaceKotlin("test.kt", "com.github.test1", "com.github.test2")
	if err != nil {
		t.Fatal("TestReplaceKotlin error", err)
	}
	bytes, _ := os.ReadFile("test.kt")
	contents := string(bytes)
	if contents != "com.github.test2" {
		t.Fatal("TestReplaceKotlin error", contents)
	}
	os.Remove("test.kt")
}
