// ### **โจทย์ 3: Palindrome Number**

// **คำอธิบาย:**
// เขียนฟังก์ชันที่ตรวจสอบว่าตัวเลขที่กำหนดเป็น Palindrome หรือไม่ (ตัวเลขที่อ่านจากซ้ายไปขวาและขวาไปซ้ายเหมือนกัน)
// **ตัวอย่าง Input:**
// 121

// **Output ที่คาดหวัง:**
// true

// **ตัวอย่าง Input:**
// -121

// **Output ที่คาดหวัง:**

// false

// **Tips:**

// - แปลงตัวเลขเป็นสตริงแล้วเปรียบเทียบสตริงย้อนกลับ
// - ห้ามใช้สตริง? ใช้คณิตศาสตร์ในการหาค่ากลับด้าน

package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := 543
	fmt.Printf("this number is Palindrom is : %v", strMethod(input))
	fmt.Printf("this number is Palindrom is : %v", numMethod(input))
}

func strMethod(i int) bool {
	str := strconv.Itoa(i)
	reverse := ""
	for _, v := range str {
		reverse = string(v) + reverse
	}
	num, err := strconv.Atoi(reverse)
	if err != nil {
		return false
	}

	if num == i {
		return true
	}

	return false
}

func numMethod(num int) int {
	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10

	}
	return res
}
