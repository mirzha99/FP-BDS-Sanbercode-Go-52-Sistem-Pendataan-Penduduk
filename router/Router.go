package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/controller/auth"
	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/controller/desa"
	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/controller/home"
	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/controller/mukim"
	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/controller/user"
	_ "github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/docs"
	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})
	// @BasePath /api/v1

	// Home godoc
	// @Summary ping example
	// @Schemes
	// @Description do ping
	// @Tags example
	// @Accept json
	// @Produce json
	// @Success 200 {string} Helloworld
	// @Router /example/helloworld [get]
	r.GET("/", home.Index)
	//auth
	r.POST("/login", auth.Login)
	r.POST("/register", user.Add)
	r.GET("/profil", middleware.ReqAuth, auth.Profil)
	//user
	r.GET("/users", middleware.ReqAuth, user.Index)
	r.GET("/user/:id", middleware.ReqAuth, user.Byid)
	r.PUT("/user/:id", middleware.ReqAuth, user.Edit)
	r.DELETE("/user/:id", middleware.ReqAuth, user.Delete)
	//mukim
	r.GET("/mukims", middleware.ReqAuth, mukim.Index)
	r.GET("/mukim/:id", middleware.ReqAuth, mukim.GetById)
	r.POST("/mukim/", middleware.ReqAuth, mukim.Add)
	r.PUT("/mukim/:id", middleware.ReqAuth, mukim.Edit)
	r.DELETE("/mukim/:id", middleware.ReqAuth, mukim.Delete)
	//desa
	r.GET("/desas", middleware.ReqAuth, desa.Index)
	r.GET("/desa/:id", middleware.ReqAuth, desa.GetById)
	r.POST("/desa/", middleware.ReqAuth, desa.Add)
	r.PUT("/desa/:id", middleware.ReqAuth, desa.Edit)
	r.DELETE("/desa/:id", middleware.ReqAuth, desa.Delete)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":3131")
}
