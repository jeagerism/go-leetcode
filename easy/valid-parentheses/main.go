// ### **โจทย์ 5: Valid Parentheses**
// **คำอธิบาย:**
// ให้ตรวจสอบว่าสตริงที่ประกอบด้วยวงเล็บ `()`, `{}`, `[]` นั้นสมดุลหรือไม่
// **ตัวอย่าง Input:**
// "()"

// **Output ที่คาดหวัง:**
// true

// **ตัวอย่าง Input:**
// "([)]"

// **Output ที่คาดหวัง:**
// false

// **Tips:**
// - ใช้ stack ในการจัดการวงเล็บเปิดและปิด

package main

import "fmt"

func main() {
	input := "()"
	fmt.Println(valid(input)) // คาดหวังผลลัพธ์: true

	input = "([)]"
	fmt.Println(valid(input)) // คาดหวังผลลัพธ์: false

	input = "{[()]}"
	fmt.Println(valid(input)) // คาดหวังผลลัพธ์: true
}

func valid(input string) bool {
	stack := []rune{}

	pair := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range input {
		if open, isClose := pair[char]; isClose {
			if len(stack) > 0 && stack[len(stack)-1] == open {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
