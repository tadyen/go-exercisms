package wordy

import (
    "fmt"
    "regexp"
    "strings"
    "strconv"
)

var Operations = map[string]string{
    "plus": "+",
    "minus": "-",
    "multiplied by": "*",
    "divided by": "/",
}

type PendingOperation struct{
    op string
    val int
}

// to be submitted
func Answer(question string) (int, bool) {
    sNumber := "-*\\d+"
    sOperation := fmt.Sprintf(" (?<op>%s) (?<val>%s)", OperationStrings(Operations), sNumber)
    
    reQuestion := regexp.MustCompile(fmt.Sprintf("What is (?<first>%s)(?<next>(%s)*)\\?", sNumber, sOperation))
    if ! reQuestion.Match([]byte(question)){
        return 0, false
    }
    
    matches := reQuestion.FindStringSubmatch(question)
    firstCapture := matches[reQuestion.SubexpIndex("first")]
    first, err := strconv.Atoi(firstCapture)
    if err != nil {
        panic(err)
    }
    if len(matches) == 1 {
        return first, true
    }
    nextCapture := matches[reQuestion.SubexpIndex("next")]
    
    reOperation := regexp.MustCompile(sOperation)
    subMatches := reOperation.FindAllStringSubmatch(nextCapture, -1)
    pendingOperations := []PendingOperation{}
    for _, match := range subMatches{
        op := match[reOperation.SubexpIndex("op")]
        val := match[reOperation.SubexpIndex("val")]
        valInt, err := strconv.Atoi(val)
        if err != nil {
            panic(err)
        }
        pendingOperations = append(pendingOperations, PendingOperation{op: op, val: valInt})
    }
    result, ok := ProcessOperations(first, pendingOperations)
    if ok {
        return result, true
    }
    return 0, false
}

func ProcessOperations(x int, ops []PendingOperation) (int,bool){
    for _,op := range ops{
            switch Operations[op.op]{
            case "+":
                x += op.val
            case "-":
                x -= op.val
            case "*":
                x *= op.val
            case "/":
                x /= op.val
            default:
                return x, false
            }
    }
    return x, true
}

func OperationStrings(operations map[string]string) string {
    var b strings.Builder
    ops := []string{}
    for k := range operations{
        ops = append(ops, k)
    }
    fmt.Fprintf(&b, "%s", strings.Join(ops,"|"))
    return b.String()
}
