package helpers

import (
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
)

func FormatTypes(types map[string]string) []byte {
	data := []byte{}
	for _, c := range "type NewType = {" {
		data = append(data, byte(c))
	}
	data = append(data, '\n')

	for key, val := range types {
		data = append(data, '\t')
		for _, c := range key {
			data = append(data, byte(c))
		}
		data = append(data, ':')
		data = append(data, ' ')
		for _, c := range val {
			data = append(data, byte(c))
		}
		data = append(data, '\n')
	}
	data = append(data, '}')

	return data
}

func RemoveDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	nestedObjects := []string{}

	for _, val := range slice {
		if _, ok := seen[val]; !ok {
			seen[val] = true
			result = append(result, val)
			if strings.Trim(val, " ")[0] == '{' {
				nestedObjects = append(nestedObjects, val)
			}
		}
	}

	fmt.Println("beffffffffore")
	fmt.Println(result)
	for i := 0; i < len(nestedObjects); i++ {
		for j := i + 1; j < len(nestedObjects); j++ {
			diff := cmp.Diff(nestedObjects[i], nestedObjects[j])
			plusIndex := strings.Index(diff, "+")
			minusIndex := strings.Index(diff, "-")
			if plusIndex == -1 || minusIndex == -1 {
				continue
			}

			lineOne := []byte{}
			lineTwo := []byte{}
			for i := plusIndex; i < len(diff); i++ {
				if diff[i] == '\n' {
					break
				}
				lineOne = append(lineOne, diff[i])
			}

			for i := minusIndex; i < len(diff); i++ {
				if diff[i] == '\n' {
					break
				}
				lineTwo = append(lineTwo, diff[i])
			}

			if strings.Contains(string(lineOne), "any") {
				fmt.Println("first any")
				result = removeElement(result, nestedObjects[j])
			} else {
				fmt.Println("second any")
				result = removeElement(result, nestedObjects[i])
			}
		}
	}

	fmt.Println("afffffffter")
	fmt.Println(result)
	return result
}

func removeElement(slice []string, value string) []string {
	result := []string{}

	for _, v := range slice {
		if v != value {
			result = append(result, v)
		}
	}

	return result
}
