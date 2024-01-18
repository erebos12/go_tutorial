package main

import "fmt"

// MessageService definiert ein Interface für einen Service, der Nachrichten sendet.
type MessageService interface {
	Send(message string, receiver string) error
}

// EmailService ist eine konkrete Implementierung von MessageService, die E-Mails sendet.
type EmailService struct{}

func (es *EmailService) Send(message string, receiver string) error {
	fmt.Printf("E-Mail an %s gesendet: %s\n", receiver, message)
	return nil
}

// SmsService ist eine konkrete Implementierung von MessageService, die SMS-Nachrichten sendet.
type SmsService struct{}

func (ss *SmsService) Send(message string, receiver string) error {
	fmt.Printf("SMS an %s gesendet: %s\n", receiver, message)
	return nil
}

// Konsument von Interface MessageService: NotificationSender benutzt ein MessageService, um Benachrichtigungen zu senden.
type NotificationSender struct {
	service MessageService
}

// Erstellt eine neue Instanz von NotificationSender und übergibt den MessageService, der für die Benachrichtigungen verwendet werden soll
func NewNotificationSender(service MessageService) *NotificationSender {
	return &NotificationSender{service: service}
}

// Benachrichtigungen werden an den MessageService gesendet.
func (ns *NotificationSender) SendNotification(message string, receiver string) error {
	return ns.service.Send(message, receiver)
}

func main() {
	// Verwenden des EmailService
	emailService := &EmailService{}
	emailSender := NewNotificationSender(emailService)
	err := emailSender.SendNotification("Willkommen bei Go!", "empfaenger@example.com")
	if err != nil {
		fmt.Println("Fehler beim Senden der E-Mail:", err)
	}

	// Verwenden des SmsService
	smsService := &SmsService{}
	smsSender := NewNotificationSender(smsService)
	err = smsSender.SendNotification("Willkommen bei Go SMS!", "123-456-7890")
	if err != nil {
		fmt.Println("Fehler beim Senden der SMS:", err)
	}
}
