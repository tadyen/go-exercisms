package matrix

import ( 
    "strings"
    "strconv"
    "errors"
)

type Number interface {
    ~int | ~float64
}

type GenericMatrix[T Number] [][]T

type Matrix = GenericMatrix[int]

func New(s string) (GenericMatrix[int], error) {
    lines := strings.Split(s, "\n")
    for i, v := range lines{
        lines[i] = strings.TrimSpace(v)
    }
    m := GenericMatrix[int]{}
    if len(lines) == 0 { return nil, errors.New("invalid string format")}
    for _, line := range lines{
        row := []int{}
        for _, v := range strings.Split(line, " "){
            num, err := Int(v)
            if err != nil { return nil, err }
            row = append(row, num)
        }
        m = append(m, row)
    }
    x,y := m.Size()
    if x == 0 || y == 0 { 
        return nil, errors.New("could not create matrix from string")
    }
    return m, nil
}

func Int(s string)(int, error) {
    i, err := strconv.ParseInt(s, 0, 0)
    return int(i), err
}

func (m GenericMatrix[T]) Size() (row, col int){
    if len(m) == 0 || len(m[0]) == 0 {
        return 0,0
    }
    // ensure same size in all rows
    size := len(m[0])
    for _, row := range m{
        if len(row) != size{
            return 0, 0
        }
    }
    return len(m), len(m[0])
}

// Cols and Rows must return the results without affecting the matrix.
func (m GenericMatrix[T]) Cols() [][]T {
    // Cols() is equivalent of transpose
    t := GenericMatrix[T]{}
    x, y := len(m), len(m[0])
    for j:=0; j<y; j++{
        t = append(t, make( []T, x ))
    }
    for i:=0; i<x; i++{
        for j:=0; j<y; j++{
            t[j][i] = m[i][j]
        }
    }
    return t
}

func (m GenericMatrix[T]) Rows() [][]T {
    // Rows() simply makes a deepcopy
    ret := GenericMatrix[T]{}
    x,y := m.Size()
    for i:=0; i<x; i++{
        ret = append(ret, make( []T, y ))
    }
    for i:=0; i<x; i++ {
        for j:=0; j<y; j++{
            ret[i][j] = m[i][j]
        }
    }
    return ret
}

func (m GenericMatrix[T]) Set(row, col int, val T) bool {
    x,y := m.Size()
    if row < 0 || row >= x { return false }
    if col < 0 || col >= y { return false }
    m[row][col] = val   // didnt think this works but it does cus its not a pointer
    return true
}
