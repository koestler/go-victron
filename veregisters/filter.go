package veregisters

type nameable interface {
	Name() string
}

func filterRegistersByName[N nameable](inp []N, exclude ...string) (oup []N) {
	oup = make([]N, 0)
	for _, i := range inp {
		if nameExcluded(i, exclude...) {
			continue
		}
		oup = append(oup, i)
	}
	return
}

func nameExcluded[N nameable](n N, exclude ...string) bool {
	for _, e := range exclude {
		if e == n.Name() {
			return true
		}
	}
	return false
}
