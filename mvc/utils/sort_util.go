package utils

import "sort"

func BubbleSort(elements []int) {
	for {
		swap := false
		for i := 1; i < len(elements); i++ {
			if elements[i-1] > elements[i] {
				elements[i-1], elements[i] = elements[i], elements[i-1]
				swap = true
			}
		}
		if !swap {
			return
		}
	}
}

func Sort(els []int) {
	if len(els) < 1000 {
		BubbleSort(els)
	} else {
		sort.Ints(els)
	}
}
