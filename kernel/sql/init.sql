--  配置表 config
create table if not exists `configs`
(
    `id`         integer primary key autoincrement not null, -- id
    `env`        varchar(16)                       not null, -- 环境(all,win,linux,mac,android,ios)
    `device`     varchar(16)                       not null, -- 设备标识
    `type`       varchar(16)                       not null, -- 类别
    `only_read`  integer  default 0                not null, -- 只读 0:可写 1:只读
    `key`        varchar(255)                      not null, -- key
    `value`      varchar(255)                      not null, -- value
    `created_at` datetime default current_timestamp,         -- 插入时间
    `updated_at` datetime default current_timestamp,         -- 更新时间
    `deleted_at` datetime default null                       -- 删除时间
);

insert into configs (env, device, type, only_read, key, value, created_at, updated_at, deleted_at)
select 'all'
     , 'all'
     , 'system'
     , 1
     , 'reflex_key_database_ver'
     , '20240930'
     , datetime('now')
     , datetime('now')
     , null
where not exists(select 1 from configs where key = 'reflex_key_database_ver');

--  账号密码表 account
create table if not exists `accounts`
(
    `id`         integer primary key autoincrement not null, -- id
    `title`      varchar(16),                                -- 标题
    `account`    varchar(255)                      not null, -- 账号
    `password`   varchar(255)                      not null, -- 密码
    `remark`     varchar(255),                               -- 备注
    `type`       varchar(255) default 'other'      not null, -- 类型
    `group`      varchar(255) default 'other'      not null, -- 分组
    `created_at` datetime     default current_timestamp,     -- 插入时间
    `updated_at` datetime     default current_timestamp,     -- 更新时间
    `deleted_at` datetime     default null                   -- 删除时间
);

