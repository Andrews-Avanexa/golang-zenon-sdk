package crypto

import (
	"math/big"
	"strconv"
)

const ed25519Curve string = "ed25519 seed"
const hardenedOffset int = 0x80000000

// typedef HashFunc = Future<Uint8List> Function(Uint8List? m);
func HashFunc() chan []int {
	H []int 
	return H
}


type KeyData struct {
	key       []int
	chainCode []int
}

func Ret_BigInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

var b int = 256
var q big.Int = *big.NewInt(Ret_BigInt("57896044618658097711785492504343953926634992332820282019728792003956564819949"))
var qm2 big.Int = *big.NewInt(Ret_BigInt("57896044618658097711785492504343953926634992332820282019728792003956564819947"))
var qp3 big.Int = *big.NewInt(Ret_BigInt("57896044618658097711785492504343953926634992332820282019728792003956564819952"))
var l big.Int = *big.NewInt(Ret_BigInt("7237005577332262213973186563042994240857116359379907606001950938285454250989"))
var d big.Int = *big.NewInt(Ret_BigInt("-4513249062541557337682894930092624173785641285191125241628941591882900924598840740"))
var I big.Int = *big.NewInt(Ret_BigInt("19681161376707505956807079304988542015446066515923890162744021073123829784752"))
var by big.Int = *big.NewInt(Ret_BigInt("46316835694926478169428394003475163141307993866256225615783033603165251855960"))
var bx big.Int = *big.NewInt(Ret_BigInt("15112221349535400772501151409588531511454012693041857206046113283949847762202"))
var un big.Int = *big.NewInt(Ret_BigInt("57896044618658097711785492504343953926634992332820282019728792003956564819967"))

var B big.Int = make([]*big.Int, 0)
B = append(B, bx % q , by % q)
var zero big.Int = *big.NewInt(0)
var one big.Int = *big.NewInt(1)
var two big.Int = *big.NewInt(2)
var eight big.Int = *big.NewInt(8)

static Future<Uint8List> hash(HashFunc f, Uint8List? m) {
    return f(m);
  }