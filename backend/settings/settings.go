package settings

import (
	"github.com/joho/godotenv"
	"fmt"
)


func GetEnvVar(key string) string{
	var myEnv map[string]string
	myEnv, err := godotenv.Read(".env")
	if err != nil {
		fmt.Printf("Error in reading env")
	}
	return myEnv[key]
}