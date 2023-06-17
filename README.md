### longhorn-cli
longhorn-cli allows to automate your routine tasks by using it inside your automation.

Run locally
```bash
go run main.go --help
go run main.go -a http://127.0.0.1:8888/v1 volumes fstrim
```

Run from container image
```bash
docker run --rm -ti xdesigns/longhorn-cli:latest -h
docker run --rm -ti xdesigns/longhorn-cli:latest -a http://127.0.0.1:8888/v1 volumes fstrim
```

Image builds can be found on [Dockerhub](https://hub.docker.com/r/xdesigns/longhorn-cli/tags)

### Development

Build Docker image
```bash
BUILD_TAG=<your_tag>
CONTAINER_REPO=<your_container_repo>
docker build -t $CONTAINER_REPO:$BUILD_TAG .
docker push $CONTAINER_REPO:$BUILD_TAG
```
