package emailservice

import (
	"fmt"
	"strings"
)

type emailService struct {
}

type IEmailService interface {
	Enviar(destinatario, assunto, corpo string)
}

func NewEmailService() IEmailService {
	return &emailService{}
}

func (es *emailService) Enviar(destinatario, assunto, corpo string) {
	err := es.validaDestinatario(destinatario)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	email := es.monta(destinatario, assunto, corpo)
	es.envia(email)
}

func (es *emailService) validaDestinatario(destinatario string) error {
	fmt.Println("📬 Validando email:", destinatario)
	if strings.Contains(destinatario, "@") {
		return nil
	}

	return fmt.Errorf("❌ Email inválido. Cancelando envio")
}

func (es *emailService) monta(destinatario, assunto, corpo string) Email {
	fmt.Println("✉️ Montando e-mail...")

	return Email{
		remetente:    "no-reply@sistema.com",
		destinatario: destinatario,
		assunto:      assunto,
		corpo:        corpo,
	}
}

func (es emailService) envia(e Email) {
	fmt.Println(e)
	fmt.Println("✅ Email enviado com sucesso!")
}
