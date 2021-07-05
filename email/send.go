package email

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"strings"
)


type Sender struct {
	Server      string
	Email       string
	Passwd      string
	Port        int
	To          []string
	Cc          []string
	Subject     string
	Body        string
	Attachment  []Attachment
	Delimeter   string
}

type Attachment struct {
	Filepath string
	Filename string
}

func (s Sender) Send() error {
	tlsConfig := tls.Config{
		ServerName:         s.Server,
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", s.Server, s.Port), &tlsConfig)
	if err != nil {
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, s.Server)
	if err != nil {
		return err
	}
	defer client.Close()

	auth := smtp.PlainAuth("", s.Email, s.Passwd, s.Server)
	if err := client.Auth(auth); err != nil {
		return err
	}

	for _, to := range s.To {
		if err := client.Rcpt(to); err != nil {
			return err
		}
	}

	writer, err := client.Data()
	if err != nil {
		return err
	}

	//basic email headers
	msg := fmt.Sprintf("From: %s\r\n", s.Email)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(s.To, ";"))
	if len(s.Cc) > 0 {
		msg += fmt.Sprintf("Cc: %s\r\n", strings.Join(s.Cc, ";"))
	}
	msg += fmt.Sprintf("Subject: %s\r\n", s.Subject)

	msg += "MIME-Version: 1.0\r\n"
	msg += fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n", s.Delimeter)

	//place HTML message
	msg += fmt.Sprintf("\r\n--%s\r\n", s.Delimeter)
	msg += "Content-Type: text/html; charset=\"utf-8\"\r\n"
	msg += "Content-Transfer-Encoding: 7bit\r\n"
	msg += fmt.Sprintf("\r\n%s", s.Body + "\r\n")

	for _, v := range s.Attachment {
		//place file
		msg += fmt.Sprintf("\r\n--%s\r\n", s.Delimeter)
		msg += "Content-Type: text/plain; charset=\"utf-8\"\r\n"
		msg += "Content-Transfer-Encoding: base64\r\n"
		msg += "Content-Disposition: attachment;filename=\"" + v.Filename + "\"\r\n"

		//read file
		file, err := ioutil.ReadFile(v.Filepath)
		if err != nil {
			return err
		}
		msg += "\r\n" + base64.StdEncoding.EncodeToString(file)
	}

	if _, err := writer.Write([]byte(msg)); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	client.Quit()
	return nil
}