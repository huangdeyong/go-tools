package hook

import "fmt"

type User struct {
	Id int64
	Name string
}

func (u *User) UpdateName(newestName string) {
	fmt.Println("UpdateName....")
	u.Name = newestName
}