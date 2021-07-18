CREATE TABLE links (
    id serial not null unique,
    long_url varchar not null unique,
    short_url varchar(10) not null unique
);