-- table yaratish
CREATE TABLE Books (
	id SERIAL PRIMARY KEY NOT NULL,
	name VARCHAR(60) NOT NULL,
	price INT NOT NULL
);

-- tablega ma'lumot yozish
INSERT INTO Books (name, price) VALUES
    ('Book1', 20),
    ('Book2', 30),
    ('Book3', 25),
    ('Book4', 40),
    ('Book5', 15);