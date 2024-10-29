package dependencyinversion

import "fmt"

type User struct {
	ID       int
	Name     string
	Email    string
	Phone    string
	DeviceID string
}
type EmailInterface interface {
	SendEmail(string, string, string) error
}

type EmailService struct{}

func (e *EmailService) SendEmail(to, subject, body string) error {
	// Implementation to send email
	fmt.Printf("Sending email to %s: %s\n", to, body)
	return nil
}

type SMSInterface interface {
	SendSMS(string, string) error
}
type SMSService struct{}

func (s *SMSService) SendSMS(to, message string) error {
	// Implementation to send SMS
	fmt.Printf("Sending SMS to %s: %s\n", to, message)
	return nil
}

type PNInterface interface {
	SendPush(string, string) error
}

type PushNotificationService struct{}

func (p *PushNotificationService) SendPush(deviceID, message string) error {
	// Implementation to send push notification
	fmt.Printf("Sending push to device %s: %s\n", deviceID, message)
	return nil
}

// NotificationManager directly depends on concrete implementations
type NotificationManager struct {
	emailService            EmailInterface
	smsService              SMSInterface
	pushNotificationService PNInterface
}

func (nm *NotificationManager) NotifyUser(user User, message string) error {
	// Directly calling concrete implementations
	if err := nm.emailService.SendEmail(user.Email, "Notification", message); err != nil {
		return err
	}

	if err := nm.smsService.SendSMS(user.Phone, message); err != nil {
		return err
	}

	if err := nm.pushNotificationService.SendPush(user.DeviceID, message); err != nil {
		return err
	}

	return nil
}
