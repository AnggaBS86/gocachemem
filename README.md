## How to use

```go
package main

import (
	"fmt"
	"time"

	cache "github.com/AnggaBS86/gocachemem"
)

func main() {
  //setup new cache
	cache := cache.NewCacheMem()

  //set `foo` cache for 5 seconds of lifetime
	cache.Set("foo", "bar", 5*time.Second)

  //sleep for 3 seconds --> check the `foo` lifetime
	fmt.Println("Sleep for 3 seconds")
	time.Sleep(3 * time.Second)

  //get the cache --> should be exists
	val, exists := cache.Get("foo")
	fmt.Println(val, exists)

  //then, sleep for 3 seconds again (total 6 seconds from start)
	fmt.Println("Sleep for 3 seconds (again)")
	time.Sleep(3 * time.Second)

  //get the cache value --> should be return nil, false
	value, ex := cache.Get("foo")
	fmt.Println(value, ex)
}

```
