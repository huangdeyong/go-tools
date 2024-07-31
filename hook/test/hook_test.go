package test

import (
	"github.com/huangdeyong/go-tools/hook"
	"testing"
)

func Test01(t *testing.T) {
	hook.HookUpdateName()

	user := hook.User{
		Id:   1,
		Name: "name1",
	}

	user.UpdateName("new-name")
}
