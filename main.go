package sort

func MergeSort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func Merge(first, second []int) []int {

	lenFirst, lenSecond := len(first), len(second)
	posFirst, posSecond := 0, 0
	maxFirst, maxSecond := lenFirst-1, lenSecond-1
	s := lenFirst + lenSecond
	merged := make([]int, s, s)

	for k := 0; k < s; k++ {
		if posFirst > maxFirst && posSecond <= maxSecond { //first slice empty
			merged[k] = second[posSecond]
			posSecond++
		} else if posSecond > maxSecond && posFirst <= maxFirst { // second slice empty
			merged[k] = first[posFirst]
			posFirst++
		} else if first[posFirst] < second[posSecond] { // first slice contains min element
			merged[k] = first[posFirst]
			posFirst++
		} else { // second slice contains min element
			merged[k] = second[posSecond]
			posSecond++
		}
	}
	return merged
}
