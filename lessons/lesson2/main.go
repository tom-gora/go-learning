package lesson2

import "fmt"

// Run ...
func Run(
	averageWill float64,
	taskDescription string,
) string {
	return fmt.Sprintf("%f3, %s", averageWill, taskDescription)
}

// TryAndSeeIfFlexibleTypeCastingIsAThingInGo ...
func TryAndSeeIfFlexibleTypeCastingIsAThingInGo(testString string) []int {
	var result []int
	for _, c := range testString {
		result = append(result, int(c))
	}
	return result
}
