package main

import "fmt"

func main() {

	// var x map[string]int --- declaration
	// x["key"] = 10 ---------- produces an error
	// fmt.Println(x)

	x := make(map[string]int)
	x["key"] = 10
	// fmt.Println(x)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	// fmt.Println(elements["Li"])

	fmt.Println(elements["Un"]) // gives a 0 element
	if elements["Un"] == "" {
		fmt.Println("key not available")
	} // basic way

	// The Way to =Go
	value, ok := elements["Un"]
	if ok {
		fmt.Println("The Value is : ", value)
	} else {
		fmt.Println("Key is not present in the map")
	}

	// A much better way to =Go
	if value, ok := elements["Un"]; ok {
		fmt.Println("Value is : ", value)
	}

	// Easier way to initialize stuff
	element := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	fmt.Println(element)

	elemention := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": {
			"name":  "Boron",
			"state": "solid",
		},
		"C": {
			"name":  "Carbon",
			"state": "solid",
		},
		"N": {
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": {
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": {
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": {
			"name":  "Neon",
			"state": "gas",
		},
	}

	fmt.Println(elemention)

}
