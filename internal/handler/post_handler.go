package handler

import (
	"errors"
	"net/http"
	"strconv"

	"go-blog/internal/model"
	"go-blog/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostHandler struct {
	db *gorm.DB
}

func NewPostHandler(db *gorm.DB) *PostHandler {
	return &PostHandler{db: db}
}

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=255"`
	Content string `json:"content" binding:"required"`
	Status  string `json:"status" binding:"omitempty,oneof=draft published"`
}

type UpdatePostRequest struct {
	Title   string `json:"title" binding:"omitempty,min=1,max=255"`
	Content string `json:"content" binding:"omitempty"`
	Status  string `json:"status" binding:"omitempty,oneof=draft published"`
}

type PostResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	AuthorID  uint   `json:"author_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func toPostResponse(post *model.Post) PostResponse {
	return PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		AuthorID:  post.AuthorID,
		Status:    post.Status,
		CreatedAt: post.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: post.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func getUserIDFromContext(c *gin.Context) (uint, error) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		return 0, errors.New("unauthorized")
	}
	userID, ok := userIDVal.(uint)
	if !ok {
		return 0, errors.New("invalid user id")
	}
	return userID, nil
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := getUserIDFromContext(c)
	if err != nil {
		if err.Error() == "unauthorized" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id"})
		}
		return
	}

	status := req.Status
	if status == "" {
		status = "draft"
	}

	post := &model.Post{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: userID,
		Status:   status,
	}

	if err := repository.CreatePost(h.db, post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "post created successfully",
		"post":    toPostResponse(post),
	})
}

func (h *PostHandler) GetPost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	post, err := repository.GetPostByID(h.db, uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": toPostResponse(post)})
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	var req UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := getUserIDFromContext(c)
	if err != nil {
		if err.Error() == "unauthorized" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id"})
		}
		return
	}

	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Content != "" {
		updates["content"] = req.Content
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no fields to update"})
		return
	}

	if err := repository.UpdatePost(h.db, uint(id), userID, updates); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found or unauthorized"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update post"})
		return
	}

	post, _ := repository.GetPostByID(h.db, uint(id))

	c.JSON(http.StatusOK, gin.H{
		"message": "post updated successfully",
		"post":    toPostResponse(post),
	})
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	userID, err := getUserIDFromContext(c)
	if err != nil {
		if err.Error() == "unauthorized" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id"})
		}
		return
	}

	if err := repository.DeletePost(h.db, uint(id), userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found or unauthorized"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post deleted successfully"})
}

func (h *PostHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/posts", h.CreatePost)
	r.PUT("/posts/:id", h.UpdatePost)
	r.DELETE("/posts/:id", h.DeletePost)
}

func (h *PostHandler) RegisterPublicRoutes(r *gin.RouterGroup) {
	r.GET("/posts/:id", h.GetPost)
}
