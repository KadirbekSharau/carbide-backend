package services

import (
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	urll "net/url"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/KadirbekSharau/carbide-backend/src/blockchain"
	"github.com/KadirbekSharau/carbide-backend/src/models"
	"github.com/KadirbekSharau/carbide-backend/src/util"
)

type DocumentService struct {
	repo               *models.DocumentRepository
	cloudStorageClient *storage.Client
}

func NewDocumentService(repo *models.DocumentRepository, client *storage.Client) *DocumentService {
	return &DocumentService{
		repo:               repo,
		cloudStorageClient: client,
	}
}

// CreateDocument creates a new document with the provided filename and encrypted data
func (s *DocumentService) CreateDocument(name string, description string, file multipart.File, userId uint) (*models.Document, error) {
	// Generate a unique ID for the document
	id := blockchain.GenerateBlockchainTransactionID()

	ctx := context.Background()
	ciphertext, err := util.EncryptData(file)
	if err != nil {
		fmt.Println(err)
	}
	// Upload the encrypted file to Google Cloud Storage
	bucket := s.cloudStorageClient.Bucket("carbide-documents")
	bucketCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	object := bucket.Object(fmt.Sprintf("documents/%d/%s", userId, id))
	wc := object.NewWriter(bucketCtx)
	if _, err := wc.Write(ciphertext); err != nil {
		return nil, err
	}
	if err := wc.Close(); err != nil {
		return nil, err
	}

	// Get the URL of the object
	objectAttrs, err := object.Attrs(context.Background())
	if err != nil {
		return nil, err
	}
	url := objectAttrs.MediaLink

	// Create the document in the database
	document := &models.Document{
		FileName: name,
		URL:      url,
		UserID:   userId,
	}
	if err := s.repo.CreateDocument(document); err != nil {
		return nil, err
	}

	return document, nil
}

func (s *DocumentService) GetUrlByUserIdAndId(userId uint64, id string) (string, error) {
	return s.repo.GetUrlByUserIdAndId(userId, id)
}

func (s *DocumentService) GetDocumentByUrl(url string) ([]byte, error) {
	bucketName, objectPath, err := splitUrl(url)
	fmt.Println(bucketName)
	fmt.Println(objectPath)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	bucket := s.cloudStorageClient.Bucket(bucketName)
	object := bucket.Object(objectPath)
	reader, err := object.NewReader(context.Background())
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func splitUrl(url string) (string, string, error) {
	u, err := urll.Parse(url)
	fmt.Println(url)
	if err != nil {
		return "", "", err
	}
	path := strings.TrimLeft(u.Path, "/")
	return u.Host, path, nil
}
