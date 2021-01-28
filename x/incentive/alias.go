package incentive

// DO NOT EDIT - generated by aliasgen tool (github.com/rhuairahrighairidh/aliasgen)

import (
	"github.com/kava-labs/kava/x/incentive/keeper"
	"github.com/kava-labs/kava/x/incentive/types"
)

const (
	AttributeKeyClaimAmount        = types.AttributeKeyClaimAmount
	AttributeKeyClaimPeriod        = types.AttributeKeyClaimPeriod
	AttributeKeyClaimType          = types.AttributeKeyClaimType
	AttributeKeyClaimedBy          = types.AttributeKeyClaimedBy
	AttributeKeyRewardPeriod       = types.AttributeKeyRewardPeriod
	AttributeValueCategory         = types.AttributeValueCategory
	BondDenom                      = types.BondDenom
	DefaultParamspace              = types.DefaultParamspace
	EventTypeClaim                 = types.EventTypeClaim
	EventTypeClaimPeriod           = types.EventTypeClaimPeriod
	EventTypeClaimPeriodExpiry     = types.EventTypeClaimPeriodExpiry
	EventTypeRewardPeriod          = types.EventTypeRewardPeriod
	HardLiquidityProviderClaimType = types.HardLiquidityProviderClaimType
	Large                          = types.Large
	Medium                         = types.Medium
	ModuleName                     = types.ModuleName
	QuerierRoute                   = types.QuerierRoute
	QueryGetClaimPeriods           = types.QueryGetClaimPeriods
	QueryGetHardRewards            = types.QueryGetHardRewards
	QueryGetParams                 = types.QueryGetParams
	QueryGetRewardPeriods          = types.QueryGetRewardPeriods
	QueryGetRewards                = types.QueryGetRewards
	QueryGetUSDXMintingRewards     = types.QueryGetUSDXMintingRewards
	RestClaimCollateralType        = types.RestClaimCollateralType
	RestClaimOwner                 = types.RestClaimOwner
	RestClaimType                  = types.RestClaimType
	RouterKey                      = types.RouterKey
	Small                          = types.Small
	StoreKey                       = types.StoreKey
	USDXMintingClaimType           = types.USDXMintingClaimType
)

var (
	// function aliases
	CalculateTimeElapsed                   = keeper.CalculateTimeElapsed
	NewKeeper                              = keeper.NewKeeper
	NewQuerier                             = keeper.NewQuerier
	DefaultGenesisState                    = types.DefaultGenesisState
	DefaultParams                          = types.DefaultParams
	GetTotalVestingPeriodLength            = types.GetTotalVestingPeriodLength
	NewGenesisAccumulationTime             = types.NewGenesisAccumulationTime
	NewGenesisState                        = types.NewGenesisState
	NewHardLiquidityProviderClaim          = types.NewHardLiquidityProviderClaim
	NewMsgClaimHardLiquidityProviderReward = types.NewMsgClaimHardLiquidityProviderReward
	NewMsgClaimUSDXMintingReward           = types.NewMsgClaimUSDXMintingReward
	NewMultiplier                          = types.NewMultiplier
	NewParams                              = types.NewParams
	NewPeriod                              = types.NewPeriod
	NewQueryHardRewardsParams              = types.NewQueryHardRewardsParams
	NewQueryRewardsParams                  = types.NewQueryRewardsParams
	NewQueryUSDXMintingRewardsParams       = types.NewQueryUSDXMintingRewardsParams
	NewRewardIndex                         = types.NewRewardIndex
	NewRewardPeriod                        = types.NewRewardPeriod
	NewUSDXMintingClaim                    = types.NewUSDXMintingClaim
	ParamKeyTable                          = types.ParamKeyTable
	RegisterCodec                          = types.RegisterCodec

	// variable aliases
	DefaultActive                                   = types.DefaultActive
	DefaultClaimEnd                                 = types.DefaultClaimEnd
	DefaultGenesisAccumulationTimes                 = types.DefaultGenesisAccumulationTimes
	DefaultHardClaims                               = types.DefaultHardClaims
	DefaultMultipliers                              = types.DefaultMultipliers
	DefaultRewardPeriods                            = types.DefaultRewardPeriods
	DefaultUSDXClaims                               = types.DefaultUSDXClaims
	ErrAccountNotFound                              = types.ErrAccountNotFound
	ErrClaimExpired                                 = types.ErrClaimExpired
	ErrClaimNotFound                                = types.ErrClaimNotFound
	ErrInsufficientModAccountBalance                = types.ErrInsufficientModAccountBalance
	ErrInvalidAccountType                           = types.ErrInvalidAccountType
	ErrInvalidClaimType                             = types.ErrInvalidClaimType
	ErrInvalidMultiplier                            = types.ErrInvalidMultiplier
	ErrNoClaimsFound                                = types.ErrNoClaimsFound
	ErrRewardPeriodNotFound                         = types.ErrRewardPeriodNotFound
	ErrZeroClaim                                    = types.ErrZeroClaim
	GovDenom                                        = types.GovDenom
	HardBorrowRewardFactorKeyPrefix                 = types.HardBorrowRewardFactorKeyPrefix
	HardDelegatorRewardFactorKeyPrefix              = types.HardDelegatorRewardFactorKeyPrefix
	HardLiquidityClaimKeyPrefix                     = types.HardLiquidityClaimKeyPrefix
	HardLiquidityRewardDenom                        = types.HardLiquidityRewardDenom
	HardSupplyRewardFactorKeyPrefix                 = types.HardSupplyRewardFactorKeyPrefix
	IncentiveMacc                                   = types.IncentiveMacc
	KeyClaimEnd                                     = types.KeyClaimEnd
	KeyHardBorrowRewardPeriods                      = types.KeyHardBorrowRewardPeriods
	KeyHardDelegatorRewardPeriods                   = types.KeyHardDelegatorRewardPeriods
	KeyHardSupplyRewardPeriods                      = types.KeyHardSupplyRewardPeriods
	KeyMultipliers                                  = types.KeyMultipliers
	KeyUSDXMintingRewardPeriods                     = types.KeyUSDXMintingRewardPeriods
	ModuleCdc                                       = types.ModuleCdc
	PreviousHardBorrowRewardAccrualTimeKeyPrefix    = types.PreviousHardBorrowRewardAccrualTimeKeyPrefix
	PreviousHardDelegatorRewardAccrualTimeKeyPrefix = types.PreviousHardDelegatorRewardAccrualTimeKeyPrefix
	PreviousHardSupplyRewardAccrualTimeKeyPrefix    = types.PreviousHardSupplyRewardAccrualTimeKeyPrefix
	PreviousUSDXMintingRewardAccrualTimeKeyPrefix   = types.PreviousUSDXMintingRewardAccrualTimeKeyPrefix
	PrincipalDenom                                  = types.PrincipalDenom
	USDXMintingClaimKeyPrefix                       = types.USDXMintingClaimKeyPrefix
	USDXMintingRewardDenom                          = types.USDXMintingRewardDenom
	USDXMintingRewardFactorKeyPrefix                = types.USDXMintingRewardFactorKeyPrefix
)

type (
	Hooks                               = keeper.Hooks
	Keeper                              = keeper.Keeper
	AccountKeeper                       = types.AccountKeeper
	BaseClaim                           = types.BaseClaim
	CDPHooks                            = types.CDPHooks
	CdpKeeper                           = types.CdpKeeper
	Claim                               = types.Claim
	Claims                              = types.Claims
	GenesisAccumulationTime             = types.GenesisAccumulationTime
	GenesisAccumulationTimes            = types.GenesisAccumulationTimes
	GenesisState                        = types.GenesisState
	HARDHooks                           = types.HARDHooks
	HardKeeper                          = types.HardKeeper
	HardLiquidityProviderClaim          = types.HardLiquidityProviderClaim
	HardLiquidityProviderClaims         = types.HardLiquidityProviderClaims
	MsgClaimHardLiquidityProviderReward = types.MsgClaimHardLiquidityProviderReward
	MsgClaimUSDXMintingReward           = types.MsgClaimUSDXMintingReward
	Multiplier                          = types.Multiplier
	MultiplierName                      = types.MultiplierName
	Multipliers                         = types.Multipliers
	Params                              = types.Params
	PostClaimReq                        = types.PostClaimReq
	QueryHardRewardsParams              = types.QueryHardRewardsParams
	QueryRewardsParams                  = types.QueryRewardsParams
	QueryUSDXMintingRewardsParams       = types.QueryUSDXMintingRewardsParams
	RewardIndex                         = types.RewardIndex
	RewardIndexes                       = types.RewardIndexes
	RewardPeriod                        = types.RewardPeriod
	RewardPeriods                       = types.RewardPeriods
	StakingKeeper                       = types.StakingKeeper
	SupplyKeeper                        = types.SupplyKeeper
	USDXMintingClaim                    = types.USDXMintingClaim
	USDXMintingClaims                   = types.USDXMintingClaims
)
