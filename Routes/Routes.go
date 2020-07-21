package Routes

import (
        "github.com/nayaksamdeep/jwt-user/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
        r.LoadHTMLGlob("templates/*");

	v1 := r.Group("/v1")
	{
		v1.GET("/", Controllers.GetHomePage)
		v1.GET("ListURL", Controllers.GetAllUrlInfo)
		v1.POST("ConvertURL", Controllers.ConvertAUrl)
		v1.GET("tinyurl/:id", Controllers.RedirectAUrl)
		v1.PUT("tinyurl/:id", Controllers.UpdateAUrl)
		v1.DELETE("tinyurl/:id", Controllers.DeleteAUrl)


		v1.POST("LoginToken", Controllers.LoginToken)
		v1.POST("RefreshToken", Controllers.RefreshToken)
		v1.POST("LogoutToken", Controllers.LogoutToken)

		v1.GET("ListUsers", Controllers.ListUsers)
		v1.POST("CreateUser", Controllers.CreateUser)
		v1.GET("user/:id", Controllers.GetUser)
		v1.PUT("user/:id", Controllers.UpdateUser)
		v1.DELETE("user/:id", Controllers.DeleteUser)

	}

	return r
}
