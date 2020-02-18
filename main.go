package firstHomework

//Подъехали первые задачки, которые надо выполнить в рамках первого модуля:
//1. Написать функцию, которая выполняет циклический сдвиг в слайсе на N элементов.
//2. Написать функцию, которая разворачивает слайс.
//Если вы напишите тесты к этому делу, то мир станет чуть лучше.
//Удобнее всего код мне сбрасывать ссылкой на персональную репу в гитлабе.

func rotate(slice []int, n int) []int {
	if n <= 0 || len(slice) == 0 {
		return slice
	}
	pos := len(slice) - n%len(slice)
	slice = append(slice[pos:], slice[:pos]...)
	return slice
}

func revert(slice []int) []int {
	n := len(slice)
	if n == 0 {
		return slice
	}
	var res = make([]int, 0, n)
	pos := n - 1
	for i := range slice {
		res = append(res, slice[pos-i])
	}
	return res
}
