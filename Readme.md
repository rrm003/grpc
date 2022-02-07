# This is Golang GRPC server with REST client. #

* Source code : https://github.com/rrm003/grpc
* Docker image : https://hub.docker.com/repository/docker/rrm003/user-app/general

## Steps to interact with server

   - #### Get the docker image for app

```
docker pull rrm003/user-app:latest
```

   - #### Run app image

```
docker run --publish 8080:8080 rrm003/user-app
```

   - #### Start interacting with app with below api:
      Detail doumentaion for API : (https://rrm003.github.io/grpc)

1. Create User

```
   curl -X 'POST' \
   'http://localhost:8080/v1/user' \
   -H 'accept: application/json' \
   -H 'Content-Type: application/json' \
   -d '{
   "fname": "string",
   "age": 0,
   "city": "string",
   "phone": 0,
   "height": 0,
   "married": true
   }'
```

2. Get User Details

```
   curl -X 'GET' \
   'http://localhost:8080/v1/user/1' \
   -H 'accept: application/json'
```

3. Get All User Details

```
   curl -X 'GET' \
   'http://localhost:8080/v1/user' \
   -H 'accept: application/json'
```
