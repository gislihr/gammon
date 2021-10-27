create schema if not exists gammon;

create table gammon.player (
    id serial primary key,
    name varchar(256) not null,
    elo integer
);