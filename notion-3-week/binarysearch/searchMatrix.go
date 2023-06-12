package binarysearch

import "fmt"

// time O(logn)
// memory O(1)
func searchMatrix(matrix [][]int, target int) bool {
	// находи середину массивов
	// проверяем первый и последний элемент
	// находим центр след диапазона
	// проверяем первый и последний элемент - если элемент находится в этом диапазоне, то проходимся по элементам массива за logn

	left, right := 0, len(matrix)-1
	for left <= right {
		med := (left + right) / 2
		ar := matrix[med]
		first, end := ar[0], ar[len(ar)-1]
		fmt.Printf("%v , %v \n", first, end)
		if target >= first && target <= end {
			l, r := 0, len(ar)-1
			for l <= r {
				m := (l + r) / 2
				v := ar[m]
				switch {
				case v == target:
					return true
				case v > target:
					r = m - 1
				case v < target:
					l = m + 1
				}
			}
			return false
		} else {
			if end > target {
				right = med - 1
			} else {
				left = med + 1
			}
		}

	}

	return false
}
