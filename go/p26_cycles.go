package main

import (
	"fmt"
	"maw/util"
	"sort"
)

type fractionPair struct {
	d        int
	quotient util.DivisionResult
}

type fractionPairs []fractionPair

func (fp fractionPairs) Len() int {
	return len(fp)
}

func (fp fractionPairs) Less(i, j int) bool {
	return len(fp[i].quotient.FractionDigits) < len(fp[j].quotient.FractionDigits)
}

func (fp fractionPairs) Swap(i, j int) {
	fp[i], fp[j] = fp[j], fp[i]
}

func getQuotients(max int) fractionPairs {
	quotients := make([]fractionPair, max)
	for d := 2; d < max; d++ {
		quotients[d] = fractionPair{d: d, quotient: util.Divide(1, d)}
	}
	return quotients
}

func getRepeatingQuotients(input fractionPairs) fractionPairs {
	var result []fractionPair
	for _, q := range input {
		if util.Repeats == q.quotient.Termination {
			result = append(result, q)
		}
	}
	return result
}

func main() {
	quotients := getQuotients(1000)
	repeatingQuotients := getRepeatingQuotients(quotients)
	fmt.Printf("number of repeating fractions %v\n", len(repeatingQuotients))
	sort.Sort(repeatingQuotients)
	last := repeatingQuotients[len(repeatingQuotients)-1]
	fmt.Printf("denominator of longest %v, with length of %v\n", last.d, len(last.quotient.FractionDigits))
}
