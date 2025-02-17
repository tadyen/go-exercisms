package strain

// test implements the following:
// got := Keep(test.list, test.filterFunc)
// got := Discard(test.list, test.filterFunc)
// test.list is the original list, got is the resulting list, filterFunc is an element-wise test against test.list

// Implement the "Keep" and "Discard" function in this file.

type Collection[T any] []T

type Collectible[T any] interface{
    Keep(compare Collection[T], filter func(x T) bool)      Collection[T]
    Discard(compare Collection[T], filter func(x T) bool)   Collection[T]
}

type filterAction int
const (
    faKeep filterAction = iota
    faDiscard
)

func (c Collection[T]) fFilter(f filterAction) (func (cmp Collection[T], ft func(x T) bool) Collection[T]) {
    return func(compare Collection[T], filter func(x T)bool) Collection[T]{
	    var test func(x T) bool
	    switch f {
	    case faKeep:
	        test = func(x T) bool { return filter(x)} 
	    case faDiscard:
	        test = func(x T) bool { return ! filter(x)} 
	    }
        var out Collection[T]
        for _, v := range compare{
            if test(v){
                out = append(out, v)
            }
        }
        return out
    }
}

func Keep[T any](c Collection[T], filter func(x T) bool) []T{
    return []T(c.fFilter(faKeep)(c, filter))
}

func Discard[T any](c Collection[T], filter func(x T) bool) []T{
    return []T(c.fFilter(faDiscard)(c, filter))
}
