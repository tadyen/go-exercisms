package dominoes

import "slices"

type Domino [2]int
type Chain []Domino

// version 2 solution, complete redo
// I thought of making it a compact recursive solution initially, but I am keeping it closer to my first iteration instead
//  where it is much more verbose and logically split out, with more robustness.
// I favour my ability to read and understand what is happening, over some black magic recursion

/*
    0.  A loop is a superset of a chain. We actually want to make loops.
    1.  A loop may be broken down into sub-loops.
	2.  A set of dominoes that forms a loop, will *always* form a loop, or set of sub-loops when chained.
	3.  This means that if a loop-forming set is filtered into a loop, and a remainder, the remainder MUST be a loop, or set of sub-loops.
	    If at any point in recursion a remainder excludes being a loop, then the parent set also cannot form a loop.
	4.  This means the break-condition is when a set is split into a result & remainder, the result is nil, but there is a non-nil remainder.
*/

// fn to be submitted
func MakeChain(input Chain) (Chain, bool) {
    if len(input) == 0 { return nil, true }
    if input.Looped() { return input, true }
    res, rem := MergeChains(Chain{}, input)

    if len(rem) == 0 {
        return res, res.Looped()    // exhausted set
    }
    if len(res) == 0 && len(rem) != 0 {
        return nil, false           // break condition
    }

    remChain, ok := MakeChain(rem)
    if !ok {
        return nil, false
    }
    res, rem = MergeChains(res, remChain)
    
    return res, res.Looped() && len(rem) == 0
}

// merge dominoes from 'b' into 'a'. 
// 'a' must result in a chain or loop, and 'b' are the remainder dominoes. 
// It may move dominoes out from 'a' into 'b' to satisfy the above.
func MergeChains(to,from Chain) (result, remainder Chain){
    newTo := append(Chain{}, to...)
    newFrom := append(Chain{}, from...)
    
    _MergeChain := func (a,b *Chain) (){
        switch{
        case !a.Chained():  // cannot insert into non-chain
            *a, *b = nil, append(*a,*b...)
        case len(*a) + len(*b) == 1:    // only 1 domino
            *a, *b = append(*a,*b...), nil
        case len(*b) == 0:              // no inputs to merge
            break;
        case a.Looped() && b.Looped():
            if res, ok := MergeLoops(*a, *b); ok{
                *a, *b = res, nil
                break;
            }
        case a.Chained():   // insert dominos from 'b' into 'a'
            Pick:
            for idx, dom := range *b{
                res, ok := Insert(*a, dom)
                if ok{
                    *a = res
                    *b = slices.Delete(*b, idx, idx+1)
                    goto Pick
                }
            }
        }
        return
    }
    _MergeChain(&newTo, &newFrom)
    return newTo, newFrom
}


// test Chain if it's chained from start to end of slice.
func (d Chain) Chained() bool{
    if len(d) <= 1 { return true }
    chained := true
    for i:=1; i<len(d) && chained; i++{
        chained = chained && ( d[i-1][1] == d[i][0] )
    }
    return chained   
}

// Test chain if it's looped, ie a chain that chains on its own ends too.
func (d Chain) Looped() bool{
    if len(d) == 0 {
        return false
    }
    return d.Chained() && ( d[0][0] == d[len(d)-1][1] )
}

// Insert domino into a chain, prioritising making a loop instead of a chain
func Insert(c Chain, dom Domino) (result Chain, ok bool){
    if len(c) == 0{
        return Chain{dom}, true
    }
    _Insert := func(c Chain, i int, d Domino) Chain {
        return append(Chain{}, slices.Insert(c,i,d)...)
    }
    for i := range c{
        if result := _Insert(c,i,dom); result.Looped(){
            return result, true
        }
        if result := _Insert(c,i,dom.Rotate()); result.Looped(){
            return result, true
        }
        if result := _Insert(c,i,dom); result.Chained(){
            return result, true
        }
        if result := _Insert(c,i,dom.Rotate()); result.Chained(){
            return result, true
        }
    }
    return c, false
}

// Merge loops together, do not need to compare domino sides. Loops are a super domino-chain with both ends sharing same value.
func MergeLoops(dst,src Chain) (Chain, bool) {
    if !(dst.Looped() && src.Looped()){ return nil, false }
    result := make(Chain, 0, len(dst)+len(src))
    copy(result,dst)
    for i:=0; i<len(src); i++{
        for j:=0; j<len(dst); j++{
            if src[i][0] == dst[j][0]{
                result = slices.Insert(result, i, slices.Concat(dst[j:], dst[0:j])... )
                return result, true
            }
        }
    }
    return nil, false
}

func (d Domino) Rotate() Domino{
    return Domino{d[1],d[0]}
}

