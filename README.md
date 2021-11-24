# keat/kit-service 📦🍽️📦🍽️

A service for meal kits, categories, categories, ingredients and such. Darn, what's excluded?: orders, payments, users

## Env File

Create a `.env` file at the project root to store the required environment variables:

```
# .env
MONGO_URI=mongodb://localhost:27017/test
MONGODB_SRV_DB=default
```

## Linting

[`golangci-lint`](https://golangci-lint.run) is used as an linter aggregator.

It runs as a [Github action](./.github/workflows/lint.yml) on the main branch. However, to run it locally, you'll need to first [install](https://golangci-lint.run/usage/install/#local-installation) and then run:

```sh
golangci-lint run
```

in the main directory.

## Running Locally

`go run` can be used for standalone execution with a `.env` file added to the project root:

```sh
go run ./cmd/campaign/main.go --env ./.env
```

## Discussion

### JWT Token

The [AWS API Gateway](https://aws.amazon.com/api-gateway/) should validate the JWT token, so the service should only extract the required information from the token (like if the requester is an _admin_ etc).

## See also

- [Hexagonal Architecture](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749)

- [Go Standard Project Layout](https://github.com/golang-standards/project-layout)

- [NewRelic Go framework integrations](https://docs.newrelic.com/docs/apm/agents/go-agent/get-started/go-agent-compatibility-requirements/#frameworks)

- A reason for using `self` or `me` in REST URI when the user id can be extracted from the JWT and the user only manage data of themselves: https://softwareengineering.stackexchange.com/a/362064/374006
