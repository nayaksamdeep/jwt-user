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

                v1.POST("RegisterUser", Controllers.RegisterUser) //New Route to add User
                v1.GET("ListUsers", Controllers.ListUsers)

                v1.POST("user/:id/LoginUser", Controllers.LoginUser)     //Provides Access Token and Refresh Token
                v1.POST("user/:id/RefreshUser", Controllers.RefreshUser) //Refreshes Token 
                v1.POST("user/:id/LogoutUser", Controllers.LogoutUser)   // Invalidates the token

		/*
                 * Use Google API for Login, We do not need to store the password
                 */
                v1.GET("GoogleLogin", Controllers.GoogleLogin) //New Route to add User
                v1.GET("auth", Controllers.GoogleAuth)

                v1.GET("user/:id", Controllers.GetUser)
                v1.PUT("user/:id", Controllers.UpdateUser)
                v1.DELETE("user/:id", Controllers.DeleteUser)

		v1.GET("ListURL", Controllers.GetAllUrlInfo)
		v1.POST("ConvertURL", Controllers.ConvertAUrl)
		v1.GET("tinyurl/:id", Controllers.RedirectAUrl)
		v1.PUT("tinyurl/:id", Controllers.UpdateAUrl)
		v1.DELETE("tinyurl/:id", Controllers.DeleteAUrl)

	}

	return r
}
