package main

import "fmt"

// Book struct สำหรับตัวอย่าง pointer
type Book struct {
	Title  string
	Author string
	Year   int
}

func main() {
	fmt.Println("=== ตัวอย่างการใช้ nil ใน Go ===\n")

	// 1. Pointer และ nil
	fmt.Println("1. Pointer และ nil:")
	var p *Book // pointer ยังไม่ได้ชี้ไปที่ไหน
	if p == nil {
		fmt.Println("   p == nil หมายถึง pointer ยังไม่ได้ชี้ไปที่ไหน")
	}

	p = &Book{Title: "Go Programming", Author: "John Doe", Year: 2024}
	if p != nil {
		fmt.Printf("   p != nil หมายถึง pointer มีค่าแล้ว: %+v\n", p)
	}
	fmt.Println()

	// 2. Map และ nil
	fmt.Println("2. Map และ nil:")
	var m map[string]int // map ยังไม่ได้สร้าง
	if m == nil {
		fmt.Println("   m == nil หมายถึง map ยังไม่ได้สร้าง")
		fmt.Println("   ไม่สามารถเพิ่มข้อมูลใน map ที่เป็น nil ได้")
	}

	m = make(map[string]int) // สร้าง map
	if m != nil {
		fmt.Println("   m != nil หมายถึง map ถูกสร้างแล้ว")
		m["score"] = 100
		fmt.Printf("   สามารถเพิ่มข้อมูลได้: %v\n", m)
	}
	fmt.Println()

	// 3. Slice และ nil
	fmt.Println("3. Slice และ nil:")
	var s []int // slice ยังไม่ได้สร้าง
	if s == nil {
		fmt.Println("   s == nil หมายถึง slice ยังไม่ได้สร้าง")
	}

	s = []int{1, 2, 3} // สร้าง slice
	if s != nil {
		fmt.Printf("   s != nil หมายถึง slice ถูกสร้างแล้ว: %v\n", s)
	}
	fmt.Println()

	// 4. Channel และ nil
	fmt.Println("4. Channel และ nil:")
	var ch chan int // channel ยังไม่ได้สร้าง
	if ch == nil {
		fmt.Println("   ch == nil หมายถึง channel ยังไม่ได้สร้าง")
	}

	ch = make(chan int) // สร้าง channel
	if ch != nil {
		fmt.Println("   ch != nil หมายถึง channel ถูกสร้างแล้ว")
	}
	close(ch) // ปิด channel
	fmt.Println()

	// 5. Function และ nil
	fmt.Println("5. Function และ nil:")
	var f func() // function ยังไม่ได้กำหนด
	if f == nil {
		fmt.Println("   f == nil หมายถึง function ยังไม่ได้กำหนด")
	}

	f = func() {
		fmt.Println("   Function ถูกเรียกใช้งาน")
	}
	if f != nil {
		fmt.Println("   f != nil หมายถึง function มีค่าแล้ว")
		f() // เรียกใช้ function
	}
	fmt.Println()

	// 6. Interface และ nil
	fmt.Println("6. Interface และ nil:")
	var i interface{} // interface ว่าง
	if i == nil {
		fmt.Println("   i == nil หมายถึง interface ยังไม่ได้กำหนดค่า")
	}

	i = "Hello"
	if i != nil {
		fmt.Printf("   i != nil หมายถึง interface มีค่าแล้ว: %v\n", i)
	}
	fmt.Println()

	// สรุป
	fmt.Println("=== สรุป ===")
	fmt.Println("ค่า nil จะเปลี่ยนไปตาม type ที่ใช้เปรียบเทียบ:")
	fmt.Println("- Slice, Pointer, Map, Channel, Function, Interface สามารถเป็น nil ได้")
	fmt.Println("- ตัวแปรธรรมดา (int, string, struct) ไม่สามารถเป็น nil ได้")
	fmt.Println()

	// ตัวอย่างที่ไม่สามารถใช้ nil ได้
	fmt.Println("ตัวอย่างตัวแปรที่ไม่สามารถเป็น nil:")
	var number int = 0
	var text string = ""
	var bookStruct Book = Book{}
	fmt.Printf("- int มีค่า default: %d\n", number)
	fmt.Printf("- string มีค่า default: \"%s\"\n", text)
	fmt.Printf("- struct มีค่า default: %+v\n", bookStruct)
}
