### server

```
./etcd -listen-client-urls="http://0.0.0.0:2379" --advertise-client-urls="http://0.0.0.0:2379"
```

### client

```
./etcdctl set  mykey "this is awesome"
./etcdctl get  mykey
```