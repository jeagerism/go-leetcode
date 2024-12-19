// ### **โจทย์ 6: Merge Two Sorted Arrays**
// **คำอธิบาย:**
// ให้รวมอาเรย์ที่เรียงลำดับแล้วสองอาเรย์เข้าเป็นอาเรย์เดียวที่เรียงลำดับ

// **ตัวอย่าง Input:**
// nums1 := []int{1, 3, 5}
// nums2 := []int{2, 4, 6}

// **Output ที่คาดหวัง:**
// []int{1, 2, 3, 4, 5, 6}

// **Tips:**
// - ใช้ pointer สองตัวสำหรับวนลูปแต่ละอาเรย์และเปรียบเทียบ

package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 3, 5}
	nums2 := []int{2, 4, 6}

	output := append(nums1, nums2...)
	fmt.Println(output)

	n := len(output)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if output[j] > output[j+1] {
				output[j+1], output[j] = output[j], output[j+1]
			}
		}
	}

	fmt.Println(output)
}
