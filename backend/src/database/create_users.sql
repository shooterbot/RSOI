drop table if exists users;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
create table if not exists users
(
	id serial primary key,
	uid uuid unique default uuid_generate_v4(),
	username text not null,
	password text not null 
);