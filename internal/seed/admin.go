package seed

import (
	"log"

	"go-blog/internal/config"
	"go-blog/internal/model"
	"go-blog/internal/utils"

	"gorm.io/gorm"
)

// SeedAdmin creates a default admin user if one doesn't exist.
// This is intended for development and demo environments only.
// The function is idempotent - safe to run multiple times.
func SeedAdmin(db *gorm.DB, cfg *config.Config) error {
	if !cfg.SeedAdmin {
		return nil
	}

	// Check if admin user already exists
	var existingUser model.User
	result := db.Where("email = ?", cfg.SeedAdminEmail).First(&existingUser)

	if result.Error == nil {
		log.Printf("Admin user already exists (%s), skipping seed", cfg.SeedAdminEmail)
		return nil
	}

	if result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(cfg.SeedAdminPass)
	if err != nil {
		return err
	}

	// Create admin user
	admin := &model.User{
		Username:     cfg.SeedAdminUser,
		Email:        cfg.SeedAdminEmail,
		PasswordHash: hashedPassword,
		Role:         "admin",
	}

	if err := db.Create(admin).Error; err != nil {
		return err
	}

	log.Printf("Default admin user created: %s / %s", cfg.SeedAdminEmail, cfg.SeedAdminPass)
	return nil
}
