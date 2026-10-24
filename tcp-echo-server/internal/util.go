package util

import (
	"fmt"
	"os"
)

func GetServerAddr() string {
	serverPort := os.Getenv("PORT")
	serverHost := "localhost"
	return fmt.Sprintf("%v:%v", serverHost, serverPort)
}
