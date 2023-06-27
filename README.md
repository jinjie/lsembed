# Litestream Embeded

Simple Go module to embed [litestream](https://github.com/benbjohnson/litestream) replication in your Go project.

Adapted from [example](https://github.com/benbjohnson/litestream-library-example/blob/1cee7706d435d241792e01c502f8f37747445d09/main.go)

## Get

`go get github.com/jinjie/lsembed`

## Example

```
// ...
replica := litestream.NewReplica(
    litestream.NewDB(app.DataDir()+"/data.db"),
    "s3",
)

replica.Client = &s3.ReplicaClient{
    AccessKeyID:     "ACCESSKEY",
    SecretAccessKey: "SECRETKEY",
    Bucket:          "litestream-test-bucket",
    Region:          "ap-southeast-1",
    Path:            "path",
}

lsdb, err := lsembed.Replicate(replica)
if err != nil {
    log.Fatal().Err(err).Msg("failed to replicate")
}

defer lsdb.Close()
//..
```
