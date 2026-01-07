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
  )
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
  )
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