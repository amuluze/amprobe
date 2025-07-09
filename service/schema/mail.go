// Package schema
// Date:   2024/10/14 16:20
// Author: Amu
// Description:
package schema

type Mail struct {
	ID       uint   `json:"id"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

type MailCreateArgs struct {
	Server   string `json:"server" validate:"required"`
	Port     int    `json:"port" validate:"required"`
	Sender   string `json:"sender" validate:"required"`
	Password string `json:"password" validate:"required"`
	Receiver string `json:"receiver"`
}

type MailUpdateArgs struct {
	ID       uint   `json:"id" validate:"required"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Sender   string `json:"sender"`
	Password string `json:"password"`
	Receiver string `json:"receiver"`
}

type MailDeleteArgs struct {
	ID uint `json:"id" validate:"required"`
}

type MailTestArgs struct {
	Receiver string `json:"receiver" validate:"required"`
}
