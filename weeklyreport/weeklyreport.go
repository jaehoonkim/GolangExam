package main 

import (
    "net/http"
    "text/template"
    "fmt"
    "database/sql"
    _ "github.com/ziutek/mymysql/godrv"
)

const (
	DB_NAME = "WeeklyReport"
	DB_USER = "root"
	DB_PASS = "1212123"
)

type user_info struct {
	Idx int
	Id string
	Pw string
	Name string
}


func openDB() *sql.DB {
	db, err := sql.Open("mymysql", "WeeklyReport/root/1212123")
	if err != nil {
		fmt.Println("err...0")
		panic(err)
	}
	return db
}

func GenerateUserInfo(idx int, id string, pw string, name string) *user_info {
	return  &user_info{ Idx : idx, Id : id, Pw : pw , Name : name }
}


// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    indexPage := template.Must(template.ParseFiles("view/index.html"))
	db := openDB()
	defer db.Close()
	var UserInfo []*user_info
	
	rows, err := db.Query("select count(*) as count from user_info")
	if err != nil {
		panic(err)
	}
	//fmt.Println(">>>>>>>>>>>>>>>")
	var count int
	var idx int
	var id, pw, name string
	rows.Next()
	scan_err := rows.Scan(&count)
	//fmt.Println(count)
	if scan_err != nil {
		panic(scan_err)
	}
	UserInfo = make([]*user_info, count)
	rows1, err1 := db.Query("select * from user_info")
	if err1 != nil {
		panic(err1)
	}

	for i:= 0; i < count; i++ {
		rows1.Next()
		scan_err = rows1.Scan(&idx, &id, &pw, &name)
		if scan_err != nil {
			panic(scan_err)
		}
		//fmt.Println(idx)
		UserInfo[i] = GenerateUserInfo(idx, id, pw, name)
	}

	indexPage.Execute(w, UserInfo)
}

func reportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("reportHandler")

		name := r.FormValue("name")
		reportid := r.FormValue("reportid")
		completedReport := r.FormValue("completed")
		todo := r.FormValue("todo")

		fmt.Println(name, reportid, completedReport, todo)

		db := openDB()
		defer db.Close()

		_, err := db.Exec("insert into report (reportid, name, completed, todo) values (?, ?, ?, ?)", reportid, name, completedReport, todo)
		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/", 301)
	}
}

func main() {
    http.HandleFunc("/", defaultHandler)
    http.HandleFunc("/report", reportHandler)
    http.ListenAndServe(":8080", nil)
}

