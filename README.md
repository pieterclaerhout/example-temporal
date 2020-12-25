# Temporal Example

Example on using [Temporal](https://temporal.io).

## Prerequisites

To run this example, you need to have [Docker](https://docker.io) and [Go](https://golang.org) installed.

## Startup

1. Startup the server by running:

    ```
    make run-server
    ```
    
2. Run the workers:

    ```
    make run-worker-transfer
    make run-worker-greeting
    ```

3. Fire off some tasks

    ```
    make run-start-transfer
    make run-start-greeting
    ```

4. Inspect via the [Web UI](http://localhost:8088).
