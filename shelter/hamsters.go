package shelter

import (
	"fmt"
	"time"
)

type Hamster struct {
	Forename string `json:"forename"`
	Surname  string `json:"surname"`
	DOB      string `json:"dateOfBirth"`
	Image    Image  `json:"image"`
	Age      time.Time
}

type Hamsters []*Hamster

type HamsterList struct {
	Values Hamsters `json:"body"`
}

func (h Hamsters) Len() int {
	return len(h)
}

func (h Hamsters) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hamsters) Less(i, j int) bool {
	return h[i].Age.Unix() < h[j].Age.Unix()
}

func (h Hamsters) ToAnimals() []Animal {
	ani := make([]Animal, 0)
	for _, hamster := range h {
		name := fmt.Sprintf("%s %s", hamster.Forename, hamster.Surname)
		now := time.Now()
		years := now.Year() - hamster.Age.Year()

		months := now.Month() - hamster.Age.Month()
		if months < 0 {
			months = 12 + months
			years--
		}

		age := fmt.Sprintf("%d Year(s) %d month(s)", years, months)
		animal := Animal{
			Name:  name,
			Age:   age,
			Image: hamster.Image.URL,
			Type:  "Hamster",
		}

		ani = append(ani, animal)
	}

	return ani
}
