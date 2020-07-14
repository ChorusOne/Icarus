package main

import (
	"github.com/ChorusOne/Icarus/atlas"
	"github.com/ChorusOne/Icarus/snapshot"
	"math/big"
)

func main() {

	x := atlas.GetEpochPaymentEvents(big.NewInt(0), big.NewInt(1000000))
	snapshot.PrettyPrint(x)
}
