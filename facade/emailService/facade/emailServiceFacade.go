package emailservice

import (
	"fmt"
	"strings"
)

type EmailService struct {
}

type IEmailService interface {
	Enviar(destinatario, assunto, corpo string)
	validaDestinatario(destinatario string) error
	monta(destinatario, assunto, corpo string) Email
	envia(e Email)
}

func NewEmailService() IEmailService {
	return &EmailService{}
}

func (es *EmailService) Enviar(destinatario, assunto, corpo string) {
	err := es.validaDestinatario(destinatario)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	email := es.monta(destinatario, assunto, corpo)
	es.envia(email)
}

func (es *EmailService) validaDestinatario(destinatario string) error {
	fmt.Println("📬 Validando email:", destinatario)
	if strings.Contains(destinatario, "@") {
		return nil
	}

	return fmt.Errorf("❌ Email inválido. Cancelando envio")
}

func (es *EmailService) monta(destinatario, assunto, corpo string) Email {
	fmt.Println("✉️ Montando e-mail...")

	return Email{
		remetente:    "no-reply@sistema.com",
		destinatario: destinatario,
		assunto:      assunto,
		corpo:        corpo,
	}
}

func (es EmailService) envia(e Email) {
	fmt.Println(e)
	fmt.Println("✅ Email enviado com sucesso!")
}
