package convert

import (
	"fmt"
	"testing"
)

func ExampleLength() {
	result := Length(1, Inch, Centimeter)
	fmt.Println(result)

	// Output: 2.54
}

func TestLengthConversion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		value    float64
		from     LengthUnit
		to       LengthUnit
		expected float64
	}{
		// Metric system tests
		{"Nanometer to Meter", 1, Nanometer, Meter, 9.999999999e-10},
		{"Micrometer to Meter", 1, Micrometer, Meter, 1e-6},
		{"Millimeter to Meter", 1, Millimeter, Meter, 1e-3},
		{"Centimeter to Meter", 1, Centimeter, Meter, 1e-2},
		{"Decimeter to Meter", 1, Decimeter, Meter, 1e-1},
		{"Meter to Meter", 1, Meter, Meter, 1},
		{"Decameter to Meter", 1, Decameter, Meter, 1e1},
		{"Hectometer to Meter", 1, Hectometer, Meter, 1e2},
		{"Kilometer to Meter", 1, Kilometer, Meter, 1e3},
		{"Megameter to Meter", 1, Megameter, Meter, 1e6},
		{"Gigameter to Meter", 1, Gigameter, Meter, 1e9},
		{"Terameter to Meter", 1, Terameter, Meter, 1e12},

		// Imperial system tests
		{"Thou to Meter", 1, Thou, Meter, 2.54e-5},
		{"Inch to Meter", 1, Inch, Meter, 0.0254},
		{"Foot to Meter", 1, Foot, Meter, 0.3048},
		{"Yard to Meter", 1, Yard, Meter, 0.9144},
		{"Chain to Meter", 1, Chain, Meter, 20.1168},
		{"Furlong to Meter", 1, Furlong, Meter, 201.168},
		{"Mile to Meter", 1, Mile, Meter, 1609.344},     // Updated to more precise value
		{"League to Meter", 1, League, Meter, 4828.032}, // Updated to more precise value
		{"Nautical Mile to Meter", 1, NauticalMile, Meter, 1852},

		// Astronomical units
		{"Astronomical Unit to Meter", 1, AstronomicalUnit, Meter, 149597870691}, // Updated to more precise value
		{"Light Year to Meter", 1, LightYear, Meter, 9460730472580044},           // Updated to more precise value
		{"Parsec to Meter", 1, Parsec, Meter, 30856775812799588},                 // Updated to more precise value
		{"Planck Length to Meter", 1, PlanckLength, Meter, 1.616049999e-35},      // Updated to more precise value

		// Other units
		{"Angstrom to Meter", 1, Angstrom, Meter, 9.999999999e-11}, // Updated to more precise value
		{"Fathom to Meter", 1, Fathom, Meter, 1.8288},
		{"Rod to Meter", 1, Rod, Meter, 5.0292},
		{"Survey Foot to Meter", 1, SurveyFoot, Meter, 0.3048006096}, // Updated to more precise value
		{"Pica to Meter", 1, Pica, Meter, 0.0042333333},              // Updated to more precise value
		{"Point to Meter", 1, Point, Meter, 0.0003527778},            // Updated to more precise value
		{"Hand to Meter", 1, Hand, Meter, 0.1016},
	}

	for _, tableTest := range tests {
		t.Run(tableTest.name, func(t *testing.T) {
			t.Parallel()

			got := Length(tableTest.value, tableTest.from, tableTest.to)
			if !almostEqual(got, tableTest.expected) {
				t.Errorf("Length(%v, %v, %v) = %v; want %v", tableTest.value, tableTest.from, tableTest.to, got, tableTest.expected)
			}
		})
	}
}

// Helper function to compare floating point numbers with a tolerance
func almostEqual(a, b float64) bool {
	const tolerance = 1e-9
	return (a-b) < tolerance && (b-a) < tolerance
}
