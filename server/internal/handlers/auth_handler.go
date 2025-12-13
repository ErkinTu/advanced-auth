package handlers

import (
	"AdvAuthGo/internal/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AssignRoleRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	RoleName string `json:"role_name" binding:"required"`
}

func setCookie(c *gin.Context, name, token string, expire time.Duration) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    token,
		Path:     "/api",
		Domain:   "",
		Expires:  time.Now().Add(expire),
		MaxAge:   int(expire.Seconds()),
		Secure:   false, // true in prod when using HTTPS
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.service.Register(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	setCookie(c, "access_token", tokens.AccessToken, 15*time.Minute)
	setCookie(c, "refresh_token", tokens.RefreshToken, 30*24*time.Hour)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered. Please activate your account.",
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	setCookie(c, "access_token", tokens.AccessToken, 15*time.Minute)
	setCookie(c, "refresh_token", tokens.RefreshToken, 30*24*time.Hour)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

func (h *AuthHandler) Activate(c *gin.Context) {
	// parse token from URL param
	token := c.Param("token")

	if err := h.service.Activate(token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account activated successfully"})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	tokens, err := h.service.Refresh(refreshToken)
	if err != nil {
		// delete cookie on invalid refresh
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "refresh_token",
			Value:    "",
			Path:     "/api",
			Domain:   "",
			Expires:  time.Unix(0, 0),
			MaxAge:   -1,
			Secure:   false,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	setCookie(c, "access_token", tokens.AccessToken, 15*time.Minute)
	setCookie(c, "refresh_token", tokens.RefreshToken, 30*24*time.Hour)

	c.JSON(http.StatusOK, gin.H{
		"message":      "Tokens refreshed",
		"access_token": tokens.AccessToken,
	})
}

func (h *AuthHandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *AuthHandler) GetUser(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, err := h.service.GetUserByToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *AuthHandler) CreateRole(c *gin.Context) {
	var body struct {
		Name string `json:"role_name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role name is required"})
		return
	}

	if err := h.service.CreateRole(body.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role created successfully"})
}

func (h *AuthHandler) DeleteRole(c *gin.Context) {
	var body struct {
		Name string `json:"role_name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.DeleteRole(body.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}

func (h *AuthHandler) AssignRoleToUser(c *gin.Context) {
	var req AssignRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AssignRoleToUser(req.UserID, req.RoleName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role assigned successfully"})
}
