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
	return DivisionResult{}
}
