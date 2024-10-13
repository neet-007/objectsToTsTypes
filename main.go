package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/neet-007/objectsToTsTypes/internal/converters"
	"github.com/neet-007/objectsToTsTypes/internal/helpers"
)

func main() {
	file, err := os.ReadFile("./test.json")
	if err != nil {
		log.Fatal("could not read file")
	}

	unmarshaledJson := make(map[string]interface{})
	if err := json.Unmarshal(file, &unmarshaledJson); err != nil {
		log.Fatal("couldn not unmarshal json")
	}

	types := converters.ConvertTypes(unmarshaledJson)

	data := helpers.FormatTypes(types)
	err = os.WriteFile("ts_types.ts", data, 0666)
	if err != nil {
		log.Fatal("could not write to file")
	}
}
