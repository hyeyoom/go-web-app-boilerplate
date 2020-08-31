package persistence

import (
	"github.com/hyeyoom/go-web-app-boilerplate/domain"
	"gorm.io/gorm"
)

type memberRepository struct {
	repository
}

func NewMemberRepository(db *gorm.DB) domain.MemberRepository {
	var mr memberRepository
	mr.db = db
	return &mr
}

func (mr *memberRepository) Create(member *domain.Member) {
	mr.db.Create(member)
}
