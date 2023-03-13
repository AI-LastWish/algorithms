func containsDuplicate(nums []int) bool {
    numsMap := map[int]int{}
    for _, n := range nums {
        if _, ok := numsMap[n]; ok {
            return true
        } else {
            numsMap[n] = 1
        }
    }
    return false
}