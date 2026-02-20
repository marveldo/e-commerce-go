package domain

type CreatePaymentDomain struct {
	CartID uint
}

type PaymentWebhookdomain struct {
	Event string
	Data  struct {
		Reference string
	}
}
