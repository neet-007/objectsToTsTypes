package converters

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/neet-007/objectsToTsTypes/internal/typesmap"
)

func ConvertInterface(structMap map[string]interface{}, padding int) string {
	keyVal := make(map[string]string)

	for key, val := range structMap {
		if nestedArr, ok := val.([]interface{}); ok {
			keyVal[key] = ConvertArray(nestedArr)
		} else if nestedStructMap, ok := val.(map[string]interface{}); ok {
			keyVal[key] = ConvertInterface(nestedStructMap, padding+1)
		} else {
			keyVal[key] = typesmap.TsTypes[reflect.TypeOf(val).String()]
		}
	}

	returnString := "{\n"
	paddingStr := strings.Repeat(string(" "), padding)
	bracePaddingStr := strings.Repeat(string(" "), padding-1)

	for key, val := range keyVal {
		returnString += fmt.Sprintf("%s%s : %s\n", paddingStr, key, val)
	}

	returnString += fmt.Sprintf("%s}", bracePaddingStr)
	return returnString
}
