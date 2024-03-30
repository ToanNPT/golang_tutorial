package main

import (
	"fmt"
)

/*
Allowed Key Types
The map key can be of any data type for which the equality operator (==) is defined. These include:

Booleans
Numbers
Strings
Arrays
Pointers
Structs
Interfaces (as long as the dynamic type supports equality)
*/
func main() {
	var userRefs map[string]string = map[string]string{"EMP001": "John Doe", "EMP002": "Jane Doe", "EMP003": "John Smith"}
	for key, value := range userRefs {
		println("Key===>", key, "Value:", value)
	}

	userRefs["EMP001"] = "ToanNPT"

	keys := []string{"EMP001", "EMP002", "EMP0034"}

	for index, key := range keys {
		if _, isExisted := userRefs[key]; isExisted {
			fmt.Printf("Index %d: Key %s: Value %s\n", index, key, userRefs[key])
		} else {
			fmt.Printf("Index %d: Key %s not found\n", index, key)
		}

	}

}
