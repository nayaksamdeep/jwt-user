package Controllers
  
import (
        "crypto/rand"
        "encoding/base64"
        "encoding/json"
        "io/ioutil"
        "log"
        "net/http"
        "os"
        "fmt"
        "github.com/nayaksamdeep/jwt-user/Models"
        "github.com/gin-gonic/gin"
        "github.com/gin-gonic/contrib/sessions"
        "golang.org/x/oauth2"
        "golang.org/x/oauth2/google"
        "github.com/go-redis/redis/v7"
)


// Credentials which stores google ids.
type Credentials struct {
        Cid     string `json:"cid"`
        Csecret string `json:"csecret"`
}

var cred Credentials
var conf *oauth2.Config

func getLoginURL(state string) string {
        return conf.AuthCodeURL(state)
}

func init() {
        //Initializing redis
        dsn := os.Getenv("REDIS_DSN")
        if len(dsn) == 0 {
                dsn = "localhost:6379"
        }
        client = redis.NewClient(&redis.Options{
                Addr: dsn, //redis port
        })
        _, err := client.Ping().Result()
        if err != nil {
                panic(err)
        }

        //Read the creds.json file
        file, err := ioutil.ReadFile("./creds.json")
        if err != nil {
                log.Printf("Missing Google Cred File error: %v\n", err)
                os.Exit(1)
        }
        if err := json.Unmarshal(file, &cred); err != nil {
                log.Println("unable to marshal data")
                return
        }

        conf = &oauth2.Config{
                ClientID:     cred.Cid,
                ClientSecret: cred.Csecret,
//                RedirectURL:  "http://127.0.0.1:9090/auth",
                RedirectURL:  "http://localhost:8080/v1/auth",
                Scopes: []string{
                        "https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
                },
                Endpoint: google.Endpoint,
        }
}

// RandToken generates a random @l length token.
func RandToken(l int) (string, error) {
        b := make([]byte, l)
        if _, err := rand.Read(b); err != nil {
                return "", err
        }
        return base64.StdEncoding.EncodeToString(b), nil
}


// AuthHandler handles authentication of a user and initiates a session.
func GoogleAuth(c *gin.Context) {
        // Handle the exchange code to initiate a transport.
        session := sessions.Default(c)
        retrievedState := session.Get("state")
        queryState := c.Request.URL.Query().Get("state")
        if retrievedState != queryState {
                log.Printf("Invalid session state: retrieved: %s; Param: %s", retrievedState, queryState)
//                c.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{"message": "Invalid session state."})
                c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid session state."})
                return
        }
        code := c.Request.URL.Query().Get("code")
        tok, err := conf.Exchange(oauth2.NoContext, code)
        if err != nil {
                log.Println(err)
//                c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Login failed. Please try again."})
                c.JSON(http.StatusBadRequest, gin.H{"message": "Login failed. Please try again."})
                return
        }

        hclient := conf.Client(oauth2.NoContext, tok)
        userinfo, err := hclient.Get("https://www.googleapis.com/oauth2/v3/userinfo")
        if err != nil {
                log.Println(err)
                c.AbortWithStatus(http.StatusBadRequest)
                return
        }
        defer userinfo.Body.Close()
        data, _ := ioutil.ReadAll(userinfo.Body)
//        u := structs.User{}
        var u Models.User
        if err = json.Unmarshal(data, &u); err != nil {
                log.Println(err)
//                c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error marshalling response. Please try agian."})
                c.JSON(http.StatusBadRequest, gin.H{"message": "Error marshalling response. Please try agian."})
                return
        }
/*
        session.Set("user-id", u.Email)
        err = session.Save()
        if err != nil {
                log.Println(err)
//                c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error while saving session. Please try again."})
                c.JSON(http.StatusBadRequest, gin.H{"message": "Error while saving session. Please try again."})
                return
        }
*/

        err = Models.CreateUser(&u)
        if err != nil {
//                c.AbortWithStatus(http.StatusNotFound)
                c.JSON(http.StatusBadRequest, gin.H{"message": "Error while saving user. Please try again."})
                fmt.Println("Request aborted with Status Not Found")
        } else {
                c.JSON(http.StatusOK, u)
                fmt.Println("Request processed with Status OK")
        }


/*
        seen := false
        db := database.MongoDBConnection{}
        if _, mongoErr := db.LoadUser(u.Email); mongoErr == nil {
                seen = true
        } else {
                err = db.SaveUser(&u)
                if err != nil {
                        log.Println(err)
                        c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error while saving user. Please try again."})
                        return
                }
        }

        c.HTML(http.StatusOK, "battle.tmpl", gin.H{"email": u.Email, "seen": seen})
*/
}

// GoogleLogin handles the login procedure.
func GoogleLogin(c *gin.Context) {
        state, err := RandToken(32)
        if err != nil {
//                c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"message": "Error while generating random data."})
                c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Error while generating random data"})
                return
        }

/*
        session := sessions.Default(c)
        session.Set("state", state)
        err = session.Save()
        if err != nil {
//                c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"message": "Error while saving session."})
                c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Error while saving session"})
                return
        }
*/
        link := getLoginURL(state)
//        c.HTML(http.StatusOK, "auth.tmpl", gin.H{"link": link})
//        c.JSON(http.StatusOK, gin.H{"link": link})
          c.Redirect(http.StatusPermanentRedirect, link)
}

