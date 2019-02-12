package main

import (
	"fmt"
	"github.com/chowanij/ioExamples/model"
	"io/ioutil"
	"os"
)

func main() {
	var osoba model.Person
	osoba.Name = "Jan"
	osoba.SecondName = "Chowaniec"
	osoba.Id = "85062612345"
	osoba.IdType = "PESEL"
	osoba.Address = "Os. Za Torem 26; Rogoźnik; 34-471 Ludźmierz"
	osoba.Age = 33
	osoba.ShowBasicInfo()
	// casting osoba struct to byte slice (stream)
	var structBytes = []byte(fmt.Sprintf("%v", osoba))

	// open output file
	fo, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("output.txt", structBytes, 0644)
	if err != nil {
		panic(err)
	}

	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

}
