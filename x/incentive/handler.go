package incentive

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/kava-labs/kava/x/incentive/keeper"
	"github.com/kava-labs/kava/x/incentive/types"
)

// NewHandler creates an sdk.Handler for incentive module messages
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgClaimReward:
			return handleMsgClaimReward(ctx, k, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)

		}
	}
}

func handleMsgClaimReward(ctx sdk.Context, k keeper.Keeper, msg types.MsgClaimReward) (*sdk.Result, error) {

	err := k.ClaimReward(ctx, msg.Sender, msg.CollateralType, types.MultiplierName(msg.MultiplierName))
	if err != nil {
		return nil, err
	}

	return &sdk.Result{
		Events: ctx.EventManager().Events(),
	}, nil
}
