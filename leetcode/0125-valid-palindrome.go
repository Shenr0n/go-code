//ASCII solution

func toLowerCase(character byte) byte {
    if character >=65 && character<=90 {
        return character + 32
    }
    return character
}


func isNonAlpha(c byte) bool {
    if (c>=32 && c<=47) || (c>=58 && c <=64) || (c>= 91 && c<=96) || (c>=123 && c<=126){
        return true
    }
    return false
}

func isPalindrome(s string) bool {
    left:=0
    right:=len(s)-1

    for left<=right{
        if isNonAlpha(s[left]){
            left++
            continue
        }

        if isNonAlpha(s[right]){
            right--
            continue
        }

        if toLowerCase(s[left]) != toLowerCase(s[right]){
            return false
        } else {
            left++
            right--
        }
    }
    return true
}
