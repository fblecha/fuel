package command

import (
  //"bufio"
	"fmt"
	//"os"
	"testing"
)

func TestSplitJsonAndMarkdown(t *testing.T) {
  fmt.Println("blah")
  // Open the file.

  if _, _, err := SplitJsonAndMarkdown("../example/content/dogs/labrador_retriever.haiku"); err != nil {
    t.Error(err)
  }
}
