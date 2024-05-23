# VLR-API

Unofficial API for [vlr.gg](https://vlr.gg), a site for Valorant Esports coverage.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

The project is deployed on an Azure VM and can be accessed using the [url](http://vlr-api.centralindia.cloudapp.azure.com/)

### Prerequisites

What things you need to install the software and how to install them

```
// go - version 1.2 or higher
https://golang.org/doc/install

// gin - for http requests and routing
go get -u github.com/gin-gonic/gin

// colly - for web scraping
go get -u github.com/gocolly/colly

// go-redis - for caching data
go get -u github.com/go-redis/redis
```

### Installing

A step by step series of examples that tell you how to get a development env running

- Clone this repository

```
git clone https://github.com/LitoleNINJA/VLR-API.git
```

- Run the below command to install all the dependencies

```
go mod tidy
```

- If you don't have docker installed, first install docker from [here](https://docs.docker.com/get-docker/). This is required to run the redis server.

- When you have docker installed, you can run the below command to start both the redis server and the API server

```
docker-compose up
```

If everything goes well, output should look like this:

```bash
[+] Running 2/0
 ✔ Container api-redis  Created                                                                                                                                                                     0.0s 
 ✔ Container api        Created                                                                                                                                                                     0.0s 
Attaching to api, api-redis
api-redis  | 1:C 23 May 2024 15:01:26.401 # WARNING Memory overcommit must be enabled! Without it, a background save or replication may fail under low memory condition. Being disabled, it can also cause failures without low memory condition, see https://github.com/jemalloc/jemalloc/issues/1328. To fix this issue add 'vm.overcommit_memory = 1' to /etc/sysctl.conf and then reboot or run the command 'sysctl vm.overcommit_memory=1' for this to take effect.
api-redis  | 1:C 23 May 2024 15:01:26.401 * oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
api-redis  | 1:C 23 May 2024 15:01:26.401 * Redis version=7.2.4, bits=64, commit=00000000, modified=0, pid=1, just started
api-redis  | 1:C 23 May 2024 15:01:26.401 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
api-redis  | 1:M 23 May 2024 15:01:26.401 * monotonic clock: POSIX clock_gettime
api-redis  | 1:M 23 May 2024 15:01:26.406 * Running mode=standalone, port=6379.
api-redis  | 1:M 23 May 2024 15:01:26.407 * Server initialized
api-redis  | 1:M 23 May 2024 15:01:26.423 * Loading RDB produced by version 7.2.4
api-redis  | 1:M 23 May 2024 15:01:26.423 * RDB age 328538 seconds
api-redis  | 1:M 23 May 2024 15:01:26.423 * RDB memory usage when created 0.83 Mb
api-redis  | 1:M 23 May 2024 15:01:26.423 * Done loading RDB, keys loaded: 0, keys expired: 0.
api-redis  | 1:M 23 May 2024 15:01:26.423 * DB loaded from disk: 0.016 seconds
api-redis  | 1:M 23 May 2024 15:01:26.423 * Ready to accept connections tcp
api        | [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
api        |
api        | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
api        |  - using env:      export GIN_MODE=release
api        |  - using code:     gin.SetMode(gin.ReleaseMode)
api        |
api        | [GIN-debug] GET    /                         --> VLR-API/server.setAPIEndpoints.func1 (5 handlers)
api        | [GIN-debug] GET    /matches                  --> VLR-API/server.getMatches (5 handlers)
api        | [GIN-debug] GET    /matches/:id              --> VLR-API/server.getMatch (5 handlers)

```

## API Endpoints
| HTTP Verbs | Endpoints | Action |
| --- | --- | --- |
| GET | /api/matches | To get all the live and upcoming matches |
| GET | /api/matches?status=completed | To get all the completed matches |
| GET | /api/matches/:id | To get a specific match details |



## Built With

* [Golang](https://golang.org/) - The language used, because it's awesome
* [Gin](https://gin-gonic.com/) - HTTP web framework for Go, used for making http requests and routing
* [Colly](http://go-colly.org/) - Web scraping framework for Go, used for scraping match data from vlr.gg
* [Redis](https://redis.io/) - In-memory data structure store, used for caching match data to reduce load on the server

## Versioning

v1.0.0 - Initial release :)

## Authors
* [Ritwik Singh](https://github.com/LitoleNINJA) 

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details