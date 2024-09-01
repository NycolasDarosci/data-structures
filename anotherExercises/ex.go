package anotherexercises

func CombineTwoCharacters(arr []string) []string {
	collection := []string{}
	for i := range arr {
		for j := range arr {
			if i != j {
				collection = append(collection, arr[i]+arr[j])
			}
		}
	}
	return collection
}
