package oracle

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Params oracle parameters
type Params struct {
	VotePeriod    sdk.Int `json:"vote_period"`    // voting period in block height; tallys and reward claim period
	VoteThreshold sdk.Dec `json:"vote_threshold"` // minimum stake power threshold to clear vote
	DropThreshold sdk.Int `json:"drop_threshold"` // tolerated drops before blacklist
}

// NewParams creates a new param instance
func NewParams(votePeriod sdk.Int, voteThreshold sdk.Dec, dropThreshold sdk.Int) Params {
	return Params{
		VotePeriod:    votePeriod,
		VoteThreshold: voteThreshold,
		DropThreshold: dropThreshold,
	}
}

// DefaultParams creates default oracle module parameters
func DefaultParams() Params {
	return NewParams(
		sdk.NewInt(900),           // 900 blocks; 15 minutes
		sdk.NewDecWithPrec(67, 2), // 67%
		sdk.NewInt(10),
	)
}

func validateParams(params Params) error {
	if params.VotePeriod.LT(sdk.ZeroInt()) {
		return fmt.Errorf("oracle parameter VotePeriod must be > 0, is %s", params.VotePeriod.String())
	}
	if params.VoteThreshold.LT(sdk.NewDecWithPrec(33, 2)) {
		return fmt.Errorf("oracle parameter VoteThreshold must be greater than 33 percent")
	}

	if params.DropThreshold.LTE(sdk.NewInt(3)) {
		return fmt.Errorf("oracle parameter DropThreshold must be greater than 3")
	}
	return nil
}

func (params Params) String() string {
	return fmt.Sprintf(`Oracle Params:
  VotePeriod:     %s
  VoteThreshold:  %s
  DropThresdhold: %s
  `, params.VotePeriod, params.VoteThreshold, params.DropThreshold)
}
