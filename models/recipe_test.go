package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestRecipeMarshalJSON(t *testing.T) {
	r := Recipe{
		Name: "Crêpe",
		SetupTime:  20,
		NumberPerson:  6,
		Details: "Mélanger des oeufs de la farine et du lait",
	}

	data, err := json.Marshal(r)
	if err != nil {
		t.Error()
	}
	fmt.Println(string(data))
	t.Log(string(data))
}