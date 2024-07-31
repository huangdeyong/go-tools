package hook

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
)

// 定义原始的 UpdateName 方法类型
type originalUpdateNameFunc func(*User, string)

// 定义保存原始函数指针的变量
var originalUpdateName originalUpdateNameFunc

// 替代的 UpdateName 方法
func myUpdateName(u *User, newestName string) {
	fmt.Println("hook -- 检查名字合法性....")

	// 调用原始的 UpdateName 方法
	originalUpdateName(u, newestName)
}

// Hook 函数
func HookUpdateName() {
	err := gohook.Hook(
		(*User).UpdateName,   // 目标方法
		myUpdateName,         // 替代方法
		&originalUpdateName,  // 原始方法指针
	)
	if err != nil {
		fmt.Printf("Failed to hook method: %v\n", err)
	} else {
		fmt.Println("Hooked UpdateName method successfully")
	}
}
