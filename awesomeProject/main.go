package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	isActive  bool
	KidsCount int
}

func main() {
	user := User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
		isActive:  true,
		KidsCount: 3,
	}
	fmt.Println("Hello World" + user.FirstName)

	//fmt.Println("User First Name:", user.FirstName)
	//fmt.Println("User Last Name:", user.LastName)
	//fmt.Println("User Age:", user.Age)
	//fmt.Println("User isActive:", user.isActive)

	fmt.Println(user.KidsCount + user.Age)

	var a int

	var e int64

	fmt.Println("int", unsafe.Sizeof(a))
	fmt.Println("int8", unsafe.Sizeof(int8(a)))
	fmt.Println("int16", unsafe.Sizeof(int16(a)))
	fmt.Println("int32", unsafe.Sizeof(int32(a)))
	fmt.Println("int64", unsafe.Sizeof(int64(a)))
	fmt.Println("--------------------------")
	fmt.Println("uint8", unsafe.Sizeof(uint8(a)))
	fmt.Println("uint16", unsafe.Sizeof(uint16(a)))
	fmt.Println("uint32", unsafe.Sizeof(uint32(a)))
	fmt.Println("uint64", unsafe.Sizeof(uint64(a)))
	fmt.Println("--------------------------")
	fmt.Println("int64 -> int8", unsafe.Sizeof(int8(e)))

}
