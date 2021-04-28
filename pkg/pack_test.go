package pkg

import (
	"fmt"
	"testing"
)

func TestPackUpdate(t *testing.T) {
	// Create a pack object
	p := Pack{}
	// Update array with packsizes
	p.addPackSize(250)
	p.addPackSize(500)
	p.addPackSize(1000)
	p.addPackSize(2000)
	p.addPackSize(5000)
	// Create a map
	myMap := p.calculateNumberOfPacks(0)
	// output
	fmt.Println("map; ", myMap)

	// access map value
	fmt.Printf("Value: %v \n", myMap[500])

	// Set new map value
	myMap[500] = 1500

	// View updated map
	fmt.Printf("Value: %v \n", myMap[500])

}

func TestCalculatePack(t *testing.T) {
	p := Pack{}
	packSizes := []int{250, 500, 1000, 2000, 5000}

	for _, e := range packSizes {
		p.addPackSize(e)
	}

	myMap := p.calculateNumberOfPacks(0)

	fmt.Println("map: ", myMap)

}
