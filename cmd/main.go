package main

import (
	"crypto/internal/client"
)

func main() {

	// uncomment to have the user point to a config
	// conf, err := configs.LoadConfigFromInput()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	client.StartPlaftormsClient()
	// http.ListenAndServe(":8080", nil)

}
