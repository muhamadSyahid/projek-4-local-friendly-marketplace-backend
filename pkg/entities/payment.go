package entities

import "time"

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusRefunded  PaymentStatus = "refunded"
)

func ParsePaymentStatus(value string) PaymentStatus {
	switch value {
	case "completed":
		return PaymentStatusCompleted
	case "failed":
		return PaymentStatusFailed
	case "refunded":
		return PaymentStatusRefunded
	default:
		return PaymentStatusPending
	}
}

func (s PaymentStatus) String() string {
	return string(s)
}

type PaymentMethod string

const (
	PaymentMethodCreditCard   PaymentMethod = "credit_card"
	PaymentMethodDebitCard    PaymentMethod = "debit_card"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
	PaymentMethodEWallet      PaymentMethod = "e_wallet"
	PaymentMethodCOD          PaymentMethod = "cod"
)

func ParsePaymentMethod(value string) PaymentMethod {
	switch value {
	case "debit_card":
		return PaymentMethodDebitCard
	case "bank_transfer":
		return PaymentMethodBankTransfer
	case "e_wallet":
		return PaymentMethodEWallet
	case "cod":
		return PaymentMethodCOD
	default:
		return PaymentMethodCreditCard
	}
}

func (m PaymentMethod) String() string {
	return string(m)
}

type Payment struct {
	ID            string        `json:"id" bson:"_id,omitempty"`
	OrderID       string        `json:"orderId" bson:"orderId"`
	Amount        float64       `json:"amount" bson:"amount"`
	Method        PaymentMethod `json:"method" bson:"method"`
	Status        PaymentStatus `json:"status" bson:"status"`
	ProofURL      *string       `json:"proofUrl,omitempty" bson:"proofUrl,omitempty"`
	TransactionID *string       `json:"transactionId,omitempty" bson:"transactionId,omitempty"`
	PaidAt        time.Time     `json:"paidAt" bson:"paidAt"`
	CreatedAt     time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time     `json:"updatedAt" bson:"updatedAt"`
}

func (p Payment) IsCompleted() bool {
	return p.Status == PaymentStatusCompleted
}

func (p Payment) IsPending() bool {
	return p.Status == PaymentStatusPending
}
