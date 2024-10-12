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

	unmarshaledJson := make(map[string]interface{})
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

func convertTypes(unmarshaledJson map[string]interface{}) map[string]string {
	fmt.Println("the json")
	for key, val := range unmarshaledJson {
		fmt.Printf("key: %s, val: %v\n", key, val)
	}
	fmt.Println()

	tsType := make(map[string]string)

	for key, val := range unmarshaledJson {
		fmt.Printf("type of temp:%v -> %v\n", val, reflect.TypeOf(val).String())

		if arr, ok := val.([]interface{}); ok {
			tsType[key] = convertArray(arr)
		} else if structMap, ok := val.(map[string]interface{}); ok {
			tsType[key] = convertStruct(structMap)
		} else {
			tsType[key] = Types[reflect.TypeOf(val).String()]
		}
	}

	return tsType
}

func convertStruct(structMap map[string]interface{}) string {
	keyVal := make(map[string]string)

	for key, val := range structMap {
		fmt.Printf("Key: %v, Value: %v\n", key, val)

		if nestedStructMap, ok := val.(map[string]interface{}); ok {
			convertStruct(nestedStructMap)
		} else {
			keyVal[key] = Types[reflect.TypeOf(val).String()]
		}
	}

	returnString := "{\n"
	for key, val := range keyVal {
		returnString += fmt.Sprintf(" %s : %s\n", key, val)
	}
	returnString += "}"
	return returnString
}

func convertArray(arr []interface{}) string {
	arrTypes := make([]string, 0)

	for _, val := range arr {
		valType := ""

		if nestedArr, ok := val.([]interface{}); ok {
			valType = fmt.Sprintf("(%s)", convertArray(nestedArr))
		} else {
			valType = fmt.Sprintf("(%s)", Types[reflect.TypeOf(val).String()])
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
