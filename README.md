# fiber-storage-memcache-modtest
 demo an issue with github.com/gofiber/storage/memcache where updating dependancies causes an error
 
 ## details
 
 Related to https://github.com/gofiber/storage/issues/591

When updating dependancies using `go get -u` `github.com/gofiber/utils@v0.1.2` is upgraded to `github.com/gofiber/utils@v1.0.0` this causes the following error:

```
# github.com/gofiber/storage/memcache
../go/pkg/mod/github.com/gofiber/storage/memcache@v0.0.0-20221027071415-dca8f183e44b/memcache.go:24:36: undefined: utils.Trim
```
Should github.com/gofiber/utils go.mod be changed to avoid an dependancy update downloading a module with breaking API changes?

```
module github.com/gofiber/utils/v1

go 1.19

...
```
