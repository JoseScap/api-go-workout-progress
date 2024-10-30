create table exercises (
  id          uuid not null primary key default gen_random_uuid(),
  created_at  timestamp without time zone not null default now(),
  updated_at  timestamp without time zone not null default now(),
  "name"      varchar(128) not null unique,
  metric      varchar(128) not null,
  category    varchar(16) not null unique 
);