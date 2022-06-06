SELECT course.course_name AS 課名, teacher.teacher_name AS 授課教師, SUM(course_record.feedback_rank) AS 教學評量總分, ROUND(AVG(course_record.feedback_rank)) AS 教學評量平均分數
    FROM course, course_record, teacher
    WHERE course.cid == course_record.cid 
        AND course.tid == teacher.tid 
        AND course.semester == "1102"
    GROUP BY course.course_name;