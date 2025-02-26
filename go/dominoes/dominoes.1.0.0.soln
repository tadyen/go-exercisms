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
func MakeChain(input []Domino) (result []Domino, ok bool) {

    // trivial, the puzzlemaster wants empty input to result with true
    if len(input) == 0 { return nil, true }

    chains := make([]Chain, 0)
    
    remainder := &input
    for {
        res, rem, ok := OneChain(*remainder)
        if res.looped {
            chains = append(chains, res)
        }
        if ok {
            break
        }
        if len(rem) <= 1 || len(res.chain) == 0{
            // no solution, otherwise remainder wouldve already been fit into the chain
            result, ok := res.chain, false

            return result,ok
        }
        remainder = &rem
    }

    switch len(chains){
    case 0:
        panic("Somehow no chains found when expected at least 1")
    case 1:
        result, ok = chains[0].chain, true
    default:
        ok = true
        for i:=1; i<len(chains) && ok; i++{
            _, ok2 := chains[0].InsertLoopLoop(&chains[i])
            ok = ok && ok2
        }
        result = chains[0].chain      
    }

    return result, ok
}

// Elementary solution
// make a chain using 1 domino at a time. May leave residue that can be further chained up
// ok is true if all inputs are used up AND result is a looped chain, else false
func OneChain(input []Domino) (result Chain, remainder []Domino, ok bool) {
    // any valid chain is ok
    // dominoes can be rotated, but cannot be reused. However duplicates may exist

    
    if len(input) == 0 { return result, nil, false }    // set to false because it's not really a chain
    
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
    
    ok = true
    if len(dominoes) > 0 || (!c.looped) {
        ok = false
    }
    result = c
    remainder = dominoes

    return result, remainder, ok
}

func (d Domino) Rotate() Domino {
    p := Domino{}
    p[0], p[1] = d[1], d[0]
    return p
}


// insert domino into chain ensuring chain isnt broken
func (c *Chain) InsertStringed(d Domino) (where int, ok bool) {
    if ! c.stringed{
        return 0, false
    }
    _insertSafe := func() (w int, o bool) {
        for i:=0; i<=len(c.chain); i++{
            c.chain = slices.Insert(c.chain, i, d)
            if c.Stringed(){
                return i, true
            }
            c.chain = slices.Delete(c.chain, i, i+1)
        }
        return 0, false
    }

    if w, o := _insertSafe(); o{
        return w, true
    }
    d = d.Rotate()
    if w, o := _insertSafe(); o{
        return w, true
    }
    return 0, false
}

// insert domino into chain ensuring loop isnt broken
func (c *Chain) InsertLooped(d Domino) (where int, ok bool){
    if ! c.looped {
        return 0, false
    }
    oldChain := make([]Domino, len(c.chain))
    copy(oldChain, c.chain)
    _insertSafe := func() (w int, o bool){
        if w, o := c.InsertStringed(d); o && c.Looped(0){
            return w, true
        }
        c.chain = oldChain
        return 0, false
    }

    if w, o := _insertSafe(); o{
        return w, true
    }
    d = d.Rotate()
    if w,o := _insertSafe(); o{
        return w, true
    }
    return 0, false
}

// insert a looped chain into another looped chain
func (to *Chain) InsertLoopLoop(from *Chain) (where int, ok bool){
    if ! (to.looped && from.looped) {
        return 0, false
    }
    
    // we need to break the loop and select a value that would fit inside. 
    // This means we have a domino with same values on each side
    // eg. [1|2], [2|3], [3|1] is a chain given by: [1..1], [2...2], or [3...3]
    // due to symmetry, testing just one side of each domino is also testing every side already
    // Additionally, since breaking a loop results in a super 1-value domino, we only need compare the values
    // we do not need to test if the domino fits inside or not, ie no domino rotation required
    for i, dom := range from.chain{
        for j, dom2 := range to.chain{
            if dom[0] == dom2[0] {
                end := len(from.chain)
                stringed := slices.Concat(from.chain[i:end], from.chain[0:i])
                to.chain = slices.Insert(to.chain, j, stringed...) 
                return j, true
            }
        }
    }
    return 0, false
}

// insert domino into chain, prioritising making a loop
func (c *Chain) Insert(d Domino) bool{
    switch len(c.chain){
    case 0:
        c.chain = append(c.chain, d)
    case 1:
        if _,ok := c.InsertStringed(d); !ok {
            return false  
        }
    default:
        switch {
        case c.looped:
            if _,ok := c.InsertLooped(d); !ok {
                return false
            }
        case c.stringed:
            if _,ok := c.InsertLooped(d); !ok {
                if _,ok := c.InsertStringed(d); !ok {
                    return false
                }
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
