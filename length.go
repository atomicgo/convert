package convert

type LengthUnit float64

const (
	// Metric system
	Nanometer  LengthUnit = 9.999999999e-10     // Nanometer to meter conversion
	Micrometer LengthUnit = 0.000001            // Micrometer to meter conversion
	Millimeter LengthUnit = 0.001               // Millimeter to meter conversion
	Centimeter LengthUnit = 0.01                // Centimeter to meter conversion
	Decimeter  LengthUnit = 0.1                 // Decimeter to meter conversion
	Meter      LengthUnit = 1                   // Meter is the base unit
	Decameter  LengthUnit = 10                  // Decameter to meter conversion
	Hectometer LengthUnit = 100                 // Hectometer to meter conversion
	Kilometer  LengthUnit = 1000                // Kilometer to meter conversion
	Megameter  LengthUnit = 1000000             // Megameter to meter conversion
	Gigameter  LengthUnit = 1000000000          // Gigameter to meter conversion
	Terameter  LengthUnit = 1000000000000       // Terameter to meter conversion
	Petameter  LengthUnit = 1000000000000000    // Petameter to meter conversion
	Exameter   LengthUnit = 1000000000000000000 // Exameter to meter
	Micron     LengthUnit = 0.000001            // Micron to meter conversion
	Picometer  LengthUnit = 1e-12               // Picometer to meter conversion
	Femtometer LengthUnit = 9.999999999e-16     // Femtometer to meter conversion
	Attometer  LengthUnit = 1e-18               // Attometer to meter conversion

	// Imperial system
	Thou         LengthUnit = 2.54e-5  // Thou (mil) to meter conversion
	Inch         LengthUnit = 0.0254   // Inch to meter conversion
	Foot         LengthUnit = 0.3048   // Foot to meter conversion
	Yard         LengthUnit = 0.9144   // Yard to meter conversion
	Chain        LengthUnit = 20.1168  // Chain to meter conversion (66 feet)
	Rope         LengthUnit = 6.096    // Rope to meter conversion (20 feet)
	Furlong      LengthUnit = 201.168  // Furlong to meter conversion (1/8th of a mile)
	Mile         LengthUnit = 1609.344 // Mile to meter conversion
	League       LengthUnit = 4828.032 // League to meter conversion (3 miles)
	NauticalMile LengthUnit = 1852     // Nautical Mile to meter conversion (international)

	// Astronomical units
	AstronomicalUnit LengthUnit = 149597870691         // Astronomical Unit (AU) to meter conversion
	LightYear        LengthUnit = 9460730472580044     // Light-year to meter conversion
	Parsec           LengthUnit = 30856775812799588    // Parsec to meter conversion
	Megaparsec       LengthUnit = 3.085677581e+22      // Megaparsec to meter
	Kiloparsec       LengthUnit = 30856775812799586000 // Kiloparsec to meter
	PlanckLength     LengthUnit = 1.616049999e-35      // Planck length to meter conversion

	// Other units
	Angstrom   LengthUnit = 9.999999999e-11 // Angstrom to meter conversion
	Fathom     LengthUnit = 1.8288          // Fathom to meter conversion (6 feet)
	Rod        LengthUnit = 5.0292          // Rod to meter conversion (16.5 feet)
	SurveyFoot LengthUnit = 0.3048006096    // Survey foot to meter conversion (used in land surveys)
	Pica       LengthUnit = 0.0042333333    // Pica to meter conversion (commonly used in typography)
	Point      LengthUnit = 0.0003527778    // Point to meter conversion (1/12th of a pica)
	Hand       LengthUnit = 0.1016          // Hand to meter conversion (used in measuring the height of horses)
)

// Length converts a length value from one unit to another.
//
// **Warning:** Due to the nature of floating point arithmetic, the result may not be exact for some conversions.
// Do not use this package for rocket science.
func Length(value float64, from, to LengthUnit) float64 {
	return value * float64(from) / float64(to)
}
