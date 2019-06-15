create table tabl_app_info (
    app_id int(11) auto_increment primary key,
    app_name varchar(1024) not null ,
    app_type varchar(64) not null ,
    create_time timestamp default current_timestamp,
    develop_path varchar(256) not null
) engine = innodb default charset = utf8 auto_increment = 1;

create table tbl_app_id(
    app_id int,
    ip varchar(64),
    key app_id_ip_index (app_id, ip)
) engine = innodb default charset =utf8 auto_increment = 1;


create table tbl_log_info(
    log_id int auto_increment primary key ,
    app_id varchar(1024) not null ,
    log_path varchar(64) not null ,
    topic  varchar(1024) not null ,
    create_time timestamp default current_timestamp,
    status tinyint default 1
) engine = innodb default charset = utf8 auto_increment=1;
