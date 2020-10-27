package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/kava-labs/kava/x/harvest/types"
)

// GetParams returns the params from the store
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var p types.Params
	k.paramSubspace.GetParamSet(ctx, &p)
	return p
}

// SetParams sets params on the store
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSubspace.SetParamSet(ctx, &params)
}

// GetLPSchedule gets the LP's schedule
func (k Keeper) GetLPSchedule(ctx sdk.Context, denom string) (types.DistributionSchedule, bool) {
	params := k.GetParams(ctx)
	for _, lps := range params.LiquidityProviderSchedules {
		if lps.DepositDenom == denom {
			return lps, true
		}
	}
	return types.DistributionSchedule{}, false
}

// GetDelegatorSchedule gets the Delgator's schedule
func (k Keeper) GetDelegatorSchedule(ctx sdk.Context, denom string) (types.DelegatorDistributionSchedule, bool) {
	params := k.GetParams(ctx)
	for _, dds := range params.DelegatorDistributionSchedules {
		if dds.DistributionSchedule.DepositDenom == denom {
			return dds, true
		}
	}
	return types.DelegatorDistributionSchedule{}, false
}

// GetBorrowLimit returns the corresponding Borrow Limit param for a specific denom
func (k Keeper) GetBorrowLimit(ctx sdk.Context, denom string) (types.BorrowLimit, bool) {
	params := k.GetParams(ctx)
	for _, bl := range params.BorrowLimits {
		if bl.Denom == denom {
			return bl, true
		}
	}
	return types.BorrowLimit{}, false
}
