package allyourbase

import (
    "slices"
    "errors"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

// LittleEndian, ie Leftmost is big. Eg base2{1,0,0,1,0} -> base10{1,8}
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
    if inputBase < 2{
        return nil, errors.New("input base must be >= 2")
    }
    if outputBase < 2{
        return nil, errors.New("output base must be >= 2")
    }
    for _,v := range inputDigits{
        if v >= inputBase || v < 0{
            return nil, errors.New("all digits must satisfy 0 <= d < input base")
        }
    }
    if inputBase == outputBase {
        return inputDigits, nil
    }
    
    // simple but weak method: convert to int, then deconvert. Minimal binary operation magix.
    // will not be able to handle large numbers
    accumulator := 0
    slices.Reverse(inputDigits)     // swap endianness
    for i,d := range inputDigits{
        accumulator += d * PowInt(inputBase, i)
        if accumulator < 0 || accumulator > MaxInt {
            return nil, errors.New("Stack Overflow")
        }
    }
    res := []int{}
    for accumulator > 0 {
        res = append([]int{accumulator % outputBase}, res...)
        accumulator = DivInt(accumulator, outputBase)
    }
    if len(res) == 0{
        res = []int{0}
    }
    return res,nil
}

// returns a ** b 
func PowInt(a,b int) int{
    res := 1
    for i:=0; i<b; i++{
        res *= a
    }
    return res
}

// returns floor(a/b)
func DivInt(a,b int) int{
    return a/b  // ammend if undefined behaviour encountered
}
