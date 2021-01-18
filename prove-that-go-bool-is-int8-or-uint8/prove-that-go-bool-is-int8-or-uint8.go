package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 声明bool变量
	b := true

	// 验证下bool占的数据宽度，虽然记忆是1，但还是打印确认下
	fmt.Printf("sizeof(bool) = %d\n", unsafe.Sizeof(b))

	// 强转内存对齐，看下值
	fmt.Printf("bool value = %d\n", *(*uint8)(unsafe.Pointer(&b)))

	// 看下负数是true or false
	*(*int8)(unsafe.Pointer(&b)) = -100

	if b {
		fmt.Printf("哈哈。。。证明负数转成bool也是true\n")
	}

}
