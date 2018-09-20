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
	Id       int64 `orm:"pk"`
	Username string
	Password string
}

/*
func init() {

	UserList = make(map[string]*User)
	u := User{"user_11111", "astaxie", "11111"}
	UserList["user_11111"] = &u
}
*/

// AddUser ...
func AddUser(u User) int64 {
	//u.Id, _ = strconv.ParseInt(strconv.FormatInt(time.Now().UnixNano(), 10), 10, 64)
	u.Id = time.Now().UnixNano()
	u.Username = "youngtip"
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
		if uu.Username != "" {
			u.Username = uu.Username
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
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
