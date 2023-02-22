package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Animal struct {
	Species string `json:"species"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
}

func zooAnimals(data [][]string) []Animal {
	var zooAnimals []Animal
	for i, line := range data {
		if i == 0 {
			continue
		}

		var rec Animal

		rec.Species = line[0]
		rec.Name = line[1]
		rec.Age, _ = strconv.Atoi(line[2])

		zooAnimals = append(zooAnimals, rec)
	}
	return zooAnimals
}

func main() {
	f, err := os.Open("/usr/local/src/wb/l1/t21/data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	zooAnimals := zooAnimals(data)

	jsonData, err := json.MarshalIndent(zooAnimals, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(jsonData))
}
