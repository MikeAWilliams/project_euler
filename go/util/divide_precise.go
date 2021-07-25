package util

type DivisionTermination int

const (
	Terminates DivisionTermination = iota
	Repeats
)

type DivisionResult struct {
	Whole          uint64
	FractionDigits []uint8
	Termination    DivisionTermination
}

func sliceContainsUint(list []uint, target uint) bool {
	for _, val := range list {
		if val == target {
			return true
		}
	}
	return false
}

func Divide(num int, denom int) DivisionResult {
	var result DivisionResult
	uNum := uint64(num)
	uDenom := uint64(denom)
	result.Whole = uNum / uDenom
	remainder := uNum % uDenom
	remainders := []uint{uint(remainder)}
	for 0 != remainder {
		newNum := remainder * 10
		digit := newNum / uDenom
		remainder = newNum % uDenom
		result.FractionDigits = append(result.FractionDigits, uint8(digit))
		if sliceContainsUint(remainders, uint(remainder)) {
			result.Termination = Repeats
			return result
		}
		remainders = append(remainders, uint(remainder))
	}
	return result
}
