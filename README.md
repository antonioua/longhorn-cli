# Work in progress. Coming soon!

### longhorn-cli
longhorn-cli allows to automate some routine tasks.

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

### Development

Build Docker image
```bash
BUILD_TAG=<your_tag>
CONTAINER_REPO=<your_container_repo>
docker build -t $CONTAINER_REPO:$BUILD_TAG .
docker push $CONTAINER_REPO:$BUILD_TAG
```
