package models

type Recipe struct {
	ID        	 string    	`json:"id"`
	Name         string 	`json:"name"`
	SetupTime    uint8 		`json:"setup_time"`
	NumberPerson uint8 		`json:"number_person"`
	Details      string 	`json:"details"`
}


