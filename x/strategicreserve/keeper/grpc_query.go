package keeper

import (
	"github.com/merlin-network/fury/x/strategicreserve/types"
)

var _ types.QueryServer = Keeper{}
