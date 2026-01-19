-- DROP TABLE IF EXISTS attendance CASCADE;
-- DROP TABLE IF EXISTS timetable CASCADE;
-- DROP TABLE IF EXISTS student CASCADE;
-- DROP TABLE IF EXISTS course CASCADE;
-- DROP TABLE IF EXISTS student_group CASCADE;
-- DROP TABLE IF EXISTS faculty CASCADE;
CREATE TABLE faculty (
    faculty_id SERIAL PRIMARY KEY,
    faculty_name VARCHAR(100) NOT NULL
);
CREATE TABLE student_group (
    group_id SERIAL PRIMARY KEY,
    faculty_id INT REFERENCES faculty(faculty_id),
    group_name VARCHAR(30) NOT NULL
);
CREATE TABLE student (
    student_id SERIAL PRIMARY KEY,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    email VARCHAR(30) UNIQUE NOT NULL,
    gender CHAR NOT NULL,
    birth_date VARCHAR(10) NOT NULL,
    group_id INT REFERENCES student_group(group_id)
);
CREATE TABLE course (
    course_id SERIAL PRIMARY KEY,
    course_name VARCHAR(100) NOT NULL
);
CREATE TABLE timetable (
    timetable_id SERIAL PRIMARY KEY,
    faculty_id INT REFERENCES faculty(faculty_id),
    group_id INT REFERENCES student_group(group_id),
    start_time TIME,
    end_time TIME,
    weekday VARCHAR(10),
    location VARCHAR(20),
    course_id INT REFERENCES course(course_id)
);
INSERT INTO course (course_id, course_name)
VALUES (1, 'ENG 101'),
    (2, 'SOC 120'),
    (3, 'KAZ 303');
CREATE TABLE attendance (
    attendance_id SERIAL PRIMARY KEY,
    student_id INT NOT NULL,
    course_id INT NOT NULL,
    visited BOOLEAN NOT NULL,
    visit_day VARCHAR(10) NOT NULL,
    FOREIGN KEY (student_id) REFERENCES student(student_id),
    FOREIGN KEY (course_id) REFERENCES course(course_id)
);
INSERT INTO faculty (faculty_id, faculty_name)
VALUES (1, 'Engineering'),
    (2, 'Гуманитарный');
INSERT INTO student_group (group_id, faculty_id, group_name)
VALUES (1, 1, 'Civil engineers'),
    (2, 1, 'Electrical engineers'),
    (3, 2, 'Sociologists'),
    (4, 2, 'Linguists');
INSERT INTO student (
        student_id,
        first_name,
        last_name,
        email,
        gender,
        birth_date,
        group_id
    )
VALUES (
        2,
        'Aisulu',
        'Saparova',
        'aisulu.saparova@nu.edu.kz',
        'F',
        '2003-04-12',
        1
    ),
    (
        3,
        'Bolat',
        'Kenesov',
        'bolat.kenesov@nu.edu.kz',
        'M',
        '2005-01-15',
        2
    ),
    (
        4,
        'Dana',
        'Maratova',
        'dana.maratova@nu.edu.kz',
        'F',
        '2005-11-08',
        1
    ),
    (
        5,
        'Yerkebulan',
        'Serikov',
        'yerkebulan.serikov@nu.edu.kz',
        'M',
        '2006-03-30',
        3
    ),
    (
        6,
        'Gulnara',
        'Zhumabaeva',
        'gulnara.zhumabaeva@nu.edu.kz',
        'F',
        '2005-07-14',
        2
    ),
    (
        7,
        'Nurlan',
        'Almasov',
        'nurlan.almasov@nu.edu.kz',
        'M',
        '2006-02-02',
        4
    ),
    (
        8,
        'Kamila',
        'Asanova',
        'kamila.asanova@nu.edu.kz',
        'F',
        '2005-07-07',
        4
    ),
    (
        9,
        'Bauyrzhan',
        'Ismailov',
        'bauyrzhan.ismailov@nu.edu.kz',
        'M',
        '2002-05-12',
        2
    ),
    (
        10,
        'Zarina',
        'Kuanysheva',
        'zarina.kuanysheva@nu.edu.kz',
        'F',
        '2003-11-09',
        3
    );
INSERT INTO timetable (
        timetable_id,
        faculty_id,
        group_id,
        start_time,
        end_time,
        weekday,
        location,
        course_id
    )
VALUES (
        1,
        1,
        1,
        '9:00:00',
        '10:00:00',
        'Monday',
        '7.103',
        1
    ),
    (
        2,
        1,
        2,
        '10:00:00',
        '11:00:00',
        'Tuesday',
        '7.103',
        1
    ),
    (
        3,
        2,
        3,
        '12:00:00',
        '13:00:00',
        'Wednesday',
        '8.109',
        2
    ),
    (
        4,
        2,
        4,
        '14:00:00',
        '15:00:00',
        'Thursday',
        '6.120',
        3
    );
SELECT *
from student
WHERE gender = 'F'
ORDER BY birth_date ASC;
ALTER TABLE timetable
ADD COLUMN professor VARCHAR(30);
UPDATE timetable
SET professor = 'Akhtar'
WHERE course_id = 1;
UPDATE timetable
SET professor = 'Donovan Cox'
WHERE course_id = 2;
UPDATE timetable
SET professor = 'Dr. Akhmetova'
WHERE course_id = 3;
-- ALTER TABLE student DROP COLUMN email;
-- homework 2
INSERT INTO student (
        student_id,
        first_name,
        last_name,
        email,
        gender,
        birth_date,
        group_id
    )
VALUES (
        11,
        'Azamat',
        'Umirzakov',
        'azamat.umirzakoff@gmail.com',
        'M',
        '2005-09-12',
        NULL
    );
INSERT INTO student (
        student_id,
        first_name,
        last_name,
        email,
        gender,
        birth_date,
        group_id
    )
VALUES (
        12,
        'Timur',
        'Tachka',
        'timur.tttachka@gmail.com',
        'M',
        '2008-01-12',
        NULL
    );
INSERT INTO student_group (group_id, faculty_id, group_name)
VALUES (5, 1, 'Mechanical engineers'),
    (6, 2, 'Anthropologists');
SELECT student.first_name,
    student.last_name,
    student_group.group_name
FROM student
    INNER JOIN student_group ON student.group_id = student_group.group_id;
SELECT student.first_name,
    student.last_name,
    student_group.group_name
FROM student
    LEFT JOIN student_group ON student.group_id = student_group.group_id;
SELECT student.first_name,
    student.last_name,
    student_group.group_name
FROM student
    RIGHT JOIN student_group ON student.group_id = student_group.group_id;
SELECT student.first_name,
    student.last_name,
    student_group.group_name
FROM student
    FULL OUTER JOIN student_group ON student.group_id = student_group.group_id;
-- users
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL
)