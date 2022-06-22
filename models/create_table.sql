create table user
(
    id          bigint auto_increment
        primary key,
    user_id     bigint                              not null,
    username    varchar(64)                         not null,
    password    varchar(64)                         not null,
    email       varchar(64)                         null,
    gender      tinyint   default 0                 not null,
    create_time timestamp default CURRENT_TIMESTAMP null,
    update_time timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint idx_user_id
        unique (user_id),
    constraint idx_username
        unique (username)
)collate = utf8mb4_general_ci;

create table community
(
    id             int auto_increment
        primary key,
    community_id   int unsigned                        not null,
    community_name varchar(128)                        not null,
    introduction   varchar(256)                        not null,
    create_time    timestamp default CURRENT_TIMESTAMP not null,
    update_time    timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    constraint idx_community_id
        unique (community_id),
    constraint idx_community_name
        unique (community_name)
)collate = utf8mb4_general_ci;