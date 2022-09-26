package discount

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewClientService(port string) DiscountServiceClient {
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to %v: %v", port, err)
	}

	return NewDiscountServiceClient(conn)
}
