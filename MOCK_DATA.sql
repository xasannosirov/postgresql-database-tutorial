-- YANGI USER VA DATABASE YARATISH
CREATE USER newuser WITH PASSWORD '1234';
CREATE DATABASE newdb;

-- POSTGRES USER BILAN NEW USERGA PERMISSION BERISH
sudo -u postgres psql
GRANT ALL PRIVILEGES ON DATABASE newdb TO newuser;
ALTER USER newuser CREATEDB;
\c newdb postgres
GRANT ALL ON SCHEMA public TO newuser;
\q 
psql -U newuser -d newdb -h localhost -W

-- USERNI O'CHIRISH
DROP USER IF EXISTS newuser;

-- DATABASE RO'YXATINI KO'RISH
\list 
\l

--DATABASE YARATISH
CREATE DATABASE newdb; 

--DATABASEGA ULANISH
\c newdb; 

-- DATABASENI O'CHIRIB YUBORISH
SELECT pg_terminate_backend(pid)
FROM pg_stat_activity
WHERE datname = 'newdb';

SELECT pg_terminate_backend(pg_stat_activity.pg_backend_pid)
FROM pg_stat_activity
WHERE pg_stat_activity.datname = 'newdb';

DROP DATABASE newdb;

-- TABLE YARATISH
CREATE TABLE IF NOT EXISTS newtalbe (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    email VARCHAR(60) NOT NULL,
    birthday DATE NOT NULL
);

-- TABLELARNI KO'RISH
\d 
\dt 

-- TABLEDAGI USTUNLARNI KO'RISH
\d newtalbe

-- TABLENI O'CHIRISH
DROP TABLE newtalbe;

-- JADVALGA MA'LUMOT YOZISH
INSERT INTO newtalbe (name, email, birthday) 
VALUES ('Ali', 'alijon92@gmail.com', DATE '1997-01-02');

--TABELGA YOZILGAN MA'LUMOTLARNI KO'RISH
SELECT * FROM newtalbe;
SELECT name AS first_name FROM newtalbe;
SELECT email, name FROM newtalbe;
SELECT birthday FROM newtalbe;

--CHIQISH
\q 

-- TOZALASH 
\! clear

-- FOYDALI SAYTLAR
-- https://www.postgresql.org/
-- https://mockaroo.com/
-- https://www.postgresqltutorial.com/ 
-- https://drawsql.app/diagrams
-- https://metanit.com/go/tutorial/10.3.php