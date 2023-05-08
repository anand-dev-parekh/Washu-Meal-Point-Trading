package models

import "golang.org/x/crypto/bcrypt"

// user struct as in DB
type User struct {
	Id        uint64 `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	AuthLevel uint8  `json:"authLevel"`
}

// hashes user obj password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)

	return err
}

// checks hashes match
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
