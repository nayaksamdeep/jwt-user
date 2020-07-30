package Controllers

import (
	"net/http"
        "fmt"
        "github.com/nayaksamdeep/jwt-user/Models"
	"github.com/gin-gonic/gin"
//	"github.com/gin-gonic/gin/binding"
        "strconv"
)

func ListUsers(c *gin.Context) {

	var userstruct []Models.User
	err := Models.ListUsers(&userstruct)
	if err != nil {
                fmt.Println("Request aborted with Status Not Found")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
                fmt.Println("Request processed with Status OK")
		c.JSON(http.StatusOK, userstruct)
	}
}

/*
 * Register User
 */
func RegisterUser(c *gin.Context) {
	var userstruct Models.User

        fmt.Println("Register User Function Enter")

        err := c.BindJSON(&userstruct)

        if err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Incorrect Field Name(s)"})
                fmt.Println("Register User aborted with Status Bad Request")
                return
        }
    
/*
        //Check if token is present
        metadata, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	userid, err := FetchAuth(metadata)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

        userstruct.ID = uint64(userid)
*/

        err = Models.CreateUser(&userstruct)
        if err != nil {
                c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
        } else {
                c.JSON(http.StatusOK, userstruct)
                fmt.Println("Request processed with Status OK")
        }

}

func GetUser(c *gin.Context) {
        var userstruct Models.User
        id := c.Params.ByName("id")

        //Check if token is present
        metadata, err := ExtractTokenMetadata(c.Request)
        if err != nil {
                c.JSON(http.StatusUnauthorized, "unauthorized")
                return
        }
        userid, err := FetchAuth(metadata)
        if err != nil {
                c.JSON(http.StatusUnauthorized, err.Error())
                return
        }
        Id := strconv.FormatInt(userid, 16)
        //compare the user from the request, with the one we defined:
        if (Id  != id) {
                c.JSON(http.StatusUnauthorized, "Please provide valid token details")
                fmt.Println("User id mismatch between the token and actuals")
                return
        }

        err = Models.GetUser(&userstruct, id)
        if err != nil {
                c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
        } else {
                fmt.Println("Request processed with Status OK")
                c.JSON(http.StatusOK, userstruct)
        }
}

func UpdateUser(c *gin.Context) {
        var userstruct Models.User
        id := c.Params.ByName("id")

        //Check if token is present
        metadata, err := ExtractTokenMetadata(c.Request)
        if err != nil {
                c.JSON(http.StatusUnauthorized, "unauthorized")
                return
        }
        userid, err := FetchAuth(metadata)
        if err != nil {
                c.JSON(http.StatusUnauthorized, err.Error())
                return
        }
        Id := strconv.FormatInt(userid, 16)
        //compare the user from the request, with the one we defined:
        if (Id  != id) {
                c.JSON(http.StatusUnauthorized, "Please provide valid token details")
                fmt.Println("User id mismatch between the token and actuals")
                return
        }

        err = Models.GetUser(&userstruct, id)
        if err != nil {
                c.JSON(http.StatusNotFound, userstruct)
        }
        c.BindJSON(&userstruct)
        err = Models.UpdateUser(&userstruct, id)
        if err != nil {
                c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
        } else {
                c.JSON(http.StatusOK, userstruct)
                fmt.Println("Request processed with Status OK")
        }
}

func DeleteUser(c *gin.Context) {
        var userstruct Models.User
        id := c.Params.ByName("id")
        err := Models.DeleteUser(&userstruct, id)
        if err != nil {
                c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
        } else {
                c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
                fmt.Println("Request processed with Status OK")
        }

}
