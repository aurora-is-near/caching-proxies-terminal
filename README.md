# Caching Proxies Terminal

## Overview
Caching Proxies Terminal is a Go-based service designed to interface with nearcore for block processing. It accepts HTTP requests to submit data chunks and leverages NATS for efficient message passing and storage. This service is a part of RPC speedup project and particularly useful for data chunks distribution to other NEAR rpc nodes.

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

- **`-use-blob-hash`** (boolean): Determines whether to use blob hash as the unique MsgID instead of combining `previous_hash_id` and `shard_id`. Default is `false`.
  ```bash
  ./caching-proxies-terminal -use-blob-hash=true
  ```
  
- **`-nats`** (string): Specifies the NATS context to use.
  ```bash
  ./caching-proxies-terminal -nats=[context]
  ```

- **`-server`** (string): Sets the NATS server to connect to.
  ```bash
  ./caching-proxies-terminal -server=[server_address]
  ```

- **`-creds`** (string): Path to the NATS credentials file.
  ```bash
  ./caching-proxies-terminal -creds=[path_to_credentials]
  ```

- **`-shard-prefix`** (string): Prefix for shard subjects.
  ```bash
  ./caching-proxies-terminal -shard-prefix=[prefix]
  ```

### Running the Service
- Start the service with default settings:
  ```bash
  ./caching-proxies-terminal
  ```
- To run with custom NATS context and shard prefix:
  ```bash
  ./caching-proxies-terminal -nats [context] -shard-prefix [prefix]
  ```
- Use the `-use-blob-hash=true` flag for unique MsgID generation based on blob hash.

### Interacting with the Service
- Submit a chunk for processing:
  ```bash
  curl -XPOST -d '[data]' localhost:1323/process
  ```

### Upcoming Changes
### Authentication Service 
- The `/process` endpoint will require a Bearer token for requests.
- Upon receiving a request, the service will:
  1. Attempt to authenticate the Bearer token by contacting the [Submission Verifier](https://github.com/aurora-is-near/caching-proxies-verifier) service.
  2. Once authenticated, the "Submission Verifier" will provide a valid NATS JWT token, which will be stored by the terminal service.
  3. This NATS JWT token will have an expiration date and will be used to store data in NATS.
  4. The terminal will not request a new JWT token if it already has a valid one. If the token is expired or missing, it will request a new one from the [Submission Verifier](https://github.com/aurora-is-near/caching-proxies-verifier).

### Configuration changes
####Deprecate
- **`-creds`** (string): Path to the NATS credentials file.
Authentication will use JWT token, aquired from verification service

####New options
- **`-verifier-server`** (string): Sets the [Submission Verifier](https://github.com/aurora-is-near/caching-proxies-verifier) server URL to connect to for Authentication.
  ```bash
  ./caching-proxies-terminal -verifier-server=[server_address]
  ```