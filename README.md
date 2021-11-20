# keat/kit-service 📦🍽️📦🍽️

A service for meal kits, categories, categories, ingredients and such. Darn, what's excluded?: orders, payments, users

## Linting

[`golangci-lint`](https://golangci-lint.run) is used as an linter aggregator.

It runs as a [Github action](./.github/workflows/lint.yml) on the main branch. However, to run it locally, you'll need to first [install](https://golangci-lint.run/usage/install/#local-installation) and then run:

```
golangci-lint run
```

in the main directory.

## JWT Token

The [AWS API Gateway](https://aws.amazon.com/api-gateway/) should validate the JWT token, so the service should only extract the required information from the token (like if the requester is an _admin_ etc).

## See also

- [Hexagonal Architecture](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749)

- [Go Standard Project Layout](https://github.com/golang-standards/project-layout)
