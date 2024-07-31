package hook

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
)

// 替代的 UpdateName 方法
// 这个方法在原始 UpdateName 方法前后执行附加逻辑
func myUpdateName(u *User, newestName string) {
	fmt.Println("hook -- 调用原方法前....")

	// 调用原始的 UpdateName 方法
	// `saveOriginalUpdateName` 在这个上下文中充当了中间函数的角色
	saveOriginalUpdateName(u, newestName)

	fmt.Println("hook -- 调用原方法后....")
}

// 定义一个中间函数来作为 `gohook.Hook` 的第三个参数
// 这个函数主要用于保存原始方法的函数指针
func saveOriginalUpdateName(u *User, newestName string) {
	// do nothing...
	// 该方法的主要目的是作为原始方法的占位符，并保存其函数指针
	// 实际上，这个方法体内的代码不会被执行
}

// HookUpdateName 进行方法挂钩
// 这个函数会将 `(*User).UpdateName` 方法挂钩到 `myUpdateName`
// `saveOriginalUpdateName` 用于保存原始方法的指针
func HookUpdateName() {
	err := gohook.Hook(
		(*User).UpdateName,       // 目标方法: 需要被挂钩的方法
		myUpdateName,             // 替代方法: 在目标方法调用时被执行的方法
		saveOriginalUpdateName,   // 中间函数: 用于保存原始方法的函数指针
	)
	if err != nil {
		fmt.Printf("Failed to hook method: %v\n", err)
	} else {
		fmt.Println("Hooked UpdateName method successfully")
	}
}
