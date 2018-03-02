package shelter

type Hamster Animal

type Hamsters []*Hamster

type HamsterList struct {
	Values []*Hamster `json:"body"`
}

// func (c Hamsters) Len() int {
// 	return len(c)
// }
//
// func (c Hamsters) Swap(i, j int) {
// 	c[i], c[j] = c[j], c[i]
// }
//
// func (c Hamsters) Less(i, j int) bool {
// 	return c[i].Age.Unix() > c[j].Age.Unix()
// }
