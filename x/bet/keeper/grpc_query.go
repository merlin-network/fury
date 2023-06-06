package keeper

import (
	"github.com/merlin-network/fury/x/bet/types"
)

var _ types.QueryServer = Keeper{}
