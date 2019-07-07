# etcd

### server

```
./etcd -listen-client-urls="http://0.0.0.0:2379" --advertise-client-urls="http://0.0.0.0:2379"
```

### client

```
./etcdctl set  mykey "this is awesome"
./etcdctl get  mykey
```

# redis

```
docker start redis-test
```

# mysql
```
create database seckill;
use seckill;
create table product
(
    id     int auto_increment
        primary key,
    name   varchar(1024) not null,
    total  int default 0 null,
    status int default 0 null
) charset = utf8;

```
