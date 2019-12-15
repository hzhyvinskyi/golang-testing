package sort

import "testing"

var elements = []int{5,0,9,2,1,6,4,8,3,7}

func TestBubbleSortASC(t *testing.T) {
	BubbleSort(elements)

	if elements[0] != 0 {
		t.Error("First element must be 0")
	}

	if elements[len(elements) - 1] != 9 {
		t.Error("Last element should be 0")
	}
}

func TestBubbleSortDESC(t *testing.T) {
	BubbleSort(elements, "desc")

	if elements[0] != 0 {
		t.Error("First element must be 0")
	}

	if elements[len(elements) - 1] != 9 {
		t.Error("Last element should be 0")
	}
}