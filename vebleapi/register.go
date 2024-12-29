package vebleapi

import "github.com/koestler/go-victron/log"

type RegisterApi struct {
	a Adapter
	l log.Logger
}

func NewRegisterApi(a Adapter, l log.Logger) *RegisterApi {
	return &RegisterApi{a: a, l: l}
}
