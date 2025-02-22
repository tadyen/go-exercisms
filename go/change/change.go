package change

import (
    "slices"
    "errors"
    "fmt"
)

/* Discussion/Examples:
Rules:
    - Panic on negative target
    - remove negative coins
    - coins are unlimited per type
    - Panic if no solution found
Least coins tricky cases:
    (1){ target, coins := 25, [10, 8, 7, 3, 2]
        naive: [10, 10, 3, 2] (greedy algorithm)
        actual: [10, 8, 7]
    }
    (2){ target, coins := 24, [10,8,2]
        incremental search: [10,10,2,2] or [10,8,2,2,2]
        actual: [8,8,8]
    }

*/

func AnyChange(coins []int, target int) ([]int, error){
    // assume coins is sorted largest to smallest, inputs already sanitised for negatives
    // TODO implement sanitisation
    
}

// This is the function to be submitted
func Change(coins []int, target int) ([]int, error) {
    fmt.Println("asdf")
    var coinbination = make( map[int]int )    // coin[value]amount
    result := []int{}
    
    //edge case asserts
    if target < 0 {
        return nil, errors.New("negative value")
    }
    if target == 0 {
        return result, nil     // Result == nil
    }
    
    // create working list of coins, sorted largest first
    largestFirst := make([]int, len(coins))
    copy(largestFirst, coins)
    slices.Sort(largestFirst)
    slices.Reverse(largestFirst)
    
    // populate map of coin -> amount
    for _,coin := range largestFirst {
        if coin > 0 {
            coinbination[coin] = 0
        }
    }
    
    // recursive function to iterate over each coin
    // currCoin 
    coinbine := func(currCoin int){
        
    }


/*
    rem := target
    for{
        coin := largestFirst[i]
        quotient := rem / coin
        rem = rem % coin
        
        
        
        // Target reached
        if rem == 0 {
            break;
        }
    } 
    */
    return result, nil
}



