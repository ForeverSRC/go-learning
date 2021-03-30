package main

import "fmt"

type user struct {
	name  string
	email string
}
// 使用一个值接收者实现方法
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// 使用一个指针接收者实现方法
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	// 值
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// 指针
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
