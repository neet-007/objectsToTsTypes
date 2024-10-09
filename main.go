package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

var types map[string]string

func init() {
	types = map[string]string{
		"string":  "string",
		"int":     "number",
		"float64": "number",
		"bool":    "boolean",
	}
}
func main() {
	file, err := os.ReadFile("./test.json")
	if err != nil {
		log.Fatal("could not read file")
	}

	unmarshaledJson := make(map[string]*json.RawMessage)
	if err := json.Unmarshal(file, &unmarshaledJson); err != nil {
		log.Fatal("couldn not unmarshal json")
	}

	fmt.Println("the json")
	for key, val := range unmarshaledJson {
		fmt.Printf("key: %s, val: %s\n", key, val)
	}
	fmt.Println()

	tsType := make(map[string]string)

	for key, val := range unmarshaledJson {
		var temp interface{}
		err := json.Unmarshal(*val, &temp)
		if err != nil {
			fmt.Printf("Error unmarshaling key %s: %v\n", key, err)
			continue
		}

		tsType[key] = types[reflect.TypeOf(temp).String()]
	}

	data := []byte{}
	for _, c := range "type NewType = {" {
		data = append(data, byte(c))
	}
	data = append(data, '\n')

	fmt.Println("the type")
	for key, val := range tsType {
		fmt.Printf("key: %s, val: %s\n", key, val)

		for _, c := range key {
			data = append(data, byte(c))
		}
		data = append(data, ':')
		for _, c := range val {
			data = append(data, byte(c))
		}
		data = append(data, '\n')
	}
	data = append(data, '}')

	err = os.WriteFile("ts_types.ts", data, 0666)
	if err != nil {
		log.Fatal("could not write to file")
	}

}
