package main

import (
	"fmt"
	"sync"
)

type DataTest struct {
	User  *User
	Order *Order
}
type User struct {
	Username string
	Email    string
}
type Order struct {
	OrderId    string
	TotalOrder int64
}

func addUser(d *DataTest, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	//var u User
	//
	//err := faker.FakeData(&u)
	//if err != nil {
	//	fmt.Println("user error : ", err)
	//}

	u := &User{
		Username: "test",
		Email:    "test@email.com",
	}

	mtx.Lock()
	d.User = u
	mtx.Unlock()

}

func addOrder(d *DataTest, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	//var o Order
	//
	//err := faker.FakeData(&o)
	//if err != nil {
	//	fmt.Println("order error : ", err)
	//}

	o := &Order{
		OrderId:    "order-1",
		TotalOrder: 5,
	}

	mtx.Lock()
	d.Order = o
	mtx.Unlock()

}

func main() {

	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	d := &DataTest{}

	wg.Add(2)

	go addUser(d, &wg, &mtx)
	go addOrder(d, &wg, &mtx)

	wg.Wait()

	fmt.Println(d.User)
	fmt.Println(d.Order)
}
