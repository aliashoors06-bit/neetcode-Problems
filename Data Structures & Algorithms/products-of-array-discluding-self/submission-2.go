func productExceptSelf(nums []int) []int {
prod,zerocount:=1,0

for _,num:=range nums {
	if num!=0 {
		prod*=num
	} else {
		zerocount++
	}
}
res:=make([]int,len(nums))
if zerocount > 1 {
	return res
}
for i, num:= range nums {
	if zerocount > 0 {
		if num==0 {
			res[i]=prod
		} else {
			res[i] = 0
		}
	} else {
		res[i]=prod/num
	}
}
return res
}
