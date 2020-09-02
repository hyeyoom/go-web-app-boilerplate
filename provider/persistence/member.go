package persistence

import (
	"github.com/hyeyoom/go-web-app-boilerplate/domain"
	"github.com/hyeyoom/go-web-app-boilerplate/provider"
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
	log := provider.GetLogger()
	log.Info(member)
	mr.db.Create(member)
}
