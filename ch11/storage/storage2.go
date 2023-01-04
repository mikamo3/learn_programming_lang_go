package storage

import "net/smtp"

var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
}
