package controllers

import (
	"errors"
	"fmt"

	"github.com/RajaSunrise/crud-fiber/database"
	"github.com/RajaSunrise/crud-fiber/models/entity"
	"github.com/RajaSunrise/crud-fiber/models/req"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []entity.User
	if err := database.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mendapatkan daftar pengguna",
		})
	}
	return c.JSON(fiber.Map{
		"users": users,
	})
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entity.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Pengguna tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map{
		"user": user,
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(req.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	validation := validator.New()
	if err := validation.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"gagal input new user",
			"error": err.Error(),
		})
	}
	
	newUser := entity.User{
		Nama: user.Nama,
		Email: user.Email,
		Umur: user.Umur,
	}
	
	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Failed to create",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Bershasil Buat User",
	})
}


func UpdateUser(c *fiber.Ctx) error {

	id := c.Params("id")

	user := new(req.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	// Validate the user input
	validation := validator.New()
	if err := validation.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to update user",
			"error":   err.Error(),
		})
	}

	var existingUser entity.User
	if err := database.DB.Where("id = ?", id).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		// Other database error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve user",
			"error":   err.Error(),
		})
	}

	existingUser.Nama = user.Nama
	existingUser.Email = user.Email
	existingUser.Umur = user.Umur

	if err := database.DB.Save(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
	})
}


func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user entity.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Pengguna tidak ditemukan",
		})
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus user",
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Berhasil menghapus user dengan ID %s", id),
	})
}