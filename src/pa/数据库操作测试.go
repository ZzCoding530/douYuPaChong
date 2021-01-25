package pa

import (
	_ "github.com/go-sql-driver/mysql"
)

type NoteInfo struct {
	ID      int
	Writer  string
	Title   string
	Content string
}

//
//var dbbb *sql.DB
//
//func initDB() (err error) {
//	dsn := "root:root@tcp(127.0.0.1:3306)/Mybatis?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err = sql.Open("mysql", dsn)
//	HandleError(err, "database")
//	err = db.Ping()
//	return err
//
//}
//func Test() {
//	println("数据库操作测试")
//	err := initDB()
//	HandleError(err, "数据库链接失败")
//	if err == nil {
//		fmt.Println("链接成功")
//	}
//
//	//chaxundanju := `select * from note where id=?`
//	//row := db.QueryRow(chaxundanju, 3)
//
//	note111 := NoteInfo{
//		Writer:  "张22总诗",
//		Title:   "Mys222ql操作",
//		Content: "Go语2222言相关的数据库操作",
//	}
//
//	_, err = dbbb.Exec("insert into note (writer,title,content) values (?,?,?);", &note111.Writer, &note111.Title, &note111.Content)
//	HandleError(err, "插入语句")
//
//}
