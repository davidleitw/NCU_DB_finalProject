package main

import (
	"database/sql"
	"log"
	"strings"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type student struct {
	sid    string
	name   string
	dept   string
	grade  int
	class  string
	status int
}

type studentTempTable struct {
	name   string
	dept   string
	grade  int
	class  string
	status string
}

func getStudentStatus(status string) int {
	switch status {
	case "在學":
		return 1
	case "休學":
		return 0
	case "退學":
		return -1
	}
	return -2
}

func createCourseLocationTable(db *sql.DB) {
	stmt := `
	INSERT INTO course_location (room_name, room_building) 
		SELECT DISTINCT course_room, course_building
			FROM course_data;	
`
	_, err := db.Exec(stmt)
	if err != nil {
		panic(err)
	}
}

func createStudentTable(db *sql.DB) {
	rows, err := db.Query("SELECT DISTINCT student_name, student_dept, student_grade, student_class, student_status FROM course_data;")
	if err != nil {
		panic(err)
	}

	students := make([]student, 0)

	defer rows.Close()
	for rows.Next() {
		var s studentTempTable

		err = rows.Scan(&s.name, &s.dept, &s.grade, &s.class, &s.status)
		if err != nil {
			panic(err)
		}
		students = append(students, student{sid: uuid.NewString(), name: s.name, dept: s.dept, grade: s.grade, class: s.class, status: getStudentStatus(s.status)})
	}

	stmt := `
	INSERT INTO student (sid, student_name, student_dept, student_grade, student_class, student_status)
		VALUES (?, ?, ?, ?, ?, ?);`

	for _, s := range students {
		switch {
		case strings.HasSuffix(s.dept, "研究所"):
			s.grade += 4
		case strings.HasSuffix(s.dept, "博士班"):
			s.grade += 6
		default:
		}

		_, err := db.Exec(stmt, &s.sid, &s.name, &s.dept, &s.grade, &s.class, &s.status)
		if err != nil {
			panic(err)
		}
	}
}

func createTeacherTable(db *sql.DB) {
	rows, err := db.Query("SELECT DISTINCT teacher_name FROM course_data")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	tNames := make([]string, 0)
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			panic(err)
		}
		tNames = append(tNames, name)
	}

	stmt := `
	INSERT INTO teacher (tid, teacher_name)
		VALUES (?, ?)`

	for _, name := range tNames {
		_, err := db.Exec(stmt, uuid.NewString(), &name)
		if err != nil {
			panic(err)
		}
	}
}

type courseTempTable struct {
	semester    string
	name        string
	cType       string
	room        sql.NullString
	time        string
	credit      int
	limit       int
	status      string
	online      string
	teacherName string
}

func getCourseType(cType string) bool {
	return cType == "必修"
}

func getCourseStatus(status string) bool {
	return status == "開課"
}

func createCourseTable(db *sql.DB) {
	rows, err := db.Query("SELECT DISTINCT semester, course_name, course_type, course_room, course_time, course_credit, course_limit, course_status, course_is_online, teacher_name FROM course_data;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	courses := make([]courseTempTable, 0)
	for rows.Next() {
		var c courseTempTable

		err := rows.Scan(&c.semester, &c.name, &c.cType, &c.room, &c.time, &c.credit, &c.limit, &c.status, &c.online, &c.teacherName)
		if err != nil {
			panic(err)
		}
		courses = append(courses, c)
	}

	stmt := `
	INSERT INTO	course (cid, course_name, course_room_name, semester, course_type, course_time, course_credit, course_limit, course_status, tid)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, (
			SELECT tid
				FROM teacher
					WHERE teacher_name == ?
		));
`

	for _, course := range courses {
		courseID := uuid.NewString()
		// 直接去除 online 欄位，如果 course.room_name 為 null, 代表線上執行
		if course.online == "是" {
			course.room = sql.NullString{String: "", Valid: false}
		}

		// true 代表必修， false 代表選修
		courseType := getCourseType(course.cType)
		// true 代表開課， false 代表停開
		courseStatus := getCourseStatus(course.status)

		_, err := db.Exec(stmt, &courseID, &course.name, &course.room.String, &course.semester, &courseType, &course.time, &course.credit,
			&course.limit, &courseStatus, &course.teacherName)
		if err != nil {
			panic(err)
		}
	}
}

// 之後可以加入備取機制
func createSelectRecordTable(db *sql.DB) {
	stme := `
	INSERT INTO select_record (cid, sid, select_result)
		SELECT course.cid, student.sid, select_result
			FROM course_data
				JOIN course
					ON course_data.course_name == course.course_name
				JOIN student
					ON course_data.student_name == student.student_name;
	`
	_, err := db.Exec(stme)
	if err != nil {
		panic(err)
	}
}

type courseRecordTempTable struct {
	studentName string
	courseName  string
	courseScore float64
	feedback    int
}

func createCourseRecordTable(db *sql.DB) {
	rows, err := db.Query("SELECT student_name, course_name, course_score, feedback_rank FROM course_data WHERE course_score IS NOT null")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	records := make([]courseRecordTempTable, 0)
	for rows.Next() {
		var record courseRecordTempTable
		err := rows.Scan(&record.studentName, &record.courseName, &record.courseScore, &record.feedback)
		if err != nil {
			panic(err)
		}

		records = append(records, record)
	}

	stmt := `
	INSERT INTO course_record (rid, sid, cid, course_score, feedback_rank)
		VALUES (?, 
				(SELECT sid 
					FROM student
						WHERE student_name == ?),
				(SELECT cid
					FROM course
						WHERE course_name == ?),
				?, ?);
				
	`
	for _, record := range records {
		id := uuid.NewString()
		_, err := db.Exec(stmt, &id, &record.studentName, &record.courseName, &record.courseScore, &record.feedback)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	createCourseLocationTable(db)
	createStudentTable(db)
	createTeacherTable(db)
	createCourseTable(db)
	createSelectRecordTable(db)
	createCourseRecordTable(db)
}
