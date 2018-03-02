package shelter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/gorilla/mux"
)

type Client struct {
	client  *http.Client
	baseURL string
}

func NewClient(c *http.Client) *Client {
	return &Client{
		c,
		"https://apigateway.test.lifeworks.com/rescue-shelter-api",
	}
}

func ServeAPI(c *http.Client) error {
	client := NewClient(c)
	r := mux.NewRouter()
	r.Methods("GET").Path("/animals").HandlerFunc(client.getAnimals)
	return http.ListenAndServe(":4000", r)
}

func (c *Client) getAnimals(w http.ResponseWriter, r *http.Request) {
	// getCats()
	dogs, err := c.getDogs()
	d, err := json.Marshal(dogs)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(d)

	// getHamster()
}

func (c *Client) getDogs() ([]Animal, error) {
	resp, err := c.client.Get(c.baseURL + "/dogs")

	if err != nil {
		log.Printf("Error doing get: %s", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	var dogs DogList
	if err = json.NewDecoder(resp.Body).Decode(&dogs); err != nil {
		log.Printf("Error reading json: %s", err.Error())
		return nil, err
	}

	for _, dog := range dogs.Values {
		t, err := time.Parse("2006-01-02", dog.DOB)
		if err != nil {
			log.Printf("Error parsing time %s", err.Error())
			return nil, err
		}

		dog.Age = t
	}

	sort.Sort(Dogs(dogs.Values))
	return dogs.Values.ToAnimals(), nil
}

// func getCats() {
// 	client := &http.Client{}
// 	resp, err := client.Get("https://apigateway.test.lifeworks.com/rescue-shelter-api/cats")
//
// 	if err != nil {
// 		log.Printf("Error doing get: %s", err.Error())
// 		return
// 	}
//
// 	defer resp.Body.Close()
//
// 	var cats CatList
//
// 	if err = json.NewDecoder(resp.Body).Decode(&cats); err != nil {
// 		log.Printf("Error reading json: %s", err.Error())
// 		return
// 	}
//
// 	group := make(map[string][]Cat)
// 	for _, c := range cats.Values {
// 		t, err := time.Parse("2006-01-02", c.DOB)
// 		if err != nil {
// 			log.Printf("Error parsing time %s", err.Error())
// 			return
// 		}
//
// 		c.Age = t
// 		switch c.Colour {
// 		case "ginger":
// 			group["ginger"] = append(group["ginger"], c)
// 		case "black":
// 			group["black"] = append(group["black"], c)
// 		default:
// 			group["other"] = append(group["other"], c)
// 		}
// 	}
//
// 	for _, value := range group {
// 		sort.Sort(Cats(value))
// 	}
//
// 	for key, v := range group {
// 		for _, i := range v {
// 			log.Println(key)
// 			log.Println(i.Forename)
// 			log.Println(i.DOB)
// 		}
// 	}
// }
//
// func getHamster() {
// 	client := &http.Client{}
// 	resp, err := client.Get("https://apigateway.test.lifeworks.com/rescue-shelter-api/hamsters")
//
// 	if err != nil {
// 		log.Printf("Error doing get: %s", err.Error())
// 		return
// 	}
//
// 	defer resp.Body.Close()
//
// 	var hamsters HamsterList
// 	if err = json.NewDecoder(resp.Body).Decode(&hamsters); err != nil {
// 		log.Printf("Error reading json: %s", err.Error())
// 		return
// 	}
//
// 	for _, dog := range hamsters.Values {
// 		t, err := time.Parse("2006-01-02", dog.DOB)
// 		if err != nil {
// 			log.Printf("Error parsing time %s", err.Error())
// 			return
// 		}
//
// 		dog.Age = t
// 	}
//
// 	sort.Sort(Hamsters(hamsters.Values))
//
// 	for _, dog := range hamsters.Values {
// 		log.Println(dog.Forename)
// 		log.Println(dog.DOB)
// 	}
// }
