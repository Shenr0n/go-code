type Stack []string

func (st *Stack) IsEmpty() bool {
    if len(*st) == 0 {
        return true
    }
    return false
}

func (st *Stack) Push(str string){
    *st = append(*st, str)
}

func (st *Stack) Pop() string{
    if st.IsEmpty(){
        return ""
    }
    index := len(*st) - 1
    value := (*st)[index]
    *st = (*st)[:index]
    return value
}

func isValid(s string) bool {
    if len(s) == 0 || len(s)%2 != 0{
        return false
    }
    var cs Stack
    for i := 0; i<=len(s)-1; i++ {
        if string(s[i]) == "(" {
            cs.Push(")")
        } else if string(s[i]) == "{" {
            cs.Push("}")
        } else if string(s[i]) == "[" {
            cs.Push("]")
        } else if cs.IsEmpty() || (cs.Pop()) != string(s[i]) {
            return false
        }
    }
    return (len(cs) == 0)
}
