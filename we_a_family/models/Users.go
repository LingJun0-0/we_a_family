package Models

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"we_a_family/we_a_family/global"
)

func Gen_title_by_num() gin.H {

	// 模拟查库
	num := rand.Intn(10)
	var value string
	if num <= 1 {
		value = "垃圾小胖"
	} else {
		value = "垃圾毛强"
	}
	return gin.H{
		"title": value,
	}
}

type Member struct {
	Id         int
	Username   string
	Password   string
	Status     string
	Created_at string
	Updated_at string
	Deleted_at sql.NullString
}

//按用户名查一个
func FindOneUser(username string) Member {
	var m Member
	err := global.DB.Where("username = ?", username).First(&m)
	if err != nil {
		fmt.Printf("findone data failed err :%s\n", err)
	}

	fmt.Printf("findone member info %v\n", m)
	return m
}

// 查询多条数据
func FindsAllUser() []Member {
	var s []Member

	rows := global.DB.Find(&s)
	if rows.Error != nil {
		fmt.Printf("findsData failed err:%s\n", rows.Error)
		return nil
	}
	//for i, row := range rows {
	//	var m Member
	//	if i < len(rows)-1 {
	//		s = append(s, row)
	//	} else {
	//		fmt.Printf("findsData scan failed err:%s\n", rows.Error)
	//		return nil
	//	}
	//	fmt.Printf("findsData User info %v\n", m)
	//	s = append(s, m)
	//}
	fmt.Printf("%v\n", s)

	return s
}

// 插入一条数据
//func InsertOneUser(username string, password string) (err error) {
//	// 增、改、删 使用Exec方法
//	exec, err := db.Exec("insert into members(username,password,status,created_at,updated_at,deleted_at) values (?,?,?,?,?,?)", username, password, 1, time.Now(), time.Now(), nil)
//	if err != nil {
//		fmt.Printf("exec insert failed err:%s\n", err)
//		return err
//	}
//	id, err := exec.LastInsertId() // 往表中最后追加一条数据
//	if err != nil {
//		fmt.Printf("exec insert failed err:%s\n", err)
//		return err
//	}
//	fmt.Printf("insert data id is : %d\n", id)
//	return nil
//}
//
//// 更新数据
//func UpdateOneUser(id int, username, password string) {
//	ret, err := db.Exec("update members set username = ?,password = ? where id = ?", username, password, id)
//	if err != nil {
//		fmt.Printf("update failed err:%s\n", err)
//	}
//	affected, err := ret.RowsAffected() // 返回受影响的函数
//	if err != nil {
//		fmt.Printf("update RowsAffected failed err:%s\n", err)
//	}
//	fmt.Printf("update success rows:%d\n", affected)
//}
//
//// 删除用户
//func DelUser(id int) {
//	ret, err := db.Exec("delete from members where id = ?", id)
//	if err != nil {
//		fmt.Printf("del failed err:%s\n", err)
//		return
//	}
//	affected, err := ret.RowsAffected()
//	if err != nil {
//		fmt.Printf("get RowsAffected failed err:%s\n", err)
//		return
//	}
//	fmt.Printf("update success rows:%d\n", affected)
//}
//
//func Close() {
//	defer db.Close()
//}
