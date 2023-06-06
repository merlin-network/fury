package keeper

import (
	"github.com/merlin-network/fury/x/sportevent/types"
)

var _ types.QueryServer = Keeper{}
