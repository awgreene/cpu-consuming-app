# A slow GoLang Server
## Build and Run Docker Image
```bash
$ docker build -t slow-golang-server .
$ docker run -it -p 8000:8000 --rm --name slow-golang-server slow-golang-server
```

```bash
$ curl -w "@curl-elapsed-time-format.txt" -o /dev/null -s "localhost:8000/sleep"
```
