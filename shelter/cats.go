package shelter

import (
	"fmt"
	"time"
)

type Cat struct {
	Forename string `json:"forename"`
	Surname  string `json:"surname"`
	DOB      string `json:"dateOfBirth"`
	Image    Image  `json:"image"`
	Colour   string `json:"colour"`
	Age      time.Time
}

type Cats []*Cat

type CatList struct {
	Values Cats `json:"body"`
}

func (c Cats) Len() int {
	return len(c)
}

func (c Cats) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Cats) Less(i, j int) bool {
	return c[i].Age.Unix() < c[j].Age.Unix()
}

func (c Cats) ToAnimals() []Animal {
	ani := make([]Animal, 0)
	for _, cat := range c {
		name := fmt.Sprintf("%s %s", cat.Forename, cat.Surname)
		now := time.Now()
		years := now.Year() - cat.Age.Year()

		months := now.Month() - cat.Age.Month()
		if months < 0 {
			months = 12 + months
			years--
		}

		age := fmt.Sprintf("%d Year(s) %d month(s)", years, months)
		animal := Animal{
			Name:  name,
			Age:   age,
			Image: cat.Image.URL,
			Type:  "Cat",
		}

		ani = append(ani, animal)
	}

	return ani
}
