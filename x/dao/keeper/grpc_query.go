package keeper

import (
	"github.com/interchainberlin/metachain/x/dao/types"
)

var _ types.QueryServer = Keeper{}
