package datastructures

import "testing"

func TestEmptyStackLength(t *testing.T) {
	// Tests the length of empty stack
	testStack := stack{}

	if testStack.getStackLength() != 0 {
		t.Error("Length of empty stack should be 0.")
	}
}

func TestEmptyStackTop(t *testing.T) {
	// Tests the top item of empty stack
	testStack := stack{}

	_, err := testStack.getTopItem()
	if err == nil {
		t.Error("An error is expected when getting the top item of an empty stack. But received none.")
	}
}

func TestEmptyStackPop(t *testing.T) {
	// Tests the pop operation on empty stack
	testStack := stack{}

	_, err := testStack.popItem()
	if err == nil {
		t.Error("An error is expected when popping an empty stack. But received none.")
	}
}

func TestStackLength(t *testing.T) {
	testStack := stack{
		items: []int{3, -5, 1, 7, -1, 2345},
	}

	if testStack.getStackLength() != 6 {
		t.Errorf("Expected length is 6, but got %d", testStack.getStackLength())
	}
}

func TestStackTop(t *testing.T) {
	testStack := stack{
		items: []int{3, -5, 1, 7, -1, 2345},
	}

	topItem, err := testStack.getTopItem()

	if err != nil {
		t.Error("Not expected an error while getting the top item of non-empty stack.")
	}

	if topItem != 2345 {
		t.Errorf("Expected 2345 to be top item, but received %d instead.", topItem)
	}
}

func TestStackPop(t *testing.T) {
	testStack := stack{
		items: []int{3, -5, 1, 7, -1, 2345},
	}

	poppedItem, err := testStack.popItem()

	if err != nil {
		t.Error("Not expected an error while popping an non-empty stack.")
	}

	if poppedItem != 2345 {
		t.Errorf("Expected 2345 to be the popped item, but popped %d instead", poppedItem)
	}
}

func TestStackPush(t *testing.T) {
	testStack := stack{
		items: []int{3, -5, 1, 7, -1, 2345},
	}

	newLength := testStack.pushItem(-103)

	if newLength != 7 {
		t.Errorf("Expected new length to be 7, but got new length as %d instead", newLength)
	}
}

func TestStackSearch(t *testing.T) {
	testStack := stack{
		items: []int{3, -5, 7, -1, 2345},
	}

	// Printing just for visibility
	testStack.printStack()

	itemIndex := testStack.searchItem(7)
	if itemIndex != 2 {
		t.Errorf("Expected index 2 for item 7, but got %d instead.", itemIndex)
	}

	itemIndex = testStack.searchItem(0)
	if itemIndex != -1 {
		t.Errorf("Expected index -1 for item 0 which doesn't exist, got %d instead", itemIndex)
	}
}
