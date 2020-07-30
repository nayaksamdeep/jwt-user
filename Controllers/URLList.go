package Controllers

import (
	"net/http"
        "fmt"
        "github.com/nayaksamdeep/jwt-user/Models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
//	"strconv"
)

func GetHomePage (c *gin.Context) {
       c.HTML(
             http.StatusOK,
             "index.html",
             gin.H{
                 "title": "Home Page",
             },
       )
}

func GetAllUrlInfo(c *gin.Context) {

	var urlstruct []Models.RedirectUrl
	err := Models.GetAllUrl(&urlstruct)
	if err != nil {
                fmt.Println("Request aborted with Status Not Found")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
                fmt.Println("Request processed with Status OK")
		c.JSON(http.StatusOK, urlstruct)
	}
}

func ConvertAUrl(c *gin.Context) {
	var urlstruct Models.RedirectUrl
	id := c.Params.ByName("id") //Newly Added

        //Check if token is present
        metadata, err := ExtractTokenMetadata(c.Request)
        if err != nil {
                fmt.Println("Anonymous User")
        	val := c.ShouldBindWith(&urlstruct, binding.FormPost);
        	fmt.Println("Binding: ", val);
        } else {
        	userid, err := FetchAuth(metadata)
        	if err != nil {
                	c.JSON(http.StatusUnauthorized, err.Error())
                	return
        	}
//        	Id := strconv.FormatInt(userid, 16)
        	val := c.ShouldBindWith(&urlstruct, binding.FormPost);
        	fmt.Println("Binding: ", val);
    		//Assign the user associated with this request
		urlstruct.USERID = userid
        }

	err = Models.ConvertAUrl(&urlstruct)
	if err != nil {
                fmt.Println("Error: " , err);
		c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
	} else {
                var idstring = fmt.Sprint(urlstruct.ID);
                urlstruct.TinyUrl = "http://localhost:8080/v1/tinyurl/" + idstring;
                urlstr := "/v1?url=" +  urlstruct.Url + "&tinyurl=Here is your shortened URL: <font color=\"green\"> <a href=\"" + urlstruct.Url + "\" target=\"_blank\">" + urlstruct.TinyUrl + "</a></font>";
	        err := Models.UpdateAUrl(&urlstruct, id)
	        if err != nil {
		    c.JSON(http.StatusNotFound, urlstruct)
	        }
                c.Redirect(302, urlstr);
                fmt.Println("Request processed with Status OK")
	}
}

func RedirectAUrl(c *gin.Context) {
	var urlstruct Models.RedirectUrl
        var AuthCheck bool
	var UserId int64

	id := c.Params.ByName("id")

        //Check if token is present
        metadata, err := ExtractTokenMetadata(c.Request)
        if err != nil {
                fmt.Println("Anonymous User")
		goto GetRedirectStruct;
        } else {
                UserId, err = FetchAuth(metadata)
                if err != nil {
                	fmt.Println("Fetch Auth Failed")
                        c.JSON(http.StatusUnauthorized, err.Error())
                        return
                }
//                UserId = strconv.FormatInt(userid, 16)
	 	AuthCheck = true	
        }

GetRedirectStruct:
	err = Models.RedirectAUrl(&urlstruct, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
	} else {
		if (AuthCheck == true) {
			if (urlstruct.USERID != UserId) {
				c.AbortWithStatus(http.StatusNotFound)
                		fmt.Println("Request aborted with Status Mismatch in auth details")
				return
			}
		}
                urlstr := "http://" +  string(urlstruct.Url)
                fmt.Println("Request processed with Status OK for ", urlstr)
                c.Redirect(http.StatusPermanentRedirect, urlstr)
	}
}

func UpdateAUrl(c *gin.Context) {
	var urlstruct Models.RedirectUrl
	var UserId int64
	id := c.Params.ByName("id")

        //Check if token is present
        metadata, err := ExtractTokenMetadata(c.Request)
        if err != nil {
                fmt.Println("Anonymous Users can not create Custom URL")
		c.AbortWithStatus(http.StatusNotFound)
		return
        } else {
                UserId, err = FetchAuth(metadata)
                if err != nil {
                        fmt.Println("Fetch Auth Failed")
                        c.JSON(http.StatusUnauthorized, err.Error())
                        return
		}
        } 

	err = Models.GetUrl(&urlstruct, id)
	if err != nil {
		c.JSON(http.StatusNotFound, urlstruct)
		return
	}
	if (urlstruct.USERID != UserId) {
		c.JSON(http.StatusNotFound, urlstruct)
		return
	}
	c.BindJSON(&urlstruct)
	err = Models.UpdateAUrl(&urlstruct, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
		return
	} else {
		c.JSON(http.StatusOK, urlstruct)
                fmt.Println("Request processed with Status OK")
	}
}

func DeleteAUrl(c *gin.Context) {
	var urlstruct Models.RedirectUrl
	var UserId int64
	id := c.Params.ByName("id")

        //Check if token is present
        metadata, err := ExtractTokenMetadata(c.Request)
        if err != nil {
                fmt.Println("Anonymous Users trying to delete the URL")
        } else {
                UserId, err = FetchAuth(metadata)
                if err != nil {
                        fmt.Println("Fetch Auth Failed")
                        c.JSON(http.StatusUnauthorized, err.Error())
                        return
                }
	}

	urlstruct.USERID = UserId

	err = Models.DeleteAUrl(&urlstruct, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
                fmt.Println("Request processed with Status OK")
	}
}
