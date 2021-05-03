use gca;

INSERT INTO authors (name, birthdate) values ("Robert Cecil Martin", '1952-12-05');
INSERT INTO books (title, released_year) values ("Clean Code", 2008);
INSERT INTO books (title, released_year) values ("Clean Architecture: A Craftsman's Guide to Software Structure and Design", 2017);
INSERT INTO books (title, released_year) values ("Clean Agile: Back to Basics", 2019);
INSERT INTO book_authors (book_id, author_id) values (1, 1);
INSERT INTO book_authors (book_id, author_id) values (2, 1);
INSERT INTO book_authors (book_id, author_id) values (3, 1);

INSERT INTO authors (name, birthdate) values ("Andy Hunt", '1964-12-27');
INSERT INTO authors (name, birthdate) values ("Dave Thomas", '1960-01-03');
INSERT INTO books (title, released_year) values ("The Pragmatic Programmer", 1999);
INSERT INTO book_authors (book_id, author_id) values (4, 2);
INSERT INTO book_authors (book_id, author_id) values (4, 3);
INSERT INTO books (title, released_year) values ("Pragmatic Thinking and Learning: Refactor Your Wetware", 2008);
INSERT INTO books (title, released_year) values ("Learn to Program using Minecraft Plugins with Bukkit", 2014);
INSERT INTO book_authors (book_id, author_id) values (5, 2);
INSERT INTO book_authors (book_id, author_id) values (6, 2);

INSERT INTO authors (name, birthdate) values ("Martin Fowler", '1963-12-18');
INSERT INTO authors (name, birthdate) values ("Kent Beck", '1961-03-31');
INSERT INTO books (title, released_year) values ("Refactoring", 1999);
INSERT INTO book_authors (book_id, author_id) values (7, 4);
INSERT INTO book_authors (book_id, author_id) values (7, 5);
INSERT INTO books (title, released_year) values ("Planning Extreme Programming", 2000);
INSERT INTO book_authors (book_id, author_id) values (8, 4);
INSERT INTO book_authors (book_id, author_id) values (8, 5);

INSERT INTO books (title, released_year) values ("Test-Driven Development by Example", 2002);
INSERT INTO books (title, released_year) values ("JUnit Pocket Guide", 2004);
INSERT INTO books (title, released_year) values ("Implementation Patterns", 2008);
INSERT INTO book_authors (book_id, author_id) values (9, 5);
INSERT INTO book_authors (book_id, author_id) values (10, 5);
INSERT INTO book_authors (book_id, author_id) values (11, 5);