package convert

import "fmt"

func ExamplePercentageToProportion() {
	fmt.Println(PercentageToProportion(50, 1)) // => 0.5x
	fmt.Println(PercentageToProportion(55, 1)) // => 0.6x
	fmt.Println(PercentageToProportion(55, 2)) // => 0.55x

	// Output:
	// 0.5x
	// 0.6x
	// 0.55x
}

func ExampleProportionToPercentage() {
	percentage, _ := ProportionToPercentage("0.5x")
	fmt.Println(percentage)

	// Output: 50
}
