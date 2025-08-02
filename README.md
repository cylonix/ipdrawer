# IPDrawer

## What is IPDrawer?

IPDrawer is IP Address Management (IPAM) application and the code is written in Golang.

## Features

* Assign IP from the predefined IP pools.
* Support gRPC and REST API interfaces.
* Support assigning the same IP for the same UUID.
* Support multi-tenancy with namespace.
* Support assigning the specified IP address if available.

## API Docs

* [REST API Reference](/pkg/server/apiclient/README.md)
* `./ipdrawer start --redis-host localhost` and access to `http://localhost:25577/swagger-ui`

## Development

### Prerequisite

* Go
* make

### Generate a binary

```bash
make
./ipdrawer --help
```

### Generate proto files

Require [buf](https://buf.build/) and docker installed.

```bash
make proto
```

### Run test

```bash
make test
```

### (Option) CI/CD Pipeline

* buildspec.yml
  * build docker image and push to ECR
* create-pr.sh
  * Create PR in another repository for deploy (GitOps)

## Author

Forked from [https://github.com/hatena/ipdrawer](https://github.com/hatena/ipdrawer)
