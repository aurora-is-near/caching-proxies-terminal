package config

import "flag"

var (
	FlagUseBlobHash             = flag.Bool("use-blob-hash", false, "Use blob hash as unique MsgID instead of previous_hash_id + shard_id. False by default.")
	FlagServer                  = flag.String("server", "", "NATS server to connect to")
	FlagShardPrefix             = flag.String("shard-prefix", "", "Prefix for shard subjects")
	FlagSubmissionsVerifierHost = flag.String("submissions-verifier-host", "", "Submissions verifier host, from where JWT tokens to authorize in NATS will come. Format must be http://host:port/authenticate")
	FlagStoragePath             = flag.String("storage-path", "", "Path to the storage directory. If not provided, the storage will be in-memory")
)
