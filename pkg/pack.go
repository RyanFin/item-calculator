package pkg

type Pack struct {
	packSize []int
}

func (p *Pack) addPackSize(newPackSize int) {
	p.packSize = append(p.packSize, newPackSize)
}

// Return the number of packs as a map
func (p *Pack) calculateNumberOfPacks(itemNumber int) map[int]int {
	var ret = make(map[int]int)

	for _, e := range p.packSize {
		ret[e] = 0
	}

	return ret
}
