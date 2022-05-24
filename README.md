# 資料庫期末報告

## 題目要求

某校教務處記錄每位學生每學期的修課情況、該科成績、教學評量，設計了以下的資料表來儲存所需資料(未進行正規化)

| #   | 欄位名稱         | 中文說明                                                         |
| --- | ---------------- | ---------------------------------------------------------------- |
| 1   | semester         | 學期別（1091意為109上學期，1092則為109下學期）                   |
| 2   | course_no        | 課程編號                                                         |
| 3   | course_name      | 課程名稱                                                         |
| 4   | course_type      | 選修別（必修／或選修）                                           |
| 5   | course_room      | 授課教室                                                         |
| 6   | course_building  | 授課地點                                                         |
| 7   | course_time      | 授課時間（一123：表示週一早上1-3節，一門課有多節次則以逗點隔開 |
| 8   | course_credit    | 學分數                                                           |
| 9   | course_limit     | 課程人數限制                                                     |
| 10  | course_status    | 課程狀態(開課/停課)                                              |
| 11  | course_is_online | 是否為線上課程                                                   |
| 12  | teacher_name     | 授課教師姓名                                                     |
| 13  | student_name     | 修課者姓名                                                       |
| 14  | student_dept     | 修課者系所                                                       |
| 15  | student_grade    | 修課者年級                                                       |
| 16  | student_status   | 學生在學狀態 (在學 / 休學 / 退學)                                |
| 17  | student_class    | 修課者班別                                                       |
| 18  | select_result    | 選課結果                                                         |
| 19  | course_score     | 該科總成績                                                       |
| 20  | feedback_rank    | 教學評量結果 (1-5分)                                             |

## 正規化結果

![](https://i.imgur.com/UDjxywL.png)

### Course

![](https://i.imgur.com/buMaFdd.png)

- 把課程抽出來，並且加上 `cid` 當作 Primary key
- 刪除 `course_is_online`，如果是線上課程， `course_room_name` 為 `null` 即可表達
- `course_type` 因為題目敘述只有表達必修/選修的用途，所以用 boolean 代替
- `course_status` 只有表達是否開課，所以同樣用 boolean 代替

### Course_location

![](https://i.imgur.com/ispE1dD.png)

- `Room_name` 作為 Primary key, 讓 `Room_building` 不用重複記，其實更有效率的作法是 `Room_name` 的字串裡面加上 building 的資訊，但是這樣感覺跟原本題目改太多了，就沒有進一步的優化

### Teacher

![](https://i.imgur.com/CbGvdzP.png)

- `tid` 作為 Primary key, 雖然題目沒有過多描述教師的欄位，但是一般情況會紀錄很多相關的基本資料，所以抽出來一個獨立的 table 並且配上 ID。

### Student

![](https://i.imgur.com/PVLKEQY.png)

- `sid` 為 Primary key
- `student_name` 的長度設 746，是目前世界紀錄擁有最長名字的人，~~防止例外事件~~
- `student_status` 在規格書中只有三種狀態(在學 / 休學 / 退學)，所以我們直接用 int 存，並且加上只能介於 -1 ~ 1 的限制。

### Select_record

![](https://i.imgur.com/BunuiIp.png)

- `select_record` 存放選課紀錄
    - 由 `cid` 以及 `sid` 分別代表課程以及學生
- `select_result` 的部份我們選擇用 int 存放，原因是正常選課會有**備取**機制，雖然這次我們的專題注重在資料庫的設計，並不需要考慮的實際邏輯的處理，但是我們預先留了可以實作**備取**機制的欄位
    - 中選 -> 存 0
    - 落選 -> -1
    - 備取 -> 大於 0 的數字代表備取順位

### Course_record

![](https://i.imgur.com/soRzWYo.png)

- `course_record` 存放的是正式選到課的紀錄
    - `rid` 為 Primary key

