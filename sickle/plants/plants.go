package plants

import (
	"fmt"
)

func GeneratePlantCapacityReport(plantCapacities ...float64) {
	for idx, c := range plantCapacities {
		fmt.Printf("Plant %d has capacity %.1f\n", idx, c)
	}

}

func GeneratePowerGridReport(activePlants []int, plantCapacities []float64, gridLoad float64) {
	capacity := 0.
	for _, plantID := range activePlants {
		capacity += plantCapacities[plantID]
	}
	fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "Load: ", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization: ", gridLoad/capacity*100)

}
