func topKFrequent(nums []int, k int) []int {
count:=make(map[int]int)
for _, num :=range nums{
	count[num]++
}
bucket := make([][]int,len(nums)+1)
for n, c:=range count{
	bucket[c]=append(bucket[c],n)
}
result := []int{}
for i := len(bucket) - 1; i >= 0; i-- {
		for _, n := range bucket[i] {
			result = append(result, n)
			if len(result) == k {
				return result
			}
		}
	}

	return result

}
