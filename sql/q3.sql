SELECT course.course_name, student.sid, student.student_name, course_record.course_score
	FROM student, course, course_record
		WHERE course_record.sid == student.sid
			AND course_record.cid == course.cid
			AND course.semester == "1102"
			AND ((student.student_dept < 5 AND course_record.course_score < 60)
			OR (student.student_dept >= 5 AND course_record.course_score < 70))
	GROUP BY student.sid;