package veregisters

func filterRegisters[R Register](inp []R, f func(r Register) bool) (oup []R) {
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
	rl.FilterRegister(func(r Register) bool {
		for _, e := range exclude {
			if e == r.Name() {
				return false
			}
		}
		return true
	})
}
