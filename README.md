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
$ curl -X POST -d "message=Hello%20Someone&to=someone" "localhost:8080/send"
{
  "body": "",
  "form": {
    "message": [
      "Hello Someone"
    ],
    "to": [
      "someone"
    ]
  },
  "headers": {
    "Accept": "*/*",
    "Content-Length": "34",
    "Content-Type": "application/x-www-form-urlencoded",
    "User-Agent": "curl/7.32.0"
  },
  "method": "POST",
  "url": "localhost:8080/send",
  "version": "HTTP/1.1"
}
```
