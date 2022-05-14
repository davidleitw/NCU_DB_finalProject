CREATE TABLE "course" (
	cid        		 VARCHAR(122)  NOT NULL UNIQUE,
	course_name      VARCHAR(255) NOT NULL,
	course_room_name VARCHAR(20), -- If course_room_name is NULL, mean that course is on Internet.
	semester         VARCHAR(4)   NOT NULL,
	course_type      BOOLEAN	  NOT NULL,
	course_time      VARCHAR(20)  NOT NULL,
	course_credit    INTEGER 	  NOT NULL,
	course_limit     INTEGER CHECK (course_limit > 0),
	course_status    BOOLEAN      NOT NULL,
	tid              VARCHAR(122)   NOT NULL
);

CREATE TABLE "course_location" (
	room_name     VARCHAR(20),
	room_building VARCHAR(20) 
);

CREATE TABLE "teacher" (
	tid          VARCHAR(122) NOT NULL UNIQUE,
	teacher_name VARCHAR(8) NOT NULL
);

 --- Reference https://reurl.cc/b29qoX
CREATE TABLE "student" (
	sid 	       VARCHAR(122)  PRIMARY KEY NOT NULL,
	student_name   VARCHAR(746) NOT NULL,
	student_dept   VARCHAR(20)  NOT NULL,
	student_grade  INTEGER      NOT NULL CHECK (student_grade > 0),
	student_class  VARCHAR(8)   NOT NULL,
	student_status INTEGER      NOT NULL CHECK (student_status BETWEEN -1 AND 1)
);

CREATE TABLE "select_record" (
	cid VARCHAR(122) NOT NULL,
	sid VARCHAR(122)  NOT NULL,
	-- @select_result
	-- @中選: 0
	-- @落選: -1
	-- @備取順位: > 0
	select_result INTEGER NOT NULL CHECK (select_result > -2)
);

CREATE TABLE "course_record" (
	rid           VARCHAR(122) PRIMARY KEY NOT NULL UNIQUE,
	sid 	      VARCHAR(122) NOT NULL,
	cid 		  VARCHAR(122) NOT NULL,
	course_score  REAL NOT NULL DEFAULT 0 CHECK (course_score BETWEEN 0.0 AND 100.0),
	feedback_rank INTEGER DEFAULT 1 CHECK (feedback_rank BETWEEN 1 AND 5)
);
