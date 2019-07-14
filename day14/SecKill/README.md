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
    id         int auto_increment
        primary key,
    new_column int           null,
    name       varchar(255)  not null,
    total      int default 0 null,
    status     int default 0 null,
    constraint product_name_uindex
        unique (name)
)
    charset = utf8;
    
    
create table `activity` (
    `id` int(11) not null auto_increment,
    `name` varchar(256) not null ,
    `product_id` int(11) default 0,
    `start_time` int(11) default 0,
    `end_time` int(11) default 0,
    `total` int(11) default 0,
    `status` int(11) default 0,
    `buy_limit` int(11) default 1,
    `sec_speed` int(11) default 100,
    `buy_rate` float default 0,
    primary key (`id`)
) engine = innodb charset = utf8    

```
