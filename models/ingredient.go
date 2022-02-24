package models

type Ingredient struct {
	ID        	string    `json:"id"`
	Name 	  	string    `json:"name"`
	Dlc  		string    `json:"dlc"`
	Quantity    uint8     `json:"quantity"`
}