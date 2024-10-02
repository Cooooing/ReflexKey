--  配置表 config
create table if not exists `config`
(
    `id`         bigint auto_increment primary key,   -- id
    `env`        varchar(16)  not null,               -- 环境(all,win,linux,mac,android,ios)
    `key`        varchar(255) not null,               -- key
    `value`      varchar(255) not null,               -- value
    `insertTime` datetime default current_timestamp,  -- 插入时间
    `updateTime` datetime on update current_timestamp -- 更新时间
);

insert into config(env, key, value)
values ('all', 'database_version', '20240930');



