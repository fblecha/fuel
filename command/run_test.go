package command

import (
  //"bufio"
	//"fmt"
	//"os"
	"testing"
)

func TestSplitJsonAndMarkdown(t *testing.T) {
  if _, _, err := SplitJsonAndMarkdown("../example/content/dogs/labrador_retriever.haiku"); err != nil {
    t.Error(err)
  }
}
