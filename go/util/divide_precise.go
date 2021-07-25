package util

type DivisionTermination int

const (
	Terminates DivisionTermination = iota
	Repeats
	Abort
)

type DivisionResult struct {
	Whole          uint64
	FractionDigits []uint8
	Termination    DivisionTermination
}

func Divide(num int, denom int) DivisionResult {
	var result DivisionResult
	uNum := uint64(num)
	uDenom := uint64(denom)
	result.Whole = uNum / uDenom
	remainder := uNum % uDenom
	for 0 != remainder {
		newNum := remainder * 10
		digit := newNum / uDenom
		remainder = newNum % uDenom
		result.FractionDigits = append(result.FractionDigits, uint8(digit))
	}
	return result
}
