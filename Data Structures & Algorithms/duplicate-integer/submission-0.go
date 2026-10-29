func hasDuplicate(nums []int) bool {
found:=map[int]bool{}
for _, num:= range nums {
	if found[num] {
		return true
	}
	found[num]=true 
}
return false
}
