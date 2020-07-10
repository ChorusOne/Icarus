package snapshot

import (
	"database/sql"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ChorusOne/Icarus/atlas"
	"github.com/ChorusOne/Icarus/blockchain"
	"github.com/ChorusOne/Icarus/contract"
)

const haltOnIntegrityCheckFailure = true

func GenerateAndStoreGenesisSnapshot(db *sql.DB, blockNumber *big.Int, rpcClient *rpc.Client) {
	log.Println("Attempting to Snapshot Genesis Block")
	var err error

	var snapshot Snapshot = GenesisSnapshotAccountLevel(db, rpcClient)
	err = DumpSnapshotData(db, snapshot, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	blockHeader, err2 := blockchain.GetHeaderLite(rpcClient, blockNumber)

	if err2 != nil {
		log.Fatal(err2)
	}
	CreateSnapshotStatusRow(db, blockNumber, blockHeader.Timestamp)
	log.Println("Genesis Block Snapshot Complete")

}

func GenerateAndStoreBlockSnapshots(db *sql.DB, blockNumber *big.Int, rpcClient *rpc.Client) {
	log.Println("Attempting To Snapshot Block #", blockNumber)
	tStart := time.Now()
	previousSnapshotBlockNumber := fetchPreviousSnapshotBlockNumber(db)

	// We will take RewardsSnapshot first as we need to pass the first output to SnapshotAccountLevel
	accountRewards, rewardsSnapshot := SnapshotRewardsAccountLevel(blockNumber, previousSnapshotBlockNumber)
	var snapshot Snapshot = SnapshotAccountLevel(db, rpcClient, blockNumber, previousSnapshotBlockNumber, accountRewards)
	var systemSnapshot SystemSnapshot = SnapshotSystemLevel(rpcClient, blockNumber)
	var votesSnapshot VotesSnapshot = VotesSnapshotAccountLevel(blockNumber)

	integrity := CheckIntegrity(snapshot, systemSnapshot)

	if !integrity && haltOnIntegrityCheckFailure {
		log.Fatal("Integrity Check Failed. Halting Icarus. ")

	}

	var err error
	tEnd := time.Now()
	log.Println("SnapshotAccount Level Time : ", tEnd.Sub(tStart))

	err = DumpRewardsSnapshotData(db, rewardsSnapshot, blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	err = DumpSnapshotData(db, snapshot, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	err = DumpVotesSnapshotData(db, votesSnapshot, blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	err = DumpSystemSnapshotData(db, systemSnapshot, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	blockHeader, err2 := blockchain.GetHeaderLite(rpcClient, blockNumber)

	if err2 != nil {
		log.Fatal(err2)
	}

	CreateSnapshotStatusRow(db, blockNumber, blockHeader.Timestamp)
	log.Println("Snapshot Complete For Block #", blockNumber)

}

func GenesisSnapshotAccountLevel(db *sql.DB, rpcClient *rpc.Client) Snapshot {

	var snapshot Snapshot = make(Snapshot)
	stateDump, err := blockchain.GetBlockState(rpcClient, big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}

	StateAddressesWithData := blockchain.GetStateAddressesWithRelevantData(stateDump)

	//Set All State Dump Columns From Storage
	for address := range StateAddressesWithData {
		row := NewSnapshotRow()
		row.GoldTokenBalance = StateAddressesWithData[address].GoldTokenBalance //AddColumnWork
		snapshot[address] = row
	}

	return snapshot
}

func SnapshotAccountLevel(
	db *sql.DB,
	rpcClient *rpc.Client,
	blockNumber *big.Int,
	previousSnapshotBlockNumber *big.Int,
	accountRewards AccountRewards) Snapshot {

	var snapshot = make(Snapshot)

	previousSnapshotAddresses, err := FetchSnapshotAddresses(db, previousSnapshotBlockNumber)
	if err != nil {
		log.Fatal(err)
	}

	stateDump, err := blockchain.GetBlockState(rpcClient, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	presentStateAddresses := blockchain.GetAddresses(stateDump)

	for _, address := range previousSnapshotAddresses {

		snapshot[address] = GenerateNewSnapshotRowWithDetails(address, blockNumber)

	}

	for _, address := range presentStateAddresses {
		if _, ok := snapshot[address]; !ok {
			snapshot[address] = GenerateNewSnapshotRowWithDetails(address, blockNumber)

		}
	}

	// Add Rewards to Table
	for address := range accountRewards {
		if _, ok := snapshot[address]; !ok {
			log.Print("Warning : This is not expected! If an address has a reward, it is expected to be in the snapshot already!")
			snapshot[address] = NewSnapshotRow()
		}

		row := snapshot[address]
		row.Reward = accountRewards[address]
		snapshot[address] = row
	}

	return snapshot

}

func GenerateNewSnapshotRowWithDetails(address string, blockNumber *big.Int) SnapshotRow {

	row := NewSnapshotRow()
	var err error

	row.GoldTokenBalance, err = contract.GoldTokenBalance(address, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	row.NonVotingLockedGoldBalance, err = contract.NonVotingLockedGoldBalance(address, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	row.LockedGoldBalance, err = contract.TotalLockedGoldBalance(address, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	row.PendingWithdrawalGoldBalance, err = contract.TotalPendingWithdrawalGoldBalance(address, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	row.CeloUSDValue, err = contract.StableTokenBalanceValue(address, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	return row
}

func VotesSnapshotAccountLevel(blockNumber *big.Int) VotesSnapshot {
	var vs VotesSnapshot = make(VotesSnapshot)

	atlas := atlas.GetModifiedAccountsLockedGold(big.NewInt(0), blockNumber)

	for ad := range atlas {
		address := ad.Hex()
		avd, err := contract.GetAccountVotesDistribution(address, blockNumber)

		if err != nil {
			log.Fatal(err)
		}

		vs[address] = avd

	}
	return vs
}

func SnapshotSystemLevel(rpcClient *rpc.Client, blockNumber *big.Int) SystemSnapshot {

	var ss SystemSnapshot
	var err error

	ss.GoldTokenSupply, err = contract.SystemGoldTokenSupply(blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	ss.TotalLockedGoldBalance, err = contract.SystemTotalLockedGoldBalance(blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	ss.NonVotingLockedGoldBalance, err = contract.SystemNonVotingLockedGoldBalance(blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	ss.TotalCeloUSDValue, err = contract.SystemStableTokenSupplyValue(blockNumber)

	if err != nil {
		log.Fatal(err)
	}
	//AddSystemColumnWork
	return ss
}

func SnapshotRewardsAccountLevel(blockNumber *big.Int, previousSnapshotBlockNumber *big.Int) (AccountRewards, RewardsSnapshot) {

	//Get RewardsInfo between previousSnapshotBlockNumber+1 and this block (Both included)

	var fromBlock *big.Int = big.NewInt(0)
	fromBlock.Add(previousSnapshotBlockNumber, big.NewInt(1))
	rewards := atlas.GetRewardEventsInfo(fromBlock, blockNumber)

	var accountRewards AccountRewards
	accountRewards = make(map[Address]*big.Int)

	var rewardsSnapshot RewardsSnapshot
	rewardsSnapshot = make([]*RewardsSnapshotRow, 0)

	for _, reward := range rewards {

		totalReward := reward.RewardValue

		if totalReward.Cmp(big.NewInt(0)) == 0 {
			continue
		}

		totalActiveUnits, err := contract.GetActiveVoteUnitsForGroup(reward.Group, blockNumber)
		if err != nil {
			log.Fatal(err)
		}

		if totalActiveUnits.Cmp(big.NewInt(0)) == 0 {
			continue
		}

		for ad := range reward.AddressAtlas {

			address := ad.String()
			addressActiveUnits, err2 := contract.GetActiveVoteUnitsForGroupByAccount(reward.Group, address, blockNumber)

			if err2 != nil {
				log.Fatal(err2)
			}

			if addressActiveUnits.Cmp(big.NewInt(0)) == 0 {
				continue
			}

			if _, ok := accountRewards[address]; !ok {
				accountRewards[address] = big.NewInt(0)
			}

			// 	addressReward := totalReward * addressActiveUnits / totalActiveUnits
			var addressReward *big.Int = big.NewInt(0)
			addressReward.Mul(totalReward, addressActiveUnits)
			addressReward.Div(addressReward, totalActiveUnits)

			// Add reward for this particular group/epoch to the total reward of this account
			accountRewards[address].Add(addressReward, accountRewards[address])

			rsr := &RewardsSnapshotRow{
				EpochBlockNumber:   reward.BlockNumber,
				Address:            address,
				Group:              reward.Group,
				AddressActiveUnits: addressActiveUnits,
				GroupActiveUnits:   totalActiveUnits,
				AddressReward:      addressReward,
				GroupReward:        totalReward,
			}
			rewardsSnapshot = append(rewardsSnapshot, rsr)
		}
	}
	return accountRewards, rewardsSnapshot
}
