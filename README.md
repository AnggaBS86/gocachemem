## How to use

```go
// setup new cache
cache := NewMemCache()

// Set cache values `foo` with 3 seconds of lifetime
cache.Set("foo", "bar", 3*time.Second)
// Set cache values `foo` with 5 seconds of lifetime
cache.Set("baz", 123, 5*time.Second)

// Getting the cache
value, found := cache.Get("foo")

fmt.Println("Cache value :", value, " found :", found)

//sleep for 6 seconds
time.Sleep(6 * time.Second)

//These cache should not exists anymore because lifetime > 3 seconds
val, cacheExists := cache.Get("foo")
fmt.Println("Cache value :", val, " Exist :", cacheExists)

//These cache should not exists anymore because lifetime > 5 seconds
valBaz, cacheExistsBaz := cache.Get("baz")
fmt.Println("Cache value :", valBaz, " Exist :", cacheExistsBaz)
```
