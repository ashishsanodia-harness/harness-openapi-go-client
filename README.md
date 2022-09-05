# Go API client for ngmanager

This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Steps to run GO client
Updated main.go by replacing <account> with Harness account value and <x-api-key-value> with x-api-key generated for same account. Also call any function (i.e swaggerCodeGenGetOrgs(), swaggerCodeGenGetAccountScopedSecret(), swaggerCodeGenGetAccountScopedSecrets() ) in main function and run
```
go mod tidy
```
```
go build . 
```
```
go run main.go
```

## Author

ashish.sanodia@harness.io
