package hdy_test

import (
	"fmt"
	"testing"
)

type Attri struct {
	Atk int64
	Df  int64
}

func Test01(t *testing.T) {
	var fix = func(dist *int64, val int64) {
		*dist = val
	}

	attri := &Attri{
		Atk: 100,
		Df:  111,
	}
	fmt.Println(fmt.Sprintf("修改前：%v", attri))

	fix(&attri.Atk, 123)

	fmt.Println(fmt.Sprintf("修改后：%v", attri))
}

