create table if not exists company.company
(
    vid        int auto_increment
    primary key,
    id         text                                null,
    number     varchar(1000)                       null,
    name       varchar(1000)                       null,
    status     varchar(100)                        null,
    updated_by text                                null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP not null,
    created_by text                                null
    );

create table if not exists company.config
(
    id         int auto_increment
    primary key,
    company_id text null,
    `key`      text null,
    value      text null
);

