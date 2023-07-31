package person

import "fmt"

type notify interface {
	send()
}

// Email 实现1
type Email struct{}

func (Email) send() {
	fmt.Println("email send")
}

// Qq 实现2
type Qq struct{}

func (Qq) send() {
	fmt.Println("qq send")
}

func Test(n notify) {
	n.send()
}
