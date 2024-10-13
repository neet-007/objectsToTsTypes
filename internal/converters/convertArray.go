package converters

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/neet-007/objectsToTsTypes/internal/helpers"
	"github.com/neet-007/objectsToTsTypes/internal/typesmap"
)

func ConvertArray(arr []interface{}) string {
	arrTypes := make([]string, 0)

	for _, val := range arr {
		valType := ""

		if nestedArr, ok := val.([]interface{}); ok {
			valType = fmt.Sprintf("(%s)", ConvertArray(nestedArr))
		} else if nestedStructMap, ok := val.(map[string]interface{}); ok {
			valType = ConvertInterface(nestedStructMap, 1)
		} else {
			valType = fmt.Sprintf("(%s)", typesmap.TsTypes[reflect.TypeOf(val).String()])
		}
		arrTypes = append(arrTypes, valType)
	}

	arrTypes = helpers.RemoveDuplicates(arrTypes)

	arrWithTypes := strings.Trim(fmt.Sprintf("%s []", strings.Join(arrTypes, " | ")), " | ")
	fmt.Printf("\n\n%s\n\n", arrWithTypes)

	return arrWithTypes
}
