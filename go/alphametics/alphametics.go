package alphametics

import (
	"fmt"
	"regexp"
	"strings"
)

func Solve(puzzle string) (result map[string]int, e error) {
	// only + operators exist, and the final word is seperated by '=' or '=='
	// Do not bother satisfying other operators (eg -, *, /)
	// TODO: assert input matches appropriate format

	// extract words (ignoring operators)
	words := regexp.MustCompile(`\(\w+\)`).FindAllString(puzzle, -1)
	for i := range words {
		words[i] = strings.ToUpper(words[i])
	}

	// sanity: final word must be the longest
	maxWordLen := len(words[len(words)-1])
	for _, w := range words[0 : len(words)-1] {
		if maxWordLen < len(w) {
			e = fmt.Errorf("final word is not the longest")
			return
		}
	}

	// ensure 10(base) unique symbols(digits) or less
	// generate digits set
	var base int = 10 // generalises the problem/soln

	// Set of digits mapping symbols to value
	digits := map[rune]int{}

	// List of symbols ordered from final word first
	symbolList := make([]rune, len(digits))

	// allocating digits and symbolList
	// final word first
	for _, word := range append(words[len(words)-1:], words[:len(words)-1]...) {
		for _, symbol := range word {
			if _, found := digits[symbol]; !found {
				symbolList = append(symbolList, symbol)
				digits[symbol] = 0
			}
		}
	}

	// sanity check: more symbols than possible digits
	if len(digits) > base {
		e = fmt.Errorf("invalid number of symbols. Found %d symbols when expected %d only", len(digits), base)
		return
	}

	// left-pad words to maxWordLen
	var pad rune = '-'
	for i, w := range words {
		words[i] = strings.Repeat(string(pad), maxWordLen-len(w)) + w
	}

	// count occurences per symbol per radix on left-side of equation
	digitCount := make([]map[rune]int, maxWordLen)
	for i := range digitCount {
		digitCount[i] = map[rune]int{}
	}
	for i := 0; i < maxWordLen; i++ {
		for _, word := range words[:len(words)-1] {
			digitCount[i][rune(word[i])] += 1
		}
	}

	// weights for each digit, ie 1,10,100, etc. in base10
	radixVect := make([]int, maxWordLen)
	radixVect[0] = 1
	for i := range radixVect[1:] {
		radixVect[i] = radixVect[i-1] * base
	}
	// reverse radixVect so that v[0] has most significant value
	for i := 0; i < (len(radixVect)&^0x1)/2; i++ {
		j := len(radixVect) - 1 - i
		temp := radixVect[i]
		radixVect[i] = radixVect[j]
		radixVect[j] = temp
	}

	// Weighted count of each symbol eg. (ABAB + AB)=>{A:1020, B:102}
	leftSymbolWeightedSum := func() (wsum map[rune]int) {
		wsum = map[rune]int{}
		for symbol := range digits {
			wsum[symbol] = 0
			for i := range digitCount {
				wsum[symbol] += digitCount[i][symbol] * radixVect[i]
			}
		}
		return wsum
	}()
	rightSymbolWeightedSum := func() (wsum map[rune]int) {
		wsum = map[rune]int{}
		for i, symbol := range words[len(words)-1] {
			wsum[symbol] += radixVect[i]
		}
		return wsum
	}()

	// calc Sums
	sumOpts := struct{ left, right string }{"left", "right"}
	sum := func(opt string) (sum int) {
		sum = 0
		var targ *map[rune]int
		switch opt {
		case sumOpts.left:
			targ = &leftSymbolWeightedSum
		case sumOpts.right:
			targ = &rightSymbolWeightedSum
		}
		for symbol, value := range digits {
			sum += (*targ)[symbol] * value
		}
		return sum
	}

	// checks
	isSumsEqual := sum(sumOpts.left) == sum(sumOpts.right)
	isLeadingZero := digits[rune(words[len(words)-1][0])] == 0
	isUniqueDigits := func() (result bool) {
		for i := 0; i < len(symbolList)-1; i++ {
			s := symbolList[i+1:]
			for _, digit := range s {
				if symbolList[i] == digit {
					return false
				}
			}
		}
		return true
	}

	if !isSumsEqual {
		e = fmt.Errorf("sums do not add up, no solution found")
		return
	}
	if !isUniqueDigits() {
		e = fmt.Errorf("symbols do not have unique values. No solution found")
		return
	}
	if isLeadingZero {
		e = fmt.Errorf("leading symbol has value 0. Invalid solution")
		return
	}
	return result, nil
}
