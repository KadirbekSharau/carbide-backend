package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/KadirbekSharau/carbide-backend/src/services"
	"github.com/gin-gonic/gin"
)

type DocumentController interface {
	CreateDocument(c *gin.Context)
	GetDocumentByUserId(ctx *gin.Context)
}

type documentController struct {
	service *services.DocumentService
}

func NewDocumentController(service *services.DocumentService) DocumentController {
	return &documentController{service: service}
}

// CreateDocument creates a new document with the provided information
func (c *documentController) CreateDocument(ctx *gin.Context) {
	name := ctx.PostForm("name")
	description := ctx.PostForm("description")
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}
	defer file.Close()

	userId, err := strconv.ParseUint(ctx.PostForm("user_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	document, err := c.service.CreateDocument(name, description, file, uint(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, document)
}

func (c *documentController) GetDocumentByUserId(ctx *gin.Context) {
	userId, err := strconv.ParseUint(ctx.Query("user_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	documentId := ctx.Param("id")
	url, err := c.service.GetUrlByUserIdAndId(userId, documentId)
	fmt.Println(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve document URL"})
		return
	}

	content, err := c.service.GetDocumentByUrl(url)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve document"})
		return
	}

	ctx.Data(http.StatusOK, "application/octet-stream", content)
}
