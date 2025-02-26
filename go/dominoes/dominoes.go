package dominoes

import "slices"

type Domino [2]int
type Chain []Domino

// version 2 solution, complete redo
func MakeChain(input Chain) (Chain, bool) {
    _MakeChain := func () (result, remainder Chain, ok bool){
        if len(dominoes) == 0 { return nil, nil, true }
        if dominoes.Looped(){ return dominoes, nil, true }
        return dominoes, nil, false
    }

    panic("")
}

func MergeChains(to,from Chain) (result, remainder Chain, ok bool){
    _MergeChain := func (a,b *Chain) (complete bool){
        switch{
        case len(*a) + len(*b) == 1:
            *a, *b = append(*a,*b...), nil
            return true
        case len(*a) + len(*b) == 0:
            return true
        case len(*b) == 0:
            if a.Chained(){
                return true
            }
            *a,*b = *b, nil
            return false
        case a.Looped() && b.Looped():
            if res, ok := MergeLoops(*a, *b); ok{
                *a, *b = res, nil
                return true
            }
        case a.Chained():
            for i,dom := range *b{
                if res, ok := Insert(*a,dom); ok{
                    *a, *b = res, slices.Delete(*b,i,i+1)
                    return true
                }
            }
        case !a.Chained():
            *a, *b = nil, append(*a,*b...)
            return true
        default:
            return false
        }
        return false
    }
}

func (d Chain) Looped() bool{
    return d.Chained() && ( d[0][0] == d[len(d)][1] )
}

func (d Chain) Chained() bool{
    if len(d) <= 1 { return true }
    chained := true
    for i:=1; i<len(d) && chained; i++{
        chained = chained && ( d[i-1][1] == d[i][0] )
    }
    return chained   
}

func Insert(c Chain, dom Domino) (Chain, bool){
    _insert := func(c Chain, i int, d Domino)(out Chain, looped, chained bool) {
        res := slices.Insert(c,i,d)
        return res, res.Looped(), res.Chained()
    }
    for i := range d{
        if chain, looped, _ := _insert(d,i,dom); looped{
            return chain, true
        }
        if chain, looped, _ := _insert(d,i,dom.Rotate()); looped{
            return chain, true
        }
        if chain, _, chained := _insert(d,i,dom); chained{
            return chain, false
        }
        if chain, _, chained := _insert(d,i,dom.Rotate()); chained{
            return chain, false
        }
    }
    return d, false
}

func MergeLoops(dst,src Chain) (Chain, bool) {
    if !(dst.Looped() && src.Looped()){ return nil, false }
    result := make(Chain, 0, len(dst)+len(src))
    copy(result,dst)
    for i:=0; i<len(src); i++{
        for j:=0; j<len(dst); j++{
            if src[i][0] == dst[j][0]{
                result = slices.Insert(result, i, slices.Concat(dst[j:len(dst)], dst[0:j])... )
                return result, true
            }
        }
    }
    return nil, false
}

func (d Domino) Rotate() Domino{
    return Domino{d[1],d[0]}
}

