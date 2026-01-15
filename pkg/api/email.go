package api

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strings"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	td "github.com/nervatura/nervatura/v6/test/driver"
)

type smtpConfig struct {
	username string
	password string
	host     string
	port     int
	tlsMin   uint16
	conn     string
	auth     string
}

func (ds *DataStore) createEmail(from string, emailTo []string, options cu.IM) (string, error) {
	const delimiter = "**=myohmy689407924327"
	var b strings.Builder
	emailOpt := cu.ToIM(options["email"], cu.IM{})

	// Write header
	fmt.Fprintf(&b, "From: %s\r\n", from)
	fmt.Fprintf(&b, "To: %s\r\n", strings.Join(emailTo, ";"))
	fmt.Fprintf(&b, "Subject: %s\r\n", cu.ToString(emailOpt["subject"], ""))
	fmt.Fprintf(&b, "MIME-Version: 1.0\r\n")
	fmt.Fprintf(&b, "Content-Type: multipart/mixed; boundary=\"%s\"\r\n", delimiter)

	// Write body
	fmt.Fprintf(&b, "\r\n--%s\r\n", delimiter)
	b.WriteString("Content-Type: text/html; charset=\"utf-8\"\r\n")
	b.WriteString("Content-Transfer-Encoding: 7bit\r\n")

	body := cu.ToString(emailOpt["html"], cu.ToString(emailOpt["text"], ""))
	fmt.Fprintf(&b, "\r\n%s\r\n", body)

	// Handle attachments
	attachments := cu.ToIMA(emailOpt["attachments"], []cu.IM{})
	if len(attachments) > 0 {
		if err := ds.addAttachments(&b, attachments, delimiter); err != nil {
			return "", err
		}
	}

	return b.String(), nil
}

func (ds *DataStore) addAttachments(b *strings.Builder, attachments []cu.IM, delimiter string) error {
	for i, attachment := range attachments {
		filename := fmt.Sprintf("docs_%d.pdf", i+1)
		if fname, found := attachment["filename"]; found {
			filename = cu.ToString(fname, filename)
		}

		fmt.Fprintf(b, "\r\n--%s\r\n", delimiter)
		fmt.Fprintf(b, "Content-Type: application/pdf; charset=\"utf-8\"\r\n"+
			"Content-Transfer-Encoding: base64\r\n"+
			"Content-Disposition: attachment;filename=\"%s\"\r\n", filename)

		params := cu.IM{"output": "pdf"}
		for _, key := range []string{"report_key", "report_id", "code", "orientation", "size", "template"} {
			if val, found := attachment[key]; found {
				params[key] = val
			}
		}

		report, err := ds.GetReport(params)
		if err != nil {
			return err
		}

		b.WriteString("\r\n")
		b.WriteString(base64.StdEncoding.EncodeToString(report["template"].([]uint8)))
	}
	return nil
}

func validateEmailOptions(options cu.IM) error {
	emailOpt, valid := options["email"].(cu.IM)
	if !valid {
		return errors.New(ut.GetMessage("missing_required_field") + ": email")
	}
	recipients := cu.ToIMA(emailOpt["recipients"], []cu.IM{})
	if len(recipients) == 0 {
		return errors.New(ut.GetMessage("missing_required_field") + ": recipients")
	}
	if cu.ToString(options["provider"], "smtp") != "smtp" {
		return errors.New(ut.GetMessage("invalid_provider"))
	}
	return nil
}

func (ds *DataStore) setupSmtpConnection(config smtpConfig) (conn net.Conn, client md.SmtpClient, err error) {
	tlsConfig := &tls.Config{
		ServerName:         config.host,
		InsecureSkipVerify: false,
		MinVersion:         tls.VersionTLS13,
	}
	if config.tlsMin > 0 {
		tlsConfig.MinVersion = config.tlsMin
	}

	connFuncs := map[string]func() (net.Conn, error){
		"net": func() (net.Conn, error) {
			return net.DialTimeout("tcp", net.JoinHostPort(config.host, fmt.Sprintf("%d", config.port)), time.Duration(2)*time.Second)
		},
		"tls": func() (net.Conn, error) {
			return tls.Dial("tcp", fmt.Sprintf("%s:%d", config.host, config.port), tlsConfig)
		},
		"test": func() (net.Conn, error) {
			return td.NewTestConn(), nil
		},
	}

	if conn, err = connFuncs[config.conn](); err == nil {
		if client, err = ds.NewSmtpClient(conn, config.host); err != nil {
			conn.Close()
		}
	}

	return conn, client, err
}

func (ds *DataStore) authenticateSmtp(client md.SmtpClient, config smtpConfig) error {
	auth := smtp.PlainAuth("", config.username, config.password, config.host)

	authFuncs := map[string]func(smtp.Auth) error{
		"auth": func(auth smtp.Auth) error { return client.Auth(auth) },
		"none": func(smtp.Auth) error { return nil },
	}

	return authFuncs[config.auth](auth)
}

func (ds *DataStore) setupEmailAddresses(client md.SmtpClient, options cu.IM, defaultFrom string) (string, []string, error) {
	emailOpt := cu.ToIM(options["email"], cu.IM{})
	from := cu.ToString(emailOpt["from"], defaultFrom)
	if err := client.Mail(from); err != nil {
		return "", nil, err
	}

	emailTo := []string{}
	recipients := cu.ToIMA(emailOpt["recipients"], []cu.IM{})
	for _, recipient := range recipients {
		email := cu.ToString(recipient["email"], "")
		emailTo = append(emailTo, email)
		if err := client.Rcpt(email); err != nil {
			return "", nil, err
		}
	}

	return from, emailTo, nil
}

func (ds *DataStore) sendEmailContent(client md.SmtpClient, from string, emailTo []string, options cu.IM) error {
	writer, err := client.Data()
	if err != nil {
		return err
	}

	emailMsg, err := ds.createEmail(from, emailTo, options)
	if err != nil {
		return err
	}

	if _, err := writer.Write([]byte(emailMsg)); err != nil {
		return err
	}

	return writer.Close()
}

func (ds *DataStore) SendEmail(options cu.IM) (result cu.IM, err error) {
	result = cu.IM{"result": "OK"}
	// Setup SMTP configuration
	smtpConfig := smtpConfig{
		username: cu.ToString(ds.Config["NT_SMTP_USER"], ""),
		password: cu.ToString(ds.Config["NT_SMTP_PASSWORD"], ""),
		host:     cu.ToString(ds.Config["NT_SMTP_HOST"], ""),
		port:     int(cu.ToInteger(ds.Config["NT_SMTP_PORT"], 465)),
		tlsMin:   uint16(cu.ToInteger(ds.Config["NT_SMTP_TLS_MIN_VERSION"], 0)),
		conn:     cu.ToString(ds.Config["NT_SMTP_CONN"], "tls"),
		auth:     cu.ToString(ds.Config["NT_SMTP_AUTH"], "auth"),
	}

	// Validate inputs
	if err := validateEmailOptions(options); err != nil {
		return result, err
	}

	// Establish connection
	conn, client, err := ds.setupSmtpConnection(smtpConfig)
	if err != nil {
		return result, err
	}
	defer conn.Close()
	defer client.Close()

	// Authenticate
	if err := ds.authenticateSmtp(client, smtpConfig); err != nil {
		return result, err
	}

	// Setup sender and recipients
	from, emailTo, err := ds.setupEmailAddresses(client, options, smtpConfig.username)
	if err != nil {
		return result, err
	}

	// Send email
	if err := ds.sendEmailContent(client, from, emailTo, options); err != nil {
		return result, err
	}

	return result, client.Quit()
}
