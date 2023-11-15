package config

import "flag"

var (
	FlagUseBlobHash = flag.Bool("use-blob-hash", false, "Use blob hash as unique MsgID instead of previous_hash_id + shard_id. False by default.")
	FlagNatsContext = flag.String("nats", "", "NATS context to use")
	FlagServer      = flag.String("server", "", "NATS server to connect to")
	FlagCreds       = flag.String("creds", "", "NATS credentials file")
	FlagShardPrefix = flag.String("shard-prefix", "", "Prefix for shard subjects")
)
