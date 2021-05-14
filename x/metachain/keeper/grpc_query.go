package keeper

import (
	"github.com/interchainberlin/metachain/x/metachain/types"
)

var _ types.QueryServer = Keeper{}
