package shelter

import (
	"encoding/json"
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
	a := make([]Animal, 0)
	errorCount := 0

	dogs, dErr := c.getDogs()
	if dErr == nil {
		a = append(a, dogs...)
	} else {
		log.Printf("Error getting dogs: %s", dErr.Error())
		errorCount++
	}

	cats, cErr := c.getCats()
	if cErr == nil {
		a = append(a, cats...)
	} else {
		log.Printf("Error getting dogs: %s", cErr.Error())
		errorCount++
	}

	hams, hErr := c.getHamsters()
	if hErr == nil {
		a = append(a, hams...)
	} else {
		log.Printf("Error getting hamsters: %s", hErr.Error())
		errorCount++
	}

	if errorCount == 3 {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	b, err := json.Marshal(a)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(b)
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

func (c *Client) getCats() ([]Animal, error) {
	resp, err := c.client.Get(c.baseURL + "/cats")
	if err != nil {
		log.Printf("Error doing get: %s", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	var cats CatList
	if err = json.NewDecoder(resp.Body).Decode(&cats); err != nil {
		log.Printf("Error reading json: %s", err.Error())
		return nil, err
	}

	group := make(map[string]Cats)

	for _, cat := range cats.Values {
		t, err := time.Parse("2006-01-02", cat.DOB)
		if err != nil {
			log.Printf("Error parsing time %s", err.Error())
			return nil, err
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

	return result, nil
}

func (c *Client) getHamsters() ([]Animal, error) {
	resp, err := c.client.Get(c.baseURL + "/hamsters")
	if err != nil {
		log.Printf("Error doing get: %s", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	var hams HamsterList
	if err = json.NewDecoder(resp.Body).Decode(&hams); err != nil {
		log.Printf("Error reading json: %s", err.Error())
		return nil, err
	}

	for _, ham := range hams.Values {
		t, err := time.Parse("2006-01-02", ham.DOB)
		if err != nil {
			log.Printf("Error parsing time %s", err.Error())
			return nil, err
		}

		ham.Age = t
	}

	sort.Sort(Hamsters(hams.Values))
	return hams.Values.ToAnimals(), nil
}
