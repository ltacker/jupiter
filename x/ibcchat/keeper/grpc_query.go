package keeper

import (
	"github.com/ltacker/jupiter/x/ibcchat/types"
)

var _ types.QueryServer = Keeper{}
