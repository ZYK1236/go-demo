package swap

func Swap(pa *int, pb *int) {
	var temp = *pa
	*pa = *pb
	*pb = temp
}
