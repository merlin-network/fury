package keeper

import (
	"github.com/merlin-network/fury/x/orderbook/types"
)

var _ types.QueryServer = Keeper{}
