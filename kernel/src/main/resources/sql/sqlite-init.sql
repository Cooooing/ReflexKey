--  配置表 config
create table if not exists `configs`
(
    `id`           integer primary key autoincrement not null, -- id
    `env`          varchar(16)                       not null, -- 环境(all,win,linux,mac,android,ios)
    `device`       varchar(16)                       not null, -- 设备标识
    `type`         varchar(16)                       not null, -- 类别
    `only_read`    integer  default 0                not null, -- 只读 0:可写 1:只读
    `key`          varchar(255)                      not null, -- key
    `value`        varchar(255)                      not null, -- value
    `create_time`  datetime default current_timestamp,         -- 插入时间
    `update_time`  datetime default current_timestamp,         -- 更新时间
    `deleted_time` datetime default null                       -- 删除时间
);

insert into `configs` (`env`, `device`, `type`, `only_read`, `key`, `value`)
values ('all', 'all', 'database', 1, 'reflex_key_database_ver', '1');

-- 生成密码历史表
create table if not exists `generate_password`
(
    `id`           integer primary key autoincrement not null, -- id
    `password`     varchar(255)                      not null, -- 密码
    `create_time`  datetime default current_timestamp,         -- 插入时间
    `update_time`  datetime default current_timestamp,         -- 更新时间
    `deleted_time` datetime default null                       -- 删除时间
);


