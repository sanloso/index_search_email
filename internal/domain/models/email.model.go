package models

const (
	EmailsFolder = "./enron_mail_20110402/maildir"
	EmailName    = "emails"
)

type Email struct {
	MessageID string `json:"message_id"`
	Date      string `json:"date"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
	Filepath  string `json:"filepath"`
}
