package shelter

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/gorilla/mux"
)

// Client is the client that connects to the rescue shelter API
type Client struct {
	client  *http.Client
	baseURL string
}

// NewClient creates a new client with a given url for the rescue shelter API
func NewClient(uri string) *Client {
	return &Client{&http.Client{}, uri}
}

// ServeAPI creates a new client and serves the route to this API
func ServeAPI(uri string) error {
	client := NewClient(uri)
	r := mux.NewRouter()
	r.Methods("GET").Path("/animals").HandlerFunc(client.getAnimals)
	return http.ListenAndServe(":4000", r)
}

// Get animals calls the three shelter API routes and sorts the results
func (c *Client) getAnimals(w http.ResponseWriter, r *http.Request) {
	a := make([]Animal, 0)
	errorCount := 0

	dogsChan := make(chan []Animal)
	hamsChan := make(chan []Animal)
	catsChan := make(chan []Animal)

	errChan := make(chan error)

	// Run each of the calls in their own thread
	go c.getDogs(dogsChan, errChan)
	go c.getCats(catsChan, errChan)
	go c.getHamsters(hamsChan, errChan)

	group := make(map[string][]Animal)

	for i := 0; i < 3; i++ {
		select {
		case dogs := <-dogsChan:
			log.Println("Dogs complete")
			group["dogs"] = dogs
		case cats := <-catsChan:
			log.Println("Cats complete")
			group["cats"] = cats
		case hamster := <-hamsChan:
			log.Println("Hamster complete")
			group["hamster"] = hamster
		case <-errChan:
			log.Println("Recieved error")
			errorCount++
		}
	}

	if errorCount == 3 {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	a = append(a, group["dogs"]...)
	a = append(a, group["cats"]...)
	a = append(a, group["hamster"]...)

	b, err := json.Marshal(a)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(b)
}

// getDogs calls the dog API and sorts the results by age
func (c *Client) getDogs(r chan<- []Animal, e chan<- error) {
	resp, err := c.client.Get(c.baseURL + "/dogs")
	if err != nil {
		log.Printf("Error doing get: %s", err.Error())
		e <- err
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Error: status code not ok")
		e <- errors.New("Status code not ok")
		return
	}

	defer resp.Body.Close()

	var dogs DogList
	if err = json.NewDecoder(resp.Body).Decode(&dogs); err != nil {
		log.Printf("Error reading json: %s", err.Error())
		e <- err
		return
	}

	for _, dog := range dogs.Values {
		t, err := time.Parse("2006-01-02", dog.DOB)
		if err != nil {
			log.Printf("Error parsing time %s", err.Error())
			e <- err
			return
		}

		dog.Age = t
	}

	sort.Sort(Dogs(dogs.Values))
	r <- dogs.Values.ToAnimals()
}

// getCats calls the cats API and sorts the results by age
func (c *Client) getCats(r chan<- []Animal, e chan<- error) {
	resp, err := c.client.Get(c.baseURL + "/cats")
	if err != nil {
		log.Printf("Error doing get: %s", err.Error())
		e <- err
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Error: status code not ok")
		e <- errors.New("Status code not ok")
		return
	}

	defer resp.Body.Close()

	var cats CatList
	if err = json.NewDecoder(resp.Body).Decode(&cats); err != nil {
		log.Printf("Error reading json: %s", err.Error())
		e <- err
		return
	}

	group := make(map[string]Cats)

	for _, cat := range cats.Values {
		t, err := time.Parse("2006-01-02", cat.DOB)
		if err != nil {
			log.Printf("Error parsing time %s", err.Error())
			e <- err
			return
		}

		cat.Age = t

		switch cat.Colour {
		case "ginger":
			group["ginger"] = append(group["ginger"], cat)
		case "black":
			group["black"] = append(group["black"], cat)
		default:
			group["other"] = append(group["other"], cat)
		}
	}

	for _, value := range group {
		sort.Sort(Cats(value))
	}

	result := make([]Animal, 0)
	result = append(result, group["ginger"].ToAnimals()...)
	result = append(result, group["black"].ToAnimals()...)
	result = append(result, group["other"].ToAnimals()...)

	r <- result
}

// getHamsters calls the hamsters API and sorts the results by age
func (c *Client) getHamsters(r chan<- []Animal, e chan<- error) {
	resp, err := c.client.Get(c.baseURL + "/hamsters")
	if err != nil {
		log.Printf("Error doing get: %s", err.Error())
		e <- err
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Error: status code not ok")
		e <- errors.New("Status code not ok")
		return
	}

	defer resp.Body.Close()

	var hams HamsterList
	if err = json.NewDecoder(resp.Body).Decode(&hams); err != nil {
		log.Printf("Error reading json: %s", err.Error())
		e <- err
		return
	}

	for _, ham := range hams.Values {
		t, err := time.Parse("2006-01-02", ham.DOB)
		if err != nil {
			log.Printf("Error parsing time %s", err.Error())
			e <- err
			return
		}

		ham.Age = t
	}

	sort.Sort(Hamsters(hams.Values))
	r <- hams.Values.ToAnimals()
}
