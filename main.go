package main

import (
	"fmt"
	"proto/pb"
)

func main() {
	o := pb.Order{}

	switch t := o.ExternalPayment.(type) {
	case *pb.Order_DigiPay:
		fmt.Printf("%v\n", t.DigiPay)
	case *pb.Order_SnappPay:
		fmt.Printf("%v\n", t.SnappPay)
	case nil:
		fmt.Println("nil")
	}
}
