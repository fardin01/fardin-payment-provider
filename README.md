# Fardin Payment Provider
This is a mock server that settles payment requests and returns a JSON object that contains `result` key with a boolean
value indicating if the payment has succeeded or not.

Being a mock/test server (and due to time constraints), it does not implement any of these critical features:
1. Authentication/Authorization
2. Input validation
3. Unit tests
4. Proper error handling
5. Proper logging
6. Configuration file
7. TLS/mTLS
8. API versioning

## Testing
At this point there are no unit tests for this code base due to time constraints. The tests can be done manually by running
a Docker container locally and calling the API. See the `Building` section below for running a local container.

## Building
A Docker image can be built locally by running `docker build .` in the root dir of this project. It will download all the 
dependencies, lint and finally compile the code.

A Docker compose file is also provided. It can be run by `docker-compose up` and you will get a container up and running 
with port `9000` exposed so you can call `localhost:9000/` to test the APIs. If you make changes to the code base, you might
need to run `docker-compose up --build` so that a new image is built with your new changes.

## Deploying

### Bash script

### Helm 3
