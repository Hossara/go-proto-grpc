package api

import (
	"context"
	"proto/pb"
)

type orderGRPCApi struct {
	pb.UnimplementedOrderServiceServer
}

func NewOrderGRPCServer() pb.OrderServiceServer {
	return new(orderGRPCApi)
}

func (s *orderGRPCApi) GetOrder(ctx context.Context, req *pb.GetOrderFilter) (*pb.GetOrderResponse, error) {
	return &pb.GetOrderResponse{
		Orders: []*pb.Order{{
			Id:        1,
			Quantity:  1,
			TotalBill: 55000000,
			IpgMethod: "Saman",
			Items: []*pb.OrderItem{
				{
					Id:              1,
					OrderId:         1,
					ItemName:        "dodnod",
					Quantity:        1,
					UnitPrice:       55000000,
					ItemDescription: "woibosibodbsodbosbdos",
				},
			},
			Status: pb.Order_PAID,
			User: &pb.User{
				Id:        1,
				FirstName: "Hossein",
				LastName:  "Araghi",
			},
			Details:         nil,
			ExternalPayment: new(pb.Order_SnappPay),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}},
	}, nil
}
