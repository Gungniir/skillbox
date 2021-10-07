package main

import "fmt"

func main() {
	fmt.Println(merge([4]int{1, 2, 3, 4}, [5]int{1, 2, 3, 4, 5}))
}

func merge(a [4]int, b [5]int) [9]int {
	var (
		result        [9]int
		aCurrentIndex = 0
		bCurrentIndex = 0
	)

	for i := 0; i < len(result); i++ {
		if aCurrentIndex == len(a) {
			result[i] = b[bCurrentIndex]
			bCurrentIndex++
			continue
		} else if bCurrentIndex == len(b) {
			result[i] = a[aCurrentIndex]
			aCurrentIndex++
			continue
		}

		if a[aCurrentIndex] < b[bCurrentIndex] {
			result[i] = a[aCurrentIndex]
			aCurrentIndex++
		} else {
			result[i] = b[bCurrentIndex]
			bCurrentIndex++
		}
	}

	return result
}
