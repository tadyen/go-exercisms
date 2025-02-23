package cipher

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

import "strings"

type charShift rune
type intShift int
type vigenere string

// shifts rune [a-zA-Z] preserving case, returns a blank otherwise
func shiftByInt(c rune, s int)(ret rune){
    var base, blank rune
    blank = 0
    cycle := 'z' - 'a' + 1
    switch {
    case 'a' <= c && c <= 'z':
        base = 'a'
    case 'A' <= c && c <= 'Z':
        base = 'A'
    default:
        base = blank   // invalid alpha
    }
    if base == blank { return blank }
    ret = base + (cycle + (c - base) + rune(s)) % cycle
    return ret
}

type direction int
const (
    increase direction = iota
    decrease
)

func shiftByRune(c rune, s rune, d direction)(ret rune){
    var shiftBy int 
    switch {
    case 'a' <= s && s <= 'z':
        shiftBy = int(s-'a')
    case 'A' <= s && s <= 'Z':
        shiftBy = int(s-'A')
    default:
        shiftBy = 0
    }

    switch d {
    case increase:
        return shiftByInt(c, shiftBy)
    case decrease:
        return shiftByInt(c, -shiftBy)
    default:
        return 0    // empty value instead of original rune
    }
}

func cleanString(input string) string{
    var b strings.Builder
    for _, c := range []rune(strings.ToLower(input)){
        if 'a' <= c && c <= 'z' {
            b.WriteRune(c)
        }
    }
    return b.String()
}

// Caesar is just Shift{a->d} ie +3
func NewCaesar() Cipher {
    return NewShift(3)
}

func NewShift(distance int) Cipher {
    if distance < -25 || distance > 25 || distance == 0 {
        return nil
    }
    return intShift(distance)
}

func (c charShift) Encode(input string) string{
    shifted := []rune(cleanString(input))
    for i, v := range shifted{
        shifted[i] = shiftByRune(v, rune(c), increase)
    }
    return cleanString(string(shifted))
}

func (c charShift) Decode(input string) string{
    shifted := []rune(cleanString(input))
    for i, v := range shifted{
        shifted[i] = shiftByRune(v, rune(c), decrease)
    }
    return cleanString(string(shifted))
}

func (x intShift) Encode(input string) string{
    shifted := []rune(cleanString(input))
    for i, v := range shifted{
        shifted[i] = shiftByInt(v, int(x))
    }
    return cleanString(string(shifted))
}

func (x intShift) Decode(input string) string{
    shifted := []rune(cleanString(input))
    for i, v := range shifted{
        shifted[i] = shiftByInt(v,-int(x))
    }
    return cleanString(string(shifted))
}

func NewVigenere(key string) Cipher {
    onlya := true
    for _,c := range []rune(key){
        if c < 'a' || c > 'z' {
            return nil
        }
        if c != 'a'{
            onlya = false
        }
    }
    if onlya { return nil }
    return vigenere(key)
}

func (v vigenere) Encode(input string) string {
    shifted := []rune(cleanString(input))
    for i, r := range shifted{
        shifted[i] = shiftByRune(r, []rune(v)[i % len(v)], increase)
    }
    return cleanString(string(shifted))
}

func (v vigenere) Decode(input string) string {
    shifted := []rune(cleanString(input))
    for i, r := range shifted{
        shifted[i] = shiftByRune(r, []rune(v)[i % len(v)], decrease)
    }
    return cleanString(string(shifted))
}
