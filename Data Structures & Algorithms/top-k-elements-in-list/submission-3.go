func topKFrequent(nums []int, k int) []int {
count:=make(map[int]int)
for _, num :=range nums{
	count[num]++
}
keys := make([]int, 0, len(count))
for num := range count {
        keys = append(keys, num)
    }
for i:=0 ; i<len(keys)-1 ; i++ {
	for j:=0 ; j<len(keys)-1-i;j++{
		if count[keys[j]] < count[keys[j+1]] {
            keys[j], keys[j+1] = keys[j+1], keys[j]
        }
	}
}
return keys[:k]
}
