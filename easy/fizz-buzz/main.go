// ### **โจทย์ 7: FizzBuzz**
// **คำอธิบาย:**
// เขียนฟังก์ชันที่คืนอาเรย์ของสตริงที่แสดงตัวเลขตั้งแต่ 1 ถึง `n` โดย:
// - แทนที่ตัวเลขที่หารด้วย 3 ลงตัวด้วย `"Fizz"`
// - แทนที่ตัวเลขที่หารด้วย 5 ลงตัวด้วย `"Buzz"`
// - แทนที่ตัวเลขที่หารด้วยทั้ง 3 และ 5 ลงตัวด้วย `"FizzBuzz"`

// **ตัวอย่าง Input:**
// n := 15

// **Output ที่คาดหวัง:**
// []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}

// **Tips:**
// - ใช้เงื่อนไข `if-else` และการเช็ค `%`

package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 15

	output := []string{}

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			output = append(output, "FizzBuzz")
		} else if i%3 == 0 {
			output = append(output, "Fizz")
		} else if i%5 == 0 {
			output = append(output, "Buzz")
		} else {
			output = append(output, strconv.Itoa(i))
		}
	}

	fmt.Println(output)
}
