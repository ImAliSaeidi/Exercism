package jedlik

import "fmt"

func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
	printResult := c.battery == 100
	result := false
	remainDistance := trackDistance

	for remainDistance > 0 && c.battery >= c.batteryDrain {
		c.Drive()
		remainDistance -= c.speed

		if printResult {
			fmt.Printf("Card: %+v\n Remain: %d\n", c, remainDistance)
		}
	}

	if remainDistance == 0 {
		result = true
	}

	return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
