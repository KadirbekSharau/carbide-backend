package models

import (
	"gorm.io/gorm"
)

type Document struct {
    gorm.Model
    UserID                  uint   `gorm:"not null"`
    Filename                string `gorm:"not null"`
    EncryptedData           []byte `gorm:"not null"`
    BlockchainTransactionID string `gorm:"not null"`
}

type DocumentRepository struct {
    db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) *DocumentRepository {
    return &DocumentRepository{db: db}
}

func (r *DocumentRepository) GetAllDocumentsForUser(userID uint) ([]Document, error) {
    var docs []Document
    result := r.db.Where("user_id = ?", userID).Find(&docs)
    if result.Error != nil {
        return nil, result.Error
    }
    return docs, nil
}

func (r *DocumentRepository) GetDocumentByIDForUser(id uint, userID uint) (*Document, error) {
    var doc Document
    result := r.db.Where("id = ? and user_id = ?", id, userID).First(&doc)
    if result.Error != nil {
        return nil, result.Error
    }
    return &doc, nil
}

func (r *DocumentRepository) CreateDocumentForUser(userID uint, doc *Document) error {
    doc.UserID = userID
    result := r.db.Create(&doc)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (r *DocumentRepository) UpdateDocumentForUser(userID uint, doc *Document) error {
    result := r.db.Model(&doc).Where("user_id = ?", userID).Updates(&doc)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (r *DocumentRepository) DeleteDocumentByIDForUser(id uint, userID uint) error {
    result := r.db.Where("id = ? and user_id = ?", id, userID).Delete(&Document{})
    if result.Error != nil {
        return result.Error
    }
    return nil
}