package services

import (
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/aerosystems/subs-service/internal/repository"
	"github.com/aerosystems/subs-service/pkg/monobank"
)

type PaymentServiceImpl struct {
	paymentMethod  models.PaymentMethod
	invoiceRepo    repository.InvoiceRepository
	monobankClient *monobank.Client
}

func NewPaymentServiceImpl(invoiceRepo repository.InvoiceRepository, monobankClient *monobank.Client) *PaymentServiceImpl {
	return &PaymentServiceImpl{
		invoiceRepo:    invoiceRepo,
		monobankClient: monobankClient,
	}
}

func (ps *PaymentServiceImpl) SetPaymentMethod(paymentMethod string) error {
	switch paymentMethod {
	case models.MonobankPaymentMethod.String():
		ps.paymentMethod = models.MonobankPaymentMethod
	}
	return nil
}

func (ps *PaymentServiceImpl) CreateInvoice(invoice *models.Invoice) error {
	monoInvoice := &monobank.Invoice{
		Amount: invoice.Amount,
	}
	_, err := ps.monobankClient.CreateInvoice(monoInvoice)
	if err != nil {
		return err
	}
	return nil
}
