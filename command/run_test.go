package command

import (
	"testing"
	//"fmt"
)

func TestSplitJsonAndMarkdown(t *testing.T) {
	if jsonMap, _, err := SplitJsonAndMarkdown("../example/content/dogs/labrador_retriever.haiku"); err != nil {
		t.Error(err)
	} else {
		if jsonMap["breed"] != "Labrador Retriever" {
			t.Fail()
		}
	}
}

func TestSeparaterButNoJson(t *testing.T) {
	jsonMap, _, err := SplitJsonAndMarkdown("../example/content/dogs/jack_russel_terrier.haiku")
	if err != nil {
		t.Error(err)
	}
	//there shouldn't be any items in the json map since we didn't have any json
	if len(jsonMap) != 0 {
		t.Fail()
	}
}

func TestNoSeparaterNoJson(t *testing.T) {
	jsonMap, markdown, err := SplitJsonAndMarkdown("../example/content/dogs/airedale_terrier.haiku");
	if err != nil {
		t.Error(err)
	}
	//there shouldn't be any items in the json map since we didn't have any json
	//fmt.Printf("test len(jsonMap) = %s \n", len(jsonMap) )

	//fmt.Printf("test markdown = %s \n", markdown)
	if len(jsonMap) == 546 {
		t.Fail()
	}
	if len(markdown) < 1 {
		t.Fail()
	}
}
