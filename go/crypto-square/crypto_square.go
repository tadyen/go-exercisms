package cryptosquare
//package main    // for testing only

import (
    "strings"
    "math"
    "fmt"
)

// to be submitted
func Encode(pt string) string {
    pt = flatten(pt)
    rec := Rectangle(pt)
    rec = rec.Transpose()
    lines := []string{}
    for _, v := range(rec){
        lines = append(lines, string(v))
    }
    var b strings.Builder
    for i, line := range lines{
        b.WriteString(line)
        if i != len(lines)-1 {
            b.WriteString(" ")
        }
    }
    return b.String()
}

func flatten(in string) string {
    var b strings.Builder
    for _, c := range []rune( strings.ToLower(in)){
        switch {
        case 'a' <= c && c <= 'z':
            b.WriteRune(c)
        case '0' <= c && c <= '9':
            b.WriteRune(c)
        }
    }
    return b.String()
}

// get bounding dimension from input length. Want col >= row
func boxSize(strLen int)(row,col int){
    row = int( math.Floor( math.Sqrt( float64(strLen) )))
    col = row
    for ; row * col < strLen; {
        switch {
        case row == col:
            col++
        case row < col:
            row++
        case col < row:
            col++
        }
    }
    return row, col
}

type rectangle[T rune] [][]T

func (r rectangle[T]) Len() (x,y int){
    if len(r) == 0 {
        return 0,0
    }
    return len(r), len(r[0])
}

func (r rectangle[T]) Transpose() rectangle[T] {
    x,y := r.Len()
    var out rectangle[T] = make([][]T, y)
    for i := range out{
        out[i] = make([]T, x)
    }
    x2, y2 := out.Len()
    if(x2 != y || y2  != x){
        panic("transpose result with invalid dimensions")
    }
    for i:=0; i<x; i++{
        for j:=0; j<y; j++{
            out[j][i] = r[i][j]
        }
    }
    return out
}

func (r rectangle[T]) Print() {
    row, col := r.Len()
    var line string
    for i:=0; i<row; i++{
        line = ""
        for j:=0; j<col; j++{
            line += fmt.Sprintf("%s ",string(r[i][j]))
        }
        fmt.Println(line)
    }
}

func Rectangle(s string)(out rectangle[rune]){
    s = flatten(s)
    row,col := boxSize(len(s))
    diff := row * col - len(s)
    for i:=0; i < diff; i++{
        s += " "
    }
    out = make([][]rune, row)
    for i := range out{
        out[i] = make([]rune, col)
    }
    
    for i:=0; i<row; i++{
        for j:=0; j<col; j++{
            out[i][j] = rune(s[i*col + j])
        }
    }
    return out
}

// testing only
/*
func main(){
    raw := "If man was meant to stay on the ground, god would have given us roots."
    expected := "imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn  sseoau "
    result := Encode(raw)
    fmt.Printf("Result: \t %s\n", result)
    fmt.Printf("Expected: \t %s\n", expected)
}
*/
