package keeper

import (
	"github.com/merlin-network/fury/x/mint/types"
)

var _ types.QueryServer = Keeper{}
