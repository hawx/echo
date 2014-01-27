# echo

Simple echoing webserver.

``` bash
$ go run echo.go
2014/01/26 21:43:59 serving on :8080
```

...and elsewheres...

``` bash
$ curl localhost:8080
{
  "body": "",
  "form": {},
  "headers": {
    "Accept": "*/*",
    "User-Agent": "curl/7.32.0"
  },
  "method": "GET",
  "url": "localhost:8080/",
  "version": "HTTP/1.1"
}
```

## Routes

`/status/:code`:
  Return response with the given status code.

`/delay/:ms`:
  Return a response after `ms` milliseconds.
