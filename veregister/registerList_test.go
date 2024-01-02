package veregister

import "sort"

type testLine struct {
	str  string
	sort int
}

func (rl *RegisterList) testStrings() []string {
	lines := make([]testLine, 0, rl.Len())

	for _, r := range rl.NumberRegisters {
		lines = append(lines, testLine{r.testString(), r.sort})
	}
	for _, r := range rl.TextRegisters {
		lines = append(lines, testLine{r.testString(), r.sort})
	}
	for _, r := range rl.EnumRegisters {
		lines = append(lines, testLine{r.testString(), r.sort})
	}
	for _, r := range rl.FieldListRegisters {
		lines = append(lines, testLine{r.testString(), r.sort})
	}

	sort.Slice(lines, func(i, j int) bool { return lines[i].sort < lines[j].sort })

	res := make([]string, 0, rl.Len())
	for _, l := range lines {
		res = append(res, l.str)
	}
	return res
}
