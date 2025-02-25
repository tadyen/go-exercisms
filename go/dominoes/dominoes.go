package dominoes

/* Discussion 
	for a valid chain, every number must appear at even numbers, ie every number has a neighbour
	ie. any odd count found => invalid
	However does being even => a valid chain must exist? (protip: No)
	counter examples: 
	     Same-sided dominoes :   [1|1]; [2|2] -> invalid chain
	     disjoint chains     :   [1|2], [3|4], [1|2], [3|4] -> invalid

	Can invalid loops be made from valid sets via following-the-tail?
    (no, as long as new members can be added into the loop from anywhere checked)
    
    Lets build the chain and pass/fail it at the end
*/

// Define the Domino type here.

import (
    "slices"
    "fmt"
)

type Domino [2]int

type Chain struct{
    chain []Domino
    looped bool
    stringed bool
}
func NewChain() *Chain{
    return &Chain{
        chain: make([]Domino, 0),
        looped: false,
        stringed: false,
    }
}

// Func to be submitted
func MakeChain(input []Domino) ([]Domino, bool) {
    // any valid chain is ok
    // dominoes can be rotated, but cannot be reused. However duplicates may exist
    fmt.Printf("MakeChain called: Input: %v\n", input)
    
    // trivial, the puzzlemaster wants it to be true
    if len(input) == 0 { return input, true }
    
    c := *NewChain()
    dominoes := make([]Domino, len(input))
    copy(dominoes, input)

    for noSolution := false; (!noSolution) && len(dominoes) > 0; {
        pickingDomino:
        for i, d := range dominoes{
            if c.Insert(d){
                dominoes = slices.Delete(dominoes, i, i+1)
                goto pickingDomino
            }
            noSolution = true
        }
    }
    if len(dominoes) > 0 || (!c.looped) {
        fmt.Printf("No solution\n")
        fmt.Printf("Chain: %v\n", c.chain)
        fmt.Printf("Remainder: %v\n\n", dominoes)
        return c.chain, false
    }

    fmt.Printf("Found solution\n")
    fmt.Printf("Chain: %v\n", c.chain)
    fmt.Printf("Remainder: %v\n\n", dominoes)
    return c.chain, true 
}

func (d Domino) Rotate() Domino {
    p := Domino{}
    p[0], p[1] = d[1], d[0]
    return p
}


// insert domino into chain ensuring chain isnt broken
func (c *Chain) InsertStringed(d Domino) bool {
    if ! c.stringed{
        return false
    }
    _insertSafe := func() bool {
        for i:=0; i<=len(c.chain); i++{
            c.chain = slices.Insert(c.chain, i, d)
            if c.Stringed(){
                return true
            }
            c.chain = slices.Delete(c.chain, i, i+1)
        }
        return false
    }

    if _insertSafe(){
        return true
    }
    d = d.Rotate()
    if _insertSafe(){
        return true
    }
    return false
}

// insert domino into chain ensuring loop isnt broken
func (c *Chain) InsertLooped(d Domino) bool{
    if ! c.looped {
        return false
    }
    oldChain := make([]Domino, len(c.chain))
    copy(oldChain, c.chain)
    _insertSafe := func() bool{
        if c.InsertStringed(d) && c.Looped(0){
            return true
        }
        c.chain = oldChain
        return false
    }

    if _insertSafe(){
        return true
    }
    d = d.Rotate()
    if _insertSafe(){
        return true
    }
    return false
}

// insert domino into chain, prioritising making a loop
func (c *Chain) Insert(d Domino) bool{
    switch len(c.chain){
    case 0:
        c.chain = append(c.chain, d)
    case 1:
        if ! c.InsertStringed(d){
            return false  
        }
    default:
        switch {
        case c.looped:
            if ! c.InsertLooped(d){
                return false
            }
        case c.stringed:
            if ! (c.InsertLooped(d) || c.InsertStringed(d)){
                return false
            }
        default:
            panic("chain is fkd. Unexpected behaviour")
        }
    }
    c.stringed = c.stringed || c.Stringed()
    c.looped = c.looped || c.Looped(0)
    return true
}

func (c *Chain) Looped(offset int) bool{
    if offset < 0 || offset > len(c.chain){
        return false
    }
    switch len(c.chain){
    case 0:
        return false
    case 1: 
        return c.chain[0][0] == c.chain[0][1]
    }

    // check from offset
    left := c.chain[0:offset]
    right := c.chain[offset:len(c.chain)]
    chain := slices.Concat(right, left)
    
    // end-to-end
    if( chain[0][0] != chain[len(chain)-1][1] ){
        return false
    }
    return c.Stringed()
}

func (c *Chain) Stringed() bool {
    switch len(c.chain){
    case 0:
        return false
    case 1:
        return true
    }
    for i:=0; i<len(c.chain)-1; i++{
        if c.chain[i][1] != c.chain[i+1][0] {
            return false
        }
    }
    return true
}
