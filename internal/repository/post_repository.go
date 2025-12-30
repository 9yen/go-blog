package repository

import (
	"go-blog/internal/model"

	"gorm.io/gorm"
)

func CreatePost(db *gorm.DB, post *model.Post) error {
	return db.Create(post).Error
}

func GetPostByID(db *gorm.DB, id uint) (*model.Post, error) {
	var post model.Post
	if err := db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func GetPostsByAuthor(db *gorm.DB, authorID uint) ([]model.Post, error) {
	var posts []model.Post
	if err := db.Where("author_id = ?", authorID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func UpdatePost(db *gorm.DB, postID uint, authorID uint, updates map[string]interface{}) error {
	result := db.Model(&model.Post{}).Where("id = ? AND author_id = ?", postID, authorID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func DeletePost(db *gorm.DB, postID uint, authorID uint) error {
	result := db.Where("id = ? AND author_id = ?", postID, authorID).Delete(&model.Post{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
