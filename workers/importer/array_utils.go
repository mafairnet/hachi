package main

func removeDuplicatesSliceString(stringSlicearray []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlicearray {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
