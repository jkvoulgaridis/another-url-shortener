package main

import (
	"api.com/url-short/mappings"
	"api.com/url-short/settings"
	"api.com/url-short/postgres"
	"net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "fmt"
    "time"
    // "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)


type GetMappingResponse struct{
	KEY string `json:"key"`
}

func CreateMapping(c *gin.Context){
	type RequestBody struct {
		OriginalUrl string `json:"url"`
	}
	var rb RequestBody
	err := c.ShouldBindJSON(&rb)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Printf("Got: %s\n", rb.OriginalUrl)
	resp := mappings.CreateMapping(rb.OriginalUrl)
	c.IndentedJSON(http.StatusOK, GetMappingResponse{KEY : resp})
}

func GetMapping(c *gin.Context){
	key := c.Params.ByName("key")
	fmt.Printf("Key to search : %s\n", key)
	resp := mappings.GetMapping(key)
	c.IndentedJSON(http.StatusOK, GetMappingResponse{KEY: resp})
}

func GoToMapping(c *gin.Context){
	key := c.Params.ByName("key")
	resp := mappings.GetMapping(key)
	c.Redirect(http.StatusFound, resp)
}

func Home(c *gin.Context){
	c.IndentedJSON(http.StatusOK, "Welcome to best URL shortener :)")
}


func prometheusHandler() gin.HandlerFunc {
    h := promhttp.Handler()

    return func(c *gin.Context) {
        h.ServeHTTP(c.Writer, c.Request)
    }
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
	    AllowOrigins:     []string{"*" },
	    AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
	    AllowHeaders:     []string{"*"},
	    ExposeHeaders:    []string{"Content-Length"},
	    AllowCredentials: true,
	    MaxAge: 12 * time.Hour,
 	}))
 	db, err := postgres.GetDbConnection()
 	if err != nil{
 		panic("failed to start db")
 	} else {
 		fmt.Println("Initializing DB, running migrations...")
	 	postgres.RunMigrations(db)
	 	postgres.CreateSamplesInDB(db)
 	}

	fmt.Printf("Hello from API :)\n")
	host := settings.GetEnvVar("HOST")
	port := settings.GetEnvVar("PORT")
	fmt.Printf(host)
	fmt.Printf(port)
	router.GET("/", Home)
	router.POST("/create-mapping/", CreateMapping)
	router.GET("/get-mapping/:key" , GetMapping)
	router.GET("go-to-mapping/:key", GoToMapping)
	router.GET("/metrics", prometheusHandler())
    router.Run(fmt.Sprintf("%s:%s", host, port))
}