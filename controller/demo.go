package controller

import (
	"demo/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Demo struct {
}

type RespJson struct {
	Name string `json:name`
}

func (a *Demo) Test(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("接收到的请求", string(reqBody))
	respJson := &RespJson{"wg_resp"}
	json, err := json.Marshal(respJson)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("resp", string(json))
	w.Write(json)
	return
}

func (a *Demo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//_, err := w.Write([]byte(r.URL.Path))
	db, err := database.GetMysqlDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	time1 := time.Now()
	rows, err := db.Query("select id from staff_select_group_online where id =?", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time.Now().Sub(time1).Microseconds())
	defer rows.Close()
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		fmt.Println(time.Now().Sub(time1).Microseconds())
		if err != nil {
			fmt.Println(err)
			return
		}
		result, err := json.Marshal(id)
		_, err = w.Write(result)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}
}
