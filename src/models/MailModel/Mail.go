package email

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
)

// Config Configuracion de Correo Electrónico
type Config struct {
	Username   string `jon:"Username,omitempty"`
	Password   string `jon:"Password,omitempty"`
	ServerHost string `jon:"ServerHost,omitempty"`
	ServerPort string `jon:"ServerPort,omitempty"`
	SenderAddr string `jon:"SenderAddr,omitempty"`
}

// TestMail prueba que la configuracion del Correo sea la adecuada.
func TestMail(cfg Config) (bool, *smtp.Client, string) {
	// Definir credenciales de autenticacion Plana
	auth := smtp.PlainAuth(
		"",
		cfg.SenderAddr,
		cfg.Password,
		cfg.ServerHost,
	)

	// Definir configuracion TLS
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         cfg.ServerHost,
	}

	Error := ""

	server := fmt.Sprintf("%v:%v", cfg.ServerHost, cfg.ServerPort)
	c, err := smtp.Dial(server)

	if err != nil {
		fmt.Println("NO SE PUEDE CONTACTAR CON EL SERVIDOR, VERIFIQUE SUS DATOS O INTENTE MAS TARDE.")
		Error = fmt.Sprintln("NO SE PUEDE CONTACTAR CON EL SERVIDOR, VERIFIQUE SUS DATOS O INTENTE MAS TARDE: ", err)
		return false, nil, Error
	}

	fmt.Println("Se consiguio un cliente SMTP.")

	err = c.StartTLS(tlsconfig)
	if err != nil {
		fmt.Println("NO SE PUEDE ESTABLECER CANAL TLS, VERIFIQUE SUS DATOS  O INTENTE MAS TARDE.")
		Error = fmt.Sprintln("NO SE PUEDE ESTABLECER CANAL TLS, VERIFIQUE SUS DATOS  O INTENTE MAS TARDE: ", err)
		return false, nil, Error
	}

	fmt.Println("Se consiguio TLS.")

	err = c.Auth(auth)
	if err != nil {
		fmt.Println("NO SE PUEDE AUTENTICAR EL USUARIO, VERIFIQUE SUS CREDENCIALES.")
		Error = fmt.Sprintln("NO SE PUEDE AUTENTICAR EL USUARIO, VERIFIQUE SUS CREDENCIALES: ", err)
		return false, nil, Error
	}

	fmt.Println("Los datos de autenticacion son correctos.")

	return true, c, "Los datos de autenticacion son correctos"
}

// SendMail Funcion para enviar sobre un cliente smtp un correo con el body, Asunto y Destino indicados
func SendMail(cfg Config, body string, subject string, receptor string) bool {
	fmt.Println("Igreso a Sender")
	EstableceConexion, c, _ := TestMail(cfg)
	if EstableceConexion {
		fmt.Println("Se pudo establecer conexion")
		headers := make(map[string]string)
		headers["From"] = cfg.SenderAddr
		headers["To"] = receptor
		headers["Subject"] = subject
		msj := ""
		for k, v := range headers {
			msj += fmt.Sprintf("%s: %s\r\n", k, v)
		}

		msj += "\r\n" + body

		err := c.Mail(cfg.SenderAddr)
		if err != nil {
			log.Panic(err)
			return false
		}
		err = c.Rcpt(receptor)
		if err != nil {
			log.Panic(err)
			return false
		}

		w, err := c.Data()
		if err != nil {
			log.Panic(err)
			return false
		}

		_, err = w.Write([]byte(msj))
		if err != nil {
			log.Panic(err)
			return false
		}

		err = w.Close()
		if err != nil {
			log.Panic(err)
			return false
		}

		c.Quit()
		return true
	}
	return false
}
