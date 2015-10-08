package command

import (
	"testing"
	//"fmt"
	"path/filepath"
	"reflect"
	"strings"
)

func TestReverse(t *testing.T) {
	stuff := []string{"a", "b", "c"}
	if !reflect.DeepEqual(Reverse(stuff), []string{"c", "b", "a"}) {
		t.Fatalf("array was not reversed")
	}
}

func TestPathToDirs(t *testing.T) {
	//path := "blah/foo/bar/monkey.md"
	const lsep = filepath.ListSeparator
	//path := string([]byte{'a', lsep, 'b'})
	path := "a/b"
	if dirs := strings.Split(path, "/"); !reflect.DeepEqual(dirs, []string{"a", "b"}) {
		t.Fatalf("dirs did not equal expected result")
	}
}
