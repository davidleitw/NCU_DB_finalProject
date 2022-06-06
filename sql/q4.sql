SELECT course.course_name AS 課名, teacher.teacher_name AS 授課教師, SUM(
    CASE 
        WHEN select_record.select_result == "中選"
        THEN 1
        ELSE 0
    END) AS 中選人次, COUNT(select_record.select_result) AS 加選人次, ROUND(SUM(
    CASE
        WHEN select_record.select_result == "中選"
        THEN 1
        ELSE 0
    END)*100.0 / COUNT(select_record.select_result), 2) AS 中選比例
    FROM course, select_record, teacher
    WHERE course.cid == select_record.cid 
        AND course.tid == teacher.tid 
        AND course.semester == "1102"
    GROUP BY course.cid;