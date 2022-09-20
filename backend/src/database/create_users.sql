CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
drop table if exists users;
create table if not exists users
(
	id serial primary key,
	uid uuid unique default uuid_generate_v4(),
	username text not null,
	password text not null 
);
drop table if exists users_likes;
create table if not exists users_likes
(
	user_id int not null,
	liked uuid not null,
	foreign key (user_id) references users(id)
);
drop table if exists users_dislikes;
create table if not exists users_dislikes
(
	user_id int not null,
	disliked uuid not null,
	foreign key (user_id) references users(id)
);