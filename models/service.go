package models

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type Service struct {
	Sid         string    `orm:"size(500);pk"`
	Key         string    `orm:"size(500);unique"`            //
	Description string    `orm:"size(500)"`                   //
	CreateAt    time.Time `orm:"auto_now_add;type(datetime)"` // first save
	CloseAt     time.Time `orm:"datetime;null"`               // s
}

func AddService(s Service) (string, error) {

	b := make([]byte, 8) //equals 16 charachters
	rand.Read(b)

	// make Id
	s.Sid = "S" + strconv.FormatInt(time.Now().UnixNano(), 10)
	s.Key = hex.EncodeToString(b)

	_, err := orm.NewOrm().Raw("INSERT INTO service (Sid, key, Description, Create_At) VALUES ($1, $2, $3, $4)", s.Sid, s.Key, s.Description, time.Now()).Exec()
	if err != nil {
		return "", err
	}
	/*
		// save to db
		sid, err := orm.NewOrm().Insert(&s)
		if err != nil {
			fmt.Println(sid, err)
			return "", err
		}
	*/

	return s.Sid, nil
}
