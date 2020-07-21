package Controllers

import (
	"net/http"
        "fmt"
        "github.com/nayaksamdeep/jwt-user/Models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

	//val := c.BindJSON(&urlstruct)
        val := c.ShouldBindWith(&urlstruct, binding.FormPost);
        fmt.Println("Binding: ", val);
	err := Models.ConvertAUrl(&urlstruct)
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
	id := c.Params.ByName("id")
	err := Models.RedirectAUrl(&urlstruct, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
	} else {
                urlstr := "http://" +  string(urlstruct.Url)
                fmt.Println("Request processed with Status OK for ", urlstr)
//		c.JSON(http.StatusOK, urlstruct)
                c.Redirect(http.StatusPermanentRedirect, urlstr)
//                c.Redirect(302, "http://www.amazon.com")
	}
}

func UpdateAUrl(c *gin.Context) {
	var urlstruct Models.RedirectUrl
	id := c.Params.ByName("id")
	err := Models.GetUrl(&urlstruct, id)
	if err != nil {
		c.JSON(http.StatusNotFound, urlstruct)
	}
	c.BindJSON(&urlstruct)
	err = Models.UpdateAUrl(&urlstruct, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
	} else {
		c.JSON(http.StatusOK, urlstruct)
                fmt.Println("Request processed with Status OK")
	}
}

func DeleteAUrl(c *gin.Context) {
	var urlstruct Models.RedirectUrl
	id := c.Params.ByName("id")
	err := Models.DeleteAUrl(&urlstruct, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
                fmt.Println("Request processed with Status OK")
	}
}
