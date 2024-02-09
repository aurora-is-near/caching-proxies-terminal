# Caching Proxies Terminal

## Overview

Caching Proxies Terminal is a Go-based service designed to interface with nearcore for block processing. It accepts HTTP
requests to submit data chunks and leverages NATS for efficient message passing and storage. This service is a part of
RPC speedup project and particularly useful for data chunks distribution to other NEAR rpc nodes.

## Installation

### Prerequisites

- Go 1.20 or higher
- Access to a NATS server

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/aurora-is-near/caching-proxies-terminal
   ```
2. Navigate to the project directory:
   ```bash
   cd caching-proxies-terminal
   ```
3. Build the project using the provided Makefile:
   ```bash
   make build
   ```

## Usage

## Configuration Options

The service offers several configuration options, which are set using command-line flags:

- **`-use-blob-hash`** (boolean): Determines whether to use blob hash as the unique MsgID instead of
  combining `previous_hash_id` and `shard_id`. Default is `false`.
  ```bash
  ./caching-proxies-terminal -use-blob-hash=true
  ```

- **`-nats`** (string): Specifies the NATS context to use.
  ```bash
  ./caching-proxies-terminal -nats=[context]
  ```

- **`-server`** (string): Sets the NATS server URL to connect to, overrides the context setting.
  ```bash
  ./caching-proxies-terminal -server=[server_address]
  ```

- **`-shard-prefix`** (string): Prefix for shard subjects. Messages will be published to `shard-prefix.shard_id` subject
  in the NATS.
  ```bash
  ./caching-proxies-terminal -shard-prefix=[prefix]
  ```

- **`-submissions-verifier-host`** (string): Sets the submissions verifier host and endpoint. Must be in format
  of `http://host:port/authenticate`
    ```bash
    ./caching-proxies-terminal -submissions-verifier-host=[verifier_host]
    ```

### Running the Service

- To run with custom NATS context and shard prefix:
  ```bash
  ./caching-proxies-terminal -nats [context] -shard-prefix [prefix] -submissions-verifier-host [verifier_host]
  ```
- Use the `-use-blob-hash=true` flag for unique MsgID generation based on blob hash.

### Interacting with the Service

- Submit a chunk for processing:
  ```bash
  curl -XPOST -H 'Authentication: Bearer [token]' -d '[data]' 'localhost:1323/process?shard_id=1&previous_hash_id=0x12348'
  ```

### Submissions verifier — Authentication Service

- The `/process` endpoint requires a Bearer token for requests.
- Upon receiving a request, the service will:
    1. Attempt to authenticate the Bearer token by contacting
       the [Submission Verifier](https://github.com/aurora-is-near/caching-proxies-verifier) service.
    2. Once authenticated, the "Submission Verifier" will provide a valid NATS JWT token, which will be used by the
       terminal service.
    3. This NATS JWT token will have an expiration date and will be used to access NATS to store the received chunk in it.
