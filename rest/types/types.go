package types

import (
	"net/http"
)

// Useful Aliases
// -----------------------------------------------------------------------------

// Address is an alias of string
type Address = string

// Height of the chain is a simple integer.
type Height = string

// Amount is any balance or vote - an abstraction over Celo quantification
type Amount = string

// Date is an alias for string (for now)
type Date = string

// Concrete API Types
// ----------------------------------------------------------------------------

// Account is a type that contains all the relevant information that Anthem seeks for an account
type Account struct {
	Address                    Address       `json:"address"`
	Height                     Height        `json:"height"`
	GoldTokenBalance           Amount        `json:"availableGoldBalance"`
	TotalLockedGoldBalance     Amount        `json:"totalLockedGoldBalance"`
	NonVotingLockedGoldBalance Amount        `json:"nonVotingLockedGoldBalance"`
	VotingLockedGoldBalance    Amount        `json:"votingLockedGoldBalance"`
	PendingWithdrawalBalance   Amount        `json:"pendingWithdrawalBalance"`
	CeloUSDValue               Amount        `json:"celoUSDValue"`
	Delegations                []*Delegation `json:"delegations"`
}

// Delegation represents what Celo calls votes
// Locked gold from an account bound to a group in the form of votes
type Delegation struct {
	Group        Address `json:"group"`
	TotalVotes   Amount  `json:"totalVotes"`
	ActiveVotes  Amount  `json:"activeVotes" `
	PendingVotes Amount  `json:"pendingVotes"`
}

// NewAccount is a constructor for Account
func NewAccount(address string) *Account {
	return &Account{
		Address:                    address,
		Height:                     "0",
		GoldTokenBalance:           "0",
		TotalLockedGoldBalance:     "0",
		NonVotingLockedGoldBalance: "0",
		VotingLockedGoldBalance:    "0",
		PendingWithdrawalBalance:   "0",
		CeloUSDValue:               "0",
		Delegations:                []*Delegation{},
	}

}

//DatedAccount represents an account and its details extracted from a daily snapshot
type DatedAccount struct {
	SnapshotDate   Date    `json:"snapshotDate"`
	Address        Address `json:"address"`
	Height         Height  `json:"height"`
	SnapshotReward Amount  `json:"snapshotReward"`

	GoldTokenBalance           Amount `json:"availableGoldBalance"`
	TotalLockedGoldBalance     Amount `json:"totalLockedGoldBalance"`
	NonVotingLockedGoldBalance Amount `json:"nonVotingLockedGoldBalance"`
	VotingLockedGoldBalance    Amount `json:"votingLockedGoldBalance"`
	PendingWithdrawalBalance   Amount `json:"pendingWithdrawalBalance"`
	CeloUSDValue               Amount `json:"celoUSDValue"`

	Delegations []*Delegation `json:"delegations"`
}

//NewDatedAccount is a constructor for DatedAccount
func NewDatedAccount(date string, address string) *DatedAccount {

	return &DatedAccount{
		SnapshotDate:   date,
		Address:        address,
		Height:         "0",
		SnapshotReward: "0",

		GoldTokenBalance:           "0",
		TotalLockedGoldBalance:     "0",
		NonVotingLockedGoldBalance: "0",
		VotingLockedGoldBalance:    "0",
		PendingWithdrawalBalance:   "0",
		CeloUSDValue:               "0",

		Delegations: []*Delegation{},
	}

}

//System stores the supply details of the blockchain at a particular block height
type System struct {
	Height                     Height `json:"height"`
	GoldTokenSupply            Amount `json:"goldTokenSupply"`
	TotalLockedGoldBalance     Amount `json:"totalLockedGoldBalance"`
	NonVotingLockedGoldBalance Amount `json:"nonVotingLockedGoldBalance"`
	TotalCeloUSDValue          Amount `json:"totalCeloUSDValue"`
}

//NewSystem is a constructor for System
func NewSystem(blockNumber string) *System {

	return &System{
		Height:                     blockNumber,
		GoldTokenSupply:            "0",
		TotalLockedGoldBalance:     "0",
		NonVotingLockedGoldBalance: "0",
		TotalCeloUSDValue:          "0"}

}

// Handler is a small type that helps to make returning closures in the functions
// below less verbose.
type Handler = func(w http.ResponseWriter, r *http.Request)
