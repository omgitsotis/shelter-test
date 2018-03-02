package shelter

import (
	"fmt"
	"time"
)

type Dog struct {
	Forename string `json:"forename"`
	Surname  string `json:"surname"`
	DOB      string `json:"dateOfBirth"`
	Image    Image  `json:"image"`
	Colour   string `json:"colour"`
	Age      time.Time
}

type Dogs []*Dog

type DogList struct {
	Values Dogs `json:"body"`
}

func (d Dogs) Len() int {
	return len(d)
}

func (d Dogs) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Dogs) Less(i, j int) bool {
	return d[i].Age.Unix() < d[j].Age.Unix()
}

func (d Dogs) ToAnimals() []Animal {
	ani := make([]Animal, 0)
	for _, dog := range d {
		name := fmt.Sprintf("%s %s", dog.Forename, dog.Surname)
		now := time.Now()
		years := now.Year() - dog.Age.Year()

		months := now.Month() - dog.Age.Month()
		if months < 0 {
			months = 12 + months
			years--
		}

		age := fmt.Sprintf("%d Year(s) %d month(s)", years, months)
		animal := Animal{
			Name:  name,
			Age:   age,
			Image: dog.Image.URL,
		}

		ani = append(ani, animal)
	}

	return ani
}
