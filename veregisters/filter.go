package veregisters

// FilterFunc is a function that returns true if the register should be kept.
type FilterFunc func(r Register) bool

// Filter removes all registers from the list for which the filter function return false.
// This is useful since the user does not need to care about the different register types.
func (rl *RegisterList) Filter(f FilterFunc) {
	rl.NumberRegisters = filterRegisters(rl.NumberRegisters, f)
	rl.TextRegisters = filterRegisters(rl.TextRegisters, f)
	rl.EnumRegisters = filterRegisters(rl.EnumRegisters, f)
}

func filterRegisters[R Register](inp []R, f FilterFunc) (oup []R) {
	oup = make([]R, 0)
	for _, r := range inp {
		if f(r) {
			oup = append(oup, r)
		}
	}
	return
}

// FilterByName removes all registers with the specified name from the list.
func (rl *RegisterList) FilterByName(exclude ...string) {
	rl.Filter(byNameFilter(exclude...))
}

func byNameFilter(exclude ...string) FilterFunc {
	return func(r Register) bool {
		for _, e := range exclude {
			if e == r.GetName() {
				return false
			}
		}
		return true
	}
}
