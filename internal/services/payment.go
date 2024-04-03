package services

import (
	"context"
	"google.golang.org/grpc"
	payments "parking-app-web-api-gateway/internal/payments/proto"
	"parking-app-web-api-gateway/internal/utils"
)

type PaymentService struct {
	URL string
}

func NewPaymentService() *PaymentService {
	url, _ := utils.GetEnvVar("PAYMENT_SERVICE_GRPC_URL")

	return &PaymentService{
		URL: url,
	}
}

func (p *PaymentService) UpdatePayment(updatePaymentRequest *payments.UpdatePaymentRequest) (*payments.UpdatePaymentResponse, error) {
	conn, err := grpc.Dial(p.URL, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	c := payments.NewPaymentsGrpcClient(conn)

	ctx := context.Background()

	r, err := c.UpdatePayment(ctx, updatePaymentRequest)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (p *PaymentService) GetPayments() (*payments.GetAllPaymentsResponse, error) {
	conn, err := grpc.Dial(p.URL, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	c := payments.NewPaymentsGrpcClient(conn)

	ctx := context.Background()
	r, err := c.GetAllPayments(ctx, &payments.GetAllPaymentsRequest{})
	if err != nil {
		return nil, err
	}

	return r, nil
}
