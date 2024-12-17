// ### **โจทย์ 1: Two Sum**

// **คำอธิบาย:**
// ให้คุณเขียนฟังก์ชันที่รับอาเรย์ของตัวเลขและตัวเลขเป้าหมาย (`target`) คืนค่าดัชนีของตัวเลขสองตัวในอาเรย์ที่เมื่อบวกกันแล้วได้ค่าตรงกับ `target`
// **ข้อกำหนด:**
// - มีตัวเลขอย่างน้อยสองตัวที่ตรงตามเงื่อนไข
// - ดัชนีเริ่มต้นที่ 0

// **ตัวอย่าง Input:**
// nums := []int{2, 7, 11, 15}
// target := 9

// **Output ที่คาดหวัง:**
// []int{0, 1}

// **Tips:**
// - ใช้ `map` เพื่อเก็บตัวเลขและดัชนีของตัวเลขที่ผ่านมา
// - สำหรับทุกตัวเลขในอาเรย์ ให้เช็คว่า `target - num` มีอยู่ใน `map` หรือไม่

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}

	for i, value := range nums {
		if j, ok := m[target-value]; ok && i != j {
			return []int{i, j}
		}
	}
	return nil
}
