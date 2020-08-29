package domain

import (
	"fmt"
	"github.com/hyeyoom/go-web-app-boilerplate/domain/base"
	"golang.org/x/crypto/bcrypt"
)

type Member struct {
	base.DefaultModel
	Email    string
	Password string
}

type MemberRepository interface {
	Create(*Member) int64
}

func NewMember(email string, plainPassword string) *Member {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	return &Member{Email: email, Password: string(hashedPassword)}
}

func (m *Member) IsCredentialVerified(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	return err == nil
}
