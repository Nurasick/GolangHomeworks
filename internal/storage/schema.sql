CREATE TABLE genders (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL
);

insert into genders(name)
values ('male'),('female');


create table groups(
	id serial primary key,
	name varchar(20) not null unique,
	direction varchar(50) not null,
	course_year int not null
);

insert into groups(name,direction,course_year)
values
	('ENG-101','Engineering',1),
	('ENG-201','Engineering',2),
	('HUM-101','Engineering',1),
	('HUM-201','Engineering',2);


create table students(
	id serial primary key,
	name varchar(100) not null,
	birth_date date,
	year_of_study smallint not null,
	gender_id int references genders(id),
	group_id int references groups(id)
);

insert into students (name, birth_date, year_of_study, gender_id, group_id)
values
    ('Aigerim', '2005-03-12', 1, 2, 1),
    ('Ruslan',  '2004-10-02', 2, 1, 2),
    ('Dana',    '2005-01-25', 1, 2, 3),
    ('Arman',   '2003-09-14', 3, 1, 2),
    ('Alina',   '2004-12-01', 2, 2, 4);

create table class_schedule(
	id serial primary key,
	group_id int references groups(id),
	subject_id int references subjects(id),
	day_of_week int not null,
	starts_at time not null,
	ends_at time not null
);

insert into class_schedule (group_id, subject_id, day_of_week, starts_at, ends_at)
values
    (3, 1, 1, '09:00', '10:30'),
    (4, 2, 1, '10:45', '12:15'),
    (1, 3, 1, '13:00', '14:30'),
    (2, 4, 1, '14:45', '16:15'),
    (3, 5, 2, '09:00', '10:30'),
    (1, 6, 2, '13:00', '14:30');

create table subjects(
	id serial primary key,
	name varchar(100) not null unique
)

insert into subjects(name)
values
	('Mathematics'),
	('Programming'),
	('Physics'),
	('History'),
	('Literature'),
	('Physical Education');

create table attendance(
	id serial primary key,
	student_id int references students(id),
	subject_id int references subjects(id),
	visit_day date not null,
	visited boolean not null,
	unique(student_id, subject_id, visit_day)
);