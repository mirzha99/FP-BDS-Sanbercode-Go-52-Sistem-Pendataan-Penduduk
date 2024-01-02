package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/controller/home"
	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/controller/user"
)

func Router() {
	r := gin.Default()

	r.GET("/", home.Index)
	r.GET("/user", user.Index)
	r.GET("/user/:id", user.Byid)
	r.POST("/user", user.Add)
	r.PUT("/user/:id", user.Edit)
	r.DELETE("/user/:id", user.Delete)
	r.Run(":3131")
}
