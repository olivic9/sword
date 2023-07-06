<h1 align="center">Sword</h1>
<p>
    <img src="https://shields.io/badge/Go-%5E1.19-red?logo=go">

</p>

## Prerequisites

<li><a href="https://golang.org/doc/install">Go</a></li>
<li><a href="https://www.docker.com/get-started">Docker</a></li>
<li><a href="https://docs.docker.com/compose/">Docker Compose</a></li>

## Local Setup

Make sure 3306 and 8190 ports are available locally


Build

Build api
```sh
make build/api
```

Build consumer
```sh
make build/consumer
```

Build both api and consumer

```sh
make build/apps
```

##Setup

Docker

Kafka and mysql ([db schema](database/schema.png)) included

make sure to import the mysql dump under database directory

```sh
make docker/up
```

## Usage

All endpoints are under http://localhost:8190 check <a href="https://documenter.getpostman.com/view/20483749/2s93zFWJro">Postman</a> for further documentation.
Valid user tokens are available [here](internal/tests/end_to_end/tokens.go)
## Run tests

```sh
make test
```

## Misc

Formatting
```sh
make format
```
Tidy
```sh
make tidy
```

## Author

👤 **Clayson Oliveira**

* Linkedin: [Clayson Oliveira](https://www.linkedin.com/in/clayson-oliveira-603a853b/)
* Github: [@olivic9](https://github.com/olivic9)