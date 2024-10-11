package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

var Types map[string]string

func init() {
	Types = map[string]string{
		"string":         "string",
		"int":            "number",
		"int8":           "number",
		"int16":          "number",
		"int32":          "number",
		"int64":          "number",
		"uint":           "number",
		"uint8":          "number",
		"uint16":         "number",
		"uint32":         "number",
		"uint64":         "number",
		"float32":        "number",
		"float64":        "number",
		"bool":           "boolean",
		"byte":           "number",
		"rune":           "string",
		"interface{}":    "any",
		"[]interface {}": "any []",
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

	types := convertTypes(unmarshaledJson)
	fmt.Println(types)

	data := formatTypes(types)
	err = os.WriteFile("ts_types.ts", data, 0666)
	if err != nil {
		log.Fatal("could not write to file")
	}
}

func convertTypes(unmarshaledJson map[string]*json.RawMessage) map[string]string {
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

		fmt.Printf("type of temp:%v -> %v\n", temp, reflect.TypeOf(temp).String())

		if reflect.TypeOf(temp).String() == "[]interface {}" {
			convertArray(val)
		}

		tsType[key] = Types[reflect.TypeOf(temp).String()]
	}

	return tsType
}

func convertStruct() {

}

func convertArray(arr *json.RawMessage) string {
	var temp []*json.RawMessage
	err := json.Unmarshal(*arr, &temp)
	if err != nil {
		fmt.Println("could not unmarshal json arr")
		return ""
	}

	arrTypes := make([]string, len(temp))

	for _, val := range temp {
		var tempVal interface{}
		err = json.Unmarshal(*val, &tempVal)
		if err != nil {
			fmt.Printf("could not unmarshal json val %v\n", val)
			return ""
		}

		valType := ""

		if reflect.TypeOf(tempVal).String() == "[]interface {}" {
			var tempValJson *json.RawMessage
			err = json.Unmarshal(*val, &tempValJson)
			if err != nil {
				fmt.Printf("could not unmarshal json val %v\n", val)
				return ""
			}

			valType = fmt.Sprintf("(%s)", convertArray(tempValJson))
		} else {
			valType = fmt.Sprintf("(%s)", Types[reflect.TypeOf(tempVal).String()])
		}
		arrTypes = append(arrTypes, valType)

	}

	arrTypes = removeDuplicates(arrTypes)

	arrWithTypes := strings.Trim(fmt.Sprintf("%s []", strings.Join(arrTypes, " | ")), " | ")
	fmt.Printf("\n\n%s\n\n", arrWithTypes)

	return arrWithTypes
}

func formatTypes(types map[string]string) []byte {
	data := []byte{}
	for _, c := range "type NewType = {" {
		data = append(data, byte(c))
	}
	data = append(data, '\n')

	fmt.Println("the type")
	for key, val := range types {
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

	return data
}

func removeDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, val := range slice {
		if _, ok := seen[val]; !ok {
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}
