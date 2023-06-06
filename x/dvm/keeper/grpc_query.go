package keeper

import (
	"github.com/merlin-network/fury/x/dvm/types"
)

var _ types.QueryServer = Keeper{}
