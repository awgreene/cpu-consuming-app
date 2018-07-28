# A GoLang Server used for performance testing
## Build and Run Docker Image
```bash
$ docker build -t performance-server .
$ docker run -it -p 8000:8000 --rm --name performance-server performance-server
```

```bash
$ curl -w "@curl-elapsed-time-format.txt" -o /dev/null -s "localhost:8000/sleep"
```
