package template

import (
	"fmt"
	"log"
	"testing"
)

func TestTemplate(t *testing.T){
	GetAllTemplate()
	orders := []Order{
		{OrderId: 1, OrderDesc: "iphone 5s", CustomerId: 100, CustomerName: "John"},
		{OrderId: 2, OrderDesc: "iphone 6", CustomerId: 101, CustomerName: "Alice"},
		{OrderId: 3, OrderDesc: "iphone 13Pro", CustomerId: 102, CustomerName: "Sarı Cizmeli"},
	}
	wo := WrappedOrder {Description: "You can find all orders gave in last two days.", Orders: orders}
	res := ParseTemplate(EmailTemplate, wo)
	if len(res) <= 0 {
		log.Fatalf("Error")
	}
	orders[0].CustomerId = 4
	res1 := ParseTemplate(EmailTemplate, wo)
	fmt.Println(res)
	fmt.Println(res1)
}