# ch-go test

```bash
go test ./
```

Result:

```
2023-09-06T17:30:56.639+0800    DEBUG   ch-go@v0.58.2/client.go:279     Flush   {"bytes": 47}
2023-09-06T17:30:56.639+0800    DEBUG   ch-go@v0.58.2/client.go:244     Packet  {"packet_code": 0, "packet": "Hello"}
2023-09-06T17:30:56.639+0800    DEBUG   ch-go@v0.58.2/handshake.go:79   Connected       {"protocol_version": 54460, "server.revision": 54465, "server.major": 23, "server.minor": 8, "server.patch": 2, "server.name": "ClickHouse (889fcc6f27e6, Asia/Shanghai) 23.8.2 (54465)", "client.protocol_version": 54460, "client.major": 0, "client.minor": 0, "client.patch": 0, "client.name": "clickhouse/ch-go (dev)"}
2023-09-06T17:30:56.639+0800    DEBUG   ch-go@v0.58.2/handshake.go:101  Writing addendum
2023-09-06T17:30:56.639+0800    DEBUG   ch-go@v0.58.2/client.go:279     Flush   {"bytes": 1}
2023-09-06T17:30:56.639+0800    DEBUG   ch-go@v0.58.2/query.go:72       sendQuery       {"query": "CREATE DATABASE IF NOT EXISTS otel ON CLUSTER default_cluster;", "query_id": "767839f5-95d5-4c28-a113-7b58fc337c65"}
2023-09-06T17:30:56.639+0800    DEBUG   ch-go@v0.58.2/client.go:279     Flush   {"bytes": 219}
2023-09-06T17:30:56.653+0800    DEBUG   ch-go@v0.58.2/client.go:244     Packet  {"packet_code": 1, "packet": "Data"}
2023-09-06T17:30:56.653+0800    DEBUG   ch-go@v0.58.2/query.go:234      Block   {"rows": 0, "columns": 6}
2023-09-06T17:30:56.753+0800    DEBUG   ch-go@v0.58.2/client.go:244     Packet  {"packet_code": 3, "packet": "Progress"}
2023-09-06T17:30:56.753+0800    DEBUG   ch-go@v0.58.2/query.go:427      Progress        {"rows": 0, "total_rows": 0, "bytes": 0, "wrote_bytes": 0, "wrote_rows": 0}
2023-09-06T17:30:56.753+0800    DEBUG   ch-go@v0.58.2/client.go:244     Packet  {"packet_code": 14, "packet": "ServerProfileEvents"}
2023-09-06T17:30:56.753+0800    DEBUG   ch-go@v0.58.2/query.go:234      Block   {"rows": 13, "columns": 6}
2023-09-06T17:30:56.753+0800    DEBUG   ch-go@v0.58.2/query.go:475      ProfileEvent    {"event.time": "2023-09-06T17:30:56.000+0800", "event.host_name": "889fcc6f27e6", "event.thread_id": 0, "event.type": "Increment", "event.name": "Query", "event.value": 1}
2023-09-06T17:30:56.753+0800    DEBUG   ch-go@v0.58.2/query.go:475      ProfileEvent    {"event.time": "2023-09-06T17:30:56.000+0800", "event.host_name": "889fcc6f27e6", "event.thread_id": 0, "event.type": "Increment", "event.name": "QueriesWithSubqueries", "event.value": 1}
2023-09-06T17:30:56.753+0800    DEBUG   ch-go@v0.58.2/query.go:475      ProfileEvent    {"event.time": "2023-09-06T17:30:56.000+0800", "event.host_name": "889fcc6f27e6", "event.thread_id": 0, "event.type": "Increment", "event.name": "NetworkReceiveElapsedMicroseconds", "event.value": 6}
2023-09-06T17:30:56.753+0800    DEBUG   ch-go@v0.58.2/query.go:475      ProfileEvent    {"event.time": "2023-09-06T17:30:56.000+0800", "event.host_name": "889fcc6f27e6", "event.thread_id": 0, "event.type": "Increment", "event.name": "NetworkSendElapsedMicroseconds", "event.value": 94}
2023-09-06T17:30:56.753+0800    DEBUG   ch-go@v0.58.2/query.go:475      ProfileEvent    {"event.time": "2023-09-06T17:30:56.000+0800", "event.host_name": "889fcc6f27e6", "event.thread_id": 0, "event.type": "Increment", "event.name": "NetworkSendBytes", "event.value": 129}
2023-09-06T17:30:56.753+0800    DEBUG   ch-go@v0.58.2/query.go:475      ProfileEvent    {"event.time": "2023-09-06T17:30:56.000+0800", "event.host_name": "889fcc6f27e6", "event.thread_id": 0, "event.type": "Increment", "event.name": "ZooKeeperTransactions", "event.value": 6}
2023-09-06T17:30:56.753+0800    DEBUG   ch-go@v0.58.2/query.go:475      ProfileEvent    {"event.time": "2023-09-06T17:30:56.000+0800", "event.host_name": "889fcc6f27e6", "event.thread_id": 0, "event.type": "Increment", "event.name": "ZooKeeperList", "event.value": 2}
2023-09-06T17:30:56.753+0800    DEBUG   ch-go@v0.58.2/query.go:475      ProfileEvent    {"event.time": "2023-09-06T17:30:56.000+0800", "event.host_name": "889fcc6f27e6", "event.thread_id": 0, "event.type": "Increment", "event.name": "ZooKeeperCreate", "event.value": 2}
2023-09-06T17:30:56.753+0800    DEBUG   ch-go@v0.58.2/query.go:475      ProfileEvent    {"event.time": "2023-09-06T17:30:56.000+0800", "event.host_name": "889fcc6f27e6", "event.thread_id": 0, "event.type": "Increment", "event.name": "ZooKeeperExists", "event.value": 1}
2023-09-06T17:30:56.753+0800    DEBUG   ch-go@v0.58.2/query.go:475      ProfileEvent    {"event.time": "2023-09-06T17:30:56.000+0800", "event.host_name": "889fcc6f27e6", "event.thread_id": 0, "event.type": "Increment", "event.name": "ZooKeeperMulti", "event.value": 1}
2023-09-06T17:30:56.754+0800    DEBUG   ch-go@v0.58.2/query.go:475      ProfileEvent    {"event.time": "2023-09-06T17:30:56.000+0800", "event.host_name": "889fcc6f27e6", "event.thread_id": 0, "event.type": "Increment", "event.name": "ContextLock", "event.value": 12}
2023-09-06T17:30:56.754+0800    DEBUG   ch-go@v0.58.2/query.go:475      ProfileEvent    {"event.time": "2023-09-06T17:30:56.000+0800", "event.host_name": "889fcc6f27e6", "event.thread_id": 0, "event.type": "Gauge", "event.name": "MemoryTrackerUsage", "event.value": 0}
2023-09-06T17:30:56.754+0800    DEBUG   ch-go@v0.58.2/query.go:475      ProfileEvent    {"event.time": "2023-09-06T17:30:56.000+0800", "event.host_name": "889fcc6f27e6", "event.thread_id": 0, "event.type": "Gauge", "event.name": "MemoryTrackerPeakUsage", "event.value": 0}
2023-09-06T17:30:56.757+0800    DEBUG   ch-go@v0.58.2/client.go:244     Packet  {"packet_code": 1, "packet": "Data"}
2023-09-06T17:30:56.757+0800    WARN    ch-go@v0.58.2/query.go:28       Cancel query
github.com/ClickHouse/ch-go.(*Client).cancelQuery
        /data/home/oliverdding/.local/share/go/pkg/mod/github.com/!click!house/ch-go@v0.58.2/query.go:28
github.com/ClickHouse/ch-go.(*Client).Do.func5
        /data/home/oliverdding/.local/share/go/pkg/mod/github.com/!click!house/ch-go@v0.58.2/query.go:691
golang.org/x/sync/errgroup.(*Group).Go.func1
        /data/home/oliverdding/.local/share/go/pkg/mod/golang.org/x/sync@v0.3.0/errgroup/errgroup.go:75
2023-09-06T17:30:56.757+0800    DEBUG   ch-go@v0.58.2/client.go:279     Flush   {"bytes": 2}
--- FAIL: TestSimplest (0.12s)
    db_test.go:34: decode block: decode block: raw block: got rows without target
FAIL
FAIL    github.com/oliverdding/ch-go-test       0.126s
FAIL
```
