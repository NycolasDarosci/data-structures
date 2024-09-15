package stackandqueue

func ReverseString(text string) string {
	stack := NewStack[string]()

	for _, value := range text {
		char := string(text[value])
		stack.Push(char)
	}

	//newStr := ""
	return ""
	//for stack.Read() {
	//		newStr += stack.Pop()
	//	}
}
