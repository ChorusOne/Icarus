package atlas

import (
	"context"

	"log"
	"math/big"

	. "github.com/ChorusOne/Icarus/common"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//GetModifiedAccountsLockedGold returns an atlas of all of the modified accounts for the LockedGold contract between a fromBlock and a toBlock
func GetModifiedAccountsLockedGold(fromBlock *big.Int, toBlock *big.Int) map[common.Address]bool {

	// Storing the Gold Locked Event Signature in logGoldLockedSigHash
	var logGoldLockedSig = []byte("GoldLocked(address,uint256)")
	var logGoldLockedSigHash = crypto.Keccak256Hash(logGoldLockedSig)

	// Storing the Gold Unlocked Event Signature in logGoldUnlockedSigHash
	var logGoldUnlockedSig = []byte("GoldUnlocked(address,uint256,uint256)")
	var logGoldUnlockedSigHash = crypto.Keccak256Hash(logGoldUnlockedSig)

	// Storing the Gold Relocked Event Signature in logGoldRelockedSigHash
	var logGoldRelockedSig = []byte("GoldRelocked(address,uint256)")
	var logGoldRelockedSigHash = crypto.Keccak256Hash(logGoldRelockedSig)

	// Storing the Gold Withdrawn Event Signature in logGoldWithdrawnSigHash
	var logGoldWithdrawnSig = []byte("GoldWithdrawn(address,uint256)")
	var logGoldWithdrawnSigHash = crypto.Keccak256Hash(logGoldWithdrawnSig)

	contractAddress := common.HexToAddress(WrapperContractDeploymentAddress[NetActive][LockedGold])

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := atlasEthClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	modifiedAccounts := make(map[common.Address]bool)

	for _, vLog := range logs {

		switch vLog.Topics[0].Hex() {

		//catching all the gold locking events
		case logGoldLockedSigHash.Hex():
			modifiedAccounts[common.HexToAddress(vLog.Topics[1].Hex())] = true

		//catching all the gold unlocking events
		case logGoldUnlockedSigHash.Hex():
			modifiedAccounts[common.HexToAddress(vLog.Topics[1].Hex())] = true

		//catching all the gold relocking events
		case logGoldRelockedSigHash.Hex():
			modifiedAccounts[common.HexToAddress(vLog.Topics[1].Hex())] = true

		//catching all the gold withdrawal events
		case logGoldWithdrawnSigHash.Hex():
			modifiedAccounts[common.HexToAddress(vLog.Topics[1].Hex())] = true

		}

	}
	return modifiedAccounts
}

//GetModifiedAccountsCeloDollar returns an atlas of all of the modified accounts for the Celo Dollar Proxy contract between a fromBlock and a toBlock
func GetModifiedAccountsCeloDollar(fromBlock *big.Int, toBlock *big.Int) map[common.Address]bool {

	// Storing the Transfer Event Signature in logTransferSigHash
	var logTransferSig = []byte("Transfer(address,address,uint256)")
	var logTransferSigHash = crypto.Keccak256Hash(logTransferSig)

	contractAddress := common.HexToAddress(WrapperContractDeploymentAddress[NetActive][StableToken])

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := atlasEthClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	modifiedAccounts := make(map[common.Address]bool)

	for _, vLog := range logs {

		switch vLog.Topics[0].Hex() {

		//catching all the transfer events
		case logTransferSigHash.Hex():
			modifiedAccounts[common.HexToAddress(vLog.Topics[1].Hex())] = true
			modifiedAccounts[common.HexToAddress(vLog.Topics[2].Hex())] = true
		}

	}
	return modifiedAccounts
}

//GetModifiedAccountsCeloGold returns an atlas of all of the modified accounts for the Celo Gold Proxy contract between a fromBlock and a toBlock
func GetModifiedAccountsCeloGold(fromBlock *big.Int, toBlock *big.Int) map[common.Address]bool {

	// Storing the Transfer Event Signature in logTransferSigHash
	var logTransferSig = []byte("Transfer(address,address,uint256)")
	var logTransferSigHash = crypto.Keccak256Hash(logTransferSig)

	contractAddress := common.HexToAddress(WrapperContractDeploymentAddress[NetActive][GoldToken])

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := atlasEthClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	modifiedAccounts := make(map[common.Address]bool)

	for _, vLog := range logs {

		switch vLog.Topics[0].Hex() {

		//catching all the transfer events
		case logTransferSigHash.Hex():
			modifiedAccounts[common.HexToAddress(vLog.Topics[1].Hex())] = true
			modifiedAccounts[common.HexToAddress(vLog.Topics[2].Hex())] = true
		}

	}
	return modifiedAccounts
}

//GetModifiedAccountsElection returns an atlas of all of the modified accounts for the Election contract between a fromBlock and a toBlock
func GetModifiedAccountsElection(fromBlock *big.Int, toBlock *big.Int) map[common.Address]bool {

	// Storing the Validator Group Vote Cast Event Signature in logValidatorGroupVoteCastSigHash
	var logValidatorGroupVoteCastSig = []byte("ValidatorGroupVoteCast(address,address,uint256)")
	var logValidatorGroupVoteCastSigHash = crypto.Keccak256Hash(logValidatorGroupVoteCastSig)

	// Storing the Validator Group Vote Activated Event Signature in logValidatorGroupVoteActivatedSigHash
	var logValidatorGroupVoteActivatedSig = []byte("ValidatorGroupVoteActivated(address,address,uint256,uint256)")
	var logValidatorGroupVoteActivatedSigHash = crypto.Keccak256Hash(logValidatorGroupVoteActivatedSig)

	// Storing the Validator Group Pending Vote Revoked Event Signature in logValidatorGroupPendingVoteRevokedSigHash
	var logValidatorGroupPendingVoteRevokedSig = []byte("ValidatorGroupPendingVoteRevoked(address,address,uint256)")
	var logValidatorGroupPendingVoteRevokedSigHash = crypto.Keccak256Hash(logValidatorGroupPendingVoteRevokedSig)

	// Storing the Validator Group Active Vote Revoked Event Signature in logValidatorGroupActiveVoteRevokedSigHash
	var logValidatorGroupActiveVoteRevokedSig = []byte("ValidatorGroupActiveVoteRevoked(address,address,uint256,uint256)")
	var logValidatorGroupActiveVoteRevokedSigHash = crypto.Keccak256Hash(logValidatorGroupActiveVoteRevokedSig)

	contractAddress := common.HexToAddress(WrapperContractDeploymentAddress[NetActive][Election])

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := atlasEthClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	modifiedAccounts := make(map[common.Address]bool)

	//I made the decision to catch four types of events, even though validator group vote cast events would have sufficed. This is over-engineering,
	//and can be reduced in scope later on.
	for _, vLog := range logs {

		switch vLog.Topics[0].Hex() {

		//catching all the validator group vote cast events
		case logValidatorGroupVoteCastSigHash.Hex():
			modifiedAccounts[common.HexToAddress(vLog.Topics[1].Hex())] = true
			modifiedAccounts[common.HexToAddress(vLog.Topics[2].Hex())] = true

		//catching all the validator group vote activated unlocking events
		case logValidatorGroupVoteActivatedSigHash.Hex():
			modifiedAccounts[common.HexToAddress(vLog.Topics[1].Hex())] = true
			modifiedAccounts[common.HexToAddress(vLog.Topics[2].Hex())] = true

		//catching all the validator group pending vote revoked events
		case logValidatorGroupPendingVoteRevokedSigHash.Hex():
			modifiedAccounts[common.HexToAddress(vLog.Topics[1].Hex())] = true
			modifiedAccounts[common.HexToAddress(vLog.Topics[2].Hex())] = true

		//catching all the validator group active vote revoked events
		case logValidatorGroupActiveVoteRevokedSigHash.Hex():
			modifiedAccounts[common.HexToAddress(vLog.Topics[1].Hex())] = true
			modifiedAccounts[common.HexToAddress(vLog.Topics[2].Hex())] = true
		}

	}
	return modifiedAccounts
}
