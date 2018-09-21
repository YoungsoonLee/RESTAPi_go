package models

import (
	"errors"
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	UserList map[string]*User
)

type User struct {
	Id                  int64     `orm:"pk"`
	Displayname         string    `orm:"size(16)";orm:"unique";`  // 4 ~ 16 letters
	Email               string    `orm:"size(100)";orm:"unique";` // max 100 letters
	Password            string    `orm:"size(1000)";orm:"null"`   // if account is provider, this column is null
	PasswordResetToken  string    `orm:"size(1000)";orm:"null"`
	PasswordResetExpire time.Time `orm:"null"`
	Confirmed           bool      `orm:"default(false)"`
	ConfirmResetToken   string    `orm:"size(1000)";orm:"null"`
	ConfirmResetExpire  time.Time `orm:"null"`
	Picture             string    `orm:"size(1000)";orm:"null"`
	Provider            string    `orm:"size(50)";orm:"null"` // google , fb
	ProviderID          string    `orm:"size(1000)";orm:"null"`
	ProviderAccessToken string    `orm:"size(1000)";orm:"null"`
	Permission          string    `orm:"size(50)";orm:"default(user)"`   // user, admin ...
	Status              string    `orm:"size(50)";orm:"default(normal)"` // normal, ban, close ...
	CreateAt            time.Time `orm:"auto_now_add;type(datetime)"`    // first save
	UpdateAt            time.Time `orm:"auto_now;type(datetime)"`        // eveytime save
}

// AddUser ...
func AddUser(u User) int64 {
	//u.Id, _ = strconv.ParseInt(strconv.FormatInt(time.Now().UnixNano(), 10), 10, 64)
	u.Id = time.Now().UnixNano()
	u.Displayname = "youngtip"
	u.Password = "1111"

	// save to db
	o := orm.NewOrm()
	_, err := o.Insert(&u)
	if err != nil {
		log.Println("insert error: ", err)
		return 0
	}

	//UserList[u.ID] = &u
	return u.Id
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Displayname != "" {
			u.Displayname = uu.Displayname
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		/*
			if uu.Profile.Age != 0 {
				u.Profile.Age = uu.Profile.Age
			}
			if uu.Profile.Address != "" {
				u.Profile.Address = uu.Profile.Address
			}
			if uu.Profile.Gender != "" {
				u.Profile.Gender = uu.Profile.Gender
			}
			if uu.Profile.Email != "" {
				u.Profile.Email = uu.Profile.Email
			}
		*/
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Displayname == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
