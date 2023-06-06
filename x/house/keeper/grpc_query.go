package keeper

import (
	"github.com/merlin-network/fury/x/house/types"
)

var _ types.QueryServer = Keeper{}
