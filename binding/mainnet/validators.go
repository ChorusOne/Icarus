// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mainnet

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ValidatorsABI is the input ABI used to generate the binding from.
const ValidatorsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeMember\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getMembershipInLastEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromCurrentSet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getValidator\",\"outputs\":[{\"name\":\"ecdsaPublicKey\",\"type\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"name\":\"affiliation\",\"type\":\"address\"},{\"name\":\"score\",\"type\":\"uint256\"},{\"name\":\"signer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorScoreParameters\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"blsKey\",\"type\":\"bytes\"},{\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"checkProofOfPossession\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"},{\"name\":\"lesser\",\"type\":\"address\"},{\"name\":\"greater\",\"type\":\"address\"}],\"name\":\"addFirstMember\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getMembershipHistory\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashingMultiplierResetPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getGroupNumMembers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumberOfBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRegisteredValidatorGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMaxGroupSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getVerifiedSealBitmapFromHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"membershipHistoryLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"ecdsaPublicKey\",\"type\":\"bytes\"}],\"name\":\"updateEcdsaPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumRegisteredValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getMembershipInLastEpochFromSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isValidatorGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxGroupSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setGroupLockedGoldRequirements\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromSet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deregisterValidatorGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"hashHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setSlashingMultiplierResetPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setCommissionUpdateDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGroupLockedGoldRequirements\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"getGroupsNumMembers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"ecdsaPublicKey\",\"type\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"updatePublicKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minQuorumSizeInCurrentSet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setValidatorLockedGoldRequirements\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uptimes\",\"type\":\"uint256[]\"}],\"name\":\"calculateGroupEpochScore\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"},{\"name\":\"groupRequirementValue\",\"type\":\"uint256\"},{\"name\":\"groupRequirementDuration\",\"type\":\"uint256\"},{\"name\":\"validatorRequirementValue\",\"type\":\"uint256\"},{\"name\":\"validatorRequirementDuration\",\"type\":\"uint256\"},{\"name\":\"validatorScoreExponent\",\"type\":\"uint256\"},{\"name\":\"validatorScoreAdjustmentSpeed\",\"type\":\"uint256\"},{\"name\":\"_membershipHistoryLength\",\"type\":\"uint256\"},{\"name\":\"_slashingMultiplierResetPeriod\",\"type\":\"uint256\"},{\"name\":\"_maxGroupSize\",\"type\":\"uint256\"},{\"name\":\"_commissionUpdateDelay\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"setNextCommissionUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberValidatorsInCurrentSet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getBlockNumberFromHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deregisterValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"getTopGroupValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uptime\",\"type\":\"uint256\"}],\"name\":\"calculateEpochScore\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"},{\"name\":\"lesserMember\",\"type\":\"address\"},{\"name\":\"greaterMember\",\"type\":\"address\"}],\"name\":\"reorderMember\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"numberValidatorsInSet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getValidatorGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"affiliate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getValidatorBlsPublicKeyFromSigner\",\"outputs\":[{\"name\":\"blsPublicKey\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resetSlashingMultiplier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCommissionUpdateDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorLockedGoldRequirements\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"updateBlsPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"uptime\",\"type\":\"uint256\"}],\"name\":\"updateValidatorScoreFromSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorLockedGoldRequirements\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"halveSlashingMultiplier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"meetsAccountLockedGoldRequirements\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"groupLockedGoldRequirements\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"addMember\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"exponent\",\"type\":\"uint256\"},{\"name\":\"adjustmentSpeed\",\"type\":\"uint256\"}],\"name\":\"setValidatorScoreParameters\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRegisteredValidatorSigners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"maxPayment\",\"type\":\"uint256\"}],\"name\":\"distributeEpochPaymentsFromSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRegisteredValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getValidatorGroupSlashingMultiplier\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLockedGoldRequirement\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commissionUpdateDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"setMaxGroupSize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validatorAccount\",\"type\":\"address\"}],\"name\":\"forceDeaffiliateIfValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"minQuorumSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateCommission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ecdsaPublicKey\",\"type\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"registerValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"groupMembershipInEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"aNumerator\",\"type\":\"uint256\"},{\"name\":\"aDenominator\",\"type\":\"uint256\"},{\"name\":\"bNumerator\",\"type\":\"uint256\"},{\"name\":\"bDenominator\",\"type\":\"uint256\"},{\"name\":\"exponent\",\"type\":\"uint256\"},{\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"fractionMulExp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"registerValidatorGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"setMembershipHistoryLength\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getParentSealBitmap\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deaffiliate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"MaxGroupSizeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"CommissionUpdateDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"exponent\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"adjustmentSpeed\",\"type\":\"uint256\"}],\"name\":\"ValidatorScoreParametersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"GroupLockedGoldRequirementsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"ValidatorLockedGoldRequirementsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"MembershipHistoryLengthSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ValidatorAffiliated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ValidatorDeaffiliated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"ecdsaPublicKey\",\"type\":\"bytes\"}],\"name\":\"ValidatorEcdsaPublicKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"blsPublicKey\",\"type\":\"bytes\"}],\"name\":\"ValidatorBlsPublicKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"score\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochScore\",\"type\":\"uint256\"}],\"name\":\"ValidatorScoreUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"ValidatorGroupRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ValidatorGroupDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorGroupMemberAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorGroupMemberRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorGroupMemberReordered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commission\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"activationBlock\",\"type\":\"uint256\"}],\"name\":\"ValidatorGroupCommissionUpdateQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"ValidatorGroupCommissionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"validatorPayment\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"groupPayment\",\"type\":\"uint256\"}],\"name\":\"ValidatorEpochPaymentDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ValidatorsBin is the compiled bytecode used for deploying new contracts.
var ValidatorsBin = "0x60806040526000620000196401000000006200006d810204565b60008054600160a060020a031916600160a060020a0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506001805562000071565b3390565b619afb80620000816000396000f3fe608060405234801561001057600080fd5b506004361061049a5760003560e060020a900480638dd31e3911610266578063cb8f98e011610155578063e50e652d116100d2578063ee09831011610096578063ee098310146116be578063eff2ea3f146116db578063f2fde38b146116f8578063facd743b1461171e578063fae8db0a14611744578063fffdfccb146117615761049a565b8063e50e652d14611518578063e7f0376614611535578063ea684f771461153d578063eb1d0b4214611651578063ec683072146116835761049a565b8063dcff4cf611610119578063dcff4cf61461149f578063df4da461146114c5578063e0e3ffe6146114cd578063e1497ff7146114d5578063e33301aa146114f25761049a565b8063cb8f98e01461141a578063d55dcbcf1461143d578063d69ef6cf14611445578063d93ab5ad14611471578063dba94fcd146114795761049a565b8063b8f93943116101e3578063c10c96ef116101a7578063c10c96ef14611398578063c22d3bba146113a0578063c54c1cd4146113c6578063c5805140146113ec578063ca6d56dc146113f45761049a565b8063b8f9394314611292578063b915f5301461129a578063bd9e9d94146112a2578063bfdb7417146112aa578063c0c6ad6f1461136c5761049a565b80639b2b592f1161022a5780639b2b592f146110ac5780639b9d5161146110c9578063a91ee0dc146111ab578063b591d3a5146111d1578063b730a299146111f75761049a565b80638dd31e391461101b5780638f32d59b1461104757806394903a971461104f578063988dcd1f1461106c5780639a7b3be7146110a45761049a565b80635779e93d1161038d578063715018a61161030a5780637b103999116102ce5780637b10399914610f2357806386d81a5a14610f2b57806387ee8a0f14610f485780638a88362614610f505780638b16b1c614610ff65780638da5cb5b146110135761049a565b8063715018a614610e1a5780637385e5da14610e2257806376c0a9ed14610e2a57806376f7425d14610e4d57806378d2545614610ebd5761049a565b80636ab951a0116103515780636ab951a014610c395780636c620d9014610c585780636fa4764714610c755780637044775414610c7d578063713ea0f314610ced5761049a565b80635779e93d14610b285780635a61d15b14610b305780635d180adb14610b5357806360fb822c14610b7657806367960e9114610b935761049a565b806339e618e81161041b5780634cd76db4116103df5780634cd76db414610a435780634e06fd8a14610a4b578063517f6d3314610ad457806351b5222514610adc57806352f13a4e14610b025761049a565b806339e618e8146108fa5780633b1eb4bf146109205780633f2708981461093d57806343d96699146109955780634b2c2f441461099d5761049a565b806319113e3b1161046257806319113e3b1461067d57806323f0ab651461069e5780633173b8db146107db57806335244f511461081357806336407b70146108e05761049a565b80630b1ca49a1461049f5780630d1312b8146104d9578063123633ea1461051b578063158ef93e146105385780631904bb2e14610540575b600080fd5b6104c5600480360360208110156104b557600080fd5b5035600160a060020a0316611769565b604080519115158252519081900360200190f35b6104ff600480360360208110156104ef57600080fd5b5035600160a060020a03166118ce565b60408051600160a060020a039092168252519081900360200190f35b6104ff6004803603602081101561053157600080fd5b5035611992565b6104c5611aa6565b6105666004803603602081101561055657600080fd5b5035600160a060020a0316611aaf565b60405180806020018060200186600160a060020a0316600160a060020a0316815260200185815260200184600160a060020a0316600160a060020a03168152602001838103835288818151815260200191508051906020019080838360005b838110156105dd5781810151838201526020016105c5565b50505050905090810190601f16801561060a5780820380516001836020036101000a031916815260200191505b50838103825287518152875160209182019189019080838360005b8381101561063d578181015183820152602001610625565b50505050905090810190601f16801561066a5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b610685611d04565b6040805192835260208301919091528051918290030190f35b6104c5600480360360608110156106b457600080fd5b600160a060020a0382351691908101906040810160208201356401000000008111156106df57600080fd5b8201836020820111156106f157600080fd5b8035906020019184600183028401116401000000008311171561071357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561076657600080fd5b82018360208201111561077857600080fd5b8035906020019184600183028401116401000000008311171561079a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611d2d945050505050565b6104c5600480360360608110156107f157600080fd5b50600160a060020a038135811691602081013582169160409091013516611eb5565b6108396004803603602081101561082957600080fd5b5035600160a060020a0316612022565b604051808060200180602001858152602001848152602001838103835287818151815260200191508051906020019060200280838360005b83811015610889578181015183820152602001610871565b50505050905001838103825286818151815260200191508051906020019060200280838360005b838110156108c85781810151838201526020016108b0565b50505050905001965050505050505060405180910390f35b6108e8612184565b60408051918252519081900360200190f35b6108e86004803603602081101561091057600080fd5b5035600160a060020a031661218a565b6108e86004803603602081101561093657600080fd5b503561220b565b610945612224565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610981578181015183820152602001610969565b505050509050019250505060405180910390f35b6108e8612287565b6108e8600480360360208110156109b357600080fd5b8101906020810181356401000000008111156109ce57600080fd5b8201836020820111156109e057600080fd5b80359060200191846001830284011164010000000083111715610a0257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061228d945050505050565b6108e86123ea565b6104c560048036036060811015610a6157600080fd5b600160a060020a038235811692602081013590911691810190606081016040820135640100000000811115610a9557600080fd5b820183602082011115610aa757600080fd5b80359060200191846001830284011164010000000083111715610ac957600080fd5b5090925090506123f0565b6108e861262d565b6104ff60048036036020811015610af257600080fd5b5035600160a060020a0316612633565b6104c560048036036020811015610b1857600080fd5b5035600160a060020a0316612721565b6108e861273f565b6104c560048036036040811015610b4657600080fd5b5080359060200135612745565b6104ff60048036036040811015610b6957600080fd5b508035906020013561285f565b6104c560048036036020811015610b8c57600080fd5b5035612975565b6108e860048036036020811015610ba957600080fd5b810190602081018135640100000000811115610bc457600080fd5b820183602082011115610bd657600080fd5b80359060200191846001830284011164010000000083111715610bf857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612c5c945050505050565b610c5660048036036020811015610c4f57600080fd5b5035612dae565b005b610c5660048036036020811015610c6e57600080fd5b5035612e54565b610685612f1f565b61094560048036036020811015610c9357600080fd5b810190602081018135640100000000811115610cae57600080fd5b820183602082011115610cc057600080fd5b80359060200191846020830284011164010000000083111715610ce257600080fd5b509092509050612f29565b6104c5600480360360a0811015610d0357600080fd5b600160a060020a038235811692602081013590911691810190606081016040820135640100000000811115610d3757600080fd5b820183602082011115610d4957600080fd5b80359060200191846001830284011164010000000083111715610d6b57600080fd5b919390929091602081019035640100000000811115610d8957600080fd5b820183602082011115610d9b57600080fd5b80359060200191846001830284011164010000000083111715610dbd57600080fd5b919390929091602081019035640100000000811115610ddb57600080fd5b820183602082011115610ded57600080fd5b80359060200191846001830284011164010000000083111715610e0f57600080fd5b509092509050612fbb565b610c566132c7565b6108e861335d565b6104c560048036036040811015610e4057600080fd5b508035906020013561336d565b6108e860048036036020811015610e6357600080fd5b810190602081018135640100000000811115610e7e57600080fd5b820183602082011115610e9057600080fd5b80359060200191846020830284011164010000000083111715610eb257600080fd5b509092509050613471565b610c566004803603610160811015610ed457600080fd5b50600160a060020a038135169060208101359060408101359060608101359060808101359060a08101359060c08101359060e08101359061010081013590610120810135906101400135613590565b6104ff61365e565b610c5660048036036020811015610f4157600080fd5b5035613672565b6108e8613899565b6108e860048036036020811015610f6657600080fd5b810190602081018135640100000000811115610f8157600080fd5b820183602082011115610f9357600080fd5b80359060200191846001830284011164010000000083111715610fb557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506139b6945050505050565b6104c56004803603602081101561100c57600080fd5b5035613b08565b6104ff613e7b565b6109456004803603604081101561103157600080fd5b50600160a060020a038135169060200135613e8a565b6104c56140cb565b6108e86004803603602081101561106557600080fd5b50356140ef565b6104c56004803603606081101561108257600080fd5b50600160a060020a038135811691602081013582169160409091013516614197565b6108e8614504565b6108e8600480360360208110156110c257600080fd5b503561450f565b6110ef600480360360208110156110df57600080fd5b5035600160a060020a031661461a565b60405180806020018881526020018781526020018681526020018060200185815260200184815260200183810383528a818151815260200191508051906020019060200280838360005b83811015611151578181015183820152602001611139565b50505050905001838103825286818151815260200191508051906020019060200280838360005b83811015611190578181015183820152602001611178565b50505050905001995050505050505050505060405180910390f35b610c56600480360360208110156111c157600080fd5b5035600160a060020a031661484a565b6104c5600480360360208110156111e757600080fd5b5035600160a060020a0316614955565b61121d6004803603602081101561120d57600080fd5b5035600160a060020a0316614c01565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561125757818101518382015260200161123f565b50505050905090810190601f1680156112845780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610c56614d8f565b6108e8614f40565b610685614f46565b6104c5600480360360408110156112c057600080fd5b8101906020810181356401000000008111156112db57600080fd5b8201836020820111156112ed57600080fd5b8035906020019184600183028401116401000000008311171561130f57600080fd5b91939092909160208101903564010000000081111561132d57600080fd5b82018360208201111561133f57600080fd5b8035906020019184600183028401116401000000008311171561136157600080fd5b509092509050614f4f565b610c566004803603604081101561138257600080fd5b50600160a060020a038135169060200135615111565b610685615171565b610c56600480360360208110156113b657600080fd5b5035600160a060020a031661517b565b6104c5600480360360208110156113dc57600080fd5b5035600160a060020a0316615353565b6106856153f5565b6104c56004803603602081101561140a57600080fd5b5035600160a060020a03166153fe565b6104c56004803603604081101561143057600080fd5b5080359060200135615519565b610945615697565b6108e86004803603604081101561145b57600080fd5b50600160a060020a0381351690602001356157bf565b610945615821565b6108e86004803603602081101561148f57600080fd5b5035600160a060020a0316615881565b6108e8600480360360208110156114b557600080fd5b5035600160a060020a0316615904565b6108e8615a08565b6108e8615af8565b6104c5600480360360208110156114eb57600080fd5b5035615afe565b610c566004803603602081101561150857600080fd5b5035600160a060020a0316615c3e565b6108e86004803603602081101561152e57600080fd5b5035615db3565b610c56615de5565b6104c56004803603606081101561155357600080fd5b81019060208101813564010000000081111561156e57600080fd5b82018360208201111561158057600080fd5b803590602001918460018302840111640100000000831117156115a257600080fd5b9193909290916020810190356401000000008111156115c057600080fd5b8201836020820111156115d257600080fd5b803590602001918460018302840111640100000000831117156115f457600080fd5b91939092909160208101903564010000000081111561161257600080fd5b82018360208201111561162457600080fd5b8035906020019184600183028401116401000000008311171561164657600080fd5b509092509050615fef565b6104ff6004803603606081101561166757600080fd5b50600160a060020a03813516906020810135906040013561644c565b610685600480360360c081101561169957600080fd5b5080359060208101359060408101359060608101359060808101359060a001356166dd565b6104c5600480360360208110156116d457600080fd5b50356168a4565b6104c5600480360360208110156116f157600080fd5b5035616c64565b610c566004803603602081101561170e57600080fd5b5035600160a060020a0316616d78565b6104c56004803603602081101561173457600080fd5b5035600160a060020a0316616dcd565b6108e86004803603602081101561175a57600080fd5b5035616dfe565b6104c5616f09565b600180548101908190556000908161177f6170c5565b600160a060020a03166364439b43336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b1580156117d757600080fd5b505afa1580156117eb573d6000803e3d6000fd5b505050506040513d602081101561180157600080fd5b5051905061180e81612721565b801561181e575061181e84616dcd565b1515611874576040805160e560020a62461bcd02815260206004820152601a60248201527f6973206e6f742067726f757020616e642076616c696461746f72000000000000604482015290519081900360640190fd5b61187e8185617189565b92505060015481146118c8576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206195c7833981519152604482015290519081900360640190fd5b50919050565b6000806118d9614504565b600160a060020a0384166000908152600460208190526040822060058101549394500191156119305761192b61191d6001846001015461748490919063ffffffff16565b83549063ffffffff6174c616565b611933565b60005b600081815260028401602052604090205490915083141561196b57815481111561196b5761196881600163ffffffff61748416565b90505b600090815260029091016020526040902060010154600160a060020a03169150505b919050565b60408051602080820184905243828401528251808303840181526060928301938490528051600094859360fa939282918401908083835b602083106119e85780518252601f1990920191602091820191016119c9565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114611a48576040519150601f19603f3d011682016040523d82523d6000602084013e611a4d565b606091505b5092509050801515611a935760405160e560020a62461bcd02815260040180806020018281038252603d8152602001806197b0603d913960400191505060405180910390fd5b611a9e826000617523565b949350505050565b60025460ff1681565b6060806000806000611ac086616dcd565b1515611b04576040805160e560020a62461bcd02815260206004820152600f6024820152600080516020619909833981519152604482015290519081900360640190fd5b600160a060020a0380871660009081526004602090815260409182902060028101548351928301909352600381015482529283926001840192911690611b499061752f565b611b516170c5565b600160a060020a0316634ce38b5f8c6040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015611ba957600080fd5b505afa158015611bbd573d6000803e3d6000fd5b505050506040513d6020811015611bd357600080fd5b50518454604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152918791830182828015611c5b5780601f10611c3057610100808354040283529160200191611c5b565b820191906000526020600020905b815481529060010190602001808311611c3e57829003601f168201915b5050875460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959a5089945092508401905082828015611ce95780601f10611cbe57610100808354040283529160200191611ce9565b820191906000526020600020905b815481529060010190602001808311611ccc57829003601f168201915b50505050509350955095509550955095505091939590929450565b600b546040805160208101909152600c5481526000918291611d259061752f565b915091509091565b60008060fb600160a060020a03168585856040516020018084600160a060020a0316600160a060020a03166c0100000000000000000000000002815260140183805190602001908083835b60208310611d975780518252601f199092019160209182019101611d78565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310611ddf5780518252601f199092019160209182019101611dc0565b6001836020036101000a03801982511681845116808217855250505050505090500193505050506040516020818303038152906040526040518082805190602001908083835b60208310611e445780518252601f199092019160209182019101611e25565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114611ea4576040519150601f19603f3d011682016040523d82523d6000602084013e611ea9565b606091505b50909695505050505050565b6001805481019081905560009081611ecb6170c5565b600160a060020a03166364439b43336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015611f2357600080fd5b505afa158015611f37573d6000803e3d6000fd5b505050506040513d6020811015611f4d57600080fd5b5051600160a060020a0381166000908152600360208190526040909120015490915015611fc4576040805160e560020a62461bcd02815260206004820152601960248201527f56616c696461746f722067726f7570206e6f7420656d70747900000000000000604482015290519081900360640190fd5b611fd081878787617533565b925050600154811461201a576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206195c7833981519152604482015290519081900360640190fd5b509392505050565b60608060008060006004600087600160a060020a0316600160a060020a03168152602001908152602001600020600401905060608160010154604051908082528060200260200182016040528015612084578160200160208202803883390190505b509050606082600101546040519080825280602002602001820160405280156120b7578160200160208202803883390190505b50905060005b836001015481101561216d5783546000906120de908363ffffffff6174c616565b600081815260028701602052604090205485519192509085908490811061210157fe5b602090810290910181019190915260008281526002870190915260409020600101548351600160a060020a039091169084908490811061213d57fe5b600160a060020a039092166020928302909101909101525061216681600163ffffffff6174c616565b90506120bd565b506003830154925491989097509195509350915050565b60105481565b600061219582612721565b15156121eb576040805160e560020a62461bcd02815260206004820152601360248201527f4e6f742076616c696461746f722067726f757000000000000000000000000000604482015290519081900360640190fd5b50600160a060020a03166000908152600360208190526040909120015490565b600061221e82612219615a08565b6179f6565b92915050565b6060600580548060200260200160405190810160405280929190818152602001828054801561227c57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161225e575b505050505090505b90565b600e5490565b60006060600060f4600160a060020a0316846040516020018082805190602001908083835b602083106122d15780518252601f1990920191602091820191016122b2565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106123345780518252601f199092019160209182019101612315565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114612394576040519150601f19603f3d011682016040523d82523d6000602084013e612399565b606091505b50925090508015156123df5760405160e560020a62461bcd0281526004018080602001828103825260388152602001806196d66038913960400191505060405180910390fd5b611a9e826000617a2a565b600d5481565b604080517f4163636f756e7473000000000000000000000000000000000000000000000000602080830191909152825180830360080181526028830180855281519183019190912060025460e060020a63dcf0aaed02909252602c84018190529351600094933393610100909304600160a060020a03169263dcf0aaed92604c80840193829003018186803b15801561248857600080fd5b505afa15801561249c573d6000803e3d6000fd5b505050506040513d60208110156124b257600080fd5b5051600160a060020a031614612512576040805160e560020a62461bcd02815260206004820152601860248201527f6f6e6c79207265676973746572656420636f6e74726163740000000000000000604482015290519081900360640190fd5b61251b86616dcd565b151561255f576040805160e560020a62461bcd02815260206004820152600f6024820152600080516020619909833981519152604482015290519081900360640190fd5b60006004600088600160a060020a0316600160a060020a0316815260200190815260200160002090506125ca81888888888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250617a9392505050565b1515612620576040805160e560020a62461bcd02815260206004820152601f60248201527f4572726f72207570646174696e67204543445341207075626c6963206b657900604482015290519081900360640190fd5b5060019695505050505050565b60065490565b60008061263e6170c5565b600160a060020a03166393c5c487846040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b15801561269657600080fd5b505afa1580156126aa573d6000803e3d6000fd5b505050506040513d60208110156126c057600080fd5b505190506126cd81616dcd565b1515612711576040805160e560020a62461bcd02815260206004820152600f6024820152600080516020619909833981519152604482015290519081900360640190fd5b61271a816118ce565b9392505050565b600160a060020a031660009081526003602052604090205460ff1690565b600e5481565b600061274f6140cb565b1515612793576040805160e560020a62461bcd02815260206004820181905260248201526000805160206198e9833981519152604482015290519081900360640190fd5b60098054841415806127a9575080600101548314155b15156127ff576040805160e560020a62461bcd02815260206004820152601e60248201527f47726f757020726571756972656d656e7473206e6f74206368616e6765640000604482015290519081900360640190fd5b60408051808201825285815260209081018590526009869055600a859055815186815290810185905281517f999f7ee1917e6d7ea08360edfe9250cda3eda859c38dcb71a92623665de64dd4929181900390910190a15060019392505050565b6040805160208082018590528183018490528251808303840181526060928301938490528051600094859360fa939282918401908083835b602083106128b65780518252601f199092019160209182019101612897565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114612916576040519150601f19603f3d011682016040523d82523d6000602084013e61291b565b606091505b50925090508015156129615760405160e560020a62461bcd0281526004018080602001828103825260368152602001806198226036913960400191505060405180910390fd5b61296c826000617523565b95945050505050565b600180548101908190556000908161298b6170c5565b600160a060020a03166364439b43336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b1580156129e357600080fd5b505afa1580156129f7573d6000803e3d6000fd5b505050506040513d6020811015612a0d57600080fd5b50519050612a1a81612721565b1515612a5e576040805160e560020a62461bcd0281526020600482015260156024820152600080516020619669833981519152604482015290519081900360640190fd5b600160a060020a0381166000908152600360208190526040909120015415612ad0576040805160e560020a62461bcd02815260206004820152601960248201527f56616c696461746f722067726f7570206e6f7420656d70747900000000000000604482015290519081900360640190fd5b600160a060020a0381166000908152600360205260409020600801805460011015612b665742612b27600960010154836001815481101515612b0e57fe5b90600052602060002001546174c690919063ffffffff16565b10612b665760405160e560020a62461bcd028152600401808060200182810382526021815260200180619a166021913960400191505060405180910390fd5b600160a060020a03821660009081526003602081905260408220805460ff19168155600181018390556002810183905590810182905560058101829055600681018290556007810182905590612bbf600883018261944d565b50600060098201819055600a90910155612bdb60058387617c19565b604051600160a060020a038316907fae7e034b0748a10a219b46074b20977a9170bf4027b156c797093773619a866990600090a260019350505060015481146118c8576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206195c7833981519152604482015290519081900360640190fd5b60006060600060f6600160a060020a0316846040516020018082805190602001908083835b60208310612ca05780518252601f199092019160209182019101612c81565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b60208310612d035780518252601f199092019160209182019101612ce4565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114612d63576040519150601f19603f3d011682016040523d82523d6000602084013e612d68565b606091505b50925090508015156123df5760405160e560020a62461bcd028152600401808060200182810382526023815260200180619a846023913960400191505060405180910390fd5b60018054810190819055612dc06140cb565b1515612e04576040805160e560020a62461bcd02815260206004820181905260248201526000805160206198e9833981519152604482015290519081900360640190fd5b60108290556001548114612e50576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206195c7833981519152604482015290519081900360640190fd5b5050565b612e5c6140cb565b1515612ea0576040805160e560020a62461bcd02815260206004820181905260248201526000805160206198e9833981519152604482015290519081900360640190fd5b600f54811415612ee45760405160e560020a62461bcd0281526004018080602001828103825260238152602001806199c26023913960400191505060405180910390fd5b600f8190556040805182815290517ff2da07d08fd8dc9c5dcf87ad6f540e306f884a47dd8de14b718a4d5395f1ca9b9181900360200190a150565b600954600a549091565b60608083839050604051908082528060200260200182016040528015612f59578160200160208202803883390190505b50905060005b8381101561201a57612f8b858583818110612f7657fe5b90506020020135600160a060020a031661218a565b8282815181101515612f9957fe5b60209081029091010152612fb481600163ffffffff6174c616565b9050612f5f565b604080517f4163636f756e7473000000000000000000000000000000000000000000000000602080830191909152825180830360080181526028830180855281519183019190912060025460e060020a63dcf0aaed02909252602c84018190529351600094933393610100909304600160a060020a03169263dcf0aaed92604c80840193829003018186803b15801561305357600080fd5b505afa158015613067573d6000803e3d6000fd5b505050506040513d602081101561307d57600080fd5b5051600160a060020a0316146130dd576040805160e560020a62461bcd02815260206004820152601860248201527f6f6e6c79207265676973746572656420636f6e74726163740000000000000000604482015290519081900360640190fd5b6130e68a616dcd565b151561312a576040805160e560020a62461bcd02815260206004820152600f6024820152600080516020619909833981519152604482015290519081900360640190fd5b6000600460008c600160a060020a0316600160a060020a031681526020019081526020016000209050613195818c8c8c8c8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250617a9392505050565b15156131eb576040805160e560020a62461bcd02815260206004820152601f60248201527f4572726f72207570646174696e67204543445341207075626c6963206b657900604482015290519081900360640190fd5b613260818c89898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8d018190048102820181019092528b815292508b91508a9081908401838280828437600092019190915250617d4692505050565b15156132b6576040805160e560020a62461bcd02815260206004820152601d60248201527f4572726f72207570646174696e6720424c53207075626c6963206b6579000000604482015290519081900360640190fd5b5060019a9950505050505050505050565b6132cf6140cb565b1515613313576040805160e560020a62461bcd02815260206004820181905260248201526000805160206198e9833981519152604482015290519081900360640190fd5b60008054604051600160a060020a03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a360008054600160a060020a0319169055565b600061336843615db3565b905090565b60006133776140cb565b15156133bb576040805160e560020a62461bcd02815260206004820181905260248201526000805160206198e9833981519152604482015290519081900360640190fd5b60078054841415806133d1575080600101548314155b15156134115760405160e560020a62461bcd0281526004018080602001828103825260228152602001806196b46022913960400191505060405180910390fd5b604080518082018252858152602090810185905260078690556008859055815186815290810185905281517f62d947118dd4c1f5ece7f787a9cad4e1127d14d403b71133e95792b473bf8389929181900390910190a15060019392505050565b60008082116134ca576040805160e560020a62461bcd02815260206004820152601260248201527f557074696d6520617272617920656d7074790000000000000000000000000000604482015290519081900360640190fd5b600e5482111561350e5760405160e560020a62461bcd02815260040180806020018281038252602b815260200180619689602b913960400191505060405180910390fd5b61351661946b565b60005b8381101561356e5761355461354761354287878581811061353657fe5b905060200201356140ef565b617eda565b839063ffffffff617ef416565b915061356781600163ffffffff6174c616565b9050613519565b50611a9e61358b61357e85617f70565b839063ffffffff617fe116565b61752f565b60025460ff16156135eb576040805160e560020a62461bcd02815260206004820152601c60248201527f636f6e747261637420616c726561647920696e697469616c697a656400000000604482015290519081900360640190fd5b6002805460ff19166001179055613601336180d0565b61360a8b61484a565b6136148a8a612745565b5061361f888861336d565b5061362a8686615519565b5061363482615afe565b5061363e81612e54565b61364784616c64565b5061365183612dae565b5050505050505050505050565b6002546101009004600160a060020a031681565b600061367c6170c5565b600160a060020a03166364439b43336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b1580156136d457600080fd5b505afa1580156136e8573d6000803e3d6000fd5b505050506040513d60208110156136fe57600080fd5b5051905061370b81612721565b151561374f576040805160e560020a62461bcd0281526020600482015260156024820152600080516020619669833981519152604482015290519081900360640190fd5b600160a060020a038116600090815260036020526040902061377261358b618175565b8311156137b35760405160e560020a62461bcd028152600401808060200182810382526025815260200180619a376025913960400191505060405180910390fd5b6040805160208101909152600582015481526137ce9061752f565b831415613825576040805160e560020a62461bcd02815260206004820152601c60248201527f436f6d6d697373696f6e206d75737420626520646966666572656e7400000000604482015290519081900360640190fd5b61382e83617eda565b516006820155600f5461384890439063ffffffff6174c616565b600782018190556040805185815260208101929092528051600160a060020a038516927f557d39a57520d9835859d4b7eda805a7f4115a59c3a374eeed488436fc62a15292908290030190a2505050565b60006060600060f9600160a060020a031643604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b602083106138f95780518252601f1990920191602091820191016138da565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114613959576040519150601f19603f3d011682016040523d82523d6000602084013e61395e565b606091505b50925090508015156139a45760405160e560020a62461bcd0281526004018080602001828103825260358152602001806197ed6035913960400191505060405180910390fd5b6139af826000617523565b9250505090565b60006060600060f7600160a060020a0316846040516020018082805190602001908083835b602083106139fa5780518252601f1990920191602091820191016139db565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b60208310613a5d5780518252601f199092019160209182019101613a3e565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114613abd576040519150601f19603f3d011682016040523d82523d6000602084013e613ac2565b606091505b5092509050801515611a935760405160e560020a62461bcd0281526004018080602001828103825260318152602001806199e56031913960400191505060405180910390fd5b6001805481019081905560009081613b1e6170c5565b600160a060020a03166364439b43336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015613b7657600080fd5b505afa158015613b8a573d6000803e3d6000fd5b505050506040513d6020811015613ba057600080fd5b50519050613bad81616dcd565b1515613bf1576040805160e560020a62461bcd02815260206004820152600f6024820152600080516020619909833981519152604482015290519081900360640190fd5b600160a060020a038082166000908152600460205260409020600281015490911615613d15576002810154600160a060020a03908116600090815260036020908152604091829020825160e060020a63542424fb028152600190910160048201529285166024840152905173__$bd3fd89075918f478685066c5f6c8518ac$__9263542424fb926044808301939192829003018186803b158015613c9457600080fd5b505af4158015613ca8573d6000803e3d6000fd5b505050506040513d6020811015613cbe57600080fd5b505115613d15576040805160e560020a62461bcd02815260206004820152601e60248201527f486173206265656e2067726f7570206d656d62657220726563656e746c790000604482015290519081900360640190fd5b6008546007820154600091613d30919063ffffffff6174c616565b9050428110613d89576040805160e560020a62461bcd02815260206004820152601c60248201527f4e6f742079657420726571756972656d656e7420656e642074696d6500000000604482015290519081900360640190fd5b613d9560068488617c19565b600160a060020a0383166000908152600460205260408120908181613dba828261947d565b613dc860018301600061947d565b505050600281018054600160a060020a031916905560006003820181905560048201819055600582018190556007909101819055604051600160a060020a038516917f51407fafe7ef9bec39c65a12a4885a274190991bf1e9057fcc384fc77ff1a7f091a26001945050505060015481146118c8576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206195c7833981519152604482015290519081900360640190fd5b600054600160a060020a031690565b600160a060020a03821660009081526003602052604080822081517fb1cfea4300000000000000000000000000000000000000000000000000000000815260019091016004820152602481018490529051606092839273__$bd3fd89075918f478685066c5f6c8518ac$__9263b1cfea4392604480840193919291829003018186803b158015613f1957600080fd5b505af4158015613f2d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015613f5657600080fd5b810190808051640100000000811115613f6e57600080fd5b82016020810184811115613f8157600080fd5b8151856020820283011164010000000082111715613f9e57600080fd5b50509291905050509050606083604051908082528060200260200182016040528015613fd4578160200160208202803883390190505b50905060005b848110156140c257613fea6170c5565b600160a060020a0316634ce38b5f848381518110151561400657fe5b906020019060200201516040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b15801561405957600080fd5b505afa15801561406d573d6000803e3d6000fd5b505050506040513d602081101561408357600080fd5b5051825183908390811061409357fe5b600160a060020a039092166020928302909101909101526140bb81600163ffffffff6174c616565b9050613fda565b50949350505050565b60008054600160a060020a03166140e0618199565b600160a060020a031614905090565b60006140fc61358b618175565b821115614153576040805160e560020a62461bcd02815260206004820181905260248201527f557074696d652063616e6e6f74206265206c6172676572207468616e206f6e65604482015290519081900360640190fd5b60008061418561416461358b618175565b61416f61358b618175565b8661417b61358b618175565b600b5460126166dd565b9092509050611a9e61358b838361819d565b60018054810190819055600090816141ad6170c5565b600160a060020a03166364439b43336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b15801561420557600080fd5b505afa158015614219573d6000803e3d6000fd5b505050506040513d602081101561422f57600080fd5b5051905061423c81612721565b1515614292576040805160e560020a62461bcd02815260206004820152600b60248201527f4e6f7420612067726f7570000000000000000000000000000000000000000000604482015290519081900360640190fd5b61429b86616dcd565b15156142df576040805160e560020a62461bcd02815260206004820152600f6024820152600080516020619909833981519152604482015290519081900360640190fd5b600160a060020a03808216600090815260036020908152604091829020825160e060020a63542424fb028152600182016004820152938a1660248501529151919273__$bd3fd89075918f478685066c5f6c8518ac$__9263542424fb92604480840193919291829003018186803b15801561435957600080fd5b505af415801561436d573d6000803e3d6000fd5b505050506040513d602081101561438357600080fd5b505115156143db576040805160e560020a62461bcd02815260206004820152601960248201527f4e6f742061206d656d626572206f66207468652067726f757000000000000000604482015290519081900360640190fd5b604080517fb2f8fe96000000000000000000000000000000000000000000000000000000008152600183016004820152600160a060020a03808a166024830152808916604483015287166064820152905173__$bd3fd89075918f478685066c5f6c8518ac$__9163b2f8fe96916084808301926000929190829003018186803b15801561446757600080fd5b505af415801561447b573d6000803e3d6000fd5b5050604051600160a060020a03808b169350851691507f38819cc49a343985b478d72f531a35b15384c398dd80fd191a14662170f895c690600090a3600193505050600154811461201a576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206195c7833981519152604482015290519081900360640190fd5b60006133684361220b565b60006060600060f9600160a060020a031684604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b6020831061456f5780518252601f199092019160209182019101614550565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d80600081146145cf576040519150601f19603f3d011682016040523d82523d6000602084013e6145d4565b606091505b5092509050801515611a935760405160e560020a62461bcd02815260040180806020018281038252602e815260200180619599602e913960400191505060405180910390fd5b60606000806000606060008061462f88612721565b1515614673576040805160e560020a62461bcd0281526020600482015260156024820152600080516020619669833981519152604482015290519081900360640190fd5b600160a060020a03881660009081526003602052604080822081517ffe3c7a8e0000000000000000000000000000000000000000000000000000000081526001820160048201529151909273__$bd3fd89075918f478685066c5f6c8518ac$__9263fe3c7a8e9260248083019392829003018186803b1580156146f557600080fd5b505af4158015614709573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561473257600080fd5b81019080805164010000000081111561474a57600080fd5b8201602081018481111561475d57600080fd5b815185602082028301116401000000008211171561477a57600080fd5b505060408051602081019091526005860154815290935061479d9250905061752f565b6040805160208101909152600684015481526147b89061752f565b600784015460408051602081019091526009860154815260088601906147dd9061752f565b600a87015482546040805160208084028201810190925282815291859183018282801561482957602002820191906000526020600020905b815481526020019060010190808311614815575b50505050509250975097509750975097509750975050919395979092949650565b6148526140cb565b1515614896576040805160e560020a62461bcd02815260206004820181905260248201526000805160206198e9833981519152604482015290519081900360640190fd5b600160a060020a03811615156148f6576040805160e560020a62461bcd02815260206004820181905260248201527f43616e6e6f7420726567697374657220746865206e756c6c2061646472657373604482015290519081900360640190fd5b6002805474ffffffffffffffffffffffffffffffffffffffff001916610100600160a060020a038416908102919091179091556040517f27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b90600090a250565b600180548101908190556000908161496b6170c5565b600160a060020a03166364439b43336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b1580156149c357600080fd5b505afa1580156149d7573d6000803e3d6000fd5b505050506040513d60208110156149ed57600080fd5b505190506149fa81616dcd565b1515614a3e576040805160e560020a62461bcd02815260206004820152600f6024820152600080516020619909833981519152604482015290519081900360640190fd5b614a4784612721565b1515614a8b576040805160e560020a62461bcd0281526020600482015260156024820152600080516020619669833981519152604482015290519081900360640190fd5b614a9481615353565b1515614ad45760405160e560020a62461bcd0281526004018080602001828103825260238152602001806195766023913960400191505060405180910390fd5b614add84615353565b1515614b33576040805160e560020a62461bcd02815260206004820152601f60248201527f47726f757020646f65736e2774206d65657420726571756972656d656e747300604482015290519081900360640190fd5b600160a060020a038082166000908152600460205260409020600281015490911615614b6557614b6381836181d5565b505b600281018054600160a060020a031916600160a060020a0387811691821790925560405190918416907f91ef92227057e201e406c3451698dd780fe7672ad74328591c88d281af31581d90600090a360019350505060015481146118c8576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206195c7833981519152604482015290519081900360640190fd5b60606000614c0d6170c5565b600160a060020a03166393c5c487846040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015614c6557600080fd5b505afa158015614c79573d6000803e3d6000fd5b505050506040513d6020811015614c8f57600080fd5b50519050614c9c81616dcd565b1515614ce0576040805160e560020a62461bcd02815260206004820152600f6024820152600080516020619909833981519152604482015290519081900360640190fd5b600160a060020a038116600090815260046020908152604091829020600190810180548451600293821615610100026000190190911692909204601f810184900484028301840190945283825290929091830182828015614d825780601f10614d5757610100808354040283529160200191614d82565b820191906000526020600020905b815481529060010190602001808311614d6557829003601f168201915b5050505050915050919050565b600180548101908190556000614da36170c5565b600160a060020a03166364439b43336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015614dfb57600080fd5b505afa158015614e0f573d6000803e3d6000fd5b505050506040513d6020811015614e2557600080fd5b50519050614e3281612721565b1515614e76576040805160e560020a62461bcd0281526020600482015260156024820152600080516020619669833981519152604482015290519081900360640190fd5b600160a060020a0381166000908152600360205260409020601054600a820154614ea59163ffffffff6174c616565b421015614ee65760405160e560020a62461bcd02815260040180806020018281038252603b81526020018061960d603b913960400191505060405180910390fd5b614eee618175565b51600990910155506001548114614f3d576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206195c7833981519152604482015290519081900360640190fd5b50565b600f5490565b60075460085482565b600080614f5a6170c5565b600160a060020a03166364439b43336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015614fb257600080fd5b505afa158015614fc6573d6000803e3d6000fd5b505050506040513d6020811015614fdc57600080fd5b50519050614fe981616dcd565b151561502d576040805160e560020a62461bcd02815260206004820152600f6024820152600080516020619909833981519152604482015290519081900360640190fd5b600160a060020a0381166000908152600460209081526040918290208251601f8901839004830281018301909352878352916150bb91839185918b908b908190840183828082843760009201919091525050604080516020601f8d018190048102820181019092528b815292508b91508a9081908401838280828437600092019190915250617d4692505050565b1515612620576040805160e560020a62461bcd02815260206004820152601d60248201527f4572726f72207570646174696e6720424c53207075626c6963206b6579000000604482015290519081900360640190fd5b3315615167576040805160e560020a62461bcd02815260206004820152601060248201527f4f6e6c7920564d2063616e2063616c6c00000000000000000000000000000000604482015290519081900360640190fd5b612e5082826182e6565b6007546008549091565b6001805481019081905561518d618522565b600160a060020a03166357601c5d336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b1580156151e557600080fd5b505afa1580156151f9573d6000803e3d6000fd5b505050506040513d602081101561520f57600080fd5b50511515615267576040805160e560020a62461bcd02815260206004820181905260248201527f4f6e6c79207265676973746572656420736c61736865722063616e2063616c6c604482015290519081900360640190fd5b61527082612721565b15156152b4576040805160e560020a62461bcd0281526020600482015260156024820152600080516020619669833981519152604482015290519081900360640190fd5b600160a060020a0382166000908152600360209081526040918290208251918201909252600982015481526152ff90613542906002906152f39061752f565b9063ffffffff6185b516565b51600982015542600a909101556001548114612e50576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206195c7833981519152604482015290519081900360640190fd5b60008061535e618522565b600160a060020a03166330ec70f5846040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b1580156153b657600080fd5b505afa1580156153ca573d6000803e3d6000fd5b505050506040513d60208110156153e057600080fd5b505190506153ed83615904565b111592915050565b600954600a5482565b60018054810190819055600090816154146170c5565b600160a060020a03166364439b43336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b15801561546c57600080fd5b505afa158015615480573d6000803e3d6000fd5b505050506040513d602081101561549657600080fd5b5051600160a060020a0381166000908152600360208190526040822001549192501061550c576040805160e560020a62461bcd02815260206004820152601560248201527f56616c696461746f722067726f757020656d7074790000000000000000000000604482015290519081900360640190fd5b61187e8185600080617533565b60006155236140cb565b1515615567576040805160e560020a62461bcd02815260206004820181905260248201526000805160206198e9833981519152604482015290519081900360640190fd5b61557261358b618175565b8211156155b35760405160e560020a62461bcd028152600401808060200182810382526028815260200180619a5c6028913960400191505060405180910390fd5b600b54831415806155ea57506040805160208101909152600c5481526155e8906155dc84617eda565b9063ffffffff6185f716565b155b151561562a5760405160e560020a62461bcd028152600401808060200182810382526029815260200180619aa76029913960400191505060405180910390fd5b604080519081016040528084815260200161564484617eda565b90528051600b5560209081015151600c556040805185815291820184905280517f4b48724280029c2ea7a445c9cea30838525342e7a9ea9468f630b52e75d6c5369281900390910190a150600192915050565b606060006156a36170c5565b905060606006805490506040519080825280602002602001820160405280156156d6578160200160208202803883390190505b50905060005b81518110156157b85782600160a060020a0316634ce38b5f60068381548110151561570357fe5b600091825260209182902001546040805160e060020a63ffffffff8616028152600160a060020a0390921660048301525160248083019392829003018186803b15801561574f57600080fd5b505afa158015615763573d6000803e3d6000fd5b505050506040513d602081101561577957600080fd5b5051825183908390811061578957fe5b600160a060020a039092166020928302909101909101526157b181600163ffffffff6174c616565b90506156dc565b5091505090565b60003315615817576040805160e560020a62461bcd02815260206004820152601060248201527f4f6e6c7920564d2063616e2063616c6c00000000000000000000000000000000604482015290519081900360640190fd5b61271a83836185fe565b6060600680548060200260200160405190810160405280929190818152602001828054801561227c57602002820191906000526020600020908154600160a060020a0316815260019091019060200180831161225e575050505050905090565b600061588c82612721565b15156158d0576040805160e560020a62461bcd0281526020600482015260156024820152600080516020619669833981519152604482015290519081900360640190fd5b600160a060020a03821660009081526003602090815260409182902082519182019092526009820154815261271a9061752f565b600061590f82616dcd565b1561591d575060075461198d565b61592682612721565b15615a0057600160a060020a038216600090815260036020819052604082200154615953906001906189b8565b600160a060020a038416600090815260036020526040812060080180549293509111156159e457805460009061599090600163ffffffff61748416565b90505b60008111156159e257426159b46009600101548484815481101515612b0e57fe5b106159ca576159c381846189b8565b92506159e2565b6159db81600163ffffffff61748416565b9050615993565b505b6009546159f7908363ffffffff6189cf16565b9250505061198d565b506000919050565b604080516000808252602082019283905281519092606092849260f89290819081908082805b60208310615a4d5780518252601f199092019160209182019101615a2e565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114615aad576040519150601f19603f3d011682016040523d82523d6000602084013e615ab2565b606091505b50925090508015156139a45760405160e560020a62461bcd0281526004018080602001828103825260258152602001806199716025913960400191505060405180910390fd5b600f5481565b6000615b086140cb565b1515615b4c576040805160e560020a62461bcd02815260206004820181905260248201526000805160206198e9833981519152604482015290519081900360640190fd5b60008211615ba4576040805160e560020a62461bcd02815260206004820152601d60248201527f4d61782067726f75702073697a652063616e6e6f74206265207a65726f000000604482015290519081900360640190fd5b600e54821415615bfe576040805160e560020a62461bcd02815260206004820152601a60248201527f4d61782067726f75702073697a65206e6f74206368616e676564000000000000604482015290519081900360640190fd5b600e8290556040805183815290517f603fe12c33c253a23da1680aa453dc70c3a0ee07763569bd5f602406ebd4e5d59181900360200190a1506001919050565b60018054810190819055615c50618522565b600160a060020a03166357601c5d336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015615ca857600080fd5b505afa158015615cbc573d6000803e3d6000fd5b505050506040513d6020811015615cd257600080fd5b50511515615d2a576040805160e560020a62461bcd02815260206004820181905260248201527f4f6e6c79207265676973746572656420736c61736865722063616e2063616c6c604482015290519081900360640190fd5b615d3382616dcd565b15615d6c57600160a060020a038083166000908152600460205260409020600281015490911615615d6a57615d6881846181d5565b505b505b6001548114612e50576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206195c7833981519152604482015290519081900360640190fd5b600061221e60036152f36002615dd96002615dcd8861450f565b9063ffffffff6189cf16565b9063ffffffff6174c616565b6000615def6170c5565b600160a060020a03166364439b43336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015615e4757600080fd5b505afa158015615e5b573d6000803e3d6000fd5b505050506040513d6020811015615e7157600080fd5b50519050615e7e81612721565b1515615ec2576040805160e560020a62461bcd0281526020600482015260156024820152600080516020619669833981519152604482015290519081900360640190fd5b600160a060020a038116600090815260036020526040902060078101541515615f35576040805160e560020a62461bcd02815260206004820152601b60248201527f4e6f20636f6d6d697373696f6e20757064617465207175657565640000000000604482015290519081900360640190fd5b6007810154431015615f7b5760405160e560020a62461bcd0281526004018080602001828103825260218152602001806196486021913960400191505060405180910390fd5b6006810180546005830181905560009182905560078301919091556040805160208101909152908152600160a060020a038316907f815d292dbc1a08dfb3103aabb6611233dd2393903e57bdf4c5b3db91198a826c90615fda9061752f565b60408051918252519081900360200190a25050565b60018054810190819055600090816160056170c5565b600160a060020a03166364439b43336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b15801561605d57600080fd5b505afa158015616071573d6000803e3d6000fd5b505050506040513d602081101561608757600080fd5b5051905061609481616dcd565b1580156160a757506160a581612721565b155b15156160fd576040805160e560020a62461bcd02815260206004820152601260248201527f416c726561647920726567697374657265640000000000000000000000000000604482015290519081900360640190fd5b6000616107618522565b600160a060020a03166330ec70f5836040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b15801561615f57600080fd5b505afa158015616173573d6000803e3d6000fd5b505050506040513d602081101561618957600080fd5b50516007549091508110156161e8576040805160e560020a62461bcd02815260206004820152601160248201527f4465706f73697420746f6f20736d616c6c000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a0382166000908152600460205260408120906162096170c5565b600160a060020a0316634ce38b5f856040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b15801561626157600080fd5b505afa158015616275573d6000803e3d6000fd5b505050506040513d602081101561628b57600080fd5b810190808051906020019092919050505090506162e08285838f8f8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250617a9392505050565b5061636782858c8c8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250617d4692505050565b506006805460018101825560009182527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f018054600160a060020a031916600160a060020a0387161790556163bd908590618a2f565b50604051600160a060020a038516907fd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c8885325190600090a260019550505050506001548114616441576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206195c7833981519152604482015290519081900360640190fd5b509695505050505050565b600061645784616dcd565b151561649b576040805160e560020a62461bcd02815260206004820152600f6024820152600080516020619909833981519152604482015290519081900360640190fd5b6164a3614504565b8311156164e45760405160e560020a62461bcd0281526004018080602001828103825260238152602001806198806023913960400191505060405180910390fd5b600160a060020a0384166000908152600460208190526040909120600581015491018054909161651a919063ffffffff6174c616565b8310616570576040805160e560020a62461bcd02815260206004820152601360248201527f696e646578206f7574206f6620626f756e647300000000000000000000000000604482015290519081900360640190fd5b80548310801590616585575060008160010154115b15156165db576040805160e560020a62461bcd02815260206004820152601360248201527f696e646578206f7574206f6620626f756e647300000000000000000000000000604482015290519081900360640190fd5b600083815260028201602052604081205460018084015491871492916166069163ffffffff61748416565b835461661990879063ffffffff61748416565b60008781526002860160205260408120549190921492508711801561666c57508660028501600061665189600163ffffffff6174c616565b815260200190815260200160002060000154118061666c5750815b905082806166775750805b15156166b75760405160e560020a62461bcd0281526004018080602001828103825260478152602001806197696047913960600191505060405180910390fd5b5050506000928352600201602052506040902060010154600160a060020a031692915050565b60008086158015906166ee57508415155b1515616744576040805160e560020a62461bcd02815260206004820152601560248201527f612064656e6f6d696e61746f72206973207a65726f0000000000000000000000604482015290519081900360640190fd5b6000806000606060fc600160a060020a03168c8c8c8c8c8c6040516020018087815260200186815260200185815260200184815260200183815260200182815260200196505050505050506040516020818303038152906040526040518082805190602001908083835b602083106167cd5780518252601f1990920191602091820191016167ae565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d806000811461682d576040519150601f19603f3d011682016040523d82523d6000602084013e616832565b606091505b5090925090508115156168795760405160e560020a62461bcd0281526004018080602001828103825260278152602001806199296027913960400191505060405180910390fd5b616884816000617523565b9350616891816020617523565b939c939b50929950505050505050505050565b600180548101908190556000906168bc61358b618175565b8311156168fd5760405160e560020a62461bcd028152600401808060200182810382526025815260200180619a376025913960400191505060405180910390fd5b60006169076170c5565b600160a060020a03166364439b43336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b15801561695f57600080fd5b505afa158015616973573d6000803e3d6000fd5b505050506040513d602081101561698957600080fd5b5051905061699681616dcd565b156169eb576040805160e560020a62461bcd02815260206004820152601f60248201527f416c726561647920726567697374657265642061732076616c696461746f7200604482015290519081900360640190fd5b6169f481612721565b15616a49576040805160e560020a62461bcd02815260206004820152601b60248201527f416c726561647920726567697374657265642061732067726f75700000000000604482015290519081900360640190fd5b6000616a53618522565b600160a060020a03166330ec70f5836040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015616aab57600080fd5b505afa158015616abf573d6000803e3d6000fd5b505050506040513d6020811015616ad557600080fd5b5051600954909150811015616b34576040805160e560020a62461bcd02815260206004820152601660248201527f4e6f7420656e6f756768206c6f636b656420676f6c6400000000000000000000604482015290519081900360640190fd5b600160a060020a0382166000908152600360205260409020805460ff19166001178155616b6086617eda565b5160058201556040805180820190915280616b79618175565b815260006020918201819052825151600985015591810151600a8401556005805460018101825592527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db09091018054600160a060020a031916600160a060020a03861690811790915560408051898152905191927fbf4b45570f1907a94775f8449817051a492a676918e38108bb762e991e6b58dc92918290030190a26001945050505060015481146118c8576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206195c7833981519152604482015290519081900360640190fd5b6000616c6e6140cb565b1515616cb2576040805160e560020a62461bcd02815260206004820181905260248201526000805160206198e9833981519152604482015290519081900360640190fd5b60008211616cf45760405160e560020a62461bcd0281526004018080602001828103825260288152602001806198586028913960400191505060405180910390fd5b600d54821415616d385760405160e560020a62461bcd0281526004018080602001828103825260258152602001806198c46025913960400191505060405180910390fd5b600d8290556040805183815290517f1c75c7fb3ee9d13d8394372d8c7cdf1702fa947faa03f6ccfa500f787b09b48a9181900360200190a1506001919050565b616d806140cb565b1515616dc4576040805160e560020a62461bcd02815260206004820181905260248201526000805160206198e9833981519152604482015290519081900360640190fd5b614f3d816180d0565b600160a060020a03166000908152600460205260408120600190810154600291811615610100026000190116041190565b60006060600060f5600160a060020a031684604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b60208310616e5e5780518252601f199092019160209182019101616e3f565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114616ebe576040519150601f19603f3d011682016040523d82523d6000602084013e616ec3565b606091505b50925090508015156123df5760405160e560020a62461bcd02815260040180806020018281038252602c815260200180619996602c913960400191505060405180910390fd5b6001805481019081905560009081616f1f6170c5565b600160a060020a03166364439b43336040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015616f7757600080fd5b505afa158015616f8b573d6000803e3d6000fd5b505050506040513d6020811015616fa157600080fd5b50519050616fae81616dcd565b1515616ff2576040805160e560020a62461bcd02815260206004820152600f6024820152600080516020619909833981519152604482015290519081900360640190fd5b600160a060020a03808216600090815260046020526040902060028101549091161515617069576040805160e560020a62461bcd02815260206004820152601b60248201527f6465616666696c696174653a206e6f7420616666696c69617465640000000000604482015290519081900360640190fd5b61707381836181d5565b5060019350505060015481146170c1576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206195c7833981519152604482015290519081900360640190fd5b5090565b600254604080517f4163636f756e7473000000000000000000000000000000000000000000000000602080830191909152825160088184030181526028830180855281519183019190912060e060020a63dcf0aaed02909152602c83015291516000936101009004600160a060020a03169263dcf0aaed92604c8082019391829003018186803b15801561715857600080fd5b505afa15801561716c573d6000803e3d6000fd5b505050506040513d602081101561718257600080fd5b5051905090565b600160a060020a0380831660008181526003602090815260408083208686168452600490925282206002015491939092911614617210576040805160e560020a62461bcd02815260206004820152601760248201527f4e6f7420616666696c696174656420746f2067726f7570000000000000000000604482015290519081900360640190fd5b6040805160e060020a63542424fb028152600183016004820152600160a060020a0385166024820152905173__$bd3fd89075918f478685066c5f6c8518ac$__9163542424fb916044808301926020929190829003018186803b15801561727657600080fd5b505af415801561728a573d6000803e3d6000fd5b505050506040513d60208110156172a057600080fd5b505115156172f8576040805160e560020a62461bcd02815260206004820152601960248201527f4e6f742061206d656d626572206f66207468652067726f757000000000000000604482015290519081900360640190fd5b604080517fe2c0c56a000000000000000000000000000000000000000000000000000000008152600183016004820152600160a060020a0385166024820152905173__$bd3fd89075918f478685066c5f6c8518ac$__9163e2c0c56a916044808301926000929190829003018186803b15801561737457600080fd5b505af4158015617388573d6000803e3d6000fd5b505050506003810154801515617413576173a0618cbc565b600160a060020a031663a8e45871866040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050600060405180830381600087803b1580156173fa57600080fd5b505af115801561740e573d6000803e3d6000fd5b505050505b61741e846000618a2f565b506174398561743483600163ffffffff6174c616565b618d4f565b83600160a060020a031685600160a060020a03167fc7666a52a66ff601ff7c0d4d6efddc9ac20a34792f6aa003d1804c9d4d5baa5760405160405180910390a3506001949350505050565b600061271a83836040805190810160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250618e0b565b60008282018381101561271a576040805160e560020a62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b600061271a8383617a2a565b5190565b600061753e85612721565b801561754e575061754e84616dcd565b15156175a4576040805160e560020a62461bcd02815260206004820152601760248201527f4e6f742076616c696461746f7220616e642067726f7570000000000000000000604482015290519081900360640190fd5b600160a060020a0385166000908152600360208190526040909120600e549181015490911161761d576040805160e560020a62461bcd02815260206004820152601f60248201527f67726f757020776f756c6420657863656564206d6178696d756d2073697a6500604482015290519081900360640190fd5b600160a060020a03858116600090815260046020526040902060020154811690871614617694576040805160e560020a62461bcd02815260206004820152601760248201527f4e6f7420616666696c696174656420746f2067726f7570000000000000000000604482015290519081900360640190fd5b6040805160e060020a63542424fb028152600183016004820152600160a060020a0387166024820152905173__$bd3fd89075918f478685066c5f6c8518ac$__9163542424fb916044808301926020929190829003018186803b1580156176fa57600080fd5b505af415801561770e573d6000803e3d6000fd5b505050506040513d602081101561772457600080fd5b50511561777b576040805160e560020a62461bcd02815260206004820152601060248201527f416c726561647920696e2067726f757000000000000000000000000000000000604482015290519081900360640190fd5b600381015460009061779490600163ffffffff6174c616565b604080517f26afac49000000000000000000000000000000000000000000000000000000008152600185016004820152600160a060020a0389166024820152905191925073__$bd3fd89075918f478685066c5f6c8518ac$__916326afac4991604480820192600092909190829003018186803b15801561781457600080fd5b505af4158015617828573d6000803e3d6000fd5b5050505061783587615353565b151561788b576040805160e560020a62461bcd02815260206004820152601a60248201527f47726f757020726571756972656d656e7473206e6f74206d6574000000000000604482015290519081900360640190fd5b61789486615353565b15156178ea576040805160e560020a62461bcd02815260206004820152601e60248201527f56616c696461746f7220726571756972656d656e7473206e6f74206d65740000604482015290519081900360640190fd5b8060011415617989576178fb618cbc565b604080517fa18fb2db000000000000000000000000000000000000000000000000000000008152600160a060020a038a81166004830152888116602483015287811660448301529151929091169163a18fb2db9160648082019260009290919082900301818387803b15801561797057600080fd5b505af1158015617984573d6000803e3d6000fd5b505050505b6179938688618a2f565b506179a98761743483600163ffffffff61748416565b85600160a060020a031687600160a060020a03167fbdf7e616a6943f81e07a7984c9d4c00197dc2f481486ce4ffa6af52a113974ad60405160405180910390a35060019695505050505050565b6000808284811515617a0457fe5b0490508284811515617a1257fe5b061515617a2057905061221e565b600101905061221e565b600081602001835110151515617a8a576040805160e560020a62461bcd02815260206004820152601460248201527f736c6963696e67206f7574206f662072616e6765000000000000000000000000604482015290519081900360640190fd5b50016020015190565b8051600090604014617aef576040805160e560020a62461bcd02815260206004820152601d60248201527f57726f6e67204543445341207075626c6963206b6579206c656e677468000000604482015290519081900360640190fd5b81516020830120600160a060020a03848116911614617b58576040805160e560020a62461bcd02815260206004820152601f60248201527f4543445341206b657920646f6573206e6f74206d61746368207369676e657200604482015290519081900360640190fd5b8151617b6a90869060208501906194c1565b5083600160a060020a03167f213377eec2c15b21fa7abcbb0cb87a67e893cdb94a2564aa4bb4d380869473c8836040518080602001828103825283818151815260200191508051906020019080838360005b83811015617bd4578181015183820152602001617bbc565b50505050905090810190601f168015617c015780820380516001836020036101000a031916815260200191505b509250505060405180910390a2506001949350505050565b825481108015617c53575081600160a060020a03168382815481101515617c3c57fe5b600091825260209091200154600160a060020a0316145b1515617c935760405160e560020a62461bcd0281526004018080602001828103825260218152602001806199506021913960400191505060405180910390fd5b8254600090617ca990600163ffffffff61748416565b90508381815481101515617cb957fe5b6000918252602090912001548454600160a060020a0390911690859084908110617cdf57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055508381815481101515617d1b57fe5b60009182526020909120018054600160a060020a031916905580617d3f858261953b565b5050505050565b8151600090606014617da2576040805160e560020a62461bcd02815260206004820152601b60248201527f57726f6e6720424c53207075626c6963206b6579206c656e6774680000000000604482015290519081900360640190fd5b8151603014617dfb576040805160e560020a62461bcd02815260206004820152601460248201527f57726f6e6720424c5320506f50206c656e677468000000000000000000000000604482015290519081900360640190fd5b617e06848484611d2d565b1515617e5c576040805160e560020a62461bcd02815260206004820152600f60248201527f496e76616c696420424c5320506f500000000000000000000000000000000000604482015290519081900360640190fd5b8251617e7190600187019060208601906194c1565b5083600160a060020a03167f36a1aabe506bbe8802233cbb9aad628e91269e77077c953f9db3e02d7092ee338460405180806020018281038252838181518152602001915080519060200190808383600083811015617bd4578181015183820152602001617bbc565b617ee261946b565b50604080516020810190915290815290565b617efc61946b565b8151835190810190811015617f5b576040805160e560020a62461bcd02815260206004820152601560248201527f616464206f766572666c6f772064657465637465640000000000000000000000604482015290519081900360640190fd5b60408051602081019091529081529392505050565b617f7861946b565b617f80618ea5565b821115617fc15760405160e560020a62461bcd02815260040180806020018281038252603681526020018061970e6036913960400191505060405180910390fd5b50604080516020810190915269d3c21bcecceda100000082028152919050565b617fe961946b565b81511515618041576040805160e560020a62461bcd02815260206004820152601160248201527f63616e2774206469766964652062792030000000000000000000000000000000604482015290519081900360640190fd5b825169d3c21bcecceda100000081810291908204146180aa576040805160e560020a62461bcd02815260206004820152601260248201527f6f766572666c6f77206174206469766964650000000000000000000000000000604482015290519081900360640190fd5b6020604051908101604052808460000151838115156180c557fe5b049052949350505050565b600160a060020a038116151561811a5760405160e560020a62461bcd0281526004018080602001828103825260268152602001806195e76026913960400191505060405180910390fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a360008054600160a060020a031916600160a060020a0392909216919091179055565b61817d61946b565b50604080516020810190915269d3c21bcecceda1000000815290565b3390565b6181a561946b565b6181ad61946b565b6181b684617f70565b90506181c061946b565b6181c984617f70565b905061296c8282617fe1565b6002820154600160a060020a039081166000818152600360209081526040808320815160e060020a63542424fb028152600182016004820152958716602487015290519294909273__$bd3fd89075918f478685066c5f6c8518ac$__9263542424fb9260448082019391829003018186803b15801561825357600080fd5b505af4158015618267573d6000803e3d6000fd5b505050506040513d602081101561827d57600080fd5b5051156182905761828e8285617189565b505b600285018054600160a060020a0319169055604051600160a060020a0383811691908616907f71815121f0622b31a3e7270eb28acb9fd10825ff418c9a18591f617bb8a31a6c90600090a3506001949350505050565b60006182f06170c5565b600160a060020a03166393c5c487846040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b15801561834857600080fd5b505afa15801561835c573d6000803e3d6000fd5b505050506040513d602081101561837257600080fd5b5051905061837f81616dcd565b15156183c3576040805160e560020a62461bcd02815260206004820152600f6024820152600080516020619909833981519152604482015290519081900360640190fd5b6183cb61946b565b6183d7613542846140ef565b90506183e161946b565b6040805160208101909152600c548152618401908363ffffffff618ec016565b905061840b61946b565b6040805160208101909152600c54815261843390618427618175565b9063ffffffff61925216565b600160a060020a0385166000908152600460209081526040918290208251918201909252600390910154815290915061847390829063ffffffff618ec016565b905061849c6135426184848561752f565b61849761358b868663ffffffff617ef416565b6192cd565b600160a060020a038516600081815260046020908152604091829020935160039094018490558151908101909152918252907fedf9f87e50e10c533bf3ae7f5a7894ae66c23e6cbbe8773d7765d20ad6f995e9906184f99061752f565b6185028661752f565b6040805192835260208301919091528051918290030190a2505050505050565b600254604080517f4c6f636b6564476f6c64000000000000000000000000000000000000000000006020808301919091528251600a818403018152602a830180855281519183019190912060e060020a63dcf0aaed02909152602e83015291516000936101009004600160a060020a03169263dcf0aaed92604e8082019391829003018186803b15801561715857600080fd5b600061271a83836040805190810160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506192dc565b5190511490565b6000806186096170c5565b600160a060020a03166393c5c487856040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b15801561866157600080fd5b505afa158015618675573d6000803e3d6000fd5b505050506040513d602081101561868b57600080fd5b5051905061869881616dcd565b15156186dc576040805160e560020a62461bcd02815260206004820152600f6024820152600080516020619909833981519152604482015290519081900360640190fd5b60006186e7826118ce565b9050600160a060020a03811615156187335760405160e560020a62461bcd0281526004018080602001828103825260258152602001806197446025913960400191505060405180910390fd5b61873c82615353565b801561874c575061874c81615353565b156189ad5761875961946b565b600160a060020a03808316600090815260036020818152604080842081518084018352600990910154815294881684526004825292839020835191820190935291015481526187ba91906187ae908189617f70565b9063ffffffff618ec016565b600160a060020a0383166000908152600360209081526040808320815192830190915260050154815291925090618801906187fc90849063ffffffff618ec016565b619349565b9050600061881e8261881285619349565b9063ffffffff61748416565b905061882861935a565b600160a060020a03166340c10f1985846040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561888a57600080fd5b505af115801561889e573d6000803e3d6000fd5b505050506040513d60208110156188b457600080fd5b506188bf905061935a565b600160a060020a03166340c10f1986836040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561892157600080fd5b505af1158015618935573d6000803e3d6000fd5b505050506040513d602081101561894b57600080fd5b505060408051828152602081018490528151600160a060020a0380881693908916927f6f5937add2ec38a0fa4959bccd86e3fcc2aafb706cd3e6c0565f87a7b36b9975929081900390910190a36189a183619349565b9550505050505061221e565b60009250505061221e565b6000818310156189c8578161271a565b5090919050565b60008215156189e05750600061221e565b8282028284828115156189ef57fe5b041461271a5760405160e560020a62461bcd0281526004018080602001828103825260218152602001806198a36021913960400191505060405180910390fd5b600160a060020a038216600090815260046020819052604082200181618a53614504565b905060008260010154600014618a9157618a8c618a7e6001856001015461748490919063ffffffff16565b84549063ffffffff6174c616565b618a94565b60005b905060008360010154118015618ab15750600160a060020a038516155b15618abd574260038401555b6000818152600284016020526040902054821415618b2a57604080518082018252928352600160a060020a0386811660208086019182526000948552600296909601909552912091518255915160019182018054600160a060020a0319169190931617909155905061221e565b60008360010154600014618b4e57618b4982600163ffffffff6174c616565b618b51565b60005b604080518082018252858152600160a060020a038981166020808401918252600086815260028b019091529390932091518255915160019182018054600160a060020a0319169190931617909155600d54908601549192501115618bce57600184810154618bc49163ffffffff6174c616565b6001850155612620565b600d5484600101541415618c1c5783546000908152600285016020526040812090815560019081018054600160a060020a03191690558454618c159163ffffffff6174c616565b8455612620565b835460009081526002850160208190526040822082815560019081018054600160a060020a03191690558654919291618c5a9163ffffffff6174c616565b81526020810191909152604001600090812090815560019081018054600160a060020a031916905584810154618c959163ffffffff61748416565b60018501558354618cad90600263ffffffff6174c616565b84555060019695505050505050565b600254604080517f456c656374696f6e000000000000000000000000000000000000000000000000602080830191909152825160088184030181526028830180855281519183019190912060e060020a63dcf0aaed02909152602c83015291516000936101009004600160a060020a03169263dcf0aaed92604c8082019391829003018186803b15801561715857600080fd5b600160a060020a03821660009081526003602052604090206008018054821415618d8d57805460018101825560008281526020902042910155618e06565b8054821015618db657428183815481101515618da557fe5b600091825260209091200155618e06565b6040805160e560020a62461bcd02815260206004820152601d60248201527f556e61626c6520746f207570646174652073697a6520686973746f7279000000604482015290519081900360640190fd5b505050565b60008184841115618e9d5760405160e560020a62461bcd0281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015618e62578181015183820152602001618e4a565b50505050905090810190601f168015618e8f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b7601357c299a88ea76a58924d52ce4f26a85af186c2b9e7490565b618ec861946b565b82511580618ed557508151155b15618eef575060408051602081019091526000815261221e565b815169d3c21bcecceda10000001415618f0957508161221e565b825169d3c21bcecceda10000001415618f2357508061221e565b600069d3c21bcecceda1000000618f39856193ed565b51811515618f4357fe5b0490506000618f5185619418565b519050600069d3c21bcecceda1000000618f6a866193ed565b51811515618f7457fe5b0490506000618f8286619418565b5190508382028415618ff057828582811515618f9a57fe5b0414618ff0576040805160e560020a62461bcd02815260206004820152601660248201527f6f766572666c6f77207831793120646574656374656400000000000000000000604482015290519081900360640190fd5b69d3c21bcecceda10000008102811561906f5769d3c21bcecceda1000000828281151561901957fe5b041461906f576040805160e560020a62461bcd02815260206004820152601f60248201527f6f766572666c6f772078317931202a2066697865643120646574656374656400604482015290519081900360640190fd5b90508084840285156190dd5784868281151561908757fe5b04146190dd576040805160e560020a62461bcd02815260206004820152601660248201527f6f766572666c6f77207832793120646574656374656400000000000000000000604482015290519081900360640190fd5b8684028715619148578488828115156190f257fe5b0414619148576040805160e560020a62461bcd02815260206004820152601660248201527f6f766572666c6f77207831793220646574656374656400000000000000000000604482015290519081900360640190fd5b619150619444565b8781151561915a57fe5b049650619165619444565b8581151561916f57fe5b04945086850287156191dd5785888281151561918757fe5b04146191dd576040805160e560020a62461bcd02815260206004820152601660248201527f6f766572666c6f77207832793220646574656374656400000000000000000000604482015290519081900360640190fd5b6191e561946b565b50604080516020818101835287825282519081019092528482529061920b908290617ef4565b90506192268160206040519081016040528086815250617ef4565b90506192418160206040519081016040528085815250617ef4565b9d9c50505050505050505050505050565b61925a61946b565b8151835110156192b4576040805160e560020a62461bcd02815260206004820152601f60248201527f737562737472616374696f6e20756e646572666c6f7720646574656374656400604482015290519081900360640190fd5b5060408051602081019091528151835103815292915050565b60008183106189c8578161271a565b6000818184116193315760405160e560020a62461bcd02815260040180806020018281038252838181518152602001915080519060200190808383600083811015618e62578181015183820152602001618e4a565b506000838581151561933f57fe5b0495945050505050565b5169d3c21bcecceda1000000900490565b600254604080517f537461626c65546f6b656e0000000000000000000000000000000000000000006020808301919091528251600b818403018152602b830180855281519183019190912060e060020a63dcf0aaed02909152602f83015291516000936101009004600160a060020a03169263dcf0aaed92604f8082019391829003018186803b15801561715857600080fd5b6193f561946b565b506040805160208101909152905169d3c21bcecceda10000009081900402815290565b61942061946b565b506040805160208101909152905169d3c21bcecceda1000000808204029003815290565b64e8d4a5100090565b5080546000825590600052602060002090810190614f3d919061955b565b60408051602081019091526000815290565b50805460018160011615610100020316600290046000825580601f106194a35750614f3d565b601f016020900490600052602060002090810190614f3d919061955b565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061950257805160ff191683800117855561952f565b8280016001018555821561952f579182015b8281111561952f578251825591602001919060010190619514565b506170c192915061955b565b815481835581811115618e0657600083815260209020618e069181019083015b61228491905b808211156170c1576000815560010161956156fe56616c696461746f7220646f65736e2774206d65657420726571756972656d656e74736572726f722063616c6c696e67206e756d62657256616c696461746f7273496e53657420707265636f6d70696c657265656e7472616e742063616c6c0000000000000000000000000000000000004f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373607265736574536c617368696e674d756c7469706c696572602063616c6c6564206265666f7265207265736574506572696f64206578706972656443616e2774206170706c7920636f6d6d697373696f6e20757064617465207965744e6f7420612076616c696461746f722067726f75700000000000000000000000557074696d65206172726179206c6172676572207468616e206d6178696d756d2067726f75702073697a6556616c696461746f7220726571756972656d656e7473206e6f74206368616e6765646572726f722063616c6c696e672067657456657269666965645365616c4269746d617046726f6d48656164657220707265636f6d70696c6563616e277420637265617465206669786964697479206e756d626572206c6172676572207468616e206d61784e65774669786564282956616c696461746f72206e6f742072656769737465726564207769746820612067726f757070726f766964656420696e64657820646f6573206e6f74206d617463682070726f76696465642065706f63684e756d62657220617420696e64657820696e20686973746f72792e6572726f722063616c6c696e672076616c696461746f725369676e65724164647265737346726f6d43757272656e7453657420707265636f6d70696c656572726f722063616c6c696e67206e756d62657256616c696461746f7273496e43757272656e7453657420707265636f6d70696c656572726f722063616c6c696e672076616c696461746f725369676e65724164647265737346726f6d53657420707265636f6d70696c654d656d6265727368697020686973746f7279206c656e6774682063616e6e6f74206265207a65726f45706f63682063616e6e6f74206265206c6172676572207468616e2063757272656e74536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f774d656d6265727368697020686973746f7279206c656e677468206e6f74206368616e6765644f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65724e6f7420612076616c696461746f7200000000000000000000000000000000006572726f722063616c6c696e67206672616374696f6e4d756c45787020707265636f6d70696c6564656c657465456c656d656e743a20696e646578206f7574206f662072616e67656572726f722063616c6c696e672067657445706f636853697a6520707265636f6d70696c656572726f722063616c6c696e6720676574506172656e745365616c4269746d617020707265636f6d70696c65636f6d6d697373696f6e207570646174652064656c6179206e6f74206368616e6765646572726f722063616c6c696e6720676574426c6f636b4e756d62657246726f6d48656164657220707265636f6d70696c654861736e2774206265656e20656d70747920666f72206c6f6e6720656e6f756768436f6d6d697373696f6e2063616e27742062652067726561746572207468616e203130302541646a7573746d656e742073706565642063616e6e6f74206265206c6172676572207468616e20316572726f722063616c6c696e67206861736848656164657220707265636f6d70696c6541646a7573746d656e7420737065656420616e64206578706f6e656e74206e6f74206368616e676564a165627a7a72305820639c7e6ee4da7ba48e238fb4afc951d6a43b206da60a6755f9080768eb6bc4340029"

// DeployValidators deploys a new Ethereum contract, binding an instance of Validators to it.
func DeployValidators(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Validators, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validators{ValidatorsCaller: ValidatorsCaller{contract: contract}, ValidatorsTransactor: ValidatorsTransactor{contract: contract}, ValidatorsFilterer: ValidatorsFilterer{contract: contract}}, nil
}

// Validators is an auto generated Go binding around an Ethereum contract.
type Validators struct {
	ValidatorsCaller     // Read-only binding to the contract
	ValidatorsTransactor // Write-only binding to the contract
	ValidatorsFilterer   // Log filterer for contract events
}

// ValidatorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorsSession struct {
	Contract     *Validators       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorsCallerSession struct {
	Contract *ValidatorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ValidatorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorsTransactorSession struct {
	Contract     *ValidatorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ValidatorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorsRaw struct {
	Contract *Validators // Generic contract binding to access the raw methods on
}

// ValidatorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorsCallerRaw struct {
	Contract *ValidatorsCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorsTransactorRaw struct {
	Contract *ValidatorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidators creates a new instance of Validators, bound to a specific deployed contract.
func NewValidators(address common.Address, backend bind.ContractBackend) (*Validators, error) {
	contract, err := bindValidators(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validators{ValidatorsCaller: ValidatorsCaller{contract: contract}, ValidatorsTransactor: ValidatorsTransactor{contract: contract}, ValidatorsFilterer: ValidatorsFilterer{contract: contract}}, nil
}

// NewValidatorsCaller creates a new read-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsCaller(address common.Address, caller bind.ContractCaller) (*ValidatorsCaller, error) {
	contract, err := bindValidators(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsCaller{contract: contract}, nil
}

// NewValidatorsTransactor creates a new write-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorsTransactor, error) {
	contract, err := bindValidators(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsTransactor{contract: contract}, nil
}

// NewValidatorsFilterer creates a new log filterer instance of Validators, bound to a specific deployed contract.
func NewValidatorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorsFilterer, error) {
	contract, err := bindValidators(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorsFilterer{contract: contract}, nil
}

// bindValidators binds a generic wrapper to an already deployed contract.
func bindValidators(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validators *ValidatorsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validators.Contract.ValidatorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validators *ValidatorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.Contract.ValidatorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validators *ValidatorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validators.Contract.ValidatorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validators *ValidatorsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validators.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validators *ValidatorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validators *ValidatorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validators.Contract.contract.Transact(opts, method, params...)
}

// CalculateEpochScore is a free data retrieval call binding the contract method 0x94903a97.
//
// Solidity: function calculateEpochScore(uint256 uptime) constant returns(uint256)
func (_Validators *ValidatorsCaller) CalculateEpochScore(opts *bind.CallOpts, uptime *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "calculateEpochScore", uptime)
	return *ret0, err
}

// CalculateEpochScore is a free data retrieval call binding the contract method 0x94903a97.
//
// Solidity: function calculateEpochScore(uint256 uptime) constant returns(uint256)
func (_Validators *ValidatorsSession) CalculateEpochScore(uptime *big.Int) (*big.Int, error) {
	return _Validators.Contract.CalculateEpochScore(&_Validators.CallOpts, uptime)
}

// CalculateEpochScore is a free data retrieval call binding the contract method 0x94903a97.
//
// Solidity: function calculateEpochScore(uint256 uptime) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) CalculateEpochScore(uptime *big.Int) (*big.Int, error) {
	return _Validators.Contract.CalculateEpochScore(&_Validators.CallOpts, uptime)
}

// CalculateGroupEpochScore is a free data retrieval call binding the contract method 0x76f7425d.
//
// Solidity: function calculateGroupEpochScore(uint256[] uptimes) constant returns(uint256)
func (_Validators *ValidatorsCaller) CalculateGroupEpochScore(opts *bind.CallOpts, uptimes []*big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "calculateGroupEpochScore", uptimes)
	return *ret0, err
}

// CalculateGroupEpochScore is a free data retrieval call binding the contract method 0x76f7425d.
//
// Solidity: function calculateGroupEpochScore(uint256[] uptimes) constant returns(uint256)
func (_Validators *ValidatorsSession) CalculateGroupEpochScore(uptimes []*big.Int) (*big.Int, error) {
	return _Validators.Contract.CalculateGroupEpochScore(&_Validators.CallOpts, uptimes)
}

// CalculateGroupEpochScore is a free data retrieval call binding the contract method 0x76f7425d.
//
// Solidity: function calculateGroupEpochScore(uint256[] uptimes) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) CalculateGroupEpochScore(uptimes []*big.Int) (*big.Int, error) {
	return _Validators.Contract.CalculateGroupEpochScore(&_Validators.CallOpts, uptimes)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) constant returns(bool)
func (_Validators *ValidatorsCaller) CheckProofOfPossession(opts *bind.CallOpts, sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "checkProofOfPossession", sender, blsKey, blsPop)
	return *ret0, err
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) constant returns(bool)
func (_Validators *ValidatorsSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _Validators.Contract.CheckProofOfPossession(&_Validators.CallOpts, sender, blsKey, blsPop)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) constant returns(bool)
func (_Validators *ValidatorsCallerSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _Validators.Contract.CheckProofOfPossession(&_Validators.CallOpts, sender, blsKey, blsPop)
}

// CommissionUpdateDelay is a free data retrieval call binding the contract method 0xe0e3ffe6.
//
// Solidity: function commissionUpdateDelay() constant returns(uint256)
func (_Validators *ValidatorsCaller) CommissionUpdateDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "commissionUpdateDelay")
	return *ret0, err
}

// CommissionUpdateDelay is a free data retrieval call binding the contract method 0xe0e3ffe6.
//
// Solidity: function commissionUpdateDelay() constant returns(uint256)
func (_Validators *ValidatorsSession) CommissionUpdateDelay() (*big.Int, error) {
	return _Validators.Contract.CommissionUpdateDelay(&_Validators.CallOpts)
}

// CommissionUpdateDelay is a free data retrieval call binding the contract method 0xe0e3ffe6.
//
// Solidity: function commissionUpdateDelay() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) CommissionUpdateDelay() (*big.Int, error) {
	return _Validators.Contract.CommissionUpdateDelay(&_Validators.CallOpts)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) constant returns(uint256, uint256)
func (_Validators *ValidatorsCaller) FractionMulExp(opts *bind.CallOpts, aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Validators.contract.Call(opts, out, "fractionMulExp", aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
	return *ret0, *ret1, err
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) constant returns(uint256, uint256)
func (_Validators *ValidatorsSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _Validators.Contract.FractionMulExp(&_Validators.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) constant returns(uint256, uint256)
func (_Validators *ValidatorsCallerSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _Validators.Contract.FractionMulExp(&_Validators.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// GetAccountLockedGoldRequirement is a free data retrieval call binding the contract method 0xdcff4cf6.
//
// Solidity: function getAccountLockedGoldRequirement(address account) constant returns(uint256)
func (_Validators *ValidatorsCaller) GetAccountLockedGoldRequirement(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getAccountLockedGoldRequirement", account)
	return *ret0, err
}

// GetAccountLockedGoldRequirement is a free data retrieval call binding the contract method 0xdcff4cf6.
//
// Solidity: function getAccountLockedGoldRequirement(address account) constant returns(uint256)
func (_Validators *ValidatorsSession) GetAccountLockedGoldRequirement(account common.Address) (*big.Int, error) {
	return _Validators.Contract.GetAccountLockedGoldRequirement(&_Validators.CallOpts, account)
}

// GetAccountLockedGoldRequirement is a free data retrieval call binding the contract method 0xdcff4cf6.
//
// Solidity: function getAccountLockedGoldRequirement(address account) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetAccountLockedGoldRequirement(account common.Address) (*big.Int, error) {
	return _Validators.Contract.GetAccountLockedGoldRequirement(&_Validators.CallOpts, account)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) constant returns(uint256)
func (_Validators *ValidatorsCaller) GetBlockNumberFromHeader(opts *bind.CallOpts, header []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getBlockNumberFromHeader", header)
	return *ret0, err
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) constant returns(uint256)
func (_Validators *ValidatorsSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _Validators.Contract.GetBlockNumberFromHeader(&_Validators.CallOpts, header)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _Validators.Contract.GetBlockNumberFromHeader(&_Validators.CallOpts, header)
}

// GetCommissionUpdateDelay is a free data retrieval call binding the contract method 0xb915f530.
//
// Solidity: function getCommissionUpdateDelay() constant returns(uint256)
func (_Validators *ValidatorsCaller) GetCommissionUpdateDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getCommissionUpdateDelay")
	return *ret0, err
}

// GetCommissionUpdateDelay is a free data retrieval call binding the contract method 0xb915f530.
//
// Solidity: function getCommissionUpdateDelay() constant returns(uint256)
func (_Validators *ValidatorsSession) GetCommissionUpdateDelay() (*big.Int, error) {
	return _Validators.Contract.GetCommissionUpdateDelay(&_Validators.CallOpts)
}

// GetCommissionUpdateDelay is a free data retrieval call binding the contract method 0xb915f530.
//
// Solidity: function getCommissionUpdateDelay() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetCommissionUpdateDelay() (*big.Int, error) {
	return _Validators.Contract.GetCommissionUpdateDelay(&_Validators.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() constant returns(uint256)
func (_Validators *ValidatorsCaller) GetEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getEpochNumber")
	return *ret0, err
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() constant returns(uint256)
func (_Validators *ValidatorsSession) GetEpochNumber() (*big.Int, error) {
	return _Validators.Contract.GetEpochNumber(&_Validators.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetEpochNumber() (*big.Int, error) {
	return _Validators.Contract.GetEpochNumber(&_Validators.CallOpts)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsCaller) GetEpochNumberOfBlock(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getEpochNumberOfBlock", blockNumber)
	return *ret0, err
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _Validators.Contract.GetEpochNumberOfBlock(&_Validators.CallOpts, blockNumber)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _Validators.Contract.GetEpochNumberOfBlock(&_Validators.CallOpts, blockNumber)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() constant returns(uint256)
func (_Validators *ValidatorsCaller) GetEpochSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getEpochSize")
	return *ret0, err
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() constant returns(uint256)
func (_Validators *ValidatorsSession) GetEpochSize() (*big.Int, error) {
	return _Validators.Contract.GetEpochSize(&_Validators.CallOpts)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetEpochSize() (*big.Int, error) {
	return _Validators.Contract.GetEpochSize(&_Validators.CallOpts)
}

// GetGroupLockedGoldRequirements is a free data retrieval call binding the contract method 0x6fa47647.
//
// Solidity: function getGroupLockedGoldRequirements() constant returns(uint256, uint256)
func (_Validators *ValidatorsCaller) GetGroupLockedGoldRequirements(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Validators.contract.Call(opts, out, "getGroupLockedGoldRequirements")
	return *ret0, *ret1, err
}

// GetGroupLockedGoldRequirements is a free data retrieval call binding the contract method 0x6fa47647.
//
// Solidity: function getGroupLockedGoldRequirements() constant returns(uint256, uint256)
func (_Validators *ValidatorsSession) GetGroupLockedGoldRequirements() (*big.Int, *big.Int, error) {
	return _Validators.Contract.GetGroupLockedGoldRequirements(&_Validators.CallOpts)
}

// GetGroupLockedGoldRequirements is a free data retrieval call binding the contract method 0x6fa47647.
//
// Solidity: function getGroupLockedGoldRequirements() constant returns(uint256, uint256)
func (_Validators *ValidatorsCallerSession) GetGroupLockedGoldRequirements() (*big.Int, *big.Int, error) {
	return _Validators.Contract.GetGroupLockedGoldRequirements(&_Validators.CallOpts)
}

// GetGroupNumMembers is a free data retrieval call binding the contract method 0x39e618e8.
//
// Solidity: function getGroupNumMembers(address account) constant returns(uint256)
func (_Validators *ValidatorsCaller) GetGroupNumMembers(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getGroupNumMembers", account)
	return *ret0, err
}

// GetGroupNumMembers is a free data retrieval call binding the contract method 0x39e618e8.
//
// Solidity: function getGroupNumMembers(address account) constant returns(uint256)
func (_Validators *ValidatorsSession) GetGroupNumMembers(account common.Address) (*big.Int, error) {
	return _Validators.Contract.GetGroupNumMembers(&_Validators.CallOpts, account)
}

// GetGroupNumMembers is a free data retrieval call binding the contract method 0x39e618e8.
//
// Solidity: function getGroupNumMembers(address account) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetGroupNumMembers(account common.Address) (*big.Int, error) {
	return _Validators.Contract.GetGroupNumMembers(&_Validators.CallOpts, account)
}

// GetGroupsNumMembers is a free data retrieval call binding the contract method 0x70447754.
//
// Solidity: function getGroupsNumMembers(address[] accounts) constant returns(uint256[])
func (_Validators *ValidatorsCaller) GetGroupsNumMembers(opts *bind.CallOpts, accounts []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getGroupsNumMembers", accounts)
	return *ret0, err
}

// GetGroupsNumMembers is a free data retrieval call binding the contract method 0x70447754.
//
// Solidity: function getGroupsNumMembers(address[] accounts) constant returns(uint256[])
func (_Validators *ValidatorsSession) GetGroupsNumMembers(accounts []common.Address) ([]*big.Int, error) {
	return _Validators.Contract.GetGroupsNumMembers(&_Validators.CallOpts, accounts)
}

// GetGroupsNumMembers is a free data retrieval call binding the contract method 0x70447754.
//
// Solidity: function getGroupsNumMembers(address[] accounts) constant returns(uint256[])
func (_Validators *ValidatorsCallerSession) GetGroupsNumMembers(accounts []common.Address) ([]*big.Int, error) {
	return _Validators.Contract.GetGroupsNumMembers(&_Validators.CallOpts, accounts)
}

// GetMaxGroupSize is a free data retrieval call binding the contract method 0x43d96699.
//
// Solidity: function getMaxGroupSize() constant returns(uint256)
func (_Validators *ValidatorsCaller) GetMaxGroupSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getMaxGroupSize")
	return *ret0, err
}

// GetMaxGroupSize is a free data retrieval call binding the contract method 0x43d96699.
//
// Solidity: function getMaxGroupSize() constant returns(uint256)
func (_Validators *ValidatorsSession) GetMaxGroupSize() (*big.Int, error) {
	return _Validators.Contract.GetMaxGroupSize(&_Validators.CallOpts)
}

// GetMaxGroupSize is a free data retrieval call binding the contract method 0x43d96699.
//
// Solidity: function getMaxGroupSize() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetMaxGroupSize() (*big.Int, error) {
	return _Validators.Contract.GetMaxGroupSize(&_Validators.CallOpts)
}

// GetMembershipHistory is a free data retrieval call binding the contract method 0x35244f51.
//
// Solidity: function getMembershipHistory(address account) constant returns(uint256[], address[], uint256, uint256)
func (_Validators *ValidatorsCaller) GetMembershipHistory(opts *bind.CallOpts, account common.Address) ([]*big.Int, []common.Address, *big.Int, *big.Int, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]common.Address)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Validators.contract.Call(opts, out, "getMembershipHistory", account)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetMembershipHistory is a free data retrieval call binding the contract method 0x35244f51.
//
// Solidity: function getMembershipHistory(address account) constant returns(uint256[], address[], uint256, uint256)
func (_Validators *ValidatorsSession) GetMembershipHistory(account common.Address) ([]*big.Int, []common.Address, *big.Int, *big.Int, error) {
	return _Validators.Contract.GetMembershipHistory(&_Validators.CallOpts, account)
}

// GetMembershipHistory is a free data retrieval call binding the contract method 0x35244f51.
//
// Solidity: function getMembershipHistory(address account) constant returns(uint256[], address[], uint256, uint256)
func (_Validators *ValidatorsCallerSession) GetMembershipHistory(account common.Address) ([]*big.Int, []common.Address, *big.Int, *big.Int, error) {
	return _Validators.Contract.GetMembershipHistory(&_Validators.CallOpts, account)
}

// GetMembershipInLastEpoch is a free data retrieval call binding the contract method 0x0d1312b8.
//
// Solidity: function getMembershipInLastEpoch(address account) constant returns(address)
func (_Validators *ValidatorsCaller) GetMembershipInLastEpoch(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getMembershipInLastEpoch", account)
	return *ret0, err
}

// GetMembershipInLastEpoch is a free data retrieval call binding the contract method 0x0d1312b8.
//
// Solidity: function getMembershipInLastEpoch(address account) constant returns(address)
func (_Validators *ValidatorsSession) GetMembershipInLastEpoch(account common.Address) (common.Address, error) {
	return _Validators.Contract.GetMembershipInLastEpoch(&_Validators.CallOpts, account)
}

// GetMembershipInLastEpoch is a free data retrieval call binding the contract method 0x0d1312b8.
//
// Solidity: function getMembershipInLastEpoch(address account) constant returns(address)
func (_Validators *ValidatorsCallerSession) GetMembershipInLastEpoch(account common.Address) (common.Address, error) {
	return _Validators.Contract.GetMembershipInLastEpoch(&_Validators.CallOpts, account)
}

// GetMembershipInLastEpochFromSigner is a free data retrieval call binding the contract method 0x51b52225.
//
// Solidity: function getMembershipInLastEpochFromSigner(address signer) constant returns(address)
func (_Validators *ValidatorsCaller) GetMembershipInLastEpochFromSigner(opts *bind.CallOpts, signer common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getMembershipInLastEpochFromSigner", signer)
	return *ret0, err
}

// GetMembershipInLastEpochFromSigner is a free data retrieval call binding the contract method 0x51b52225.
//
// Solidity: function getMembershipInLastEpochFromSigner(address signer) constant returns(address)
func (_Validators *ValidatorsSession) GetMembershipInLastEpochFromSigner(signer common.Address) (common.Address, error) {
	return _Validators.Contract.GetMembershipInLastEpochFromSigner(&_Validators.CallOpts, signer)
}

// GetMembershipInLastEpochFromSigner is a free data retrieval call binding the contract method 0x51b52225.
//
// Solidity: function getMembershipInLastEpochFromSigner(address signer) constant returns(address)
func (_Validators *ValidatorsCallerSession) GetMembershipInLastEpochFromSigner(signer common.Address) (common.Address, error) {
	return _Validators.Contract.GetMembershipInLastEpochFromSigner(&_Validators.CallOpts, signer)
}

// GetNumRegisteredValidators is a free data retrieval call binding the contract method 0x517f6d33.
//
// Solidity: function getNumRegisteredValidators() constant returns(uint256)
func (_Validators *ValidatorsCaller) GetNumRegisteredValidators(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getNumRegisteredValidators")
	return *ret0, err
}

// GetNumRegisteredValidators is a free data retrieval call binding the contract method 0x517f6d33.
//
// Solidity: function getNumRegisteredValidators() constant returns(uint256)
func (_Validators *ValidatorsSession) GetNumRegisteredValidators() (*big.Int, error) {
	return _Validators.Contract.GetNumRegisteredValidators(&_Validators.CallOpts)
}

// GetNumRegisteredValidators is a free data retrieval call binding the contract method 0x517f6d33.
//
// Solidity: function getNumRegisteredValidators() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetNumRegisteredValidators() (*big.Int, error) {
	return _Validators.Contract.GetNumRegisteredValidators(&_Validators.CallOpts)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) constant returns(bytes32)
func (_Validators *ValidatorsCaller) GetParentSealBitmap(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getParentSealBitmap", blockNumber)
	return *ret0, err
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) constant returns(bytes32)
func (_Validators *ValidatorsSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _Validators.Contract.GetParentSealBitmap(&_Validators.CallOpts, blockNumber)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) constant returns(bytes32)
func (_Validators *ValidatorsCallerSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _Validators.Contract.GetParentSealBitmap(&_Validators.CallOpts, blockNumber)
}

// GetRegisteredValidatorGroups is a free data retrieval call binding the contract method 0x3f270898.
//
// Solidity: function getRegisteredValidatorGroups() constant returns(address[])
func (_Validators *ValidatorsCaller) GetRegisteredValidatorGroups(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getRegisteredValidatorGroups")
	return *ret0, err
}

// GetRegisteredValidatorGroups is a free data retrieval call binding the contract method 0x3f270898.
//
// Solidity: function getRegisteredValidatorGroups() constant returns(address[])
func (_Validators *ValidatorsSession) GetRegisteredValidatorGroups() ([]common.Address, error) {
	return _Validators.Contract.GetRegisteredValidatorGroups(&_Validators.CallOpts)
}

// GetRegisteredValidatorGroups is a free data retrieval call binding the contract method 0x3f270898.
//
// Solidity: function getRegisteredValidatorGroups() constant returns(address[])
func (_Validators *ValidatorsCallerSession) GetRegisteredValidatorGroups() ([]common.Address, error) {
	return _Validators.Contract.GetRegisteredValidatorGroups(&_Validators.CallOpts)
}

// GetRegisteredValidatorSigners is a free data retrieval call binding the contract method 0xd55dcbcf.
//
// Solidity: function getRegisteredValidatorSigners() constant returns(address[])
func (_Validators *ValidatorsCaller) GetRegisteredValidatorSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getRegisteredValidatorSigners")
	return *ret0, err
}

// GetRegisteredValidatorSigners is a free data retrieval call binding the contract method 0xd55dcbcf.
//
// Solidity: function getRegisteredValidatorSigners() constant returns(address[])
func (_Validators *ValidatorsSession) GetRegisteredValidatorSigners() ([]common.Address, error) {
	return _Validators.Contract.GetRegisteredValidatorSigners(&_Validators.CallOpts)
}

// GetRegisteredValidatorSigners is a free data retrieval call binding the contract method 0xd55dcbcf.
//
// Solidity: function getRegisteredValidatorSigners() constant returns(address[])
func (_Validators *ValidatorsCallerSession) GetRegisteredValidatorSigners() ([]common.Address, error) {
	return _Validators.Contract.GetRegisteredValidatorSigners(&_Validators.CallOpts)
}

// GetRegisteredValidators is a free data retrieval call binding the contract method 0xd93ab5ad.
//
// Solidity: function getRegisteredValidators() constant returns(address[])
func (_Validators *ValidatorsCaller) GetRegisteredValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getRegisteredValidators")
	return *ret0, err
}

// GetRegisteredValidators is a free data retrieval call binding the contract method 0xd93ab5ad.
//
// Solidity: function getRegisteredValidators() constant returns(address[])
func (_Validators *ValidatorsSession) GetRegisteredValidators() ([]common.Address, error) {
	return _Validators.Contract.GetRegisteredValidators(&_Validators.CallOpts)
}

// GetRegisteredValidators is a free data retrieval call binding the contract method 0xd93ab5ad.
//
// Solidity: function getRegisteredValidators() constant returns(address[])
func (_Validators *ValidatorsCallerSession) GetRegisteredValidators() ([]common.Address, error) {
	return _Validators.Contract.GetRegisteredValidators(&_Validators.CallOpts)
}

// GetTopGroupValidators is a free data retrieval call binding the contract method 0x8dd31e39.
//
// Solidity: function getTopGroupValidators(address account, uint256 n) constant returns(address[])
func (_Validators *ValidatorsCaller) GetTopGroupValidators(opts *bind.CallOpts, account common.Address, n *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getTopGroupValidators", account, n)
	return *ret0, err
}

// GetTopGroupValidators is a free data retrieval call binding the contract method 0x8dd31e39.
//
// Solidity: function getTopGroupValidators(address account, uint256 n) constant returns(address[])
func (_Validators *ValidatorsSession) GetTopGroupValidators(account common.Address, n *big.Int) ([]common.Address, error) {
	return _Validators.Contract.GetTopGroupValidators(&_Validators.CallOpts, account, n)
}

// GetTopGroupValidators is a free data retrieval call binding the contract method 0x8dd31e39.
//
// Solidity: function getTopGroupValidators(address account, uint256 n) constant returns(address[])
func (_Validators *ValidatorsCallerSession) GetTopGroupValidators(account common.Address, n *big.Int) ([]common.Address, error) {
	return _Validators.Contract.GetTopGroupValidators(&_Validators.CallOpts, account, n)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address account) constant returns(bytes ecdsaPublicKey, bytes blsPublicKey, address affiliation, uint256 score, address signer)
func (_Validators *ValidatorsCaller) GetValidator(opts *bind.CallOpts, account common.Address) (struct {
	EcdsaPublicKey []byte
	BlsPublicKey   []byte
	Affiliation    common.Address
	Score          *big.Int
	Signer         common.Address
}, error) {
	ret := new(struct {
		EcdsaPublicKey []byte
		BlsPublicKey   []byte
		Affiliation    common.Address
		Score          *big.Int
		Signer         common.Address
	})
	out := ret
	err := _Validators.contract.Call(opts, out, "getValidator", account)
	return *ret, err
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address account) constant returns(bytes ecdsaPublicKey, bytes blsPublicKey, address affiliation, uint256 score, address signer)
func (_Validators *ValidatorsSession) GetValidator(account common.Address) (struct {
	EcdsaPublicKey []byte
	BlsPublicKey   []byte
	Affiliation    common.Address
	Score          *big.Int
	Signer         common.Address
}, error) {
	return _Validators.Contract.GetValidator(&_Validators.CallOpts, account)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address account) constant returns(bytes ecdsaPublicKey, bytes blsPublicKey, address affiliation, uint256 score, address signer)
func (_Validators *ValidatorsCallerSession) GetValidator(account common.Address) (struct {
	EcdsaPublicKey []byte
	BlsPublicKey   []byte
	Affiliation    common.Address
	Score          *big.Int
	Signer         common.Address
}, error) {
	return _Validators.Contract.GetValidator(&_Validators.CallOpts, account)
}

// GetValidatorBlsPublicKeyFromSigner is a free data retrieval call binding the contract method 0xb730a299.
//
// Solidity: function getValidatorBlsPublicKeyFromSigner(address signer) constant returns(bytes blsPublicKey)
func (_Validators *ValidatorsCaller) GetValidatorBlsPublicKeyFromSigner(opts *bind.CallOpts, signer common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getValidatorBlsPublicKeyFromSigner", signer)
	return *ret0, err
}

// GetValidatorBlsPublicKeyFromSigner is a free data retrieval call binding the contract method 0xb730a299.
//
// Solidity: function getValidatorBlsPublicKeyFromSigner(address signer) constant returns(bytes blsPublicKey)
func (_Validators *ValidatorsSession) GetValidatorBlsPublicKeyFromSigner(signer common.Address) ([]byte, error) {
	return _Validators.Contract.GetValidatorBlsPublicKeyFromSigner(&_Validators.CallOpts, signer)
}

// GetValidatorBlsPublicKeyFromSigner is a free data retrieval call binding the contract method 0xb730a299.
//
// Solidity: function getValidatorBlsPublicKeyFromSigner(address signer) constant returns(bytes blsPublicKey)
func (_Validators *ValidatorsCallerSession) GetValidatorBlsPublicKeyFromSigner(signer common.Address) ([]byte, error) {
	return _Validators.Contract.GetValidatorBlsPublicKeyFromSigner(&_Validators.CallOpts, signer)
}

// GetValidatorGroup is a free data retrieval call binding the contract method 0x9b9d5161.
//
// Solidity: function getValidatorGroup(address account) constant returns(address[], uint256, uint256, uint256, uint256[], uint256, uint256)
func (_Validators *ValidatorsCaller) GetValidatorGroup(opts *bind.CallOpts, account common.Address) ([]common.Address, *big.Int, *big.Int, *big.Int, []*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new([]*big.Int)
		ret5 = new(*big.Int)
		ret6 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _Validators.contract.Call(opts, out, "getValidatorGroup", account)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetValidatorGroup is a free data retrieval call binding the contract method 0x9b9d5161.
//
// Solidity: function getValidatorGroup(address account) constant returns(address[], uint256, uint256, uint256, uint256[], uint256, uint256)
func (_Validators *ValidatorsSession) GetValidatorGroup(account common.Address) ([]common.Address, *big.Int, *big.Int, *big.Int, []*big.Int, *big.Int, *big.Int, error) {
	return _Validators.Contract.GetValidatorGroup(&_Validators.CallOpts, account)
}

// GetValidatorGroup is a free data retrieval call binding the contract method 0x9b9d5161.
//
// Solidity: function getValidatorGroup(address account) constant returns(address[], uint256, uint256, uint256, uint256[], uint256, uint256)
func (_Validators *ValidatorsCallerSession) GetValidatorGroup(account common.Address) ([]common.Address, *big.Int, *big.Int, *big.Int, []*big.Int, *big.Int, *big.Int, error) {
	return _Validators.Contract.GetValidatorGroup(&_Validators.CallOpts, account)
}

// GetValidatorGroupSlashingMultiplier is a free data retrieval call binding the contract method 0xdba94fcd.
//
// Solidity: function getValidatorGroupSlashingMultiplier(address account) constant returns(uint256)
func (_Validators *ValidatorsCaller) GetValidatorGroupSlashingMultiplier(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getValidatorGroupSlashingMultiplier", account)
	return *ret0, err
}

// GetValidatorGroupSlashingMultiplier is a free data retrieval call binding the contract method 0xdba94fcd.
//
// Solidity: function getValidatorGroupSlashingMultiplier(address account) constant returns(uint256)
func (_Validators *ValidatorsSession) GetValidatorGroupSlashingMultiplier(account common.Address) (*big.Int, error) {
	return _Validators.Contract.GetValidatorGroupSlashingMultiplier(&_Validators.CallOpts, account)
}

// GetValidatorGroupSlashingMultiplier is a free data retrieval call binding the contract method 0xdba94fcd.
//
// Solidity: function getValidatorGroupSlashingMultiplier(address account) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetValidatorGroupSlashingMultiplier(account common.Address) (*big.Int, error) {
	return _Validators.Contract.GetValidatorGroupSlashingMultiplier(&_Validators.CallOpts, account)
}

// GetValidatorLockedGoldRequirements is a free data retrieval call binding the contract method 0xc10c96ef.
//
// Solidity: function getValidatorLockedGoldRequirements() constant returns(uint256, uint256)
func (_Validators *ValidatorsCaller) GetValidatorLockedGoldRequirements(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Validators.contract.Call(opts, out, "getValidatorLockedGoldRequirements")
	return *ret0, *ret1, err
}

// GetValidatorLockedGoldRequirements is a free data retrieval call binding the contract method 0xc10c96ef.
//
// Solidity: function getValidatorLockedGoldRequirements() constant returns(uint256, uint256)
func (_Validators *ValidatorsSession) GetValidatorLockedGoldRequirements() (*big.Int, *big.Int, error) {
	return _Validators.Contract.GetValidatorLockedGoldRequirements(&_Validators.CallOpts)
}

// GetValidatorLockedGoldRequirements is a free data retrieval call binding the contract method 0xc10c96ef.
//
// Solidity: function getValidatorLockedGoldRequirements() constant returns(uint256, uint256)
func (_Validators *ValidatorsCallerSession) GetValidatorLockedGoldRequirements() (*big.Int, *big.Int, error) {
	return _Validators.Contract.GetValidatorLockedGoldRequirements(&_Validators.CallOpts)
}

// GetValidatorScoreParameters is a free data retrieval call binding the contract method 0x19113e3b.
//
// Solidity: function getValidatorScoreParameters() constant returns(uint256, uint256)
func (_Validators *ValidatorsCaller) GetValidatorScoreParameters(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Validators.contract.Call(opts, out, "getValidatorScoreParameters")
	return *ret0, *ret1, err
}

// GetValidatorScoreParameters is a free data retrieval call binding the contract method 0x19113e3b.
//
// Solidity: function getValidatorScoreParameters() constant returns(uint256, uint256)
func (_Validators *ValidatorsSession) GetValidatorScoreParameters() (*big.Int, *big.Int, error) {
	return _Validators.Contract.GetValidatorScoreParameters(&_Validators.CallOpts)
}

// GetValidatorScoreParameters is a free data retrieval call binding the contract method 0x19113e3b.
//
// Solidity: function getValidatorScoreParameters() constant returns(uint256, uint256)
func (_Validators *ValidatorsCallerSession) GetValidatorScoreParameters() (*big.Int, *big.Int, error) {
	return _Validators.Contract.GetValidatorScoreParameters(&_Validators.CallOpts)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) constant returns(bytes32)
func (_Validators *ValidatorsCaller) GetVerifiedSealBitmapFromHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getVerifiedSealBitmapFromHeader", header)
	return *ret0, err
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) constant returns(bytes32)
func (_Validators *ValidatorsSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _Validators.Contract.GetVerifiedSealBitmapFromHeader(&_Validators.CallOpts, header)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) constant returns(bytes32)
func (_Validators *ValidatorsCallerSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _Validators.Contract.GetVerifiedSealBitmapFromHeader(&_Validators.CallOpts, header)
}

// GroupLockedGoldRequirements is a free data retrieval call binding the contract method 0xc5805140.
//
// Solidity: function groupLockedGoldRequirements() constant returns(uint256 value, uint256 duration)
func (_Validators *ValidatorsCaller) GroupLockedGoldRequirements(opts *bind.CallOpts) (struct {
	Value    *big.Int
	Duration *big.Int
}, error) {
	ret := new(struct {
		Value    *big.Int
		Duration *big.Int
	})
	out := ret
	err := _Validators.contract.Call(opts, out, "groupLockedGoldRequirements")
	return *ret, err
}

// GroupLockedGoldRequirements is a free data retrieval call binding the contract method 0xc5805140.
//
// Solidity: function groupLockedGoldRequirements() constant returns(uint256 value, uint256 duration)
func (_Validators *ValidatorsSession) GroupLockedGoldRequirements() (struct {
	Value    *big.Int
	Duration *big.Int
}, error) {
	return _Validators.Contract.GroupLockedGoldRequirements(&_Validators.CallOpts)
}

// GroupLockedGoldRequirements is a free data retrieval call binding the contract method 0xc5805140.
//
// Solidity: function groupLockedGoldRequirements() constant returns(uint256 value, uint256 duration)
func (_Validators *ValidatorsCallerSession) GroupLockedGoldRequirements() (struct {
	Value    *big.Int
	Duration *big.Int
}, error) {
	return _Validators.Contract.GroupLockedGoldRequirements(&_Validators.CallOpts)
}

// GroupMembershipInEpoch is a free data retrieval call binding the contract method 0xeb1d0b42.
//
// Solidity: function groupMembershipInEpoch(address account, uint256 epochNumber, uint256 index) constant returns(address)
func (_Validators *ValidatorsCaller) GroupMembershipInEpoch(opts *bind.CallOpts, account common.Address, epochNumber *big.Int, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "groupMembershipInEpoch", account, epochNumber, index)
	return *ret0, err
}

// GroupMembershipInEpoch is a free data retrieval call binding the contract method 0xeb1d0b42.
//
// Solidity: function groupMembershipInEpoch(address account, uint256 epochNumber, uint256 index) constant returns(address)
func (_Validators *ValidatorsSession) GroupMembershipInEpoch(account common.Address, epochNumber *big.Int, index *big.Int) (common.Address, error) {
	return _Validators.Contract.GroupMembershipInEpoch(&_Validators.CallOpts, account, epochNumber, index)
}

// GroupMembershipInEpoch is a free data retrieval call binding the contract method 0xeb1d0b42.
//
// Solidity: function groupMembershipInEpoch(address account, uint256 epochNumber, uint256 index) constant returns(address)
func (_Validators *ValidatorsCallerSession) GroupMembershipInEpoch(account common.Address, epochNumber *big.Int, index *big.Int) (common.Address, error) {
	return _Validators.Contract.GroupMembershipInEpoch(&_Validators.CallOpts, account, epochNumber, index)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) constant returns(bytes32)
func (_Validators *ValidatorsCaller) HashHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "hashHeader", header)
	return *ret0, err
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) constant returns(bytes32)
func (_Validators *ValidatorsSession) HashHeader(header []byte) ([32]byte, error) {
	return _Validators.Contract.HashHeader(&_Validators.CallOpts, header)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) constant returns(bytes32)
func (_Validators *ValidatorsCallerSession) HashHeader(header []byte) ([32]byte, error) {
	return _Validators.Contract.HashHeader(&_Validators.CallOpts, header)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Validators *ValidatorsCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Validators *ValidatorsSession) Initialized() (bool, error) {
	return _Validators.Contract.Initialized(&_Validators.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Validators *ValidatorsCallerSession) Initialized() (bool, error) {
	return _Validators.Contract.Initialized(&_Validators.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Validators *ValidatorsCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Validators *ValidatorsSession) IsOwner() (bool, error) {
	return _Validators.Contract.IsOwner(&_Validators.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Validators *ValidatorsCallerSession) IsOwner() (bool, error) {
	return _Validators.Contract.IsOwner(&_Validators.CallOpts)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address account) constant returns(bool)
func (_Validators *ValidatorsCaller) IsValidator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "isValidator", account)
	return *ret0, err
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address account) constant returns(bool)
func (_Validators *ValidatorsSession) IsValidator(account common.Address) (bool, error) {
	return _Validators.Contract.IsValidator(&_Validators.CallOpts, account)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address account) constant returns(bool)
func (_Validators *ValidatorsCallerSession) IsValidator(account common.Address) (bool, error) {
	return _Validators.Contract.IsValidator(&_Validators.CallOpts, account)
}

// IsValidatorGroup is a free data retrieval call binding the contract method 0x52f13a4e.
//
// Solidity: function isValidatorGroup(address account) constant returns(bool)
func (_Validators *ValidatorsCaller) IsValidatorGroup(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "isValidatorGroup", account)
	return *ret0, err
}

// IsValidatorGroup is a free data retrieval call binding the contract method 0x52f13a4e.
//
// Solidity: function isValidatorGroup(address account) constant returns(bool)
func (_Validators *ValidatorsSession) IsValidatorGroup(account common.Address) (bool, error) {
	return _Validators.Contract.IsValidatorGroup(&_Validators.CallOpts, account)
}

// IsValidatorGroup is a free data retrieval call binding the contract method 0x52f13a4e.
//
// Solidity: function isValidatorGroup(address account) constant returns(bool)
func (_Validators *ValidatorsCallerSession) IsValidatorGroup(account common.Address) (bool, error) {
	return _Validators.Contract.IsValidatorGroup(&_Validators.CallOpts, account)
}

// MaxGroupSize is a free data retrieval call binding the contract method 0x5779e93d.
//
// Solidity: function maxGroupSize() constant returns(uint256)
func (_Validators *ValidatorsCaller) MaxGroupSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "maxGroupSize")
	return *ret0, err
}

// MaxGroupSize is a free data retrieval call binding the contract method 0x5779e93d.
//
// Solidity: function maxGroupSize() constant returns(uint256)
func (_Validators *ValidatorsSession) MaxGroupSize() (*big.Int, error) {
	return _Validators.Contract.MaxGroupSize(&_Validators.CallOpts)
}

// MaxGroupSize is a free data retrieval call binding the contract method 0x5779e93d.
//
// Solidity: function maxGroupSize() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) MaxGroupSize() (*big.Int, error) {
	return _Validators.Contract.MaxGroupSize(&_Validators.CallOpts)
}

// MeetsAccountLockedGoldRequirements is a free data retrieval call binding the contract method 0xc54c1cd4.
//
// Solidity: function meetsAccountLockedGoldRequirements(address account) constant returns(bool)
func (_Validators *ValidatorsCaller) MeetsAccountLockedGoldRequirements(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "meetsAccountLockedGoldRequirements", account)
	return *ret0, err
}

// MeetsAccountLockedGoldRequirements is a free data retrieval call binding the contract method 0xc54c1cd4.
//
// Solidity: function meetsAccountLockedGoldRequirements(address account) constant returns(bool)
func (_Validators *ValidatorsSession) MeetsAccountLockedGoldRequirements(account common.Address) (bool, error) {
	return _Validators.Contract.MeetsAccountLockedGoldRequirements(&_Validators.CallOpts, account)
}

// MeetsAccountLockedGoldRequirements is a free data retrieval call binding the contract method 0xc54c1cd4.
//
// Solidity: function meetsAccountLockedGoldRequirements(address account) constant returns(bool)
func (_Validators *ValidatorsCallerSession) MeetsAccountLockedGoldRequirements(account common.Address) (bool, error) {
	return _Validators.Contract.MeetsAccountLockedGoldRequirements(&_Validators.CallOpts, account)
}

// MembershipHistoryLength is a free data retrieval call binding the contract method 0x4cd76db4.
//
// Solidity: function membershipHistoryLength() constant returns(uint256)
func (_Validators *ValidatorsCaller) MembershipHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "membershipHistoryLength")
	return *ret0, err
}

// MembershipHistoryLength is a free data retrieval call binding the contract method 0x4cd76db4.
//
// Solidity: function membershipHistoryLength() constant returns(uint256)
func (_Validators *ValidatorsSession) MembershipHistoryLength() (*big.Int, error) {
	return _Validators.Contract.MembershipHistoryLength(&_Validators.CallOpts)
}

// MembershipHistoryLength is a free data retrieval call binding the contract method 0x4cd76db4.
//
// Solidity: function membershipHistoryLength() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) MembershipHistoryLength() (*big.Int, error) {
	return _Validators.Contract.MembershipHistoryLength(&_Validators.CallOpts)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsCaller) MinQuorumSize(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "minQuorumSize", blockNumber)
	return *ret0, err
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _Validators.Contract.MinQuorumSize(&_Validators.CallOpts, blockNumber)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _Validators.Contract.MinQuorumSize(&_Validators.CallOpts, blockNumber)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() constant returns(uint256)
func (_Validators *ValidatorsCaller) MinQuorumSizeInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "minQuorumSizeInCurrentSet")
	return *ret0, err
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() constant returns(uint256)
func (_Validators *ValidatorsSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _Validators.Contract.MinQuorumSizeInCurrentSet(&_Validators.CallOpts)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _Validators.Contract.MinQuorumSizeInCurrentSet(&_Validators.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() constant returns(uint256)
func (_Validators *ValidatorsCaller) NumberValidatorsInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "numberValidatorsInCurrentSet")
	return *ret0, err
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() constant returns(uint256)
func (_Validators *ValidatorsSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _Validators.Contract.NumberValidatorsInCurrentSet(&_Validators.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _Validators.Contract.NumberValidatorsInCurrentSet(&_Validators.CallOpts)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsCaller) NumberValidatorsInSet(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "numberValidatorsInSet", blockNumber)
	return *ret0, err
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _Validators.Contract.NumberValidatorsInSet(&_Validators.CallOpts, blockNumber)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _Validators.Contract.NumberValidatorsInSet(&_Validators.CallOpts, blockNumber)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Validators *ValidatorsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Validators *ValidatorsSession) Owner() (common.Address, error) {
	return _Validators.Contract.Owner(&_Validators.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Validators *ValidatorsCallerSession) Owner() (common.Address, error) {
	return _Validators.Contract.Owner(&_Validators.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Validators *ValidatorsCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Validators *ValidatorsSession) Registry() (common.Address, error) {
	return _Validators.Contract.Registry(&_Validators.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Validators *ValidatorsCallerSession) Registry() (common.Address, error) {
	return _Validators.Contract.Registry(&_Validators.CallOpts)
}

// SlashingMultiplierResetPeriod is a free data retrieval call binding the contract method 0x36407b70.
//
// Solidity: function slashingMultiplierResetPeriod() constant returns(uint256)
func (_Validators *ValidatorsCaller) SlashingMultiplierResetPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "slashingMultiplierResetPeriod")
	return *ret0, err
}

// SlashingMultiplierResetPeriod is a free data retrieval call binding the contract method 0x36407b70.
//
// Solidity: function slashingMultiplierResetPeriod() constant returns(uint256)
func (_Validators *ValidatorsSession) SlashingMultiplierResetPeriod() (*big.Int, error) {
	return _Validators.Contract.SlashingMultiplierResetPeriod(&_Validators.CallOpts)
}

// SlashingMultiplierResetPeriod is a free data retrieval call binding the contract method 0x36407b70.
//
// Solidity: function slashingMultiplierResetPeriod() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) SlashingMultiplierResetPeriod() (*big.Int, error) {
	return _Validators.Contract.SlashingMultiplierResetPeriod(&_Validators.CallOpts)
}

// ValidatorLockedGoldRequirements is a free data retrieval call binding the contract method 0xbd9e9d94.
//
// Solidity: function validatorLockedGoldRequirements() constant returns(uint256 value, uint256 duration)
func (_Validators *ValidatorsCaller) ValidatorLockedGoldRequirements(opts *bind.CallOpts) (struct {
	Value    *big.Int
	Duration *big.Int
}, error) {
	ret := new(struct {
		Value    *big.Int
		Duration *big.Int
	})
	out := ret
	err := _Validators.contract.Call(opts, out, "validatorLockedGoldRequirements")
	return *ret, err
}

// ValidatorLockedGoldRequirements is a free data retrieval call binding the contract method 0xbd9e9d94.
//
// Solidity: function validatorLockedGoldRequirements() constant returns(uint256 value, uint256 duration)
func (_Validators *ValidatorsSession) ValidatorLockedGoldRequirements() (struct {
	Value    *big.Int
	Duration *big.Int
}, error) {
	return _Validators.Contract.ValidatorLockedGoldRequirements(&_Validators.CallOpts)
}

// ValidatorLockedGoldRequirements is a free data retrieval call binding the contract method 0xbd9e9d94.
//
// Solidity: function validatorLockedGoldRequirements() constant returns(uint256 value, uint256 duration)
func (_Validators *ValidatorsCallerSession) ValidatorLockedGoldRequirements() (struct {
	Value    *big.Int
	Duration *big.Int
}, error) {
	return _Validators.Contract.ValidatorLockedGoldRequirements(&_Validators.CallOpts)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) constant returns(address)
func (_Validators *ValidatorsCaller) ValidatorSignerAddressFromCurrentSet(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "validatorSignerAddressFromCurrentSet", index)
	return *ret0, err
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) constant returns(address)
func (_Validators *ValidatorsSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _Validators.Contract.ValidatorSignerAddressFromCurrentSet(&_Validators.CallOpts, index)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) constant returns(address)
func (_Validators *ValidatorsCallerSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _Validators.Contract.ValidatorSignerAddressFromCurrentSet(&_Validators.CallOpts, index)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) constant returns(address)
func (_Validators *ValidatorsCaller) ValidatorSignerAddressFromSet(opts *bind.CallOpts, index *big.Int, blockNumber *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "validatorSignerAddressFromSet", index, blockNumber)
	return *ret0, err
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) constant returns(address)
func (_Validators *ValidatorsSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _Validators.Contract.ValidatorSignerAddressFromSet(&_Validators.CallOpts, index, blockNumber)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) constant returns(address)
func (_Validators *ValidatorsCallerSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _Validators.Contract.ValidatorSignerAddressFromSet(&_Validators.CallOpts, index, blockNumber)
}

// AddFirstMember is a paid mutator transaction binding the contract method 0x3173b8db.
//
// Solidity: function addFirstMember(address validator, address lesser, address greater) returns(bool)
func (_Validators *ValidatorsTransactor) AddFirstMember(opts *bind.TransactOpts, validator common.Address, lesser common.Address, greater common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "addFirstMember", validator, lesser, greater)
}

// AddFirstMember is a paid mutator transaction binding the contract method 0x3173b8db.
//
// Solidity: function addFirstMember(address validator, address lesser, address greater) returns(bool)
func (_Validators *ValidatorsSession) AddFirstMember(validator common.Address, lesser common.Address, greater common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddFirstMember(&_Validators.TransactOpts, validator, lesser, greater)
}

// AddFirstMember is a paid mutator transaction binding the contract method 0x3173b8db.
//
// Solidity: function addFirstMember(address validator, address lesser, address greater) returns(bool)
func (_Validators *ValidatorsTransactorSession) AddFirstMember(validator common.Address, lesser common.Address, greater common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddFirstMember(&_Validators.TransactOpts, validator, lesser, greater)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address validator) returns(bool)
func (_Validators *ValidatorsTransactor) AddMember(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "addMember", validator)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address validator) returns(bool)
func (_Validators *ValidatorsSession) AddMember(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddMember(&_Validators.TransactOpts, validator)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address validator) returns(bool)
func (_Validators *ValidatorsTransactorSession) AddMember(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddMember(&_Validators.TransactOpts, validator)
}

// Affiliate is a paid mutator transaction binding the contract method 0xb591d3a5.
//
// Solidity: function affiliate(address group) returns(bool)
func (_Validators *ValidatorsTransactor) Affiliate(opts *bind.TransactOpts, group common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "affiliate", group)
}

// Affiliate is a paid mutator transaction binding the contract method 0xb591d3a5.
//
// Solidity: function affiliate(address group) returns(bool)
func (_Validators *ValidatorsSession) Affiliate(group common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Affiliate(&_Validators.TransactOpts, group)
}

// Affiliate is a paid mutator transaction binding the contract method 0xb591d3a5.
//
// Solidity: function affiliate(address group) returns(bool)
func (_Validators *ValidatorsTransactorSession) Affiliate(group common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Affiliate(&_Validators.TransactOpts, group)
}

// Deaffiliate is a paid mutator transaction binding the contract method 0xfffdfccb.
//
// Solidity: function deaffiliate() returns(bool)
func (_Validators *ValidatorsTransactor) Deaffiliate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "deaffiliate")
}

// Deaffiliate is a paid mutator transaction binding the contract method 0xfffdfccb.
//
// Solidity: function deaffiliate() returns(bool)
func (_Validators *ValidatorsSession) Deaffiliate() (*types.Transaction, error) {
	return _Validators.Contract.Deaffiliate(&_Validators.TransactOpts)
}

// Deaffiliate is a paid mutator transaction binding the contract method 0xfffdfccb.
//
// Solidity: function deaffiliate() returns(bool)
func (_Validators *ValidatorsTransactorSession) Deaffiliate() (*types.Transaction, error) {
	return _Validators.Contract.Deaffiliate(&_Validators.TransactOpts)
}

// DeregisterValidator is a paid mutator transaction binding the contract method 0x8b16b1c6.
//
// Solidity: function deregisterValidator(uint256 index) returns(bool)
func (_Validators *ValidatorsTransactor) DeregisterValidator(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "deregisterValidator", index)
}

// DeregisterValidator is a paid mutator transaction binding the contract method 0x8b16b1c6.
//
// Solidity: function deregisterValidator(uint256 index) returns(bool)
func (_Validators *ValidatorsSession) DeregisterValidator(index *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.DeregisterValidator(&_Validators.TransactOpts, index)
}

// DeregisterValidator is a paid mutator transaction binding the contract method 0x8b16b1c6.
//
// Solidity: function deregisterValidator(uint256 index) returns(bool)
func (_Validators *ValidatorsTransactorSession) DeregisterValidator(index *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.DeregisterValidator(&_Validators.TransactOpts, index)
}

// DeregisterValidatorGroup is a paid mutator transaction binding the contract method 0x60fb822c.
//
// Solidity: function deregisterValidatorGroup(uint256 index) returns(bool)
func (_Validators *ValidatorsTransactor) DeregisterValidatorGroup(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "deregisterValidatorGroup", index)
}

// DeregisterValidatorGroup is a paid mutator transaction binding the contract method 0x60fb822c.
//
// Solidity: function deregisterValidatorGroup(uint256 index) returns(bool)
func (_Validators *ValidatorsSession) DeregisterValidatorGroup(index *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.DeregisterValidatorGroup(&_Validators.TransactOpts, index)
}

// DeregisterValidatorGroup is a paid mutator transaction binding the contract method 0x60fb822c.
//
// Solidity: function deregisterValidatorGroup(uint256 index) returns(bool)
func (_Validators *ValidatorsTransactorSession) DeregisterValidatorGroup(index *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.DeregisterValidatorGroup(&_Validators.TransactOpts, index)
}

// DistributeEpochPaymentsFromSigner is a paid mutator transaction binding the contract method 0xd69ef6cf.
//
// Solidity: function distributeEpochPaymentsFromSigner(address signer, uint256 maxPayment) returns(uint256)
func (_Validators *ValidatorsTransactor) DistributeEpochPaymentsFromSigner(opts *bind.TransactOpts, signer common.Address, maxPayment *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "distributeEpochPaymentsFromSigner", signer, maxPayment)
}

// DistributeEpochPaymentsFromSigner is a paid mutator transaction binding the contract method 0xd69ef6cf.
//
// Solidity: function distributeEpochPaymentsFromSigner(address signer, uint256 maxPayment) returns(uint256)
func (_Validators *ValidatorsSession) DistributeEpochPaymentsFromSigner(signer common.Address, maxPayment *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.DistributeEpochPaymentsFromSigner(&_Validators.TransactOpts, signer, maxPayment)
}

// DistributeEpochPaymentsFromSigner is a paid mutator transaction binding the contract method 0xd69ef6cf.
//
// Solidity: function distributeEpochPaymentsFromSigner(address signer, uint256 maxPayment) returns(uint256)
func (_Validators *ValidatorsTransactorSession) DistributeEpochPaymentsFromSigner(signer common.Address, maxPayment *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.DistributeEpochPaymentsFromSigner(&_Validators.TransactOpts, signer, maxPayment)
}

// ForceDeaffiliateIfValidator is a paid mutator transaction binding the contract method 0xe33301aa.
//
// Solidity: function forceDeaffiliateIfValidator(address validatorAccount) returns()
func (_Validators *ValidatorsTransactor) ForceDeaffiliateIfValidator(opts *bind.TransactOpts, validatorAccount common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "forceDeaffiliateIfValidator", validatorAccount)
}

// ForceDeaffiliateIfValidator is a paid mutator transaction binding the contract method 0xe33301aa.
//
// Solidity: function forceDeaffiliateIfValidator(address validatorAccount) returns()
func (_Validators *ValidatorsSession) ForceDeaffiliateIfValidator(validatorAccount common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ForceDeaffiliateIfValidator(&_Validators.TransactOpts, validatorAccount)
}

// ForceDeaffiliateIfValidator is a paid mutator transaction binding the contract method 0xe33301aa.
//
// Solidity: function forceDeaffiliateIfValidator(address validatorAccount) returns()
func (_Validators *ValidatorsTransactorSession) ForceDeaffiliateIfValidator(validatorAccount common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ForceDeaffiliateIfValidator(&_Validators.TransactOpts, validatorAccount)
}

// HalveSlashingMultiplier is a paid mutator transaction binding the contract method 0xc22d3bba.
//
// Solidity: function halveSlashingMultiplier(address account) returns()
func (_Validators *ValidatorsTransactor) HalveSlashingMultiplier(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "halveSlashingMultiplier", account)
}

// HalveSlashingMultiplier is a paid mutator transaction binding the contract method 0xc22d3bba.
//
// Solidity: function halveSlashingMultiplier(address account) returns()
func (_Validators *ValidatorsSession) HalveSlashingMultiplier(account common.Address) (*types.Transaction, error) {
	return _Validators.Contract.HalveSlashingMultiplier(&_Validators.TransactOpts, account)
}

// HalveSlashingMultiplier is a paid mutator transaction binding the contract method 0xc22d3bba.
//
// Solidity: function halveSlashingMultiplier(address account) returns()
func (_Validators *ValidatorsTransactorSession) HalveSlashingMultiplier(account common.Address) (*types.Transaction, error) {
	return _Validators.Contract.HalveSlashingMultiplier(&_Validators.TransactOpts, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x78d25456.
//
// Solidity: function initialize(address registryAddress, uint256 groupRequirementValue, uint256 groupRequirementDuration, uint256 validatorRequirementValue, uint256 validatorRequirementDuration, uint256 validatorScoreExponent, uint256 validatorScoreAdjustmentSpeed, uint256 _membershipHistoryLength, uint256 _slashingMultiplierResetPeriod, uint256 _maxGroupSize, uint256 _commissionUpdateDelay) returns()
func (_Validators *ValidatorsTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address, groupRequirementValue *big.Int, groupRequirementDuration *big.Int, validatorRequirementValue *big.Int, validatorRequirementDuration *big.Int, validatorScoreExponent *big.Int, validatorScoreAdjustmentSpeed *big.Int, _membershipHistoryLength *big.Int, _slashingMultiplierResetPeriod *big.Int, _maxGroupSize *big.Int, _commissionUpdateDelay *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "initialize", registryAddress, groupRequirementValue, groupRequirementDuration, validatorRequirementValue, validatorRequirementDuration, validatorScoreExponent, validatorScoreAdjustmentSpeed, _membershipHistoryLength, _slashingMultiplierResetPeriod, _maxGroupSize, _commissionUpdateDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x78d25456.
//
// Solidity: function initialize(address registryAddress, uint256 groupRequirementValue, uint256 groupRequirementDuration, uint256 validatorRequirementValue, uint256 validatorRequirementDuration, uint256 validatorScoreExponent, uint256 validatorScoreAdjustmentSpeed, uint256 _membershipHistoryLength, uint256 _slashingMultiplierResetPeriod, uint256 _maxGroupSize, uint256 _commissionUpdateDelay) returns()
func (_Validators *ValidatorsSession) Initialize(registryAddress common.Address, groupRequirementValue *big.Int, groupRequirementDuration *big.Int, validatorRequirementValue *big.Int, validatorRequirementDuration *big.Int, validatorScoreExponent *big.Int, validatorScoreAdjustmentSpeed *big.Int, _membershipHistoryLength *big.Int, _slashingMultiplierResetPeriod *big.Int, _maxGroupSize *big.Int, _commissionUpdateDelay *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.Initialize(&_Validators.TransactOpts, registryAddress, groupRequirementValue, groupRequirementDuration, validatorRequirementValue, validatorRequirementDuration, validatorScoreExponent, validatorScoreAdjustmentSpeed, _membershipHistoryLength, _slashingMultiplierResetPeriod, _maxGroupSize, _commissionUpdateDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x78d25456.
//
// Solidity: function initialize(address registryAddress, uint256 groupRequirementValue, uint256 groupRequirementDuration, uint256 validatorRequirementValue, uint256 validatorRequirementDuration, uint256 validatorScoreExponent, uint256 validatorScoreAdjustmentSpeed, uint256 _membershipHistoryLength, uint256 _slashingMultiplierResetPeriod, uint256 _maxGroupSize, uint256 _commissionUpdateDelay) returns()
func (_Validators *ValidatorsTransactorSession) Initialize(registryAddress common.Address, groupRequirementValue *big.Int, groupRequirementDuration *big.Int, validatorRequirementValue *big.Int, validatorRequirementDuration *big.Int, validatorScoreExponent *big.Int, validatorScoreAdjustmentSpeed *big.Int, _membershipHistoryLength *big.Int, _slashingMultiplierResetPeriod *big.Int, _maxGroupSize *big.Int, _commissionUpdateDelay *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.Initialize(&_Validators.TransactOpts, registryAddress, groupRequirementValue, groupRequirementDuration, validatorRequirementValue, validatorRequirementDuration, validatorScoreExponent, validatorScoreAdjustmentSpeed, _membershipHistoryLength, _slashingMultiplierResetPeriod, _maxGroupSize, _commissionUpdateDelay)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xea684f77.
//
// Solidity: function registerValidator(bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsTransactor) RegisterValidator(opts *bind.TransactOpts, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "registerValidator", ecdsaPublicKey, blsPublicKey, blsPop)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xea684f77.
//
// Solidity: function registerValidator(bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsSession) RegisterValidator(ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.Contract.RegisterValidator(&_Validators.TransactOpts, ecdsaPublicKey, blsPublicKey, blsPop)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xea684f77.
//
// Solidity: function registerValidator(bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsTransactorSession) RegisterValidator(ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.Contract.RegisterValidator(&_Validators.TransactOpts, ecdsaPublicKey, blsPublicKey, blsPop)
}

// RegisterValidatorGroup is a paid mutator transaction binding the contract method 0xee098310.
//
// Solidity: function registerValidatorGroup(uint256 commission) returns(bool)
func (_Validators *ValidatorsTransactor) RegisterValidatorGroup(opts *bind.TransactOpts, commission *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "registerValidatorGroup", commission)
}

// RegisterValidatorGroup is a paid mutator transaction binding the contract method 0xee098310.
//
// Solidity: function registerValidatorGroup(uint256 commission) returns(bool)
func (_Validators *ValidatorsSession) RegisterValidatorGroup(commission *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.RegisterValidatorGroup(&_Validators.TransactOpts, commission)
}

// RegisterValidatorGroup is a paid mutator transaction binding the contract method 0xee098310.
//
// Solidity: function registerValidatorGroup(uint256 commission) returns(bool)
func (_Validators *ValidatorsTransactorSession) RegisterValidatorGroup(commission *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.RegisterValidatorGroup(&_Validators.TransactOpts, commission)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address validator) returns(bool)
func (_Validators *ValidatorsTransactor) RemoveMember(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "removeMember", validator)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address validator) returns(bool)
func (_Validators *ValidatorsSession) RemoveMember(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.RemoveMember(&_Validators.TransactOpts, validator)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address validator) returns(bool)
func (_Validators *ValidatorsTransactorSession) RemoveMember(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.RemoveMember(&_Validators.TransactOpts, validator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validators *ValidatorsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validators *ValidatorsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Validators.Contract.RenounceOwnership(&_Validators.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validators *ValidatorsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Validators.Contract.RenounceOwnership(&_Validators.TransactOpts)
}

// ReorderMember is a paid mutator transaction binding the contract method 0x988dcd1f.
//
// Solidity: function reorderMember(address validator, address lesserMember, address greaterMember) returns(bool)
func (_Validators *ValidatorsTransactor) ReorderMember(opts *bind.TransactOpts, validator common.Address, lesserMember common.Address, greaterMember common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "reorderMember", validator, lesserMember, greaterMember)
}

// ReorderMember is a paid mutator transaction binding the contract method 0x988dcd1f.
//
// Solidity: function reorderMember(address validator, address lesserMember, address greaterMember) returns(bool)
func (_Validators *ValidatorsSession) ReorderMember(validator common.Address, lesserMember common.Address, greaterMember common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ReorderMember(&_Validators.TransactOpts, validator, lesserMember, greaterMember)
}

// ReorderMember is a paid mutator transaction binding the contract method 0x988dcd1f.
//
// Solidity: function reorderMember(address validator, address lesserMember, address greaterMember) returns(bool)
func (_Validators *ValidatorsTransactorSession) ReorderMember(validator common.Address, lesserMember common.Address, greaterMember common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ReorderMember(&_Validators.TransactOpts, validator, lesserMember, greaterMember)
}

// ResetSlashingMultiplier is a paid mutator transaction binding the contract method 0xb8f93943.
//
// Solidity: function resetSlashingMultiplier() returns()
func (_Validators *ValidatorsTransactor) ResetSlashingMultiplier(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "resetSlashingMultiplier")
}

// ResetSlashingMultiplier is a paid mutator transaction binding the contract method 0xb8f93943.
//
// Solidity: function resetSlashingMultiplier() returns()
func (_Validators *ValidatorsSession) ResetSlashingMultiplier() (*types.Transaction, error) {
	return _Validators.Contract.ResetSlashingMultiplier(&_Validators.TransactOpts)
}

// ResetSlashingMultiplier is a paid mutator transaction binding the contract method 0xb8f93943.
//
// Solidity: function resetSlashingMultiplier() returns()
func (_Validators *ValidatorsTransactorSession) ResetSlashingMultiplier() (*types.Transaction, error) {
	return _Validators.Contract.ResetSlashingMultiplier(&_Validators.TransactOpts)
}

// SetCommissionUpdateDelay is a paid mutator transaction binding the contract method 0x6c620d90.
//
// Solidity: function setCommissionUpdateDelay(uint256 delay) returns()
func (_Validators *ValidatorsTransactor) SetCommissionUpdateDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setCommissionUpdateDelay", delay)
}

// SetCommissionUpdateDelay is a paid mutator transaction binding the contract method 0x6c620d90.
//
// Solidity: function setCommissionUpdateDelay(uint256 delay) returns()
func (_Validators *ValidatorsSession) SetCommissionUpdateDelay(delay *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetCommissionUpdateDelay(&_Validators.TransactOpts, delay)
}

// SetCommissionUpdateDelay is a paid mutator transaction binding the contract method 0x6c620d90.
//
// Solidity: function setCommissionUpdateDelay(uint256 delay) returns()
func (_Validators *ValidatorsTransactorSession) SetCommissionUpdateDelay(delay *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetCommissionUpdateDelay(&_Validators.TransactOpts, delay)
}

// SetGroupLockedGoldRequirements is a paid mutator transaction binding the contract method 0x5a61d15b.
//
// Solidity: function setGroupLockedGoldRequirements(uint256 value, uint256 duration) returns(bool)
func (_Validators *ValidatorsTransactor) SetGroupLockedGoldRequirements(opts *bind.TransactOpts, value *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setGroupLockedGoldRequirements", value, duration)
}

// SetGroupLockedGoldRequirements is a paid mutator transaction binding the contract method 0x5a61d15b.
//
// Solidity: function setGroupLockedGoldRequirements(uint256 value, uint256 duration) returns(bool)
func (_Validators *ValidatorsSession) SetGroupLockedGoldRequirements(value *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetGroupLockedGoldRequirements(&_Validators.TransactOpts, value, duration)
}

// SetGroupLockedGoldRequirements is a paid mutator transaction binding the contract method 0x5a61d15b.
//
// Solidity: function setGroupLockedGoldRequirements(uint256 value, uint256 duration) returns(bool)
func (_Validators *ValidatorsTransactorSession) SetGroupLockedGoldRequirements(value *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetGroupLockedGoldRequirements(&_Validators.TransactOpts, value, duration)
}

// SetMaxGroupSize is a paid mutator transaction binding the contract method 0xe1497ff7.
//
// Solidity: function setMaxGroupSize(uint256 size) returns(bool)
func (_Validators *ValidatorsTransactor) SetMaxGroupSize(opts *bind.TransactOpts, size *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setMaxGroupSize", size)
}

// SetMaxGroupSize is a paid mutator transaction binding the contract method 0xe1497ff7.
//
// Solidity: function setMaxGroupSize(uint256 size) returns(bool)
func (_Validators *ValidatorsSession) SetMaxGroupSize(size *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetMaxGroupSize(&_Validators.TransactOpts, size)
}

// SetMaxGroupSize is a paid mutator transaction binding the contract method 0xe1497ff7.
//
// Solidity: function setMaxGroupSize(uint256 size) returns(bool)
func (_Validators *ValidatorsTransactorSession) SetMaxGroupSize(size *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetMaxGroupSize(&_Validators.TransactOpts, size)
}

// SetMembershipHistoryLength is a paid mutator transaction binding the contract method 0xeff2ea3f.
//
// Solidity: function setMembershipHistoryLength(uint256 length) returns(bool)
func (_Validators *ValidatorsTransactor) SetMembershipHistoryLength(opts *bind.TransactOpts, length *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setMembershipHistoryLength", length)
}

// SetMembershipHistoryLength is a paid mutator transaction binding the contract method 0xeff2ea3f.
//
// Solidity: function setMembershipHistoryLength(uint256 length) returns(bool)
func (_Validators *ValidatorsSession) SetMembershipHistoryLength(length *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetMembershipHistoryLength(&_Validators.TransactOpts, length)
}

// SetMembershipHistoryLength is a paid mutator transaction binding the contract method 0xeff2ea3f.
//
// Solidity: function setMembershipHistoryLength(uint256 length) returns(bool)
func (_Validators *ValidatorsTransactorSession) SetMembershipHistoryLength(length *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetMembershipHistoryLength(&_Validators.TransactOpts, length)
}

// SetNextCommissionUpdate is a paid mutator transaction binding the contract method 0x86d81a5a.
//
// Solidity: function setNextCommissionUpdate(uint256 commission) returns()
func (_Validators *ValidatorsTransactor) SetNextCommissionUpdate(opts *bind.TransactOpts, commission *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setNextCommissionUpdate", commission)
}

// SetNextCommissionUpdate is a paid mutator transaction binding the contract method 0x86d81a5a.
//
// Solidity: function setNextCommissionUpdate(uint256 commission) returns()
func (_Validators *ValidatorsSession) SetNextCommissionUpdate(commission *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetNextCommissionUpdate(&_Validators.TransactOpts, commission)
}

// SetNextCommissionUpdate is a paid mutator transaction binding the contract method 0x86d81a5a.
//
// Solidity: function setNextCommissionUpdate(uint256 commission) returns()
func (_Validators *ValidatorsTransactorSession) SetNextCommissionUpdate(commission *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetNextCommissionUpdate(&_Validators.TransactOpts, commission)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Validators *ValidatorsTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Validators *ValidatorsSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Validators.Contract.SetRegistry(&_Validators.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Validators *ValidatorsTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Validators.Contract.SetRegistry(&_Validators.TransactOpts, registryAddress)
}

// SetSlashingMultiplierResetPeriod is a paid mutator transaction binding the contract method 0x6ab951a0.
//
// Solidity: function setSlashingMultiplierResetPeriod(uint256 value) returns()
func (_Validators *ValidatorsTransactor) SetSlashingMultiplierResetPeriod(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setSlashingMultiplierResetPeriod", value)
}

// SetSlashingMultiplierResetPeriod is a paid mutator transaction binding the contract method 0x6ab951a0.
//
// Solidity: function setSlashingMultiplierResetPeriod(uint256 value) returns()
func (_Validators *ValidatorsSession) SetSlashingMultiplierResetPeriod(value *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetSlashingMultiplierResetPeriod(&_Validators.TransactOpts, value)
}

// SetSlashingMultiplierResetPeriod is a paid mutator transaction binding the contract method 0x6ab951a0.
//
// Solidity: function setSlashingMultiplierResetPeriod(uint256 value) returns()
func (_Validators *ValidatorsTransactorSession) SetSlashingMultiplierResetPeriod(value *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetSlashingMultiplierResetPeriod(&_Validators.TransactOpts, value)
}

// SetValidatorLockedGoldRequirements is a paid mutator transaction binding the contract method 0x76c0a9ed.
//
// Solidity: function setValidatorLockedGoldRequirements(uint256 value, uint256 duration) returns(bool)
func (_Validators *ValidatorsTransactor) SetValidatorLockedGoldRequirements(opts *bind.TransactOpts, value *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setValidatorLockedGoldRequirements", value, duration)
}

// SetValidatorLockedGoldRequirements is a paid mutator transaction binding the contract method 0x76c0a9ed.
//
// Solidity: function setValidatorLockedGoldRequirements(uint256 value, uint256 duration) returns(bool)
func (_Validators *ValidatorsSession) SetValidatorLockedGoldRequirements(value *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetValidatorLockedGoldRequirements(&_Validators.TransactOpts, value, duration)
}

// SetValidatorLockedGoldRequirements is a paid mutator transaction binding the contract method 0x76c0a9ed.
//
// Solidity: function setValidatorLockedGoldRequirements(uint256 value, uint256 duration) returns(bool)
func (_Validators *ValidatorsTransactorSession) SetValidatorLockedGoldRequirements(value *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetValidatorLockedGoldRequirements(&_Validators.TransactOpts, value, duration)
}

// SetValidatorScoreParameters is a paid mutator transaction binding the contract method 0xcb8f98e0.
//
// Solidity: function setValidatorScoreParameters(uint256 exponent, uint256 adjustmentSpeed) returns(bool)
func (_Validators *ValidatorsTransactor) SetValidatorScoreParameters(opts *bind.TransactOpts, exponent *big.Int, adjustmentSpeed *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setValidatorScoreParameters", exponent, adjustmentSpeed)
}

// SetValidatorScoreParameters is a paid mutator transaction binding the contract method 0xcb8f98e0.
//
// Solidity: function setValidatorScoreParameters(uint256 exponent, uint256 adjustmentSpeed) returns(bool)
func (_Validators *ValidatorsSession) SetValidatorScoreParameters(exponent *big.Int, adjustmentSpeed *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetValidatorScoreParameters(&_Validators.TransactOpts, exponent, adjustmentSpeed)
}

// SetValidatorScoreParameters is a paid mutator transaction binding the contract method 0xcb8f98e0.
//
// Solidity: function setValidatorScoreParameters(uint256 exponent, uint256 adjustmentSpeed) returns(bool)
func (_Validators *ValidatorsTransactorSession) SetValidatorScoreParameters(exponent *big.Int, adjustmentSpeed *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetValidatorScoreParameters(&_Validators.TransactOpts, exponent, adjustmentSpeed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Validators *ValidatorsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Validators *ValidatorsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Validators.Contract.TransferOwnership(&_Validators.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Validators *ValidatorsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Validators.Contract.TransferOwnership(&_Validators.TransactOpts, newOwner)
}

// UpdateBlsPublicKey is a paid mutator transaction binding the contract method 0xbfdb7417.
//
// Solidity: function updateBlsPublicKey(bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsTransactor) UpdateBlsPublicKey(opts *bind.TransactOpts, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "updateBlsPublicKey", blsPublicKey, blsPop)
}

// UpdateBlsPublicKey is a paid mutator transaction binding the contract method 0xbfdb7417.
//
// Solidity: function updateBlsPublicKey(bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsSession) UpdateBlsPublicKey(blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.Contract.UpdateBlsPublicKey(&_Validators.TransactOpts, blsPublicKey, blsPop)
}

// UpdateBlsPublicKey is a paid mutator transaction binding the contract method 0xbfdb7417.
//
// Solidity: function updateBlsPublicKey(bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsTransactorSession) UpdateBlsPublicKey(blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.Contract.UpdateBlsPublicKey(&_Validators.TransactOpts, blsPublicKey, blsPop)
}

// UpdateCommission is a paid mutator transaction binding the contract method 0xe7f03766.
//
// Solidity: function updateCommission() returns()
func (_Validators *ValidatorsTransactor) UpdateCommission(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "updateCommission")
}

// UpdateCommission is a paid mutator transaction binding the contract method 0xe7f03766.
//
// Solidity: function updateCommission() returns()
func (_Validators *ValidatorsSession) UpdateCommission() (*types.Transaction, error) {
	return _Validators.Contract.UpdateCommission(&_Validators.TransactOpts)
}

// UpdateCommission is a paid mutator transaction binding the contract method 0xe7f03766.
//
// Solidity: function updateCommission() returns()
func (_Validators *ValidatorsTransactorSession) UpdateCommission() (*types.Transaction, error) {
	return _Validators.Contract.UpdateCommission(&_Validators.TransactOpts)
}

// UpdateEcdsaPublicKey is a paid mutator transaction binding the contract method 0x4e06fd8a.
//
// Solidity: function updateEcdsaPublicKey(address account, address signer, bytes ecdsaPublicKey) returns(bool)
func (_Validators *ValidatorsTransactor) UpdateEcdsaPublicKey(opts *bind.TransactOpts, account common.Address, signer common.Address, ecdsaPublicKey []byte) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "updateEcdsaPublicKey", account, signer, ecdsaPublicKey)
}

// UpdateEcdsaPublicKey is a paid mutator transaction binding the contract method 0x4e06fd8a.
//
// Solidity: function updateEcdsaPublicKey(address account, address signer, bytes ecdsaPublicKey) returns(bool)
func (_Validators *ValidatorsSession) UpdateEcdsaPublicKey(account common.Address, signer common.Address, ecdsaPublicKey []byte) (*types.Transaction, error) {
	return _Validators.Contract.UpdateEcdsaPublicKey(&_Validators.TransactOpts, account, signer, ecdsaPublicKey)
}

// UpdateEcdsaPublicKey is a paid mutator transaction binding the contract method 0x4e06fd8a.
//
// Solidity: function updateEcdsaPublicKey(address account, address signer, bytes ecdsaPublicKey) returns(bool)
func (_Validators *ValidatorsTransactorSession) UpdateEcdsaPublicKey(account common.Address, signer common.Address, ecdsaPublicKey []byte) (*types.Transaction, error) {
	return _Validators.Contract.UpdateEcdsaPublicKey(&_Validators.TransactOpts, account, signer, ecdsaPublicKey)
}

// UpdatePublicKeys is a paid mutator transaction binding the contract method 0x713ea0f3.
//
// Solidity: function updatePublicKeys(address account, address signer, bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsTransactor) UpdatePublicKeys(opts *bind.TransactOpts, account common.Address, signer common.Address, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "updatePublicKeys", account, signer, ecdsaPublicKey, blsPublicKey, blsPop)
}

// UpdatePublicKeys is a paid mutator transaction binding the contract method 0x713ea0f3.
//
// Solidity: function updatePublicKeys(address account, address signer, bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsSession) UpdatePublicKeys(account common.Address, signer common.Address, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.Contract.UpdatePublicKeys(&_Validators.TransactOpts, account, signer, ecdsaPublicKey, blsPublicKey, blsPop)
}

// UpdatePublicKeys is a paid mutator transaction binding the contract method 0x713ea0f3.
//
// Solidity: function updatePublicKeys(address account, address signer, bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsTransactorSession) UpdatePublicKeys(account common.Address, signer common.Address, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.Contract.UpdatePublicKeys(&_Validators.TransactOpts, account, signer, ecdsaPublicKey, blsPublicKey, blsPop)
}

// UpdateValidatorScoreFromSigner is a paid mutator transaction binding the contract method 0xc0c6ad6f.
//
// Solidity: function updateValidatorScoreFromSigner(address signer, uint256 uptime) returns()
func (_Validators *ValidatorsTransactor) UpdateValidatorScoreFromSigner(opts *bind.TransactOpts, signer common.Address, uptime *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "updateValidatorScoreFromSigner", signer, uptime)
}

// UpdateValidatorScoreFromSigner is a paid mutator transaction binding the contract method 0xc0c6ad6f.
//
// Solidity: function updateValidatorScoreFromSigner(address signer, uint256 uptime) returns()
func (_Validators *ValidatorsSession) UpdateValidatorScoreFromSigner(signer common.Address, uptime *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.UpdateValidatorScoreFromSigner(&_Validators.TransactOpts, signer, uptime)
}

// UpdateValidatorScoreFromSigner is a paid mutator transaction binding the contract method 0xc0c6ad6f.
//
// Solidity: function updateValidatorScoreFromSigner(address signer, uint256 uptime) returns()
func (_Validators *ValidatorsTransactorSession) UpdateValidatorScoreFromSigner(signer common.Address, uptime *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.UpdateValidatorScoreFromSigner(&_Validators.TransactOpts, signer, uptime)
}

// ValidatorsCommissionUpdateDelaySetIterator is returned from FilterCommissionUpdateDelaySet and is used to iterate over the raw logs and unpacked data for CommissionUpdateDelaySet events raised by the Validators contract.
type ValidatorsCommissionUpdateDelaySetIterator struct {
	Event *ValidatorsCommissionUpdateDelaySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsCommissionUpdateDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsCommissionUpdateDelaySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsCommissionUpdateDelaySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsCommissionUpdateDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsCommissionUpdateDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsCommissionUpdateDelaySet represents a CommissionUpdateDelaySet event raised by the Validators contract.
type ValidatorsCommissionUpdateDelaySet struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCommissionUpdateDelaySet is a free log retrieval operation binding the contract event 0xf2da07d08fd8dc9c5dcf87ad6f540e306f884a47dd8de14b718a4d5395f1ca9b.
//
// Solidity: event CommissionUpdateDelaySet(uint256 delay)
func (_Validators *ValidatorsFilterer) FilterCommissionUpdateDelaySet(opts *bind.FilterOpts) (*ValidatorsCommissionUpdateDelaySetIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "CommissionUpdateDelaySet")
	if err != nil {
		return nil, err
	}
	return &ValidatorsCommissionUpdateDelaySetIterator{contract: _Validators.contract, event: "CommissionUpdateDelaySet", logs: logs, sub: sub}, nil
}

// WatchCommissionUpdateDelaySet is a free log subscription operation binding the contract event 0xf2da07d08fd8dc9c5dcf87ad6f540e306f884a47dd8de14b718a4d5395f1ca9b.
//
// Solidity: event CommissionUpdateDelaySet(uint256 delay)
func (_Validators *ValidatorsFilterer) WatchCommissionUpdateDelaySet(opts *bind.WatchOpts, sink chan<- *ValidatorsCommissionUpdateDelaySet) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "CommissionUpdateDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsCommissionUpdateDelaySet)
				if err := _Validators.contract.UnpackLog(event, "CommissionUpdateDelaySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommissionUpdateDelaySet is a log parse operation binding the contract event 0xf2da07d08fd8dc9c5dcf87ad6f540e306f884a47dd8de14b718a4d5395f1ca9b.
//
// Solidity: event CommissionUpdateDelaySet(uint256 delay)
func (_Validators *ValidatorsFilterer) ParseCommissionUpdateDelaySet(log types.Log) (*ValidatorsCommissionUpdateDelaySet, error) {
	event := new(ValidatorsCommissionUpdateDelaySet)
	if err := _Validators.contract.UnpackLog(event, "CommissionUpdateDelaySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsGroupLockedGoldRequirementsSetIterator is returned from FilterGroupLockedGoldRequirementsSet and is used to iterate over the raw logs and unpacked data for GroupLockedGoldRequirementsSet events raised by the Validators contract.
type ValidatorsGroupLockedGoldRequirementsSetIterator struct {
	Event *ValidatorsGroupLockedGoldRequirementsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsGroupLockedGoldRequirementsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsGroupLockedGoldRequirementsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsGroupLockedGoldRequirementsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsGroupLockedGoldRequirementsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsGroupLockedGoldRequirementsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsGroupLockedGoldRequirementsSet represents a GroupLockedGoldRequirementsSet event raised by the Validators contract.
type ValidatorsGroupLockedGoldRequirementsSet struct {
	Value    *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGroupLockedGoldRequirementsSet is a free log retrieval operation binding the contract event 0x999f7ee1917e6d7ea08360edfe9250cda3eda859c38dcb71a92623665de64dd4.
//
// Solidity: event GroupLockedGoldRequirementsSet(uint256 value, uint256 duration)
func (_Validators *ValidatorsFilterer) FilterGroupLockedGoldRequirementsSet(opts *bind.FilterOpts) (*ValidatorsGroupLockedGoldRequirementsSetIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "GroupLockedGoldRequirementsSet")
	if err != nil {
		return nil, err
	}
	return &ValidatorsGroupLockedGoldRequirementsSetIterator{contract: _Validators.contract, event: "GroupLockedGoldRequirementsSet", logs: logs, sub: sub}, nil
}

// WatchGroupLockedGoldRequirementsSet is a free log subscription operation binding the contract event 0x999f7ee1917e6d7ea08360edfe9250cda3eda859c38dcb71a92623665de64dd4.
//
// Solidity: event GroupLockedGoldRequirementsSet(uint256 value, uint256 duration)
func (_Validators *ValidatorsFilterer) WatchGroupLockedGoldRequirementsSet(opts *bind.WatchOpts, sink chan<- *ValidatorsGroupLockedGoldRequirementsSet) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "GroupLockedGoldRequirementsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsGroupLockedGoldRequirementsSet)
				if err := _Validators.contract.UnpackLog(event, "GroupLockedGoldRequirementsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGroupLockedGoldRequirementsSet is a log parse operation binding the contract event 0x999f7ee1917e6d7ea08360edfe9250cda3eda859c38dcb71a92623665de64dd4.
//
// Solidity: event GroupLockedGoldRequirementsSet(uint256 value, uint256 duration)
func (_Validators *ValidatorsFilterer) ParseGroupLockedGoldRequirementsSet(log types.Log) (*ValidatorsGroupLockedGoldRequirementsSet, error) {
	event := new(ValidatorsGroupLockedGoldRequirementsSet)
	if err := _Validators.contract.UnpackLog(event, "GroupLockedGoldRequirementsSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsMaxGroupSizeSetIterator is returned from FilterMaxGroupSizeSet and is used to iterate over the raw logs and unpacked data for MaxGroupSizeSet events raised by the Validators contract.
type ValidatorsMaxGroupSizeSetIterator struct {
	Event *ValidatorsMaxGroupSizeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsMaxGroupSizeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsMaxGroupSizeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsMaxGroupSizeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsMaxGroupSizeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsMaxGroupSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsMaxGroupSizeSet represents a MaxGroupSizeSet event raised by the Validators contract.
type ValidatorsMaxGroupSizeSet struct {
	Size *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMaxGroupSizeSet is a free log retrieval operation binding the contract event 0x603fe12c33c253a23da1680aa453dc70c3a0ee07763569bd5f602406ebd4e5d5.
//
// Solidity: event MaxGroupSizeSet(uint256 size)
func (_Validators *ValidatorsFilterer) FilterMaxGroupSizeSet(opts *bind.FilterOpts) (*ValidatorsMaxGroupSizeSetIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "MaxGroupSizeSet")
	if err != nil {
		return nil, err
	}
	return &ValidatorsMaxGroupSizeSetIterator{contract: _Validators.contract, event: "MaxGroupSizeSet", logs: logs, sub: sub}, nil
}

// WatchMaxGroupSizeSet is a free log subscription operation binding the contract event 0x603fe12c33c253a23da1680aa453dc70c3a0ee07763569bd5f602406ebd4e5d5.
//
// Solidity: event MaxGroupSizeSet(uint256 size)
func (_Validators *ValidatorsFilterer) WatchMaxGroupSizeSet(opts *bind.WatchOpts, sink chan<- *ValidatorsMaxGroupSizeSet) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "MaxGroupSizeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsMaxGroupSizeSet)
				if err := _Validators.contract.UnpackLog(event, "MaxGroupSizeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMaxGroupSizeSet is a log parse operation binding the contract event 0x603fe12c33c253a23da1680aa453dc70c3a0ee07763569bd5f602406ebd4e5d5.
//
// Solidity: event MaxGroupSizeSet(uint256 size)
func (_Validators *ValidatorsFilterer) ParseMaxGroupSizeSet(log types.Log) (*ValidatorsMaxGroupSizeSet, error) {
	event := new(ValidatorsMaxGroupSizeSet)
	if err := _Validators.contract.UnpackLog(event, "MaxGroupSizeSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsMembershipHistoryLengthSetIterator is returned from FilterMembershipHistoryLengthSet and is used to iterate over the raw logs and unpacked data for MembershipHistoryLengthSet events raised by the Validators contract.
type ValidatorsMembershipHistoryLengthSetIterator struct {
	Event *ValidatorsMembershipHistoryLengthSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsMembershipHistoryLengthSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsMembershipHistoryLengthSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsMembershipHistoryLengthSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsMembershipHistoryLengthSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsMembershipHistoryLengthSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsMembershipHistoryLengthSet represents a MembershipHistoryLengthSet event raised by the Validators contract.
type ValidatorsMembershipHistoryLengthSet struct {
	Length *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMembershipHistoryLengthSet is a free log retrieval operation binding the contract event 0x1c75c7fb3ee9d13d8394372d8c7cdf1702fa947faa03f6ccfa500f787b09b48a.
//
// Solidity: event MembershipHistoryLengthSet(uint256 length)
func (_Validators *ValidatorsFilterer) FilterMembershipHistoryLengthSet(opts *bind.FilterOpts) (*ValidatorsMembershipHistoryLengthSetIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "MembershipHistoryLengthSet")
	if err != nil {
		return nil, err
	}
	return &ValidatorsMembershipHistoryLengthSetIterator{contract: _Validators.contract, event: "MembershipHistoryLengthSet", logs: logs, sub: sub}, nil
}

// WatchMembershipHistoryLengthSet is a free log subscription operation binding the contract event 0x1c75c7fb3ee9d13d8394372d8c7cdf1702fa947faa03f6ccfa500f787b09b48a.
//
// Solidity: event MembershipHistoryLengthSet(uint256 length)
func (_Validators *ValidatorsFilterer) WatchMembershipHistoryLengthSet(opts *bind.WatchOpts, sink chan<- *ValidatorsMembershipHistoryLengthSet) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "MembershipHistoryLengthSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsMembershipHistoryLengthSet)
				if err := _Validators.contract.UnpackLog(event, "MembershipHistoryLengthSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMembershipHistoryLengthSet is a log parse operation binding the contract event 0x1c75c7fb3ee9d13d8394372d8c7cdf1702fa947faa03f6ccfa500f787b09b48a.
//
// Solidity: event MembershipHistoryLengthSet(uint256 length)
func (_Validators *ValidatorsFilterer) ParseMembershipHistoryLengthSet(log types.Log) (*ValidatorsMembershipHistoryLengthSet, error) {
	event := new(ValidatorsMembershipHistoryLengthSet)
	if err := _Validators.contract.UnpackLog(event, "MembershipHistoryLengthSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Validators contract.
type ValidatorsOwnershipTransferredIterator struct {
	Event *ValidatorsOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsOwnershipTransferred represents a OwnershipTransferred event raised by the Validators contract.
type ValidatorsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Validators *ValidatorsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ValidatorsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsOwnershipTransferredIterator{contract: _Validators.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Validators *ValidatorsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ValidatorsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsOwnershipTransferred)
				if err := _Validators.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Validators *ValidatorsFilterer) ParseOwnershipTransferred(log types.Log) (*ValidatorsOwnershipTransferred, error) {
	event := new(ValidatorsOwnershipTransferred)
	if err := _Validators.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the Validators contract.
type ValidatorsRegistrySetIterator struct {
	Event *ValidatorsRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsRegistrySet represents a RegistrySet event raised by the Validators contract.
type ValidatorsRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Validators *ValidatorsFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*ValidatorsRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsRegistrySetIterator{contract: _Validators.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Validators *ValidatorsFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *ValidatorsRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsRegistrySet)
				if err := _Validators.contract.UnpackLog(event, "RegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegistrySet is a log parse operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Validators *ValidatorsFilterer) ParseRegistrySet(log types.Log) (*ValidatorsRegistrySet, error) {
	event := new(ValidatorsRegistrySet)
	if err := _Validators.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorAffiliatedIterator is returned from FilterValidatorAffiliated and is used to iterate over the raw logs and unpacked data for ValidatorAffiliated events raised by the Validators contract.
type ValidatorsValidatorAffiliatedIterator struct {
	Event *ValidatorsValidatorAffiliated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorAffiliatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorAffiliated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorAffiliated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorAffiliatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorAffiliatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorAffiliated represents a ValidatorAffiliated event raised by the Validators contract.
type ValidatorsValidatorAffiliated struct {
	Validator common.Address
	Group     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAffiliated is a free log retrieval operation binding the contract event 0x91ef92227057e201e406c3451698dd780fe7672ad74328591c88d281af31581d.
//
// Solidity: event ValidatorAffiliated(address indexed validator, address indexed group)
func (_Validators *ValidatorsFilterer) FilterValidatorAffiliated(opts *bind.FilterOpts, validator []common.Address, group []common.Address) (*ValidatorsValidatorAffiliatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorAffiliated", validatorRule, groupRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorAffiliatedIterator{contract: _Validators.contract, event: "ValidatorAffiliated", logs: logs, sub: sub}, nil
}

// WatchValidatorAffiliated is a free log subscription operation binding the contract event 0x91ef92227057e201e406c3451698dd780fe7672ad74328591c88d281af31581d.
//
// Solidity: event ValidatorAffiliated(address indexed validator, address indexed group)
func (_Validators *ValidatorsFilterer) WatchValidatorAffiliated(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorAffiliated, validator []common.Address, group []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorAffiliated", validatorRule, groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorAffiliated)
				if err := _Validators.contract.UnpackLog(event, "ValidatorAffiliated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorAffiliated is a log parse operation binding the contract event 0x91ef92227057e201e406c3451698dd780fe7672ad74328591c88d281af31581d.
//
// Solidity: event ValidatorAffiliated(address indexed validator, address indexed group)
func (_Validators *ValidatorsFilterer) ParseValidatorAffiliated(log types.Log) (*ValidatorsValidatorAffiliated, error) {
	event := new(ValidatorsValidatorAffiliated)
	if err := _Validators.contract.UnpackLog(event, "ValidatorAffiliated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorBlsPublicKeyUpdatedIterator is returned from FilterValidatorBlsPublicKeyUpdated and is used to iterate over the raw logs and unpacked data for ValidatorBlsPublicKeyUpdated events raised by the Validators contract.
type ValidatorsValidatorBlsPublicKeyUpdatedIterator struct {
	Event *ValidatorsValidatorBlsPublicKeyUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorBlsPublicKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorBlsPublicKeyUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorBlsPublicKeyUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorBlsPublicKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorBlsPublicKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorBlsPublicKeyUpdated represents a ValidatorBlsPublicKeyUpdated event raised by the Validators contract.
type ValidatorsValidatorBlsPublicKeyUpdated struct {
	Validator    common.Address
	BlsPublicKey []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorBlsPublicKeyUpdated is a free log retrieval operation binding the contract event 0x36a1aabe506bbe8802233cbb9aad628e91269e77077c953f9db3e02d7092ee33.
//
// Solidity: event ValidatorBlsPublicKeyUpdated(address indexed validator, bytes blsPublicKey)
func (_Validators *ValidatorsFilterer) FilterValidatorBlsPublicKeyUpdated(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsValidatorBlsPublicKeyUpdatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorBlsPublicKeyUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorBlsPublicKeyUpdatedIterator{contract: _Validators.contract, event: "ValidatorBlsPublicKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorBlsPublicKeyUpdated is a free log subscription operation binding the contract event 0x36a1aabe506bbe8802233cbb9aad628e91269e77077c953f9db3e02d7092ee33.
//
// Solidity: event ValidatorBlsPublicKeyUpdated(address indexed validator, bytes blsPublicKey)
func (_Validators *ValidatorsFilterer) WatchValidatorBlsPublicKeyUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorBlsPublicKeyUpdated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorBlsPublicKeyUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorBlsPublicKeyUpdated)
				if err := _Validators.contract.UnpackLog(event, "ValidatorBlsPublicKeyUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorBlsPublicKeyUpdated is a log parse operation binding the contract event 0x36a1aabe506bbe8802233cbb9aad628e91269e77077c953f9db3e02d7092ee33.
//
// Solidity: event ValidatorBlsPublicKeyUpdated(address indexed validator, bytes blsPublicKey)
func (_Validators *ValidatorsFilterer) ParseValidatorBlsPublicKeyUpdated(log types.Log) (*ValidatorsValidatorBlsPublicKeyUpdated, error) {
	event := new(ValidatorsValidatorBlsPublicKeyUpdated)
	if err := _Validators.contract.UnpackLog(event, "ValidatorBlsPublicKeyUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorDeaffiliatedIterator is returned from FilterValidatorDeaffiliated and is used to iterate over the raw logs and unpacked data for ValidatorDeaffiliated events raised by the Validators contract.
type ValidatorsValidatorDeaffiliatedIterator struct {
	Event *ValidatorsValidatorDeaffiliated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorDeaffiliatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorDeaffiliated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorDeaffiliated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorDeaffiliatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorDeaffiliatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorDeaffiliated represents a ValidatorDeaffiliated event raised by the Validators contract.
type ValidatorsValidatorDeaffiliated struct {
	Validator common.Address
	Group     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeaffiliated is a free log retrieval operation binding the contract event 0x71815121f0622b31a3e7270eb28acb9fd10825ff418c9a18591f617bb8a31a6c.
//
// Solidity: event ValidatorDeaffiliated(address indexed validator, address indexed group)
func (_Validators *ValidatorsFilterer) FilterValidatorDeaffiliated(opts *bind.FilterOpts, validator []common.Address, group []common.Address) (*ValidatorsValidatorDeaffiliatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorDeaffiliated", validatorRule, groupRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorDeaffiliatedIterator{contract: _Validators.contract, event: "ValidatorDeaffiliated", logs: logs, sub: sub}, nil
}

// WatchValidatorDeaffiliated is a free log subscription operation binding the contract event 0x71815121f0622b31a3e7270eb28acb9fd10825ff418c9a18591f617bb8a31a6c.
//
// Solidity: event ValidatorDeaffiliated(address indexed validator, address indexed group)
func (_Validators *ValidatorsFilterer) WatchValidatorDeaffiliated(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorDeaffiliated, validator []common.Address, group []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorDeaffiliated", validatorRule, groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorDeaffiliated)
				if err := _Validators.contract.UnpackLog(event, "ValidatorDeaffiliated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorDeaffiliated is a log parse operation binding the contract event 0x71815121f0622b31a3e7270eb28acb9fd10825ff418c9a18591f617bb8a31a6c.
//
// Solidity: event ValidatorDeaffiliated(address indexed validator, address indexed group)
func (_Validators *ValidatorsFilterer) ParseValidatorDeaffiliated(log types.Log) (*ValidatorsValidatorDeaffiliated, error) {
	event := new(ValidatorsValidatorDeaffiliated)
	if err := _Validators.contract.UnpackLog(event, "ValidatorDeaffiliated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorDeregisteredIterator is returned from FilterValidatorDeregistered and is used to iterate over the raw logs and unpacked data for ValidatorDeregistered events raised by the Validators contract.
type ValidatorsValidatorDeregisteredIterator struct {
	Event *ValidatorsValidatorDeregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorDeregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorDeregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorDeregistered represents a ValidatorDeregistered event raised by the Validators contract.
type ValidatorsValidatorDeregistered struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeregistered is a free log retrieval operation binding the contract event 0x51407fafe7ef9bec39c65a12a4885a274190991bf1e9057fcc384fc77ff1a7f0.
//
// Solidity: event ValidatorDeregistered(address indexed validator)
func (_Validators *ValidatorsFilterer) FilterValidatorDeregistered(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsValidatorDeregisteredIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorDeregistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorDeregisteredIterator{contract: _Validators.contract, event: "ValidatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchValidatorDeregistered is a free log subscription operation binding the contract event 0x51407fafe7ef9bec39c65a12a4885a274190991bf1e9057fcc384fc77ff1a7f0.
//
// Solidity: event ValidatorDeregistered(address indexed validator)
func (_Validators *ValidatorsFilterer) WatchValidatorDeregistered(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorDeregistered, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorDeregistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorDeregistered)
				if err := _Validators.contract.UnpackLog(event, "ValidatorDeregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorDeregistered is a log parse operation binding the contract event 0x51407fafe7ef9bec39c65a12a4885a274190991bf1e9057fcc384fc77ff1a7f0.
//
// Solidity: event ValidatorDeregistered(address indexed validator)
func (_Validators *ValidatorsFilterer) ParseValidatorDeregistered(log types.Log) (*ValidatorsValidatorDeregistered, error) {
	event := new(ValidatorsValidatorDeregistered)
	if err := _Validators.contract.UnpackLog(event, "ValidatorDeregistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorEcdsaPublicKeyUpdatedIterator is returned from FilterValidatorEcdsaPublicKeyUpdated and is used to iterate over the raw logs and unpacked data for ValidatorEcdsaPublicKeyUpdated events raised by the Validators contract.
type ValidatorsValidatorEcdsaPublicKeyUpdatedIterator struct {
	Event *ValidatorsValidatorEcdsaPublicKeyUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorEcdsaPublicKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorEcdsaPublicKeyUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorEcdsaPublicKeyUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorEcdsaPublicKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorEcdsaPublicKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorEcdsaPublicKeyUpdated represents a ValidatorEcdsaPublicKeyUpdated event raised by the Validators contract.
type ValidatorsValidatorEcdsaPublicKeyUpdated struct {
	Validator      common.Address
	EcdsaPublicKey []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorEcdsaPublicKeyUpdated is a free log retrieval operation binding the contract event 0x213377eec2c15b21fa7abcbb0cb87a67e893cdb94a2564aa4bb4d380869473c8.
//
// Solidity: event ValidatorEcdsaPublicKeyUpdated(address indexed validator, bytes ecdsaPublicKey)
func (_Validators *ValidatorsFilterer) FilterValidatorEcdsaPublicKeyUpdated(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsValidatorEcdsaPublicKeyUpdatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorEcdsaPublicKeyUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorEcdsaPublicKeyUpdatedIterator{contract: _Validators.contract, event: "ValidatorEcdsaPublicKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorEcdsaPublicKeyUpdated is a free log subscription operation binding the contract event 0x213377eec2c15b21fa7abcbb0cb87a67e893cdb94a2564aa4bb4d380869473c8.
//
// Solidity: event ValidatorEcdsaPublicKeyUpdated(address indexed validator, bytes ecdsaPublicKey)
func (_Validators *ValidatorsFilterer) WatchValidatorEcdsaPublicKeyUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorEcdsaPublicKeyUpdated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorEcdsaPublicKeyUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorEcdsaPublicKeyUpdated)
				if err := _Validators.contract.UnpackLog(event, "ValidatorEcdsaPublicKeyUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorEcdsaPublicKeyUpdated is a log parse operation binding the contract event 0x213377eec2c15b21fa7abcbb0cb87a67e893cdb94a2564aa4bb4d380869473c8.
//
// Solidity: event ValidatorEcdsaPublicKeyUpdated(address indexed validator, bytes ecdsaPublicKey)
func (_Validators *ValidatorsFilterer) ParseValidatorEcdsaPublicKeyUpdated(log types.Log) (*ValidatorsValidatorEcdsaPublicKeyUpdated, error) {
	event := new(ValidatorsValidatorEcdsaPublicKeyUpdated)
	if err := _Validators.contract.UnpackLog(event, "ValidatorEcdsaPublicKeyUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorEpochPaymentDistributedIterator is returned from FilterValidatorEpochPaymentDistributed and is used to iterate over the raw logs and unpacked data for ValidatorEpochPaymentDistributed events raised by the Validators contract.
type ValidatorsValidatorEpochPaymentDistributedIterator struct {
	Event *ValidatorsValidatorEpochPaymentDistributed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorEpochPaymentDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorEpochPaymentDistributed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorEpochPaymentDistributed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorEpochPaymentDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorEpochPaymentDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorEpochPaymentDistributed represents a ValidatorEpochPaymentDistributed event raised by the Validators contract.
type ValidatorsValidatorEpochPaymentDistributed struct {
	Validator        common.Address
	ValidatorPayment *big.Int
	Group            common.Address
	GroupPayment     *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorEpochPaymentDistributed is a free log retrieval operation binding the contract event 0x6f5937add2ec38a0fa4959bccd86e3fcc2aafb706cd3e6c0565f87a7b36b9975.
//
// Solidity: event ValidatorEpochPaymentDistributed(address indexed validator, uint256 validatorPayment, address indexed group, uint256 groupPayment)
func (_Validators *ValidatorsFilterer) FilterValidatorEpochPaymentDistributed(opts *bind.FilterOpts, validator []common.Address, group []common.Address) (*ValidatorsValidatorEpochPaymentDistributedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorEpochPaymentDistributed", validatorRule, groupRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorEpochPaymentDistributedIterator{contract: _Validators.contract, event: "ValidatorEpochPaymentDistributed", logs: logs, sub: sub}, nil
}

// WatchValidatorEpochPaymentDistributed is a free log subscription operation binding the contract event 0x6f5937add2ec38a0fa4959bccd86e3fcc2aafb706cd3e6c0565f87a7b36b9975.
//
// Solidity: event ValidatorEpochPaymentDistributed(address indexed validator, uint256 validatorPayment, address indexed group, uint256 groupPayment)
func (_Validators *ValidatorsFilterer) WatchValidatorEpochPaymentDistributed(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorEpochPaymentDistributed, validator []common.Address, group []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorEpochPaymentDistributed", validatorRule, groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorEpochPaymentDistributed)
				if err := _Validators.contract.UnpackLog(event, "ValidatorEpochPaymentDistributed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorEpochPaymentDistributed is a log parse operation binding the contract event 0x6f5937add2ec38a0fa4959bccd86e3fcc2aafb706cd3e6c0565f87a7b36b9975.
//
// Solidity: event ValidatorEpochPaymentDistributed(address indexed validator, uint256 validatorPayment, address indexed group, uint256 groupPayment)
func (_Validators *ValidatorsFilterer) ParseValidatorEpochPaymentDistributed(log types.Log) (*ValidatorsValidatorEpochPaymentDistributed, error) {
	event := new(ValidatorsValidatorEpochPaymentDistributed)
	if err := _Validators.contract.UnpackLog(event, "ValidatorEpochPaymentDistributed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorGroupCommissionUpdateQueuedIterator is returned from FilterValidatorGroupCommissionUpdateQueued and is used to iterate over the raw logs and unpacked data for ValidatorGroupCommissionUpdateQueued events raised by the Validators contract.
type ValidatorsValidatorGroupCommissionUpdateQueuedIterator struct {
	Event *ValidatorsValidatorGroupCommissionUpdateQueued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorGroupCommissionUpdateQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorGroupCommissionUpdateQueued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorGroupCommissionUpdateQueued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorGroupCommissionUpdateQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorGroupCommissionUpdateQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorGroupCommissionUpdateQueued represents a ValidatorGroupCommissionUpdateQueued event raised by the Validators contract.
type ValidatorsValidatorGroupCommissionUpdateQueued struct {
	Group           common.Address
	Commission      *big.Int
	ActivationBlock *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupCommissionUpdateQueued is a free log retrieval operation binding the contract event 0x557d39a57520d9835859d4b7eda805a7f4115a59c3a374eeed488436fc62a152.
//
// Solidity: event ValidatorGroupCommissionUpdateQueued(address indexed group, uint256 commission, uint256 activationBlock)
func (_Validators *ValidatorsFilterer) FilterValidatorGroupCommissionUpdateQueued(opts *bind.FilterOpts, group []common.Address) (*ValidatorsValidatorGroupCommissionUpdateQueuedIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorGroupCommissionUpdateQueued", groupRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorGroupCommissionUpdateQueuedIterator{contract: _Validators.contract, event: "ValidatorGroupCommissionUpdateQueued", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupCommissionUpdateQueued is a free log subscription operation binding the contract event 0x557d39a57520d9835859d4b7eda805a7f4115a59c3a374eeed488436fc62a152.
//
// Solidity: event ValidatorGroupCommissionUpdateQueued(address indexed group, uint256 commission, uint256 activationBlock)
func (_Validators *ValidatorsFilterer) WatchValidatorGroupCommissionUpdateQueued(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorGroupCommissionUpdateQueued, group []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorGroupCommissionUpdateQueued", groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorGroupCommissionUpdateQueued)
				if err := _Validators.contract.UnpackLog(event, "ValidatorGroupCommissionUpdateQueued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorGroupCommissionUpdateQueued is a log parse operation binding the contract event 0x557d39a57520d9835859d4b7eda805a7f4115a59c3a374eeed488436fc62a152.
//
// Solidity: event ValidatorGroupCommissionUpdateQueued(address indexed group, uint256 commission, uint256 activationBlock)
func (_Validators *ValidatorsFilterer) ParseValidatorGroupCommissionUpdateQueued(log types.Log) (*ValidatorsValidatorGroupCommissionUpdateQueued, error) {
	event := new(ValidatorsValidatorGroupCommissionUpdateQueued)
	if err := _Validators.contract.UnpackLog(event, "ValidatorGroupCommissionUpdateQueued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorGroupCommissionUpdatedIterator is returned from FilterValidatorGroupCommissionUpdated and is used to iterate over the raw logs and unpacked data for ValidatorGroupCommissionUpdated events raised by the Validators contract.
type ValidatorsValidatorGroupCommissionUpdatedIterator struct {
	Event *ValidatorsValidatorGroupCommissionUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorGroupCommissionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorGroupCommissionUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorGroupCommissionUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorGroupCommissionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorGroupCommissionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorGroupCommissionUpdated represents a ValidatorGroupCommissionUpdated event raised by the Validators contract.
type ValidatorsValidatorGroupCommissionUpdated struct {
	Group      common.Address
	Commission *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupCommissionUpdated is a free log retrieval operation binding the contract event 0x815d292dbc1a08dfb3103aabb6611233dd2393903e57bdf4c5b3db91198a826c.
//
// Solidity: event ValidatorGroupCommissionUpdated(address indexed group, uint256 commission)
func (_Validators *ValidatorsFilterer) FilterValidatorGroupCommissionUpdated(opts *bind.FilterOpts, group []common.Address) (*ValidatorsValidatorGroupCommissionUpdatedIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorGroupCommissionUpdated", groupRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorGroupCommissionUpdatedIterator{contract: _Validators.contract, event: "ValidatorGroupCommissionUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupCommissionUpdated is a free log subscription operation binding the contract event 0x815d292dbc1a08dfb3103aabb6611233dd2393903e57bdf4c5b3db91198a826c.
//
// Solidity: event ValidatorGroupCommissionUpdated(address indexed group, uint256 commission)
func (_Validators *ValidatorsFilterer) WatchValidatorGroupCommissionUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorGroupCommissionUpdated, group []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorGroupCommissionUpdated", groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorGroupCommissionUpdated)
				if err := _Validators.contract.UnpackLog(event, "ValidatorGroupCommissionUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorGroupCommissionUpdated is a log parse operation binding the contract event 0x815d292dbc1a08dfb3103aabb6611233dd2393903e57bdf4c5b3db91198a826c.
//
// Solidity: event ValidatorGroupCommissionUpdated(address indexed group, uint256 commission)
func (_Validators *ValidatorsFilterer) ParseValidatorGroupCommissionUpdated(log types.Log) (*ValidatorsValidatorGroupCommissionUpdated, error) {
	event := new(ValidatorsValidatorGroupCommissionUpdated)
	if err := _Validators.contract.UnpackLog(event, "ValidatorGroupCommissionUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorGroupDeregisteredIterator is returned from FilterValidatorGroupDeregistered and is used to iterate over the raw logs and unpacked data for ValidatorGroupDeregistered events raised by the Validators contract.
type ValidatorsValidatorGroupDeregisteredIterator struct {
	Event *ValidatorsValidatorGroupDeregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorGroupDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorGroupDeregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorGroupDeregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorGroupDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorGroupDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorGroupDeregistered represents a ValidatorGroupDeregistered event raised by the Validators contract.
type ValidatorsValidatorGroupDeregistered struct {
	Group common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupDeregistered is a free log retrieval operation binding the contract event 0xae7e034b0748a10a219b46074b20977a9170bf4027b156c797093773619a8669.
//
// Solidity: event ValidatorGroupDeregistered(address indexed group)
func (_Validators *ValidatorsFilterer) FilterValidatorGroupDeregistered(opts *bind.FilterOpts, group []common.Address) (*ValidatorsValidatorGroupDeregisteredIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorGroupDeregistered", groupRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorGroupDeregisteredIterator{contract: _Validators.contract, event: "ValidatorGroupDeregistered", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupDeregistered is a free log subscription operation binding the contract event 0xae7e034b0748a10a219b46074b20977a9170bf4027b156c797093773619a8669.
//
// Solidity: event ValidatorGroupDeregistered(address indexed group)
func (_Validators *ValidatorsFilterer) WatchValidatorGroupDeregistered(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorGroupDeregistered, group []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorGroupDeregistered", groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorGroupDeregistered)
				if err := _Validators.contract.UnpackLog(event, "ValidatorGroupDeregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorGroupDeregistered is a log parse operation binding the contract event 0xae7e034b0748a10a219b46074b20977a9170bf4027b156c797093773619a8669.
//
// Solidity: event ValidatorGroupDeregistered(address indexed group)
func (_Validators *ValidatorsFilterer) ParseValidatorGroupDeregistered(log types.Log) (*ValidatorsValidatorGroupDeregistered, error) {
	event := new(ValidatorsValidatorGroupDeregistered)
	if err := _Validators.contract.UnpackLog(event, "ValidatorGroupDeregistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorGroupMemberAddedIterator is returned from FilterValidatorGroupMemberAdded and is used to iterate over the raw logs and unpacked data for ValidatorGroupMemberAdded events raised by the Validators contract.
type ValidatorsValidatorGroupMemberAddedIterator struct {
	Event *ValidatorsValidatorGroupMemberAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorGroupMemberAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorGroupMemberAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorGroupMemberAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorGroupMemberAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorGroupMemberAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorGroupMemberAdded represents a ValidatorGroupMemberAdded event raised by the Validators contract.
type ValidatorsValidatorGroupMemberAdded struct {
	Group     common.Address
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupMemberAdded is a free log retrieval operation binding the contract event 0xbdf7e616a6943f81e07a7984c9d4c00197dc2f481486ce4ffa6af52a113974ad.
//
// Solidity: event ValidatorGroupMemberAdded(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) FilterValidatorGroupMemberAdded(opts *bind.FilterOpts, group []common.Address, validator []common.Address) (*ValidatorsValidatorGroupMemberAddedIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorGroupMemberAdded", groupRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorGroupMemberAddedIterator{contract: _Validators.contract, event: "ValidatorGroupMemberAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupMemberAdded is a free log subscription operation binding the contract event 0xbdf7e616a6943f81e07a7984c9d4c00197dc2f481486ce4ffa6af52a113974ad.
//
// Solidity: event ValidatorGroupMemberAdded(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) WatchValidatorGroupMemberAdded(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorGroupMemberAdded, group []common.Address, validator []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorGroupMemberAdded", groupRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorGroupMemberAdded)
				if err := _Validators.contract.UnpackLog(event, "ValidatorGroupMemberAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorGroupMemberAdded is a log parse operation binding the contract event 0xbdf7e616a6943f81e07a7984c9d4c00197dc2f481486ce4ffa6af52a113974ad.
//
// Solidity: event ValidatorGroupMemberAdded(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) ParseValidatorGroupMemberAdded(log types.Log) (*ValidatorsValidatorGroupMemberAdded, error) {
	event := new(ValidatorsValidatorGroupMemberAdded)
	if err := _Validators.contract.UnpackLog(event, "ValidatorGroupMemberAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorGroupMemberRemovedIterator is returned from FilterValidatorGroupMemberRemoved and is used to iterate over the raw logs and unpacked data for ValidatorGroupMemberRemoved events raised by the Validators contract.
type ValidatorsValidatorGroupMemberRemovedIterator struct {
	Event *ValidatorsValidatorGroupMemberRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorGroupMemberRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorGroupMemberRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorGroupMemberRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorGroupMemberRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorGroupMemberRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorGroupMemberRemoved represents a ValidatorGroupMemberRemoved event raised by the Validators contract.
type ValidatorsValidatorGroupMemberRemoved struct {
	Group     common.Address
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupMemberRemoved is a free log retrieval operation binding the contract event 0xc7666a52a66ff601ff7c0d4d6efddc9ac20a34792f6aa003d1804c9d4d5baa57.
//
// Solidity: event ValidatorGroupMemberRemoved(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) FilterValidatorGroupMemberRemoved(opts *bind.FilterOpts, group []common.Address, validator []common.Address) (*ValidatorsValidatorGroupMemberRemovedIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorGroupMemberRemoved", groupRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorGroupMemberRemovedIterator{contract: _Validators.contract, event: "ValidatorGroupMemberRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupMemberRemoved is a free log subscription operation binding the contract event 0xc7666a52a66ff601ff7c0d4d6efddc9ac20a34792f6aa003d1804c9d4d5baa57.
//
// Solidity: event ValidatorGroupMemberRemoved(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) WatchValidatorGroupMemberRemoved(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorGroupMemberRemoved, group []common.Address, validator []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorGroupMemberRemoved", groupRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorGroupMemberRemoved)
				if err := _Validators.contract.UnpackLog(event, "ValidatorGroupMemberRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorGroupMemberRemoved is a log parse operation binding the contract event 0xc7666a52a66ff601ff7c0d4d6efddc9ac20a34792f6aa003d1804c9d4d5baa57.
//
// Solidity: event ValidatorGroupMemberRemoved(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) ParseValidatorGroupMemberRemoved(log types.Log) (*ValidatorsValidatorGroupMemberRemoved, error) {
	event := new(ValidatorsValidatorGroupMemberRemoved)
	if err := _Validators.contract.UnpackLog(event, "ValidatorGroupMemberRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorGroupMemberReorderedIterator is returned from FilterValidatorGroupMemberReordered and is used to iterate over the raw logs and unpacked data for ValidatorGroupMemberReordered events raised by the Validators contract.
type ValidatorsValidatorGroupMemberReorderedIterator struct {
	Event *ValidatorsValidatorGroupMemberReordered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorGroupMemberReorderedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorGroupMemberReordered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorGroupMemberReordered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorGroupMemberReorderedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorGroupMemberReorderedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorGroupMemberReordered represents a ValidatorGroupMemberReordered event raised by the Validators contract.
type ValidatorsValidatorGroupMemberReordered struct {
	Group     common.Address
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupMemberReordered is a free log retrieval operation binding the contract event 0x38819cc49a343985b478d72f531a35b15384c398dd80fd191a14662170f895c6.
//
// Solidity: event ValidatorGroupMemberReordered(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) FilterValidatorGroupMemberReordered(opts *bind.FilterOpts, group []common.Address, validator []common.Address) (*ValidatorsValidatorGroupMemberReorderedIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorGroupMemberReordered", groupRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorGroupMemberReorderedIterator{contract: _Validators.contract, event: "ValidatorGroupMemberReordered", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupMemberReordered is a free log subscription operation binding the contract event 0x38819cc49a343985b478d72f531a35b15384c398dd80fd191a14662170f895c6.
//
// Solidity: event ValidatorGroupMemberReordered(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) WatchValidatorGroupMemberReordered(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorGroupMemberReordered, group []common.Address, validator []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorGroupMemberReordered", groupRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorGroupMemberReordered)
				if err := _Validators.contract.UnpackLog(event, "ValidatorGroupMemberReordered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorGroupMemberReordered is a log parse operation binding the contract event 0x38819cc49a343985b478d72f531a35b15384c398dd80fd191a14662170f895c6.
//
// Solidity: event ValidatorGroupMemberReordered(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) ParseValidatorGroupMemberReordered(log types.Log) (*ValidatorsValidatorGroupMemberReordered, error) {
	event := new(ValidatorsValidatorGroupMemberReordered)
	if err := _Validators.contract.UnpackLog(event, "ValidatorGroupMemberReordered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorGroupRegisteredIterator is returned from FilterValidatorGroupRegistered and is used to iterate over the raw logs and unpacked data for ValidatorGroupRegistered events raised by the Validators contract.
type ValidatorsValidatorGroupRegisteredIterator struct {
	Event *ValidatorsValidatorGroupRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorGroupRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorGroupRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorGroupRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorGroupRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorGroupRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorGroupRegistered represents a ValidatorGroupRegistered event raised by the Validators contract.
type ValidatorsValidatorGroupRegistered struct {
	Group      common.Address
	Commission *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupRegistered is a free log retrieval operation binding the contract event 0xbf4b45570f1907a94775f8449817051a492a676918e38108bb762e991e6b58dc.
//
// Solidity: event ValidatorGroupRegistered(address indexed group, uint256 commission)
func (_Validators *ValidatorsFilterer) FilterValidatorGroupRegistered(opts *bind.FilterOpts, group []common.Address) (*ValidatorsValidatorGroupRegisteredIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorGroupRegistered", groupRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorGroupRegisteredIterator{contract: _Validators.contract, event: "ValidatorGroupRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupRegistered is a free log subscription operation binding the contract event 0xbf4b45570f1907a94775f8449817051a492a676918e38108bb762e991e6b58dc.
//
// Solidity: event ValidatorGroupRegistered(address indexed group, uint256 commission)
func (_Validators *ValidatorsFilterer) WatchValidatorGroupRegistered(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorGroupRegistered, group []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorGroupRegistered", groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorGroupRegistered)
				if err := _Validators.contract.UnpackLog(event, "ValidatorGroupRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorGroupRegistered is a log parse operation binding the contract event 0xbf4b45570f1907a94775f8449817051a492a676918e38108bb762e991e6b58dc.
//
// Solidity: event ValidatorGroupRegistered(address indexed group, uint256 commission)
func (_Validators *ValidatorsFilterer) ParseValidatorGroupRegistered(log types.Log) (*ValidatorsValidatorGroupRegistered, error) {
	event := new(ValidatorsValidatorGroupRegistered)
	if err := _Validators.contract.UnpackLog(event, "ValidatorGroupRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorLockedGoldRequirementsSetIterator is returned from FilterValidatorLockedGoldRequirementsSet and is used to iterate over the raw logs and unpacked data for ValidatorLockedGoldRequirementsSet events raised by the Validators contract.
type ValidatorsValidatorLockedGoldRequirementsSetIterator struct {
	Event *ValidatorsValidatorLockedGoldRequirementsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorLockedGoldRequirementsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorLockedGoldRequirementsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorLockedGoldRequirementsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorLockedGoldRequirementsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorLockedGoldRequirementsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorLockedGoldRequirementsSet represents a ValidatorLockedGoldRequirementsSet event raised by the Validators contract.
type ValidatorsValidatorLockedGoldRequirementsSet struct {
	Value    *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterValidatorLockedGoldRequirementsSet is a free log retrieval operation binding the contract event 0x62d947118dd4c1f5ece7f787a9cad4e1127d14d403b71133e95792b473bf8389.
//
// Solidity: event ValidatorLockedGoldRequirementsSet(uint256 value, uint256 duration)
func (_Validators *ValidatorsFilterer) FilterValidatorLockedGoldRequirementsSet(opts *bind.FilterOpts) (*ValidatorsValidatorLockedGoldRequirementsSetIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorLockedGoldRequirementsSet")
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorLockedGoldRequirementsSetIterator{contract: _Validators.contract, event: "ValidatorLockedGoldRequirementsSet", logs: logs, sub: sub}, nil
}

// WatchValidatorLockedGoldRequirementsSet is a free log subscription operation binding the contract event 0x62d947118dd4c1f5ece7f787a9cad4e1127d14d403b71133e95792b473bf8389.
//
// Solidity: event ValidatorLockedGoldRequirementsSet(uint256 value, uint256 duration)
func (_Validators *ValidatorsFilterer) WatchValidatorLockedGoldRequirementsSet(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorLockedGoldRequirementsSet) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorLockedGoldRequirementsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorLockedGoldRequirementsSet)
				if err := _Validators.contract.UnpackLog(event, "ValidatorLockedGoldRequirementsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorLockedGoldRequirementsSet is a log parse operation binding the contract event 0x62d947118dd4c1f5ece7f787a9cad4e1127d14d403b71133e95792b473bf8389.
//
// Solidity: event ValidatorLockedGoldRequirementsSet(uint256 value, uint256 duration)
func (_Validators *ValidatorsFilterer) ParseValidatorLockedGoldRequirementsSet(log types.Log) (*ValidatorsValidatorLockedGoldRequirementsSet, error) {
	event := new(ValidatorsValidatorLockedGoldRequirementsSet)
	if err := _Validators.contract.UnpackLog(event, "ValidatorLockedGoldRequirementsSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the Validators contract.
type ValidatorsValidatorRegisteredIterator struct {
	Event *ValidatorsValidatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorRegistered represents a ValidatorRegistered event raised by the Validators contract.
type ValidatorsValidatorRegistered struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_Validators *ValidatorsFilterer) FilterValidatorRegistered(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsValidatorRegisteredIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorRegisteredIterator{contract: _Validators.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_Validators *ValidatorsFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorRegistered, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorRegistered)
				if err := _Validators.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRegistered is a log parse operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_Validators *ValidatorsFilterer) ParseValidatorRegistered(log types.Log) (*ValidatorsValidatorRegistered, error) {
	event := new(ValidatorsValidatorRegistered)
	if err := _Validators.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorScoreParametersSetIterator is returned from FilterValidatorScoreParametersSet and is used to iterate over the raw logs and unpacked data for ValidatorScoreParametersSet events raised by the Validators contract.
type ValidatorsValidatorScoreParametersSetIterator struct {
	Event *ValidatorsValidatorScoreParametersSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorScoreParametersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorScoreParametersSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorScoreParametersSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorScoreParametersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorScoreParametersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorScoreParametersSet represents a ValidatorScoreParametersSet event raised by the Validators contract.
type ValidatorsValidatorScoreParametersSet struct {
	Exponent        *big.Int
	AdjustmentSpeed *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorScoreParametersSet is a free log retrieval operation binding the contract event 0x4b48724280029c2ea7a445c9cea30838525342e7a9ea9468f630b52e75d6c536.
//
// Solidity: event ValidatorScoreParametersSet(uint256 exponent, uint256 adjustmentSpeed)
func (_Validators *ValidatorsFilterer) FilterValidatorScoreParametersSet(opts *bind.FilterOpts) (*ValidatorsValidatorScoreParametersSetIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorScoreParametersSet")
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorScoreParametersSetIterator{contract: _Validators.contract, event: "ValidatorScoreParametersSet", logs: logs, sub: sub}, nil
}

// WatchValidatorScoreParametersSet is a free log subscription operation binding the contract event 0x4b48724280029c2ea7a445c9cea30838525342e7a9ea9468f630b52e75d6c536.
//
// Solidity: event ValidatorScoreParametersSet(uint256 exponent, uint256 adjustmentSpeed)
func (_Validators *ValidatorsFilterer) WatchValidatorScoreParametersSet(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorScoreParametersSet) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorScoreParametersSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorScoreParametersSet)
				if err := _Validators.contract.UnpackLog(event, "ValidatorScoreParametersSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorScoreParametersSet is a log parse operation binding the contract event 0x4b48724280029c2ea7a445c9cea30838525342e7a9ea9468f630b52e75d6c536.
//
// Solidity: event ValidatorScoreParametersSet(uint256 exponent, uint256 adjustmentSpeed)
func (_Validators *ValidatorsFilterer) ParseValidatorScoreParametersSet(log types.Log) (*ValidatorsValidatorScoreParametersSet, error) {
	event := new(ValidatorsValidatorScoreParametersSet)
	if err := _Validators.contract.UnpackLog(event, "ValidatorScoreParametersSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorScoreUpdatedIterator is returned from FilterValidatorScoreUpdated and is used to iterate over the raw logs and unpacked data for ValidatorScoreUpdated events raised by the Validators contract.
type ValidatorsValidatorScoreUpdatedIterator struct {
	Event *ValidatorsValidatorScoreUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorScoreUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorScoreUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorScoreUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorScoreUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorScoreUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorScoreUpdated represents a ValidatorScoreUpdated event raised by the Validators contract.
type ValidatorsValidatorScoreUpdated struct {
	Validator  common.Address
	Score      *big.Int
	EpochScore *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorScoreUpdated is a free log retrieval operation binding the contract event 0xedf9f87e50e10c533bf3ae7f5a7894ae66c23e6cbbe8773d7765d20ad6f995e9.
//
// Solidity: event ValidatorScoreUpdated(address indexed validator, uint256 score, uint256 epochScore)
func (_Validators *ValidatorsFilterer) FilterValidatorScoreUpdated(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsValidatorScoreUpdatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorScoreUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorScoreUpdatedIterator{contract: _Validators.contract, event: "ValidatorScoreUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorScoreUpdated is a free log subscription operation binding the contract event 0xedf9f87e50e10c533bf3ae7f5a7894ae66c23e6cbbe8773d7765d20ad6f995e9.
//
// Solidity: event ValidatorScoreUpdated(address indexed validator, uint256 score, uint256 epochScore)
func (_Validators *ValidatorsFilterer) WatchValidatorScoreUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorScoreUpdated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorScoreUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorScoreUpdated)
				if err := _Validators.contract.UnpackLog(event, "ValidatorScoreUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorScoreUpdated is a log parse operation binding the contract event 0xedf9f87e50e10c533bf3ae7f5a7894ae66c23e6cbbe8773d7765d20ad6f995e9.
//
// Solidity: event ValidatorScoreUpdated(address indexed validator, uint256 score, uint256 epochScore)
func (_Validators *ValidatorsFilterer) ParseValidatorScoreUpdated(log types.Log) (*ValidatorsValidatorScoreUpdated, error) {
	event := new(ValidatorsValidatorScoreUpdated)
	if err := _Validators.contract.UnpackLog(event, "ValidatorScoreUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
