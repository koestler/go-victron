package dataflow

//go:generate mockgen -source registerFilter.go -destination mock/registerFilter_mock.go

type RegisterFilterFunc func(Register) bool

type RegisterFilterConf interface {
	IncludeRegisters() []string
	SkipRegisters() []string
	IncludeCategories() []string
	SkipCategories() []string
	DefaultInclude() bool
}

func RegisterFilter(registerFilter RegisterFilterConf) RegisterFilterFunc {
	includeRegistersMap := sliceToMap(registerFilter.IncludeRegisters())
	skipRegistersMap := sliceToMap(registerFilter.SkipRegisters())
	includeCategoriesMap := sliceToMap(registerFilter.IncludeCategories())
	skipCategoriesMap := sliceToMap(registerFilter.SkipCategories())
	defaultInclude := registerFilter.DefaultInclude()

	return func(reg Register) bool {
		regName := reg.Name()
		if _, ok := includeRegistersMap[regName]; ok {
			return true
		}

		if _, ok := skipRegistersMap[regName]; ok {
			return false
		}

		categoryName := reg.Category()
		if _, ok := includeCategoriesMap[categoryName]; ok {
			return true
		}

		if _, ok := skipCategoriesMap[categoryName]; ok {
			return false
		}

		return defaultInclude
	}
}

var AllRegisterFilter RegisterFilterFunc = func(Register) bool {
	return true
}

func sliceToMap[T comparable](inp []T) map[T]struct{} {
	oup := make(map[T]struct{}, len(inp))
	for _, v := range inp {
		oup[v] = struct{}{}
	}
	return oup
}