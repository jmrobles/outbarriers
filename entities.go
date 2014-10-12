package outbarriers

import (
	"log"
	"time"
)

type User struct {
	Id       int64
	Email    string
	Password string // TODO: hashed password
}
type Session struct {
	Id        int64
	UserId    int64
	Token     string `sql:"unique; not null"`
	CreatedAt time.Time
}

func (ctx *Context) AuthUser(email, password string) *User {

	var user User
	ctx.DB.Where("email = ? and password = ?", email, password).First(&user)
	return &user
}

func (ctx *Context) CheckAdmin() {

	// Create admin user if not exists
	var user User
	ctx.DB.Where(User{Email: ADMIN_EMAIL, Password: ADMIN_PASSWORD}).FirstOrInit(&user)
	if ctx.DB.NewRecord(user) {
		log.Printf("Admin user not exists, creating...")
		ctx.DB.Create(&user)
	}
}

func (ctx *Context) LoginUser(user *User) string {

	token := RandomString(64)
	session := &Session{UserId: user.Id, Token: token, CreatedAt: time.Now()}
	ctx.DB.Create(session)
	return token
}

func (ctx *Context) GetSessionByToken(token string) *Session {

	var session Session
	ctx.DB.Where("token = ?", token).First(&session)
	if session.Id == 0 {
		return nil
	}
	return &session
}
