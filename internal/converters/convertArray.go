package converters

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/neet-007/objectsToTsTypes/internal/helpers"
	"github.com/neet-007/objectsToTsTypes/internal/typesmap"
)

func ConvertArray(arr []interface{}, padding int) string {
	arrTypes := make([]string, 0)

	for _, val := range arr {
		valType := ""

		if nestedArr, ok := val.([]interface{}); ok {
			valType = fmt.Sprintf("(%s)", ConvertArray(nestedArr, padding+1))
		} else if nestedStructMap, ok := val.(map[string]interface{}); ok {
			valType = ConvertInterface(nestedStructMap, padding+1, "arr")
		} else {
			valType = fmt.Sprintf("(%s)", typesmap.TsTypes[reflect.TypeOf(val).String()])
		}
		arrTypes = append(arrTypes, valType)
	}

	arrTypes = helpers.RemoveDuplicates(arrTypes)
	var typesToStr string

	if len(arrTypes) == 0 {
		typesToStr = "any"
	} else {
		typesToStr = strings.Join(arrTypes, " | ")
	}

	arrWithTypes := strings.Trim(fmt.Sprintf("%s []", typesToStr), " | ")

	return arrWithTypes
}
