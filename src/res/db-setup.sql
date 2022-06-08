create schema if not exists `go_bookstore`;

use `go_bookstore`;

drop table if exists books;
create table books(
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `author` varchar(255) NOT NULL,
  `publication` varchar(255) NOT NULL
);