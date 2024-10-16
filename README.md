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

	cache := cache.GoCacheMem()

	cache.Set("foo", "bar", 5*time.Second)

	fmt.Println("Sleep for 3 seconds")
	time.Sleep(3 * time.Second)

	val, exists := cache.Get("foo")
	fmt.Println(val, exists)

	fmt.Println("Sleep for 3 seconds (again)")
	time.Sleep(3 * time.Second)

	value, ex := cache.Get("foo")
	fmt.Println(value, ex)

}
```
