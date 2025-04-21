package treje

import (
	"github.com/rojack96/treje/mapset"
	mtype "github.com/rojack96/treje/mapset/types"
	"github.com/rojack96/treje/set"
	stypes "github.com/rojack96/treje/set/types"
)

func NewSet() stypes.Set {
	return set.New()
}

func NewMapSet() mtype.MapSet {
	return mapset.New()
}
