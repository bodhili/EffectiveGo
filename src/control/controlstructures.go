package control

import "fmt"

// while

func while() {
	for {
		fmt.Println("1")
	}
}

func while0() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func while2() {
	params := make(map[string]string)
	params["A"] = "A"
	params["B"] = "B"
	params["C"] = "C"

	for key, v := range params {
		fmt.Printf("key: %s value:%v", key, v)
	}
}

func switch0(c string) {
	var t interface{}
	t = c
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T", t)
	case bool:
		fmt.Printf("boolean %t\n", t)
	case int:
		fmt.Printf("integer %d\n", t)
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t)
	case *int:
		fmt.Printf("pointer to integer %d\n", *t)
	}
}
