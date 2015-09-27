package command

import (
	"testing"
	//"github.com/fblecha/haiku/command"
)

func TestSplitJsonAndMarkdown(t *testing.T) {
	if _, _, err := SplitJsonAndMarkdown("../example/content/dogs/labrador_retriever.haiku"); err != nil {
		t.Error(err)
	}
}

func TestSeparaterButNoJson(t *testing.T) {
	if _, _, err := SplitJsonAndMarkdown("../example/content/dogs/jack_russel_terrier.haiku"); err != nil {
		t.Error(err)
	}
}

func TestNoSeparaterNoJson(t *testing.T) {
	if _, _, err := SplitJsonAndMarkdown("../example/content/dogs/airedale_terrier.haiku"); err != nil {
		t.Error(err)
	}
}
