Build commands 

```
docker buildx build --no-cache --platform linux/amd64,linux/arm/v6 --push -t shabykov12/catboost-c-api:1.1.1-alpine -f Dockerfile.alpine .
```

```
docker buildx build --no-cache --platform linux/amd64,linux/arm/v5 --push -t shabykov12/catboost-c-api:1.1.1-golang-1.16.3-buster -f Dockerfile.golang.1.16.3-buster .
```

```
docker buildx build --no-cache --platform linux/amd64,linux/arm/v7,linux/arm64/v8 --push -t shabykov12/catboost-c-api:1.1.1-golang-1.16.3-streach -f Dockerfile.golang.1.16.3-streach .  
```
