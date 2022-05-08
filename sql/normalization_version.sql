CREATE TABLE course (
	course_no        varchar(20)  NOT NULL UNIQUE,
	course_name      varchar(255) NOT NULL,
	course_roomid    varchar(12), -- If course_room_name is NULL, mean that course is on Internet.
	semester         varchar(4)   NOT NULL,
	course_type      BOOLEAN	  NOT NULL,
	course_time      varchar(20)  NOT NULL,
	course_credit    INTEGER NOT NULL,
	course_limit     INTEGER CHECK (course_limit > 0),
	course_status    BOOLEAN      NOT NULL,
	tid              varchar(8)   NOT NULL,
);

CREATE TABLE course_location (
	room_id       varchar(12) NOT NULL UNIQUE,
	room_name     varchar(20),
	room_building varchar(20), 
);

CREATE TABLE teacher (
	tid          varchar(8) NOT NULL UNIQUE,
	teacher_name varchar(8) NOT NULL,
);

CREATE TABLE student (
	sid 	       varchar(12) PRIMARY KEY NOT NULL UNIQUE,
	student_name   varchar(20) NOT NULL,
	student_dept   varchar(20) NOT NULL,
	student_grade  INTEGER     NOT NULL CHECK (student_grade > 0),
	student_class  varchar(8)  NOT NULL,
	student_status INTEGER     NOT NULL CHECK (student_status BETWEEN 0 AND 2),
);

CREATE TABLE select_record (
	c_no varchar(20) NOT NULL,
	sid  varchar(8)  NOT NULL,
	-- @select_result
	-- @中選: 0
	-- @落選: -1 
	-- @備取順位: > 0
	select_result INTEGER NOT NULL CHECK (select_result > -2),
);

CREATE TABLE course_record (
	rid varchar(20) PRIMARY KEY NOT NULL UNIQUE,
	sid varchar(12) NOT NULL,
	cno varchar(20) NOT NULL,
	course_score INTEGER NOT NULL DEFAULT 0 CHECK (course_score BETWEEN 0 AND 100),
	feedback_rank INTEGER DEFAULT 1 CHECK (feedback BETWEEN 1 AND 5), 
);


