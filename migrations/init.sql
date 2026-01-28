CREATE TABLE faculty (
    faculty_id SERIAL PRIMARY KEY,
    faculty_name VARCHAR(100) NOT NULL
);
CREATE TABLE student_group (
    group_id SERIAL PRIMARY KEY,
    faculty_id INT REFERENCES faculty(faculty_id) NOT NULL,
    group_name VARCHAR(30) NOT NULL
);
CREATE TABLE student (
    student_id SERIAL PRIMARY KEY,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    email VARCHAR(30) UNIQUE NOT NULL,
    gender CHAR(1) NOT NULL CHECK (gender IN ('M', 'F')),
    birth_date DATE NOT NULL,
    group_id INT REFERENCES student_group(group_id) NOT NULL
);
CREATE TABLE professor (
    professor_id SERIAL PRIMARY KEY,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    email VARCHAR(30) UNIQUE NOT NULL,
    faculty_id INT REFERENCES faculty(faculty_id) NOT NULL
);
CREATE TABLE course (
    course_id SERIAL PRIMARY KEY,
    course_name VARCHAR(100) NOT NULL,
    faculty_id INT REFERENCES faculty(faculty_id) NOT NULL,
    professor_id INT REFERENCES professor(professor_id) NOT NULL
);
CREATE TYPE weekday_enum AS ENUM (
    'Monday',
    'Tuesday',
    'Wednesday',
    'Thursday',
    'Friday',
    'Saturday',
    'Sunday'
);
CREATE TABLE timetable (
    timetable_id SERIAL PRIMARY KEY,
    faculty_id INT REFERENCES faculty(faculty_id) NOT NULL,
    group_id INT REFERENCES student_group(group_id) NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    weekday weekday_enum NOT NULL,
    location VARCHAR(20),
    course_id INT REFERENCES course(course_id) NOT NULL
);
CREATE TABLE attendance (
    attendance_id SERIAL PRIMARY KEY,
    visited BOOLEAN NOT NULL,
    visit_day DATE NOT NULL,
    student_id INT NOT NULL REFERENCES student(student_id),
    course_id INT NOT NULL REFERENCES course(course_id),
    timetable_id INT NOT NULL REFERENCES timetable(timetable_id)
);
CREATE TABLE course_enrollment (
    enrollment_id SERIAL PRIMARY KEY,
    student_id INT NOT NULL REFERENCES student(student_id),
    course_id INT NOT NULL REFERENCES course(course_id)
);
INSERT INTO faculty (faculty_name)
VALUES ('School of Computing'),
    ('Business School');
INSERT INTO student_group (faculty_id, group_name)
VALUES (1, 'CS-2024-A'),
    (2, 'BA-2024-B');
INSERT INTO student (
        first_name,
        last_name,
        email,
        gender,
        birth_date,
        group_id
    )
VALUES (
        'Alice',
        'Smith',
        'alice.s@uni.edu',
        'F',
        '2003-05-15',
        1
    ),
    (
        'Bob',
        'Jones',
        'bob.j@uni.edu',
        'M',
        '2002-11-20',
        1
    ),
    (
        'Charlie',
        'Brown',
        'c.brown@uni.edu',
        'M',
        '2003-01-10',
        2
    );
INSERT INTO professor (
        first_name,
        last_name,
        email,
        faculty_id
    )
VALUES (
        'Dr. Emily',
        'Johnson',
        'emily.johnson@uni.edu',
        1
    ),
    (
        'Dr. Michael',
        'Williams',
        'michael.williams@uni.edu',
        2
    );
INSERT INTO course (course_name, faculty_id, professor_id)
VALUES ('Introduction to SQL', 1, 1),
    ('Marketing 101', 2, 2);
INSERT INTO timetable (
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
        '09:00:00',
        '11:00:00',
        'Monday',
        'Room 404',
        1
    ),
    (
        2,
        2,
        '13:00:00',
        '15:00:00',
        'Tuesday',
        'Hall B',
        2
    );
INSERT INTO attendance (
        visited,
        visit_day,
        student_id,
        course_id,
        timetable_id
    )
VALUES (TRUE, '2026-01-19', 1, 1, 1),
    (FALSE, '2026-01-19', 2, 1, 1),
    (TRUE, '2026-01-20', 3, 2, 2);
INSERT INTO course_enrollment (student_id, course_id)
VALUES (1, 1),
    -- Alice → Introduction to SQL
    (2, 1),
    -- Bob → Introduction to SQL
    (3, 2),
    -- Charlie → Marketing 101
    (3, 1);
-- Charlie also takes SQL
-- users
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL
);
INSERT INTO users (email, password)
VALUES ('admin@uni.edu', 'hashed_password_123'),
    ('alice.s@uni.edu', 'student_pass_456');