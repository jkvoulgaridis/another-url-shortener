package mappings


import (
	"fmt"
	"api.com/url-short/redis_client"
	"github.com/xyproto/randomstring"
)

func CreateMapping(val string) string {
	key := randomstring.HumanFriendlyString(10)
	fmt.Printf("Key: %s\n", key)
	fmt.Printf("Val: %s\n", val)
	err := redis_client.SetItem(key, val)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create Mapping...\n")
	return key
}

func GetMapping(key string) string {
	val := redis_client.GetItem(key)
	fmt.Printf("Retrieving Mapping %s\n", val)
	return val
}
