# Single Responsibility Principle (SRP) Problem

## Problem Statement
Create a user management system that currently has a User struct and a UserManager struct that handles multiple responsibilities:
- User data management
- Email notifications  
- Data validation
- File export

## Challenge
Refactor this code to follow the Single Responsibility Principle by:
- Separating the different responsibilities into their own structs/services
- Making each component responsible for only one specific task
- Implementing proper interfaces to maintain loose coupling

## Code that Violates SRP

```go
package user

type User struct {
    ID       int
    Name     string
    Email    string
    Password string
}

type UserManager struct {
    users []User
}

func (um *UserManager) AddUser(user User) error {
    // Validate user data
    if user.Name == "" || user.Email == "" {
        return errors.New("invalid user data")
    }

    // Store user
    um.users = append(um.users, user)

    // Send welcome email
    err := um.sendWelcomeEmail(user)
    if err != nil {
        return err
    }

    // Export users to CSV
    err = um.exportToCSV()
    if err != nil {
        return err
    }

    return nil
}

func (um *UserManager) sendWelcomeEmail(user User) error {
    // Email sending logic
    return nil
}

func (um *UserManager) exportToCSV() error {
    // CSV export logic
    return nil
}
```