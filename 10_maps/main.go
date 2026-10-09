package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// if key doesn't exist in the map then it returns "zero" value
	fmt.Println(m["name"], m["area"], m["phone"])

	fmt.Println(len(m))

	delete(m, "area")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)

	mpp := map[string]int{
		"price":  12,
		"phones": 300,
	}

	mpp2 := map[string]int{
		"price":  12,
		"phones": 300,
	}

	fmt.Println(mpp)

	_, exist := mpp["price"]

	if exist {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	fmt.Println(maps.Equal(mpp, mpp2))
}
