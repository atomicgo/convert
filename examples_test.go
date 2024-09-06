package convert_test

import (
	"fmt"

	"atomicgo.dev/convert"
)

func Example_demo() {
}

func ExampleLength() {
	result := convert.Length(1, convert.Inch, convert.Centimeter)
	fmt.Println(result)

	// Output: 2.54
}
