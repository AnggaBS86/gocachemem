# Gocachemem
Memory based cache library for Golang

## How to use

```go
package main

import (
	"fmt"
	"time"

	cache "github.com/AnggaBS86/gocachemem"
)

func main() {
	// initialize cache
	cache := cache.GoCacheMem()

	// set cache name `foo` with value `bar` for 5 seconds
	cache.Set("foo", "bar", 5*time.Second)

	// sleep for 3 seconds
	fmt.Println("Sleep for 3 seconds")
	time.Sleep(3 * time.Second)

	// check the cache lifetime --> cache name `foo` should be still exists
	val, exists := cache.Get("foo")
	fmt.Println(val, exists)

	//sleep again for 3 seconds --> total 6 seconds from the start
	fmt.Println("Sleep for 3 seconds (again)")
	time.Sleep(3 * time.Second)

	// check the cache lifetime --> the cache name `foo` should be not exists
	value, ex := cache.Get("foo")
	fmt.Println(value, ex) //return `<nil> false`
}
```
