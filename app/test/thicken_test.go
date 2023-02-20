package test

import (
	"crypto/md5"
	"crypto/sha1"

	"io"
	"testing"
)

func TestWithMD5(t *testing.T) {
	thicken := md5.New()
	old_str := "Hello World"
	io.WriteString(thicken, old_str)
	new_str := thicken.Sum(nil)
	t.Logf("old: %s\t new %x", old_str, new_str)
}

func TestWithSHA1(t *testing.T) {
	thicken := sha1.New()
	old_str := "Hello World"
	io.WriteString(thicken, old_str)
	new_str := thicken.Sum(nil)
	t.Logf("old: %s\t new %x", old_str, new_str)
}
