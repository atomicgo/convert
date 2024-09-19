package convert

import "fmt"

// PercentageToProportion converts a percentage to a proportion with a given precision.
func PercentageToProportion(percentage float64, precision int) string {
	proportion := percentage / 100.0
	format := fmt.Sprintf("%%.%dfx", precision)

	return fmt.Sprintf(format, proportion)
}

// ProportionToPercentage converts a proportion to a percentage.
func ProportionToPercentage(proportion string) (float64, error) {
	var value float64

	_, err := fmt.Sscanf(proportion, "%fx", &value)
	if err != nil {
		return 0, fmt.Errorf("could not parse proportion: %w", err)
	}

	percentage := value * 100

	return percentage, nil
}
