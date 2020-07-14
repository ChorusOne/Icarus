package snapshot

import (
	"github.com/ChorusOne/Icarus/contract"
	"math/big"
)

type Address = string

type Snapshot map[Address]SnapshotRow

type SnapshotRow struct {
	GoldTokenBalance             *big.Int
	LockedGoldBalance            *big.Int
	NonVotingLockedGoldBalance   *big.Int
	PendingWithdrawalGoldBalance *big.Int
	CeloUSDValue                 *big.Int
	Reward                       *big.Int
	Commission                   *big.Int
	//AddColumnWork
}

// This type acts as a bridge between Postgres row and Golang SnapshotRow
type SnapshotPostgresRow struct {
	Address                      string
	GoldTokenBalance             string
	LockedGoldBalance            string
	NonVotingLockedGoldBalance   string
	PendingWithdrawalGoldBalance string
	CeloUSDValue                 string
	Reward                       string
	Commission                   string
	//AddColumnWork
}

func AddRowToSnapshot(snapshot *Snapshot, spr *SnapshotPostgresRow) {
	sr := NewSnapshotRow()
	//AddColumnWork
	sr.GoldTokenBalance, _ = new(big.Int).SetString(spr.GoldTokenBalance, 10)
	sr.LockedGoldBalance, _ = new(big.Int).SetString(spr.LockedGoldBalance, 10)
	sr.PendingWithdrawalGoldBalance, _ = new(big.Int).SetString(spr.PendingWithdrawalGoldBalance, 10)
	sr.NonVotingLockedGoldBalance, _ = new(big.Int).SetString(spr.NonVotingLockedGoldBalance, 10)
	sr.CeloUSDValue, _ = new(big.Int).SetString(spr.CeloUSDValue, 10)
	sr.Reward, _ = new(big.Int).SetString(spr.Reward, 10)
	sr.Commission, _ = new(big.Int).SetString(spr.Commission, 10)
	address := spr.Address
	(*snapshot)[address] = sr
}

//Use this to set trivial values
func NewSnapshotRow() SnapshotRow {
	var sr SnapshotRow
	sr.GoldTokenBalance = big.NewInt(0)
	sr.LockedGoldBalance = big.NewInt(0)
	sr.NonVotingLockedGoldBalance = big.NewInt(0)
	sr.PendingWithdrawalGoldBalance = big.NewInt(0)
	sr.CeloUSDValue = big.NewInt(0)
	sr.Reward = big.NewInt(0)
	sr.Commission = big.NewInt(0)
	//AddColumnWork
	return sr

}

// ------ System Types ----

type SystemSnapshot struct {
	GoldTokenSupply            *big.Int
	TotalLockedGoldBalance     *big.Int
	NonVotingLockedGoldBalance *big.Int
	TotalCeloUSDValue          *big.Int
	//AddSystemColumnWork

}

type AccountVotesDistribution = contract.AccountVotesDistribution
type VotesSnapshot = map[Address]*AccountVotesDistribution

type AccountRewards = map[Address]*big.Int

type RewardsSnapshot = []*RewardsSnapshotRow

type RewardsSnapshotRow struct {
	EpochBlockNumber   *big.Int
	Address            string
	Group              string
	AddressActiveUnits *big.Int
	GroupActiveUnits   *big.Int
	AddressReward      *big.Int
	GroupReward        *big.Int
}

type AccountCommissions = map[Address]*big.Int

type CommissionsSnapshot = []*CommissionsSnapshotRow

type CommissionsSnapshotRow struct {
	EpochBlockNumber *big.Int
	Validator        string
	ValidatorPayment *big.Int
	Group            string
	GroupPayment     *big.Int
}

// Types and Structs useful for API

// SnapshotInfo is a tuple representing a snapshot's date and block number
type SnapshotDate struct {
	Date        string
	BlockNumber string
}

type BlockNumber = string

//AccountSnapshot is helper type for fetching historical snapshot data of an account
type AccountSnapshot map[BlockNumber]*AccountSnapshotRow

//AccountSnapshotRow is helper type for fetching historical snapshot data of an account
type AccountSnapshotRow struct {
	//AddColumnWork
	GoldTokenBalance             string
	LockedGoldBalance            string
	NonVotingLockedGoldBalance   string
	PendingWithdrawalGoldBalance string
	CeloUSDValue                 string
	Reward                       string
	Commission                   string
}

type AccountVotesSnapshot map[BlockNumber][]*DelegationRow

type DelegationRow struct {
	Group        string
	TotalVotes   string
	ActiveVotes  string
	PendingVotes string
}

type SystemSnapshotRow struct {
	BlockNumber                string
	SnapshotDate               string
	GoldTokenSupply            string
	TotalLockedGoldBalance     string
	NonVotingLockedGoldBalance string
	TotalCeloUSDValue          string
	//AddSystemColumnWork

}
