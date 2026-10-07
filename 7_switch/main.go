package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("unknown")
	}

	// multiple-condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("work-day")
	}

	// type-switch
	// interface means any type can be assign => interface can be replaced with any
	whoAmI := func (i any)  {
		switch i.(type) { // return variable type
		case int:
			fmt.Println("integer")
		case string:
			fmt.Println("string")
		default:
			fmt.Printf("%T \n", i) // -> to see the type
		}  
	}

	whoAmI("golang")
	whoAmI(true)
}
