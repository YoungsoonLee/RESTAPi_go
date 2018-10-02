package models

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"strconv"
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
	Displayname         string     `orm:"size(30);unique"`  // 4 ~ 16 letters for local,
	Email               string     `orm:"size(100);unique"` // max 100 letters
	Password            string     `orm:"null"`             // if account is provider, this column is null
	Salt                string     `orm:"null"`
	PasswordResetToken  string     `orm:"size(1000);null"`
	PasswordResetExpire *time.Time `orm:"null"`
	Confirmed           bool       `orm:"default(false)"`
	ConfirmResetToken   string     `orm:"size(1000);null"`
	ConfirmResetExpire  time.Time  `orm:"null"`
	Picture             string     `orm:"size(1000);null"`
	Provider            string     `orm:"size(50);null"` // google , facebook
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

	//fmt.Println(fmt.Sprintf("%x", h), password, salt)

	return fmt.Sprintf("%x", h), nil
}

// CheckPass compare input password.
func (u *User) CheckPass(pass string) (ok bool, err error) {
	//fmt.Println(pass, u.Salt)
	hash, err := generatePassHash(pass, u.Salt)
	if err != nil {
		return false, err
	}

	//fmt.Println(u.Password, hash, u.Password == hash)

	return u.Password == hash, nil
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

	//TODO: wallet, ?? transaction

	// send confirm mail async
	go libs.MakeMail(u.Email, "confirm", u.ConfirmResetToken)

	return u.Id, nil
}

// AddSocialUser ...
func AddSocialUser(u User) (int64, string, error) {
	// make Id
	u.Id = time.Now().UnixNano()
	u.Displayname = "FB" + strconv.FormatInt(time.Now().UnixNano(), 10)
	u.Confirmed = true

	// save to db
	o := orm.NewOrm()
	_, err := o.Insert(&u)
	if err != nil {
		return 0, "", err
	}

	//TODO: wallet, ?? transaction

	return u.Id, u.Displayname, nil
}

// UpdateSocialInfo ...
func UpdateSocialInfo(u User) (int64, string, error) {

	u.Confirmed = true

	o := orm.NewOrm()
	if _, err := o.Update(&u, "Provider", "ProviderAccessToken", "ProviderID", "Picture", "Confirmed"); err != nil {
		return 0, "", err
	}

	return u.Id, u.Displayname, nil
}

// FindAuthByDisplayname ...
// using for auth
func FindAuthByDisplayname(displayname string) (User, error) {
	var user User
	o := orm.NewOrm()
	err := o.Raw("SELECT Id, Displayname, Password, Salt FROM \"user\" WHERE Displayname = ?", displayname).QueryRow(&user)
	//fmt.Println(user.Salt)
	return user, err
}

// FindByDisplayname ...
// TODO: add balance
func FindByDisplayname(displayname string) (User, error) {
	var user User
	o := orm.NewOrm()
	err := o.Raw("SELECT Id, Displayname , Confirmed, Picture, Provider, Permission, Status, Create_At, Update_At   FROM \"user\" WHERE Displayname = ?", displayname).QueryRow(&user)
	//fmt.Println(user.Salt)
	return user, err
}

// FindByEmail ...
// TODO: add balance
func FindByEmail(email string) (User, error) {
	var user User
	o := orm.NewOrm()
	err := o.Raw("SELECT Id, Displayname, Confirmed, Picture, Provider, Permission, Status, Create_At, Update_At FROM \"user\" WHERE Email = ?", email).QueryRow(&user)

	return user, err
}

// FindByProvider ...
func FindByProvider(provider string, accessToken string, providerID string) bool {
	/*
		var user User
		o := orm.NewOrm()
		err := o.Raw("SELECT Id, Displayname, Email, Confirmed, Picture, Provider, Permission, Status, Create_At, Update_At FROM \"user\" WHERE provider = ? and accessToken= ?", provider, accessToken).QueryRow(&user)
		return user, err
	*/
	o := orm.NewOrm()
	exist := o.QueryTable("user").Filter("Provider", provider).Filter("ProviderAccessToken", accessToken).Filter("ProviderID", providerID).Exist()

	return exist
}

// CheckConfirmEmailToken ...
func CheckConfirmEmailToken(token string) (*User, *libs.ControllerError) {
	var user *User
	o := orm.NewOrm()

	// already confirmed
	err := o.Raw("select Id, Displayname, Confirmed from \"user\" where Confirm_Reset_Token =? and Confirmed = true", token).QueryRow(&user)
	if err == nil {
		// already confirmed or wrong token
		beego.Info("CheckConfirmEmailToken (Already confirmed): ", token, " , ", err)
		return user, libs.ErrAlreadyConfirmed
	}

	// wrong token
	err = o.Raw("select Id, Displayname, Confirmed from \"user\" where Confirm_Reset_Token =? and Confirmed = false", token).QueryRow(&user)
	if err != nil {
		// already confirmed or wrong token
		beego.Error("error CheckConfirmEmailToken(wrong token): ", token, " , ", err)
		return user, libs.ErrWrongToken
	}

	//  expired token
	err = o.Raw("select Id, Displayname, Confirmed from \"user\" where Confirm_Reset_Token =? and Confirm_Reset_Expire <= ?", token, time.Now()).QueryRow(&user)
	if err == nil {
		// expire token
		beego.Error("error CheckConfirmEmailToken(expired token): ", token, " , ", err)
		return user, libs.ErrExpiredToken
	}

	return user, nil
}

// Confirm Email ...
//func ConfirmEmail(token string) (User, error) {
func ConfirmEmail(u User) (User, error) {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE \"user\" SET Confirmed = ?, Confirm_Reset_Expire=?", true, nil).Exec()
	if err != nil {
		return User{}, err
	}

	return u, err
}

func ResendConfirmEmail(u User) (User, error) {
	// make email confirm token
	u2, err := uuid.NewV4()
	if err != nil {
		return User{}, err
	}

	u.ConfirmResetToken = u2.String()
	u.ConfirmResetExpire = time.Now().Add(1 * time.Hour)
	u.Confirmed = false

	o := orm.NewOrm()
	if _, err := o.Update(&u, "Confirmed", "ConfirmResetToken", "ConfirmResetExpire"); err != nil {
		return User{}, err
	}

	// send confirm mail async
	go libs.MakeMail(u.Email, "confirm", u.ConfirmResetToken)

	return u, nil
}

func SendPasswordResetToken(u User) (User, error) {
	// make forgot password token
	u2, err := uuid.NewV4()
	if err != nil {
		return User{}, err
	}

	u.PasswordResetToken = u2.String()
	ct := time.Now().Add(1 * time.Hour)
	u.PasswordResetExpire = &ct

	o := orm.NewOrm()
	if _, err := o.Update(&u, "PasswordResetToken", "PasswordResetExpire"); err != nil {
		return User{}, err
	}

	// send confirm mail async
	go libs.MakeMail(u.Email, "forgotPassword", u.PasswordResetToken)

	return u, nil
}

func CheckResetPasswordToken(resetToken string) (*User, *libs.ControllerError) {
	var user *User

	o := orm.NewOrm()
	// wrong token
	err := o.Raw("select Id, Displayname, Confirmed from \"user\" where Password_Reset_Token =?", resetToken).QueryRow(&user)
	if err != nil {
		// already confirmed or wrong token
		beego.Error("error CheckResetPasswordToken(wrong token): ", resetToken, " , ", err)
		return user, libs.ErrWrongToken
	}

	//  expired token
	err = o.Raw("select Id, Displayname, Confirmed from \"user\" where Password_Reset_Token =? and Password_Reset_Expire <= ?", resetToken, time.Now()).QueryRow(&user)
	if err == nil {
		// expire token
		beego.Error("error CheckResetPasswordToken(expired token): ", resetToken, " , ", err)
		return user, libs.ErrExpiredToken
	}

	return user, nil
}

// ConfirmResetPasswordToken ...
func ConfirmResetPasswordToken(u User) (User, error) {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE \"user\" SET Password_Reset_Token = ?, Password_Reset_Expire=?", nil, nil).Exec()
	if err != nil {
		return User{}, err
	}

	return u, err
}

// ---------------------------------------------------------------------------------------------------------------
// Not use maybe ...
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

func DeleteUser(uid string) {
	delete(UserList, uid)
}
