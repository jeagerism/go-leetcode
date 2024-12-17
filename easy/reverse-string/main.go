// ### **โจทย์ 2: Reverse String**
// **คำอธิบาย:**
// เขียนฟังก์ชันที่รับสตริงและคืนค่าสตริงที่เรียงลำดับย้อนกลับ

// **ตัวอย่าง Input:**
// "hello"

// **Output ที่คาดหวัง:**
// "olleh"

// **Tips:**
// - ใช้ slice ของตัวอักษร (`[]rune`) และวนลูปสลับตำแหน่งหัว-ท้าย

package main

import "fmt"

func main() {
	input := "hello"
	runeInput := []rune(input)
	n := len(input)
	for i := 0; i < n; i++ {
		fmt.Printf("%c", runeInput[n-i-1])
	}

}
