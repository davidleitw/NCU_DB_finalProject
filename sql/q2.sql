SELECT student.sid, student.student_name, student.student_dept
    FROM student, course, course_record
        WHERE course_record.sid == student.sid 
            AND course_record.cid == course.cid
            AND course.course_name == "計算機概論"
            AND course.semester == "1102";