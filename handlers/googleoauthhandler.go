package handlers

import (
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
    "github.com/gin-gonic/gin"
	"net/http"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
	"github.com/joho/godotenv"
)

type Config struct {
    GoogleLoginConfig oauth2.Config
}

type OAuthData struct {
	ID string `json:id`
	Email string `json:email`
	VerifiedEmail bool `json:verified_email`
	Name string `json:name`
	GivenName string `json:given_name`
	FamilyName string `json:family_name`
	Picture string `json:picture`
}

func GoogleConfig() oauth2.Config {
	var AppConfig Config
	err := godotenv.Load(".env")
	if err != nil {
        log.Fatalf("Some error occured. Err: %s", err)
    }
    AppConfig.GoogleLoginConfig = oauth2.Config{
        RedirectURL:  "http://localhost:8080/google_callback",
		ClientID:     os.Getenv("ClientID"),
        ClientSecret: os.Getenv("ClientSecret"),
        Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
            "https://www.googleapis.com/auth/userinfo.profile"},
        Endpoint: google.Endpoint,
    }

    return AppConfig.GoogleLoginConfig
}


func (config *Config) GoogleLogin(c *gin.Context) {

    url := config.GoogleLoginConfig.AuthCodeURL("randomstate")

    c.Status(http.StatusSeeOther)
    c.Redirect(http.StatusSeeOther,url)
    c.JSON(http.StatusSeeOther,url)
	return
}

func (config *Config) GoogleCallback(c *gin.Context) {
	var UData OAuthData
    state := c.Query("state")
    if state != "randomstate" {
        c.String(http.StatusOK,"States don't Match!!")
    }

    code := c.Query("code")

    googlecon := config.GoogleLoginConfig

    token, err := googlecon.Exchange(c, code)
    if err != nil {
        c.String(http.StatusOK,"Code-Token Exchange Failed")
		return 
    }

    resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
    if err != nil {
        c.String(http.StatusOK,"User Data Fetch Failed")
		return 
    }

    userData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        c.String(http.StatusOK,"JSON Parsing Failed")
		return 
    }
	json.Unmarshal(userData,&UData)
    c.String(http.StatusOK,UData.Name)
	return
}