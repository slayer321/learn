package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	//jsonMarshal()
	//exercise125()
	//exercise126()
	//exercise127()
	//sortExercse()
	exercise129()
}

type user struct {
	First   string `json:"first_name"`
	Age     int    `json:"age"`
	Surname string `json:"surname,omitempty"`
}

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func exercise129() {
	u1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := person{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []person{u1, u2, u3}
	//fmt.Println(users)
	sort.Slice(users, func(i, j int) bool {
		if users[i].Age < users[j].Age {
			return true
		}
		return false
	})

	fmt.Println(users)

	sort.Slice(users, func(i, j int) bool {
		if users[i].Last < users[j].Last {
			return true
		}
		return false
	})

	fmt.Println(users)

}

func sortExercse() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)

}

func exercise127() {
	u1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := person{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []person{u1, u2, u3}

	//fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		log.Printf("error: can't encode - %s", err)
	}
	//fmt.Println(users)
}

func exercise126() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	//fmt.Println(s)
	var users []person
	err := json.Unmarshal([]byte(s), &users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", users)
}

func exercise125() {
	u1 := user{
		First:   "James",
		Age:     32,
		Surname: "Bond",
	}

	u2 := user{
		First:   "Moneypenny",
		Age:     27,
		Surname: "G",
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	bytes, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}

type FruitBasket struct {
	Name    string
	Fruit   []string
	Id      int64  `json:"ref"`
	private string // An unexported field is not encoded.
	Created time.Time
}

func jsonMarshal() {
	basket := FruitBasket{
		Name:    "Standard",
		Fruit:   []string{"Apple", "Banana", "Orange"},
		Id:      999,
		private: "Second-rate",
		Created: time.Now(),
	}

	var jsonData []byte
	jsonData, err := json.Marshal(basket)

	//jsonData, err := json.MarshalIndent(basket, "", "  ")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Marshal Output: %+v\n", string(jsonData))

	var basket2 FruitBasket

	err = json.Unmarshal(jsonData, &basket2)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Unmarshal Output: %+v\n", basket2)
}
