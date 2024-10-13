package helpers

func FormatTypes(types map[string]string) []byte {
	data := []byte{}
	for _, c := range "type NewType = {" {
		data = append(data, byte(c))
	}
	data = append(data, '\n')

	for key, val := range types {
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

func RemoveDuplicates(slice []string) []string {
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
