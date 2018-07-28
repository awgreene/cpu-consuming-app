# A GoLang Server used for performance testing
A simple server with endpoints that simulate (very) simple performance problems such as lag or high cpu utilization.

## Build and Run Docker Image
```bash
$ docker build -t performance-server .
$ docker run -it -p 8000:8000 --rm --name performance-server performance-server
```
