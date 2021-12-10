package stack

type Item interface{}

type Stack struct {
	items []Item
}

func NewStack() *Stack {
	return &Stack{
		items: nil,
	}
}

func NewStackWithData(items []Item) *Stack {
	return &Stack{
		items: items,
	}
}

func (stack *Stack) Push(item Item) {
	stack.items = append(stack.items, item)
}

func (stack *Stack) Pop() Item {
	if len(stack.items) == 0 {
		return nil
	}

	lastItem := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]

	return lastItem
}

func (stack *Stack) Top() Item {
	if len(stack.items) == 0 {
		return nil
	}

	return stack.items[len(stack.items)-1]
}

func (stack *Stack) IsEmpty() bool {
	if len(stack.items) == 0 {
		return true
	}

	return false
}

func (stack *Stack) Size() int {
	return len(stack.items)
}

func (stack *Stack) Clear() {
	stack.items = nil
}
