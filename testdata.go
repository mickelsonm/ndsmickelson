package ndsmickelson

import (
	"encoding/json"
)

//People is a collection of persons
type People []Person

//Person is a person
type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

//GetTestPeople gives us some test data of people
func GetTestPeople() People {
	return People{
		Person{ID: 18, Name: "Frank", Email: "frank@hotmail.com"},
		Person{ID: 23, Name: "Susan", Email: "susan@gmail.com"},
		Person{ID: 123, Name: "Marshall", Email: "marshall@wisc.net"},
		Person{ID: 777, Name: "Jack", Email: "jack@beanstalk.com"},
		Person{ID: 812, Name: "Michelle"},
	}
}

//ToJSON the collection of people
func (p *People) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}
