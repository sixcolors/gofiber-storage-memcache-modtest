# fiber-storage-memcache-modtest
 demo an issue with github.com/gofiber/storage/memcache where updating dependancies causes an error
 
 ## details
 
 Related to https://github.com/gofiber/storage/issues/591

When updating dependancies using `go get -u` `github.com/gofiber/utils@v0.1.2` is upgraded to `github.com/gofiber/utils@v1.0.0` this causes the following error:

```
# github.com/gofiber/storage/memcache
../go/pkg/mod/github.com/gofiber/storage/memcache@v0.0.0-20221027071415-dca8f183e44b/memcache.go:24:36: undefined: utils.Trim
```

## the cause

v1.0.0 github.com/gofiber/utils has breaking API changes but the go.mod was not changed to avoid a dependancy update

```
module github.com/gofiber/utils/

go 1.19
...
```

it should have been

```
module github.com/gofiber/utils/v1

go 1.19
...
```

## testing a fix

@efectn proposed we release `github.com/gofiber/utils@v1.0.1` at the commit of tag `v0.1.2` and a create a new releae `github.com/gofiber/utils/v2@v2.0.0-beta.1` with the gofiber v3 changes.

v2 was chosen as v1 has already been released as a final release publically.

```
module github.com/gofiber/utils/v2

go 1.19
...
```

Tested and working for memcache sotrage package and new utils in this test project.
