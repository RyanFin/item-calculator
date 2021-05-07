package pkg

import (
	"fmt"
	"sort"
)

type Pack struct {
	//  define an array of pack sizes
	packSize []int
}

func (p *Pack) addPackSize(newPackSize int) {
	p.packSize = append(p.packSize, newPackSize)
}

// @desc Return the number of packs as a map
// @params number of items ordered
// @return map displaying the key -> pack size and the amount of packs required to fulfil the order
// func (p *Pack) calculateNumberOfPacks(itemsOrdered int) map[int]int {
// 	var ret = make(map[int]int)

// 	for _, e := range p.packSize {
// 		ret[e] = 0
// 	}

// 	if itemsOrdered < 1 {
// 		// return immediately if no items
// 		return ret
// 	}

// 	// e.g. itemNum = 5. 5 is less than 250
// 	if itemsOrdered >= 1 && itemsOrdered <= p.packSize[0] {
// 		// increment map frequency value by 1
// 		ret[p.packSize[0]] += 1
// 	}

// 	if itemsOrdered >= p.packSize[0]+1 && itemsOrdered <= p.packSize[1] {
// 		// increment map frequency value by 1
// 		ret[p.packSize[1]] += 1
// 	}

// 	return ret
// }

func (p *Pack) calculateNumberOfPacks(itemsOrdered int) map[int]int {
	var ret = make(map[int]int)

	for _, e := range p.packSize {
		ret[e] = 0
	}

	totalPacked := 0

	// Sort packsize array
	sort.Ints(p.packSize)

	for totalPacked < itemsOrdered {
		fmt.Println("totalPacked: ", totalPacked)
		for i := 0; i < len(p.packSize); i++ {
			if p.packSize[i] > itemsOrdered {
				fmt.Println("p.packSize[i-1]: ", p.packSize[i-1])
				ret[p.packSize[i-1]] += 1
				totalPacked += p.packSize[i-1]
				break
			}
		}

		// special case if itemsOrdered larger than largest pack size
		if itemsOrdered > p.packSize[len(p.packSize)-1] {
			ret[p.packSize[len(p.packSize)-1]] += 1
			totalPacked += p.packSize[len(p.packSize)-1]
		}

		fmt.Println("totalPacked: ", totalPacked)
		remainder := itemsOrdered - totalPacked
		if remainder > 0 {
			ret[p.packSize[0]] += 1
		}
		totalPacked += remainder

	}

	return ret
}
