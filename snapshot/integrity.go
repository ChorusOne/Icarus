package snapshot

import (
	"log"
	"math/big"

	"github.com/oleiade/reflections"
)

var integrityTests = [](func(Snapshot, SystemSnapshot) bool){

	totalLockedGoldBalanceIntegrity,
	nonVotingLockedGoldBalanceIntegrity,
	goldTokenSupplyIntegrity,
	totalCeloUSDValueIntegrity,
}

var lockedGoldToleranceThreshold = big.NewInt(10000)

const (
	pass = "PASS"
	fail = "FAIL"
)

func CheckIntegrity(as Snapshot, ss SystemSnapshot) bool {

	for _, testFunction := range integrityTests {
		testFunction(as, ss)
	}
	return true
}

func totalLockedGoldBalanceIntegrity(as Snapshot, ss SystemSnapshot) bool {

	sumOf := snapshotColumnSum(as, "LockedGoldBalance")
	systemMetric := ss.TotalLockedGoldBalance

	message := "Integrity Test - TotalLockedGoldBalance - %s - Summation = %s, SystemWideValue = %s  : Delta = %s, Threshold = %s \n"

	//Because of precision loss incurred during conversions in LockedGold.sol, this test
	//accomodates for the same with some tolerance built into the test.

	//ColumnSum is allowed to differ from the system metric by atmost`lockedGoldToleranceThreshold` units
	//Beyond this tolerance threshold, the test fails

	delta := big.NewInt(0)
	delta = delta.Sub(sumOf, systemMetric)
	delta = delta.Abs(delta)

	// If Delta > Threshold
	if delta.Cmp(lockedGoldToleranceThreshold) > 0 {

		log.Printf(message, fail, sumOf, systemMetric, delta, lockedGoldToleranceThreshold)
		log.Printf("Delta is Greater than Tolerance Threshold for LockedGold")
		return false

	}

	log.Printf(message, pass, sumOf, systemMetric, delta, lockedGoldToleranceThreshold)
	return true

}

func nonVotingLockedGoldBalanceIntegrity(as Snapshot, ss SystemSnapshot) bool {

	sumOf := snapshotColumnSum(as, "NonVotingLockedGoldBalance")
	systemMetric := ss.NonVotingLockedGoldBalance
	message := "Integrity Test - NonVotingLockedGoldBalance - %s - Summation = %s, SystemWideValue = %s \n"

	if sumOf.Cmp(systemMetric) != 0 {

		log.Printf(message, fail, sumOf, systemMetric)
		log.Printf("Delta SystemWideValue - Summation = %s", big.NewInt(0).Sub(systemMetric, sumOf))
		return false
	}

	log.Printf(message, pass, sumOf, systemMetric)
	return true
}

func goldTokenSupplyIntegrity(as Snapshot, ss SystemSnapshot) bool {
	sumOf := snapshotColumnSum(as, "GoldTokenBalance")
	systemMetric := ss.GoldTokenSupply
	message := "Integrity Test - GoldTokenSupply - %s - Summation = %s, SystemWideValue = %s \n"

	if sumOf.Cmp(systemMetric) != 0 {
		log.Printf(message, fail, sumOf, systemMetric)
		log.Printf("Delta SystemWideValue - Summation = %s", big.NewInt(0).Sub(systemMetric, sumOf))
		return false
	}

	log.Printf(message, pass, sumOf, systemMetric)
	return true

}

func totalCeloUSDValueIntegrity(as Snapshot, ss SystemSnapshot) bool {
	sumOf := snapshotColumnSum(as, "CeloUSDValue")
	systemMetric := ss.TotalCeloUSDValue
	message := "Integrity Test - TotalCeloUSDValue - %s - Summation = %s, SystemWideValue = %s \n"

	if sumOf.Cmp(systemMetric) != 0 {
		log.Printf(message, fail, sumOf, systemMetric)
		log.Printf("Delta SystemWideValue - Summation = %s", big.NewInt(0).Sub(systemMetric, sumOf))
		return false
	}

	log.Printf(message, pass, sumOf, systemMetric)
	return true

}

func snapshotColumnSum(as Snapshot, columnName string) *big.Int {
	sum := big.NewInt(0)
	for address := range as {
		val, _ := reflections.GetField(as[address], columnName)
		//Asserting val as *big.Int
		sum.Add(sum, val.(*big.Int))
	}
	return sum
}
