create database app;
use app;

create table if not exists users (
  id int  not null,
  username char not null,
  password char not null,
  created_at timestamp not null default NOW(),
  updated_at timestamp not null default NOW(),

  primary key (id)
);

create table if not exists metrics (
  id int not null,
  status int,
  updated_at timestamp not null default NOW(),

  primary key (id)
);


create table if not exists notifications (
  id int not null,

  primary key (id)
);
