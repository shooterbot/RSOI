drop table if exists books;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table if not exists books
(
	id serial primary key,
	uid uuid unique default uuid_generate_v4(),
	name varchar(50),	
	publisher varchar(50),
	year int,
	likes int,
	dislikes int,
	status bool
);

drop table if exists books_tags;
create table if not exists books_tags
(
	book_id int,
	tag varchar(20)
);

insert into books(id, name, publisher, year, likes, dislikes, status) values(1, 'Город Грехов', 'Publisher1', 1991, 5, 1, true);
insert into books_tags(book_id, tag) values(1, 'Боевик');
insert into books(id, name, publisher, year, likes, dislikes, status) values(2, 'Cкотт Пилигрим', 'Publisher2', 2004, 120, 5, true);
insert into books_tags(book_id, tag) values(2, 'Юмор');
insert into books(id, name, publisher, year, likes, dislikes, status) values(3, 'Танкистка', 'Publisher3', 1988, 4, 1, true);
insert into books_tags(book_id, tag) values(3, 'Фантастика');
insert into books(id, name, publisher, year, likes, dislikes, status) values(4, 'Аннарасуманара', 'Publisher4', 2010, 95, 5, true);
insert into books_tags(book_id, tag) values(4, 'Детектив');
insert into books(id, name, publisher, year, likes, dislikes, status) values(5, 'Всеведущий читатель', 'Publisher5', 2004, 85, 2, false);
insert into books_tags(book_id, tag) values(5, 'Магия');
