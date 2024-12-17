// ### **โจทย์ 4: Remove Duplicates from Sorted Array**
// **คำอธิบาย:**
// ให้คุณเขียนฟังก์ชันที่รับอาเรย์ที่เรียงลำดับแล้ว และคืนค่าความยาวของอาเรย์ใหม่ที่ไม่มีตัวเลขซ้ำ (เปลี่ยนอาเรย์เดิมได้)

// **ข้อกำหนด:**
// - ไม่ต้องสร้างอาเรย์ใหม่
// - เปลี่ยนค่าภายในอาเรย์ที่รับเข้า

// **ตัวอย่าง Input:**
// []int{1, 1, 2}

// **Output ที่คาดหวัง:**
// 2 // อาเรย์ใหม่คือ [1, 2, _]

// **Tips:**
// - ใช้ pointer สองตัว (`i` และ `j`) ในการเปรียบเทียบและย้ายค่าที่ไม่ซ้ำไปยังตำแหน่งใหม่

package main

import "fmt"

func main() {
	input := []int{1, 1, 2, 2, 3, 3, 3, 4, 4}
	n := len(input)
	i := 0
	for j := 1; j < n; j++ {
		if input[i] != input[j] {
			i++
			input[i] = input[j]
		}
	}

	fmt.Println(input[:i+1])
}
