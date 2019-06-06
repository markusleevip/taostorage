package db_test

import (
	"fmt"
	"github.com/markusleevip/taostorage/db"
)

func ExampleClient() {
	client, _ := db.New(":7398")

	s := client.Set("test", []byte("hello"))
	value, _ := client.Get("test")

	fmt.Println(s)
	fmt.Println(string(value[:]))
	s = client.Del("test")
	fmt.Println(s)
	// Output:
	// <nil>
	// hello
	// <nil>

}
