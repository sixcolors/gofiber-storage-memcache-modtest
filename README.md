# fiber-storage-memcache-modtest
 demo an issue with github.com/gofiber/storage/memcache where updating mods causes an error

When updating packages using `go get -u` `github.com/gofiber/utils@v0.1.2` is upgraded to `github.com/gofiber/utils@v1.0.0` this causes the following error:

```
# github.com/gofiber/storage/memcache
../go/pkg/mod/github.com/gofiber/storage/memcache@v0.0.0-20221027071415-dca8f183e44b/memcache.go:24:36: undefined: utils.Trim
```
