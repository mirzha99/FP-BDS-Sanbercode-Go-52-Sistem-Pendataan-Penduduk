package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/controller/auth"
	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/controller/home"
	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/controller/user"
	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/middleware"
)

func Router() {
	r := gin.Default()
	r.GET("/", home.Index)

	r.POST("/login", auth.Login)
	r.POST("/register", user.Add)
	r.GET("/profil", middleware.ReqAuth, auth.Profil)

	r.GET("/users", middleware.ReqAuth, user.Index)
	r.GET("/user/:id", middleware.ReqAuth, user.Byid)
	r.PUT("/user/:id", middleware.ReqAuth, user.Edit)
	r.DELETE("/user/:id", middleware.ReqAuth, user.Delete)

	r.Run(":3131")
}
