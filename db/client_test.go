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

func ExampleClientquery(){
	client, _ := db.New(":7398")

	s := client.Set("test", []byte("hello"))
	s = client.Set("test1", []byte("hello"))
	s = client.Set("test2", []byte("hello"))
	list ,_:= client.Iterator("test")
	fmt.Println(list)

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

