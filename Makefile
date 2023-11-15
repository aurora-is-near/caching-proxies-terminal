build:
	go build main.go

example-run:
	./main -nats lol -shard-prefix shards

example-run-hashed:
	./main -nats lol -shard-prefix shards -use-blob-hash=true

example-curl:
	 curl -XPOST -d 'hello' localhost:1323/process\?previous_hash_id=kek\&shard_id=1

example-subscribe:
	nats subscribe shards.1