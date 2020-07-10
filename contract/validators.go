package contract

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	binding "github.com/ChorusOne/Icarus/binding/mainnet"
	. "github.com/ChorusOne/Icarus/common"
)

var validatorsContractInstance *binding.Validators

func InitValidatorsContractInstance() {

	instance, err := binding.NewValidators(common.HexToAddress(WrapperContractDeploymentAddress[NetActive][Validators]), clientInstance)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Validators Contract is loaded")
	validatorsContractInstance = instance

}

//GetRegisteredValidatorGroups returns the registered validator groups at the given block height
func GetRegisteredValidatorGroups(atBlockNumber *big.Int) ([]common.Address, error) {
	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Validators])

	if !check {
		blankSlice := []common.Address{}
		return blankSlice, nil

	}

	vgs, err := validatorsContractInstance.GetRegisteredValidatorGroups(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err != nil {
		log.Println(err)
	}
	return vgs, err

}

func GetValidatorScore(valAddress common.Address, atBlockNumber *big.Int) (*big.Int, error) {

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Validators])
	score := big.NewInt(0)
	if !check {
		return score, nil
	}

	vDetails, err := validatorsContractInstance.GetValidator(&bind.CallOpts{BlockNumber: atBlockNumber}, valAddress)

	if err != nil {
		log.Println(err)
		return score, err
	}

	score = vDetails.Score
	return score, nil
}

func GetValidatorGroup(groupAddressHex string, atBlockNumber *big.Int) ([]common.Address, *big.Int, *big.Int, *big.Int, []*big.Int, *big.Int, *big.Int, error) {

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Validators])

	if !check {
		blankSlice := []common.Address{}
		zero := big.NewInt(0)
		zeroSlice := []*big.Int{}
		return blankSlice, zero, zero, zero, zeroSlice, zero, zero, nil

	}

	group := common.HexToAddress(groupAddressHex)
	return validatorsContractInstance.GetValidatorGroup(&bind.CallOpts{BlockNumber: atBlockNumber}, group)

}

//GetRegisteredValidators returns the registered validators at the given block height
func GetRegisteredValidators(atBlockNumber *big.Int) ([]common.Address, error) {
	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Validators])

	if !check {
		blankSlice := []common.Address{}
		return blankSlice, nil

	}

	validators, err := validatorsContractInstance.GetRegisteredValidators(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err != nil {
		log.Println(err)
	}
	return validators, err

}

func GetDetailsOfValidatorGroups(atBlockNumber *big.Int, groups []common.Address) ([]*ValidatorGroup, error) {

	result := []*ValidatorGroup{}

	check := heightSanityCheck(atBlockNumber, DeploymentBlockNumber[NetActive][Validators])

	if !check {
		blankSlice := []*ValidatorGroup{}
		return blankSlice, nil

	}

	totalVotes, err := GetTotalVotes(atBlockNumber)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	totalVotesMap, err2 := getTotalVotesForEligibleValidatorGroups(atBlockNumber)

	if err2 != nil {
		log.Println(err2)
		return nil, err2
	}

	for _, address := range groups {

		name, err1 := GetName(address.Hex(), atBlockNumber)

		if err1 != nil {
			log.Println(err1)
			return nil, err1
		}
		metadata, err2 := GetMetadataURL(address.Hex(), atBlockNumber)
		if err2 != nil {
			log.Println(err2)
			return nil, err2
		}

		tvg, keyPresent := totalVotesMap[address]
		if !keyPresent {
			tvg = big.NewInt(0)
		}

		tvgFloat := new(big.Float).SetInt(tvg)
		totalVotesFloat := new(big.Float).SetInt(totalVotes)

		votingPowerFraction := new(big.Float).Quo(tvgFloat, totalVotesFloat)

		capacityAvailable, err3 := GetNumVotesReceivable(address.Hex(), atBlockNumber)
		if err3 != nil {
			log.Println(err3)
			return nil, err3
		}

		totalCapacity := new(big.Int).Add(capacityAvailable, tvg)

		memberKeys, commission, _, _, _, multiplier, _, err4 := GetValidatorGroup(address.Hex(), atBlockNumber)

		if err4 != nil {
			log.Println(err4)
			return nil, err4
		}

		valDetails := []*ValidatorDetails{}
		valScoreSummation := big.NewInt(0)
		valCount := big.NewInt(int64(len(memberKeys)))
		groupScore := big.NewInt(0)

		if len(memberKeys) > 0 {
			for _, valAddress := range memberKeys {

				score, err := GetValidatorScore(valAddress, atBlockNumber)

				if err != nil {
					log.Println(err)
					return nil, err
				}
				valScoreSummation.Add(valScoreSummation, score)
				valDetails = append(valDetails, &ValidatorDetails{Address: valAddress, Score: score})
			}

			groupScore = big.NewInt(0).Div(valScoreSummation, valCount)
		}

		vg := &ValidatorGroup{
			Address:             address,
			Name:                name,
			Metadata:            metadata,
			VotingPower:         tvg,
			VotingPowerFraction: votingPowerFraction,
			CapacityAvailable:   capacityAvailable,
			TotalCapacity:       totalCapacity,
			Multiplier:          multiplier,
			GroupShare:          commission,
			ValidatorAddresses:  memberKeys,
			BlockNumber:         atBlockNumber,
			ValidatorDetails:    valDetails,
			GroupScore:          groupScore,
		}

		result = append(result, vg)
	}

	return result, nil

}

type ValidatorGroup struct {
	Address             common.Address      `json:"group"`
	Name                string              `json:"name"`
	Metadata            string              `json:"metadataURL"`
	BlockNumber         *big.Int            `json:"blockNumber"`
	VotingPower         *big.Int            `json:"votingPower"`
	VotingPowerFraction *big.Float          `json:"votingPowerFraction"`
	CapacityAvailable   *big.Int            `json:"capacityAvailable"`
	TotalCapacity       *big.Int            `json:"totalCapacity"`
	Multiplier          *big.Int            `json:"multiplier"`
	GroupShare          *big.Int            `json:"groupShare"`
	ValidatorAddresses  []common.Address    `json:"validatorAddresses"`
	ValidatorDetails    []*ValidatorDetails `json:"validatorDetails"`
	GroupScore          *big.Int            `json:"groupScore"`
}

type ValidatorDetails struct {
	Address common.Address `json:"validatorAddress"`
	Score   *big.Int       `json:"validator_score"`
}
