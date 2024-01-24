CREATE TABLE IF NOT EXISTS Courses (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(60) NOT NULL,
    price INT NOT NULL,
    teacher VARCHAR(60) NOT NULL
);

CREATE TABLE IF NOT EXISTS Students (
    id SERIAL PRIMARY KEY NOT NULL,
    first_name VARCHAR(32) NOT NULL,
    last_name VARCHAR(32) NOT NULL,
    age INT NOT NULL
);

CREATE TABLE IF NOT EXISTS StudentCourse (
    student_id INT,
    course_id INT,
    FOREIGN KEY (student_id) REFERENCES Students(id),
    FOREIGN KEY (course_id) REFERENCES Courses(id)
);

INSERT INTO Courses (name, price, teacher) VALUES
('Syber Security', 3000000,'Saud Abdulwahed'), --1
('HR Suhbatlari', 2000000, 'Ali Zairov'), --2
('SMM', 1200000, 'Furqat Toxsanov'), --3
('Social Media Design', 1000000, 'Abdulloh Suratov'), --4
('Node JS Telegram Bot', 2500000, 'Abu Bakr Asadullayev'); --5

INSERT INTO Students (first_name, last_name, age) VALUES
('Asadbek', 'Faxriddinov', 17), --1
('Ahrorbek', 'Olimjonov', 21), --2
('Tohirjon','Odilov', 20), --3
('Akramjon', 'Abduvahobov', 16), --4
('Xasan', 'Nosirov', 17); --5

INSERT INTO StudentCourse (student_id, course_id) VALUES
(1, 1), --1
(1, 2), --2
(2, 2), --3
(2, 3), --4
(3, 3), --5
(3, 4), --6
(4, 4), --7
(4, 5), --8
(5, 5), --9
(5, 1); --10

-- kurs talabalari

SELECT 
    s.first_name AS StudentFirstName,
    s.last_name AS StudentLastName,
    s.age AS Age,
    c.name AS CourseNmae,
    c.teacher AS CourseTeacher,
    c.price AS CoursePrice
FROM
    Students s 
INNER JOIN
    StudentCourse sc ON s.id = sc.student_id
INNER JOIN 
    Courses c ON c.id = sc.course_id
WHERE
    c.name = 'SMM';

-- talaba kurslari

SELECT 
    s.first_name AS StudentFirstName,
    s.last_name AS StudentLastName,
    s.age AS Age,
    c.name AS CourseNmae,
    c.teacher AS CourseTeacher,
    c.price AS CoursePrice
FROM
    Students s 
INNER JOIN
    StudentCourse sc ON s.id = sc.student_id
INNER JOIN 
    Courses c ON c.id = sc.course_id
WHERE
    s.id = 1;