-- psql -U newuser -d newdb -h localhost -W

-- Books nomli table
CREATE TABLE Books (
	id SERIAL PRIMARY KEY NOT NULL,
	name VARCHAR(60) NOT NULL,
	price INT NOT NULL
);

-- Authors nomli table
CREATE TABLE Authors (
	id SERIAL PRIMARY KEY NOT NULL,
	last_name VARCHAR(32) NOT NULL,
    first_name VARCHAR(32) NOT NULL
);

-- BookAuthor nomli table
CREATE TABLE BookAuthor (
    book_id INT NOT NULL,
    author_id INT NOT NULL,
    FOREIGN KEY (book_id) REFERENCES Books(id),
    FOREIGN KEY (author_id) REFERENCES Authors(id)
);

-- Books nomli table ga ma'lumot kiritish
INSERT INTO Books (name, price) VALUES
    ('Book1', 20),
    ('Book2', 30),
    ('Book3', 25),
    ('Book4', 40),
    ('Book5', 15);

-- Authors nomli table ga ma'lumot kiritish
INSERT INTO Authors (last_name, first_name) VALUES
    ('Author1LastName', 'Author1FirstName'),
    ('Author2LastName', 'Author2FirstName'),
    ('Author3LastName', 'Author3FirstName'),
    ('Author4LastName', 'Author4FirstName'),
    ('Author5LastName', 'Author5FirstName');

-- BookAuthor nomli table ga ma'lumot kiritish (Many-to-Many bog'lanish)
INSERT INTO BookAuthor (book_id, author_id) VALUES
(1, 1),
(1, 2),
(2, 2),
(2, 3),
(3, 3),
(3, 4),
(4, 4),
(4, 5),
(5, 1),
(5, 3);

-- Bir authorni yozgan barcha kitoblarni chiqarish
SELECT 
    Books.*
FROM 
    Books
INNER JOIN 
    BookAuthor ON Books.id = BookAuthor.book_id
INNER JOIN 
    Authors ON BookAuthor.author_id = Authors.id
WHERE 
    Authors.first_name = 'Author1FirstName' 
AND 
    Authors.last_name = 'Author1LastName';


-- Bir kitobni yozgan barcha mualliflarni chiqarish
SELECT 
    Authors.*
FROM 
    Authors
INNER JOIN 
    BookAuthor ON Authors.id = BookAuthor.author_id
INNER JOIN 
    Books ON BookAuthor.book_id = Books.id
WHERE 
    Books.name = 'Book1';
