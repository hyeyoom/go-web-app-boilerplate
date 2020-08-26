package domain

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Member struct {
	Email    string
	Password string
}

type MemberRepository interface {
}

func NewMember(Email string, PlainPassword string) *Member {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(PlainPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	return &Member{Email: Email, Password: string(hashedPassword)}
}

func (m *Member) IsCredentialVerified(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	return err == nil
}
