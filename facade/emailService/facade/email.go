package emailservice

import "fmt"

type Email struct {
	remetente    string
	destinatario string
	assunto      string
	corpo        string
}

func (e Email) String() string {
	return fmt.Sprintf(`
	De: %s
	Para: %s
	Assunto: %s
	Corpo: %s
	
	`, e.remetente, e.destinatario, e.assunto, e.corpo)
}
