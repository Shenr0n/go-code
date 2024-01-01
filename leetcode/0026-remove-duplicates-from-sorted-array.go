func removeDuplicates(nums []int) int {

    previous := nums[0]
    unique := 1

    for i := 1; i < len(nums); i++ {
        if nums[i] != previous{
            nums[unique] = nums[i]
            unique+=1
        }
        previous = nums[i]
    }
    return unique
}
