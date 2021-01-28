# fiber_graphql

Example project fiber api with [gqlgen](https://github.com/99designs/gqlgen) (graphql)

## Step

1. init by go mod
    ``` shell
    go mod init
    ```
1. get gqlgen
    ```
    go get github.com/99designs/gqlgen
    ```
1. Init project skeleton skeleton
    ```
    go run github.com/99designs/gqlgen init
    ```
1. Implement the code in resolver
1. If change the schema (xxxx.graphqls) run this command
    ```
    go run github.com/99designs/gqlgen generate
    ```
    