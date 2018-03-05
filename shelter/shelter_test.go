package shelter

import (
    "net/http"
    "testing"
    "net/http/httptest"
    "time"
    "encoding/json"
)

func TestGetAnimals(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
        case "/dogs":
            w.Write(createMockDogs())
        case "/cats":
            w.Write(createMockCats())
        case "/hamsters":
            w.Write(createMockHamsters())
        }
	}))

    defer ts.Close()

    c := NewClient(ts.URL)

    req := httptest.NewRequest("GET", "http://localhost:4000/animals", nil)
    w := httptest.NewRecorder()
    c.getAnimals(w, req)

    resp := w.Result()

    if resp.StatusCode != http.StatusOK {
        t.Errorf("Incorrect status code: expected [%d], got [%d]",
            http.StatusOK, resp.StatusCode)
    }

    defer resp.Body.Close()

    var animals []Animal
    if err := json.NewDecoder(resp.Body).Decode(&animals); err != nil {
	    t.Error(err)
	}

    if len(animals) != 8 {
        t.Errorf("Incorrect number of items: expected [%d], got [%d]",
            8, len(animals))
    }

    if animals[0].Name != "Older Dog" {
        t.Errorf("Incorrect dog: expected [%s], got [%s]",
            "Older Dog", animals[0].Name)
    }

    if animals[1].Name != "Younger Dog" {
        t.Errorf("Incorrect dog: expected [%s], got [%s]",
            "Younger Dog", animals[1].Name)
    }

    if animals[2].Name != "Ginger Cat" {
        t.Errorf("Incorrect cat: expected [%s], got [%s]",
            "Ginger Cat", animals[2].Name)
    }

    if animals[3].Name != "Older black Cat" {
        t.Errorf("Incorrect cat: expected [%s], got [%s]",
            "Older black Cat", animals[3].Name)
    }

    if animals[4].Name != "Younger black Cat" {
        t.Errorf("Incorrect cat: expected [%s], got [%s]",
            "Younger black Cat", animals[0].Name)
    }

    if animals[5].Name != "White Cat" {
        t.Errorf("Incorrect cat: expected [%s], got [%s]",
            "White Cat", animals[5].Name)
    }

    if animals[6].Name != "Younger Hamster" {
        t.Errorf("Incorrect hamster: expected [%s], got [%s]",
            "Younger Hamster", animals[6].Name)
    }

    if animals[7].Name != "Older Hamster" {
        t.Errorf("Incorrect hamster: expected [%s], got [%s]",
            "Older Hamster", animals[7].Name)
    }
}

func TestGetAnimalsNoDogs(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
        case "/dogs":
            w.WriteHeader(http.StatusBadRequest)
        case "/cats":
            w.Write(createMockCats())
        case "/hamsters":
            w.Write(createMockHamsters())
        }
	}))

    defer ts.Close()

    c := NewClient(ts.URL)

    req := httptest.NewRequest("GET", "http://localhost:4000/animals", nil)
    w := httptest.NewRecorder()
    c.getAnimals(w, req)

    resp := w.Result()

    if resp.StatusCode != http.StatusOK {
        t.Errorf("Incorrect status code: expected [%d], got [%d]",
            http.StatusOK, resp.StatusCode)
    }

    defer resp.Body.Close()

    var animals []Animal
    if err := json.NewDecoder(resp.Body).Decode(&animals); err != nil {
	    t.Error(err)
	}

    if len(animals) != 6 {
        t.Errorf("Incorrect number of items: expected [%d], got [%d]",
            8, len(animals))
    }

    if animals[0].Name != "Ginger Cat" {
        t.Errorf("Incorrect cat: expected [%s], got [%s]",
            "Ginger Cat", animals[0].Name)
    }

    if animals[1].Name != "Older black Cat" {
        t.Errorf("Incorrect cat: expected [%s], got [%s]",
            "Older black Cat", animals[1].Name)
    }

    if animals[2].Name != "Younger black Cat" {
        t.Errorf("Incorrect cat: expected [%s], got [%s]",
            "Younger black Cat", animals[2].Name)
    }

    if animals[3].Name != "White Cat" {
        t.Errorf("Incorrect cat: expected [%s], got [%s]",
            "White Cat", animals[3].Name)
    }

    if animals[4].Name != "Younger Hamster" {
        t.Errorf("Incorrect hamster: expected [%s], got [%s]",
            "Younger Hamster", animals[4].Name)
    }

    if animals[5].Name != "Older Hamster" {
        t.Errorf("Incorrect hamster: expected [%s], got [%s]",
            "Older Hamster", animals[5].Name)
    }
}

func TestGetAnimalsNoCats(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
        case "/dogs":
            w.Write(createMockDogs())
        case "/cats":
            w.WriteHeader(http.StatusBadRequest)
        case "/hamsters":
            w.Write(createMockHamsters())
        }
	}))

    defer ts.Close()

    c := NewClient(ts.URL)

    req := httptest.NewRequest("GET", "http://localhost:4000/animals", nil)
    w := httptest.NewRecorder()
    c.getAnimals(w, req)

    resp := w.Result()

    if resp.StatusCode != http.StatusOK {
        t.Errorf("Incorrect status code: expected [%d], got [%d]",
            http.StatusOK, resp.StatusCode)
    }

    defer resp.Body.Close()

    var animals []Animal
    if err := json.NewDecoder(resp.Body).Decode(&animals); err != nil {
	    t.Error(err)
	}

    if len(animals) != 4 {
        t.Errorf("Incorrect number of items: expected [%d], got [%d]",
            8, len(animals))
    }

    if animals[0].Name != "Older Dog" {
        t.Errorf("Incorrect dog: expected [%s], got [%s]",
            "Older Dog", animals[0].Name)
    }

    if animals[1].Name != "Younger Dog" {
        t.Errorf("Incorrect dog: expected [%s], got [%s]",
            "Younger Dog", animals[1].Name)
    }

    if animals[2].Name != "Younger Hamster" {
        t.Errorf("Incorrect hamster: expected [%s], got [%s]",
            "Younger Hamster", animals[4].Name)
    }

    if animals[3].Name != "Older Hamster" {
        t.Errorf("Incorrect hamster: expected [%s], got [%s]",
            "Older Hamster", animals[5].Name)
    }
}

func TestGetAnimalsNoHamsters(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
        case "/dogs":
            w.Write(createMockDogs())
        case "/cats":
            w.Write(createMockCats())
        case "/hamsters":
            w.WriteHeader(http.StatusBadRequest)
        }
	}))

    defer ts.Close()

    c := NewClient(ts.URL)

    req := httptest.NewRequest("GET", "http://localhost:4000/animals", nil)
    w := httptest.NewRecorder()
    c.getAnimals(w, req)

    resp := w.Result()

    if resp.StatusCode != http.StatusOK {
        t.Errorf("Incorrect status code: expected [%d], got [%d]",
            http.StatusOK, resp.StatusCode)
    }

    defer resp.Body.Close()

    var animals []Animal
    if err := json.NewDecoder(resp.Body).Decode(&animals); err != nil {
	    t.Error(err)
	}

    if len(animals) != 6 {
        t.Errorf("Incorrect number of items: expected [%d], got [%d]",
            8, len(animals))
    }

    if animals[0].Name != "Older Dog" {
        t.Errorf("Incorrect dog: expected [%s], got [%s]",
            "Older Dog", animals[0].Name)
    }

    if animals[1].Name != "Younger Dog" {
        t.Errorf("Incorrect dog: expected [%s], got [%s]",
            "Younger Dog", animals[1].Name)
    }

    if animals[2].Name != "Ginger Cat" {
        t.Errorf("Incorrect cat: expected [%s], got [%s]",
            "Ginger Cat", animals[2].Name)
    }

    if animals[3].Name != "Older black Cat" {
        t.Errorf("Incorrect cat: expected [%s], got [%s]",
            "Older black Cat", animals[3].Name)
    }

    if animals[4].Name != "Younger black Cat" {
        t.Errorf("Incorrect cat: expected [%s], got [%s]",
            "Younger black Cat", animals[4].Name)
    }

    if animals[5].Name != "White Cat" {
        t.Errorf("Incorrect cat: expected [%s], got [%s]",
            "White Cat", animals[5].Name)
    }
}

func TestGetAnimalsAllFail(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		 w.WriteHeader(http.StatusBadRequest)
	}))

    defer ts.Close()

    c := NewClient(ts.URL)

    req := httptest.NewRequest("GET", "http://localhost:4000/animals", nil)
    w := httptest.NewRecorder()
    c.getAnimals(w, req)

    resp := w.Result()

    if resp.StatusCode != http.StatusBadGateway {
        t.Errorf("Incorrect status code: expected [%d], got [%d]",
            http.StatusBadGateway, resp.StatusCode)
    }

}

func createMockDogs() []byte {
    d1 := Dog{
        Forename:"Younger",
        Surname: "Dog",
        DOB: time.Now().AddDate(-1, 0, 0).Format("2006-01-02"),
        Image: Image{"youngerdog.png"},
    }

    d2 := Dog{
        Forename: "Older",
        Surname: "Dog",
        DOB: time.Now().AddDate(-1, -3, 0).Format("2006-01-02"),
        Image: Image{"olderdog.png"},
    }

    dogs := make([]*Dog, 0)
    dogs = append(dogs, &d1, &d2)
    dogList := DogList{dogs}

    b, _ := json.Marshal(dogList)
    return b
}

func createMockCats() []byte {
    c1 := Cat{
        Forename: "Younger black",
        Surname: "Cat",
        DOB: time.Now().AddDate(-1, 0, 0).Format("2006-01-02"),
        Image: Image{"youngerblackcat.png"},
        Colour: "black",
    }

    c2 := Cat{
        Forename: "Ginger",
        Surname: "Cat",
        DOB: time.Now().AddDate(-1, 0, 0).Format("2006-01-02"),
        Image: Image{"gingercat.png"},
        Colour: "ginger",
    }

    c3 := Cat{
        Forename: "White",
        Surname: "Cat",
        DOB: time.Now().AddDate(-1, 0, 0).Format("2006-01-02"),
        Image: Image{"whitecat.png"},
        Colour: "white",
    }

    c4 := Cat{
        Forename: "Older black",
        Surname: "Cat",
        DOB: time.Now().AddDate(-1, -3, 0).Format("2006-01-02"),
        Image: Image{"olderblackcat.png"},
        Colour: "black",
    }

    cats := make([]*Cat, 0)
    cats = append(cats, &c1, &c2, &c3, &c4)
    catList := CatList{cats}

    b, _ := json.Marshal(catList)
    return b
}

func createMockHamsters() []byte {
    h1 := Hamster{
        Forename: "Younger",
        Surname: "Hamster",
        DOB: time.Now().AddDate(-1, 0, 0).Format("2006-01-02"),
        Image: Image{"youngerhamster.png"},
    }

    h2 := Hamster{
        Forename: "Older",
        Surname: "Hamster",
        DOB: time.Now().AddDate(-1, -3, 0).Format("2006-01-02"),
        Image: Image{"olderhamster.png"},
    }

    hamsters := make([]*Hamster, 0)
    hamsters = append(hamsters, &h1, &h2)
    hamsterList := HamsterList{hamsters}

    b, _ := json.Marshal(hamsterList)
    return b
}
