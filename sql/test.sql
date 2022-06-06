select COUNT(*), course.cid, student.sid, select_result
    from course_data
        join course
            on course_data.course_name == course.course_name
        join student
            on course_data.student_name == student.student_name;
