package command

import (
	//"bufio"
	//"fmt"
	//"os"
	"testing"
)

func TestSplitJsonAndMarkdown(t *testing.T) {
	if _, _, err := splitJsonAndMarkdown("../example/content/dogs/labrador_retriever.haiku"); err != nil {
		t.Error(err)
	}
}
