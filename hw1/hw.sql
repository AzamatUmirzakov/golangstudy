DROP TABLE student;
DROP TABLE student_group;
DROP TABLE faculty;
DROP TABLE timetable;
CREATE TABLE faculty (
    faculty_id INT PRIMARY KEY,
    faculty_name VARCHAR(100) NOT NULL
);
CREATE TABLE student_group (
    group_id INT PRIMARY KEY,
    faculty_id INT REFERENCES faculty(faculty_id),
    group_name VARCHAR(30) NOT NULL
);
CREATE TABLE student (
    student_id INT PRIMARY KEY,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    email VARCHAR(30) UNIQUE NOT NULL,
    gender CHAR NOT NULL,
    birth_date DATE,
    group_id INT REFERENCES student_group(group_id)
);
CREATE TABLE timetable (
    timetable_id INT PRIMARY KEY,
    faculty_id INT,
    group_id INT,
    time VARCHAR(20),
    weekday VARCHAR(10),
    location VARCHAR(20),
    subject VARCHAR(50)
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
        time,
        weekday,
        location,
        subject
    )
VALUES (
        1,
        1,
        1,
        '9:00 - 10:00',
        'Monday',
        '7.103',
        'ENG 101'
    ),
    (
        2,
        1,
        2,
        '10:00 - 11:00',
        'Tuesday',
        '7.103',
        'ENG 101'
    ),
    (
        3,
        2,
        3,
        '12:00 - 13:00',
        'Wednesday',
        '8.109',
        'SOC 120'
    ),
    (
        4,
        2,
        4,
        '14:00 - 15:00',
        'Thursday',
        '6.120',
        'KAZ 303'
    );
SELECT *
from student
WHERE gender = 'F'
ORDER BY birth_date ASC;
ALTER TABLE timetable
ADD COLUMN professor VARCHAR(30);
UPDATE timetable
SET professor = 'Akhtar'
WHERE subject = 'ENG 101';
UPDATE timetable
SET professor = 'Donovan Cox'
WHERE subject = 'SOC 120';
UPDATE timetable
SET professor = 'Dr. Akhmetova'
WHERE subject = 'KAZ 303';
-- ALTER TABLE student DROP COLUMN email;