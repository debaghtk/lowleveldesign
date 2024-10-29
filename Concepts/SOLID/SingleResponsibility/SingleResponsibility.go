package user

import "errors"

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
	err := validateUser(user)
	if err != nil {
		return err
	}

	// Store user
	um.users = append(um.users, user)

	// Send welcome email
	err = um.sendWelcomeEmail(user)
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
	err := um.sendWelcomeEmail(user)
	if err != nil {
		return err
	}

	return nil
}

func (um *UserManager) exportToCSV() error {
	// CSV export logic
	err := um.exportToCSV()
	if err != nil {
		return err
	}
	return nil
}

func validateUser(user User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("invalid user data")
	}
	return nil
}
