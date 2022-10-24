create database jwt_sample;
\c jwt_sample;
create user penachett with encrypted password 'samplesapp908365';
create table users (
    id serial primary key,
    login text not null,
    password text not null,
    access_token text default '',
    access_token_expire bigint,
    refresh_token text default '',
    refresh_token_expire bigint);

grant all privileges on all tables in schema public to penachett;
grant all privileges on all sequences in schema public to penachett;
grant all privileges on all functions in schema public to penachett;
