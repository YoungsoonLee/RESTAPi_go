package models

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/astaxie/beego"

	"github.com/YoungsoonLee/RESTAPi_go/libs"
	"github.com/astaxie/beego/orm"
	uuid "github.com/satori/go.uuid"
	"github.com/zoonman/gravatar"
	"golang.org/x/crypto/scrypt"
)

var (
	UserList map[string]*User
)

type User struct {
	Id                  int64      `orm:"pk"`
	Displayname         string     `orm:"size(16);unique"`  // 4 ~ 16 letters
	Email               string     `orm:"size(100);unique"` // max 100 letters
	Password            string     `orm:"null"`             // if account is provider, this column is null
	Salt                string     `orm:"null"`
	PasswordResetToken  string     `orm:"size(1000);null"`
	PasswordResetExpire *time.Time `orm:"null"`
	Confirmed           bool       `orm:"default(false)"`
	ConfirmResetToken   string     `orm:"size(1000);null"`
	ConfirmResetExpire  time.Time  `orm:"null"`
	Picture             string     `orm:"size(1000);null"`
	Provider            string     `orm:"size(50);null"` // google , fb
	ProviderID          string     `orm:"size(1000);null"`
	ProviderAccessToken string     `orm:"size(1000);null"`
	Permission          string     `orm:"size(50);default(user)"`      // user, admin ...
	Status              string     `orm:"size(50);default(normal)"`    // normal, ban, close ...
	CreateAt            time.Time  `orm:"auto_now_add;type(datetime)"` // first save
	UpdateAt            time.Time  `orm:"auto_now;type(datetime)"`     // eveytime save
}

const pwHashBytes = 64

func generateSalt() (salt string, err error) {
	buf := make([]byte, pwHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", buf), nil
}

func generatePassHash(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, pwHashBytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h), nil
}

// AddUser ...
func AddUser(u User) (int64, error) {
	// make Id
	u.Id = time.Now().UnixNano()

	// make hashed password
	salt, err := generateSalt()
	if err != nil {
		return 0, err
	}
	hash, err := generatePassHash(u.Password, salt)
	if err != nil {
		return 0, err
	}

	// set password & salt
	u.Password = hash
	u.Salt = salt

	// get gravatar
	u.Picture = gravatar.Avatar(u.Email, 80)

	// make email confirm token
	u2, err := uuid.NewV4()
	if err != nil {
		return 0, err
	}
	u.ConfirmResetToken = u2.String()

	// set email confirm expire time +1 hour
	//addTime := time.Now().Add(1 * time.Hour)
	u.ConfirmResetExpire = time.Now().Add(1 * time.Hour)

	// save to db
	o := orm.NewOrm()
	_, err = o.Insert(&u)
	if err != nil {
		return 0, err
	}

	// send confirm mail async
	go libs.MakeMail(u.Email, "confirm", u.ConfirmResetToken)

	return u.Id, nil
}

// FindByDisplayname ...
func FindByDisplayname(displayname string) (User, error) {
	var user User
	o := orm.NewOrm()
	err := o.Raw("SELECT Id, Displayname FROM \"user\" WHERE Displayname = ?", displayname).QueryRow(&user)

	return user, err
}

// FindByEmail ...
func FindByEmail(email string) (User, error) {
	var user User
	o := orm.NewOrm()
	err := o.Raw("SELECT Id, Email FROM \"user\" WHERE Email = ?", email).QueryRow(&user)

	return user, err
}

// CheckEmailConfirmToken ...
func CheckEmailConfirmToken(token string) (User, *libs.ControllerError) {
	var user User
	o := orm.NewOrm()

	// already confirmed
	err := o.Raw("select Id, Displayname, Confirmed from \"user\" where Confirm_Reset_Token =? and confirmed = true", token).QueryRow(&user)
	if err != nil {
		// already confirmed or wrong token
		beego.Error("error CheckEmailConfirmToken(already confirm or wrong token): ", token, " , ", err)
		return user, libs.ErrAlreadyConfirmedOrWrongToken
	}

	//  expired token
	err = o.Raw("select Id, Displayname, Confirmed from \"user\" where Confirm_Reset_Token =? and Confirm_Reset_Expire <= ?", token, time.Now()).QueryRow(&user)
	if err != nil {
		beego.Error("error CheckEmailConfirmToken(expired token): ", token, " , ", err)
		return user, libs.ErrExpiredToken
	}

	return user, nil

}

// Confirm Email ...
//func ConfirmEmail(token string) (User, error) {
func ConfirmEmail(user User) (User, error) {
	o := orm.NewOrm()

	user.Confirmed = true
	user.ConfirmResetToken = ""
	user.ConfirmResetExpire = time.Time{}

	num, err := o.Update(&user)
	if err == nil {
		fmt.Println(num)
	}
	/*
		var user User
		o := orm.NewOrm()
		err := o.Raw("select Id, Displayname, Confirmed from \"user\" where Confirm_Reset_Token =? and Confirm_Reset_Expire >= ?", token, time.Now()).QueryRow(&user)
		// if err == nil, exists
		if err == nil {
			// update Confirm=true, ConfirmResetToken=nil, ConfirmResetExpire=nil
			user.Confirmed = true
			user.ConfirmResetToken = ""
			user.ConfirmResetExpire = time.Time{}
			if num, err := o.Update(&user); err == nil {
				fmt.Println(num)
			}
		}
	*/

	return user, err

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
