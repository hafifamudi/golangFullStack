package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//tangkap input dari user
	//mapping input dari user ke struct RegisterUserInput
	//struct di atas di passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.ApiResponse("Account Registered failed!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Account Registered failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//generate token
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.ApiResponse("Account Registered failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)
	response := helper.ApiResponse("Account has been created!", 200, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	/**
	user memasukan input (email & password)
	input di tangkap handler
	mapping dari input user ke input struct
	input struct passing ke service
	di service akan mencari dengan bantuan repository user dengan email
	kemudia jika di temui cocokan password
	**/
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Login failed!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ApiResponse("Login failed!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//generate token
	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.ApiResponse("Account Registered failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, token)

	response := helper.ApiResponse("Succesfully Logedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	//input email dari user
	//input di mapping ke dalam struct
	//struct input di passing ke dalam service
	//service memanggil repository -> check email sudah ada atau belum
	//repository - db

	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Email checking failed!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}

		response := helper.ApiResponse("Email checking failed!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": IsEmailAvailable,
	}

	metaMessage := "Email has been Registered!"

	if IsEmailAvailable {
		metaMessage = "Email is Available"
	}

	response := helper.ApiResponse(metaMessage, http.StatusUnprocessableEntity, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	//input dari user
	//simpan gambar pada direktory images
	//pada service panggil repo
	//JWT -> utk sementara hardcode (user ID = 1)
	//repo -> ambil user ID == 1
	//repo update lokasi file

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.ApiResponse("Failed to upload Avatar image!", http.StatusUnprocessableEntity, "error", data)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	extension := file.Filename
	path := "images/" + uuid.New().String() + extension

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uplaoded": false}

		response := helper.ApiResponse("Failed to upload Avatar image!", http.StatusUnprocessableEntity, "error", data)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userID := 1

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uplaoded": false}

		response := helper.ApiResponse("Failed to upload Avatar image!", http.StatusUnprocessableEntity, "error", data)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_uplaoded": true}

	response := helper.ApiResponse("Avatar successfully uploaded", http.StatusOK, "error", data)

	c.JSON(http.StatusOK, response)

}
