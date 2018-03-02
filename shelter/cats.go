package shelter

type Cat struct {
	Animal
	Colour   string `json:"colour"`
}

type Cats []Cat

type CatList struct {
	Values Cats `json:"body"`
}

// func (c Cats) Len() int {
// 	return len(c)
// }
//
// func (c Cats) Swap(i, j int) {
// 	c[i], c[j] = c[j], c[i]
// }
//
// func (c Cats) Less(i, j int) bool {
// 	return c[i].Age.Unix() < c[j].Age.Unix()
// }
