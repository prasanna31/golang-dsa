package Stack

import "testing"

// Stack Init Test
func TestStackInit(t *testing.T) {
	stack := InitStack[int]()

	t.Run("Stack size should be empty on Init", func(t *testing.T) {
		if isEmptyOnInit := stack.IsEmpty(); isEmptyOnInit != true {
			t.Error("Stack is not empty on Initialization")
		}
	})

	t.Run("Clear operation should empty the stack", func(t *testing.T) {
		stack.Push(78)
		stack.Push(90)
		stack.Pop()

		if len := stack.Length(); len == 0 {
			t.Errorf("Stack length should not be 0 after pushing elements")
		}

		stack.Clear()
		if shouldBeEmpty := stack.IsEmpty(); shouldBeEmpty != true {
			t.Errorf("Stack is not cleared even after calling Clear() fn")
		}
	})

	t.Run("Pushed elements should be in LIFO order", func(t *testing.T) {
		// clear before inserting values
		stack.Clear()

		stack.Push(1)
		stack.Push(2)
		stack.Push(76)

		if lenAfterPushing3Elements := stack.Length(); lenAfterPushing3Elements != 3 {
			t.Errorf("Elements not pushed to stack properly")
		}

		if topElement, err := stack.Peek(); topElement != 76 || err != nil {
			t.Errorf("Got %d expected %d", topElement, 76)
		}

		stack.Clear()
	})

	t.Run("Stack POP Operation should work as expected", func(t *testing.T) {
		stack.Clear()
		stack.Push(7)
		if val, _ := stack.Pop(); val != 7 {
			t.Errorf("Got %d expected %d. Last Pushed value and Popped values are not same", val, 10)
		}

		stack.Clear()
	})
}
