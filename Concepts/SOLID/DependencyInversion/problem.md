# Dependency Inversion Principle Problem Statement

## Background
A notification system needs to send messages to users through different channels (email, SMS, push notifications). Currently, the system directly depends on specific notification implementations.

## Problem
The current implementation violates DIP:

```go
type User struct {
    ID       int
    Name     string
    Email    string
    Phone    string
    DeviceID string
}

type EmailService struct {}

func (e *EmailService) SendEmail(to, subject, body string) error {
    // Implementation to send email
    fmt.Printf("Sending email to %s: %s\n", to, body)
    return nil
}

type SMSService struct {}

func (s *SMSService) SendSMS(to, message string) error {
    // Implementation to send SMS
    fmt.Printf("Sending SMS to %s: %s\n", to, message)
    return nil
}

type PushNotificationService struct {}

func (p *PushNotificationService) SendPush(deviceID, message string) error {
    // Implementation to send push notification
    fmt.Printf("Sending push to device %s: %s\n", deviceID, message)
    return nil
}

// NotificationManager directly depends on concrete implementations
type NotificationManager struct {
    emailService          *EmailService
    smsService           *SMSService
    pushNotificationService *PushNotificationService
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
```

## Issues
1. High-level module (`NotificationManager`) depends on low-level modules
2. Concrete implementations are tightly coupled
3. Difficult to test due to direct dependencies
4. Hard to add new notification channels
5. Cannot easily mock services for testing

## Challenge
Refactor the notification system to follow the Dependency Inversion Principle so that:
1. High-level modules depend on abstractions
2. Low-level modules depend on abstractions
3. The system is loosely coupled
4. New notification channels can be easily added
5. The system is testable with mocks

## Requirements
1. Support existing notification channels (email, SMS, push)
2. Make it easy to add new notification channels
3. Allow for easy testing with mocks
4. Support notification channel configuration
5. Handle errors appropriately

## Bonus Challenge
Add a new notification channel (e.g., Slack or Discord) to demonstrate how your solution supports extension through abstractions.

## Expected Usage
The solution should allow code like this to work:

```go
func main() {
    // Initialize with different implementations
    notifier := NewNotificationManager([]NotificationService{
        NewEmailService(),
        NewSMSService(),
        NewPushService(),
    })

    // Send notification through all channels
    user := User{
        Email: "user@example.com",
        Phone: "1234567890",
        DeviceID: "device123",
    }
    
    err := notifier.NotifyUser(user, "Hello!")
}

// Testing should be easy with mocks
func TestNotificationManager(t *testing.T) {
    mockService := &MockNotificationService{}
    notifier := NewNotificationManager([]NotificationService{mockService})
    
    // Test notification sending
    err := notifier.NotifyUser(user, "Test message")
    // Assert expectations
}
```
