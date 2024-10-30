create table exercises (
    id          uuid not null primary key,
    created_at  timestamp without time zone not null,
    updated_at  timestamp without time zone not null,
    "name"        varchar(128) not null unique,
    metric      varchar(128)
);