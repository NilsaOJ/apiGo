package models


import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestIngredientMarshalJSON(t *testing.T) {
	i := Ingredient{
		Name: "Lait",
		Dlc: "12/12/2022",
		Quantity: 1,
	}

	data, err := json.Marshal(i)
	if err != nil {
		t.Error()
	}
	fmt.Println(string(data))
	t.Log(string(data))
}