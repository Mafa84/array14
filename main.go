package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Old array: ", arr)
	Array14(arr, len(arr))
}

/*
Дан массив A размера N. Вывести вначале его элементы с четными
номерами (в порядке возрастания номеров), а затем — элементы с нечетными номерами (также в порядке возрастания номеров):
A2, A4, A6, . . ., A1, A3, A5, . . . .
*/

func Array14(array []int, n int) {
	var newArr []int
	for i := 1; i < n; i += 2 {
		newArr = append(newArr, array[i])
	}

	for i := 0; i < n; i += 2 {
		newArr = append(newArr, array[i])
	}

	fmt.Println("New array: ", newArr)
}
