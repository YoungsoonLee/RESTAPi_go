package models

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type Service struct {
	SID         string    `orm:"column(SID);size(500);pk"`    // service id
	Key         string    `orm:"size(500);unique"`            // key for encrypt
	Description string    `orm:"size(500)"`                   //
	CreateAt    time.Time `orm:"type(datetime);auto_now_add"` // first save
	CloseAt     time.Time `orm:"type(datetime);auto_now"`     // eveytime save
}

func AddService(s Service) (string, error) {

	b := make([]byte, 8) //equals 16 charachters
	rand.Read(b)

	// make Id
	s.SID = "S" + strconv.FormatInt(time.Now().UnixNano(), 10)
	s.Key = hex.EncodeToString(b)

	_, err := orm.NewOrm().Raw("INSERT INTO service (SID, key, Description, Create_At) VALUES ($1, $2, $3, $4)", s.SID, s.Key, s.Description, time.Now()).Exec()
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

	return s.SID, nil
}
