package flatten

func Flatten(nested interface{}) []interface{} {
    flat := []interface{}{}

    switch t := nested.(type){
    case []interface{}:
        for _, v := range t{
            flat = append(flat, Flatten(v)...)
        }
    case interface{}:
        flat = append(flat, t)
    }
    return flat
}

