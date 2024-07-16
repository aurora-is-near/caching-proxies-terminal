build:
	go build .

example-run:
	./caching-proxies-terminal -nats lol -shard-prefix shards

example-run-hashed:
	./caching-proxies-terminal -nats lol -shard-prefix shards -use-blob-hash=true

example-curl:
	 curl -XPOST -d 'hello' localhost:1323/process\?previous_hash_id=kek\&shard_id=1\&token=token

example-subscribe:
	nats subscribe shards.1