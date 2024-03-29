package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/config"
	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/models/Muser"
	"golang.org/x/crypto/bcrypt"
)

// SuccessResponse adalah struktur respons sukses
type SuccessResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// ErrorResponse adalah struktur respons kesalahan
type ErrorResponse struct {
	Error  string `json:"error"`
	Detail string `json:"detail"`
}

// Login godoc
// @Summary Login pengguna
// @Description Melakukan login pengguna dan mengembalikan token otentikasi
// @Accept json
// @Produce json

// @Param loginInput body Muser.LoginInput true "User login credentials"
// @Success 200 {object} SuccessResponse "Sukses"

// @Router /login [post]
func Login(ctx *gin.Context) {
	var loginInput Muser.LoginInput
	var user Muser.User
	if err := ctx.ShouldBindJSON(&loginInput); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Where("username = ?", loginInput.Username).First(&user).Error; err != nil {
		ctx.JSON(400, gin.H{"error": "Username Not Found"})
		return
	}
	//compare hash password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Password Invalid"})
		return
	}
	//jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_login": user.PublicUser(),
		"exp":        time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenResult, err := token.SignedString([]byte(os.Getenv("key_secret")))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Token Fail Cretaed", "detail": err.Error()})
		return
	}

	//json
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("jwt-token", tokenResult, 3600*24*30, "", "", false, true)
	ctx.JSON(200, SuccessResponse{
		Message: "Welcome " + user.Nama,
		Token:   tokenResult,
	})

}
func Profil(ctx *gin.Context) {
	profil, err := ctx.Get("user")

	if !err {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	ctx.JSON(200, profil)
}
