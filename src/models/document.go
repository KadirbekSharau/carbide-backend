package models

import (
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	UserID                  uint   `gorm:"not null"`
	FileName                string `gorm:"not null"`
	URL                     string `gorm:"not null"`
	BlockchainTransactionID string `gorm:"not null"`
}

type DocumentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) *DocumentRepository {
	return &DocumentRepository{db: db}
}

func (r *DocumentRepository) CreateDocument(document *Document) error {
	return r.db.Create(document).Error
}

func (r *DocumentRepository) GetUrlByUserIdAndId(userId uint64, id string) (string, error) {
	var url string
	if err := r.db.Table("documents").
		Where("user_id = ? AND id = ?", userId, id).
		Pluck("url", &url).Error; err != nil {
		return "", err
	}
	return url, nil
}
