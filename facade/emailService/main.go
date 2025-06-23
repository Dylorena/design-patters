package main

import s "emailService/facade"

func main() {
	email := s.NewEmailService()
	email.Enviar("usuario@exemplo.com", "Bem-vindo!", "Obrigado por se registrar.")
}
