//Declaring a type Key to store the frequency of characters in a string, as the key

type Key [26] byte

//Fucntion to convert string to the key value so that it can be appended to the map and eventually the result to be returned
func stringToKey(str string) (key Key) {
    for i:= range str {
        key[str[i]-'a']++
    }
    return
}


func groupAnagrams(strs []string) [][]string {
    strMap := make(map[Key][]string)

    for _,strv := range strs {
        //Get the key of the value (string) from the array
        tempKey:= stringToKey(strv)

        //Append the strv (string value) to the corresponding key in the strMap
        //If key is not present, a new key value gets added, otherwise just appended
        strMap[tempKey] = append(strMap[tempKey], strv)
    }

    //2d array to store the anagrams grouped together using the keys and values in the map
    //Initial length is 0 and capacity is len(strMap)
    result := make([][]string, 0, len(strMap))

    for _,val := range strMap {
        result = append(result, val)
    }
    return result
}
