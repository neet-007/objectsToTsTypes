package converters

import (
	"reflect"

	"github.com/neet-007/objectsToTsTypes/internal/typesmap"
)

func ConvertTypes(unmarshaledJson map[string]interface{}) map[string]string {
	tsType := make(map[string]string)

	for key, val := range unmarshaledJson {
		if arr, ok := val.([]interface{}); ok {
			tsType[key] = ConvertArray(arr, 2)
		} else if structMap, ok := val.(map[string]interface{}); ok {
			tsType[key] = ConvertInterface(structMap, 2, key)
		} else {
			tsType[key] = typesmap.TsTypes[reflect.TypeOf(val).String()]
		}
	}

	return tsType
}
