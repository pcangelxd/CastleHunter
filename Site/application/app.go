package application

import (
	"../config"
	"./storage"
	"fmt"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

type Application struct {
	Config				*config.Config
	Storage				storage.Storage
	HashPassRecovery	map[string]string
}

func NewApplication(configPath string) Application {
	cfg := config.NewConfig(configPath)
	store := storage.NewStorageSQLite()

	return Application{
		Config:  cfg,
		Storage: store,
		HashPassRecovery: map[string]string{},
	}
}

func (app *Application) Run() {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	path, _ := os.Getwd()
	router.LoadHTMLGlob(strings.Replace(path, "\\", "/", -1) + "/application/templates/*")
	router.Use(static.Serve("/static/", static.LocalFile(strings.Replace(path, "\\", "/", -1) + "/application/static/", true)))


	router.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		isLogin := session.Get("isLogin")
		ru := session.Get("ru")
		en := session.Get("en")
		if ru != true && en != true {
			ru = true
		}
		hashPassRecovery := session.Get("hashPassRecovery")

		if hashPassRecovery != nil {
			userHashPassRecovery := session.Get("userHashPassRecovery")
			session.Delete("userHashPassRecovery")
			session.Delete("hashPassRecovery")
			delete(app.HashPassRecovery, fmt.Sprintf("%v", hashPassRecovery))
			session.Save()

			c.HTML(
				http.StatusOK,
				"main.html",
				gin.H{
					"email": userHashPassRecovery,
					"pass_recovery": true,
					"ru": ru,
					"en": en,
					"cost7Day": app.Config.Cost7Day,
					"cost14Day": app.Config.Cost14Day,
					"cost30Day": app.Config.Cost30Day,
					"bankDetails": app.Config.BankDetails,
				},
			)
		} else {
			if isLogin == true {
				email := session.Get("email")
				errorRu := session.Get("errorRu")
				ErrorEn := session.Get("ErrorEn")
				notificationRu := session.Get("notificationRu")
				notificationEn := session.Get("notificationEn")
				requestInfo := session.Get("requestInfo")
				requestCost := session.Get("requestCost")

				session.Delete("errorRu")
				session.Delete("ErrorEn")
				session.Delete("notificationRu")
				session.Delete("notificationEn")
				session.Save()

				u, ok := app.Storage.GetUser(fmt.Sprintf("%v", email))
				admin := u.Role == 1

				if ok {
					accessFind := false
					if u.DurationOfThePrivilege != "00.00.0000" {
						date, _ := time.Parse("2006/1/2", u.DurationOfThePrivilege)
						accessFind = date.Before(time.Now())
					}

					if requestInfo != nil {
						if requestInfo.(int64) > 0 {
							session.Delete("requestInfo")
							session.Delete("requestCost")
							session.Save()
							req, _ := app.GetRequest(requestInfo.(int64))
							c.HTML(
								http.StatusOK,
								"main(logged_in).html",
								gin.H{
									"user": u,
									"errorRuState": errorRu != nil,
									"errorRu": fmt.Sprintf("%v", errorRu),
									"ErrorEnState": ErrorEn != nil,
									"ErrorEn": fmt.Sprintf("%v", ErrorEn),
									"notificationRuState": notificationRu != nil,
									"notificationRu": fmt.Sprintf("%v", notificationRu),
									"notificationEnState": notificationEn != nil,
									"notificationEn": fmt.Sprintf("%v", notificationEn),
									"requestInfo": req,
									"requestCost": requestCost,
									"requestInfoState": true,
									"accessFind": accessFind,
									"ru": ru,
									"en": en,
									"admin": admin,
									"cost7Day": app.Config.Cost7Day,
									"cost14Day": app.Config.Cost14Day,
									"cost30Day": app.Config.Cost30Day,
									"bankDetails": app.Config.BankDetails,
								},
							)
						}
					} else {
						c.HTML(
							http.StatusOK,
							"main(logged_in).html",
							gin.H{
								"user": u,
								"errorRuState": errorRu != nil,
								"errorRu": fmt.Sprintf("%v", errorRu),
								"ErrorEnState": ErrorEn != nil,
								"ErrorEn": fmt.Sprintf("%v", ErrorEn),
								"notificationRuState": notificationRu != nil,
								"notificationRu": fmt.Sprintf("%v", notificationRu),
								"notificationEnState": notificationEn != nil,
								"notificationEn": fmt.Sprintf("%v", notificationEn),
								"accessFind": accessFind,
								"ru": ru,
								"en": en,
								"admin": admin,
								"cost7Day": app.Config.Cost7Day,
								"cost14Day": app.Config.Cost14Day,
								"cost30Day": app.Config.Cost30Day,
								"bankDetails": app.Config.BankDetails,
							},
						)
					}
				} else {
					session.Set("errorRu", "Пользователь не найден.")
					session.Set("ErrorEn", "(EN) Пользователь не найден.")
					session.Save()
				}
			} else {
				errorRu := session.Get("errorRu")
				errorEn := session.Get("errorEn")
				notificationRu := session.Get("notificationRu")
				notificationEn := session.Get("notificationEn")
				session.Delete("errorRu")
				session.Delete("errorEn")
				session.Delete("notificationRu")
				session.Delete("notificationEn")
				session.Save()

				c.HTML(
					http.StatusOK,
					"main.html",
					gin.H{
						"errorRuState": errorRu != nil,
						"errorRu": fmt.Sprintf("%v", errorRu),
						"errorEnState": errorEn != nil,
						"errorEn": fmt.Sprintf("%v", errorEn),
						"notificationRuState": notificationRu != nil,
						"notificationRu": fmt.Sprintf("%v", notificationRu),
						"notificationEnState": notificationEn != nil,
						"notificationEn": fmt.Sprintf("%v", notificationEn),
						"ru": ru,
						"en": en,
						"cost7Day": app.Config.Cost7Day,
						"cost14Day": app.Config.Cost14Day,
						"cost30Day": app.Config.Cost30Day,
						"bankDetails": app.Config.BankDetails,
					},
				)
			}
		}
	})

	router.GET("/history", func(c *gin.Context) {
		session := sessions.Default(c)
		isLogin := session.Get("isLogin")
		errorRu := session.Get("errorRu")
		errorEn := session.Get("errorEn")
		ru := session.Get("ru")
		en := session.Get("en")
		if ru != true && en != true {
			ru = true
		}
		notificationRu := session.Get("notificationRu")
		notificationEn := session.Get("notificationEn")
		session.Delete("errorRu")
		session.Delete("errorEn")
		session.Delete("notificationRu")
		session.Delete("notificationEn")
		session.Save()

		if isLogin == true {
			email := session.Get("email")
			requestInfo := session.Get("requestInfo")
			requestCost := session.Get("requestCost")
			u, ok := app.Storage.GetUser(fmt.Sprintf("%v", email))
			admin := u.Role == 1

			requests := app.GetRequests(email.(string))

			if ok {
				accessFind := false
				if u.DurationOfThePrivilege != "00.00.0000" {
					date, _ := time.Parse("2006/1/2", u.DurationOfThePrivilege)
					accessFind = date.Before(time.Now())
				}

				if requestInfo != nil {
					if requestInfo.(int64) > 0 {
						session.Delete("requestInfo")
						session.Delete("requestCost")
						session.Save()
						req, _ := app.GetRequest(requestInfo.(int64))
						c.HTML(
							http.StatusOK,
							"history(logged_in).html",
							gin.H{
								"user":              u,
								"errorRuState":        errorRu != nil,
								"errorRu":             fmt.Sprintf("%v", errorRu),
								"errorEnState":        errorEn != nil,
								"errorEn":             fmt.Sprintf("%v", errorEn),
								"notificationRuState": notificationRu != nil,
								"notificationRu":      fmt.Sprintf("%v", notificationRu),
								"notificationEnState": notificationEn != nil,
								"notificationEn":      fmt.Sprintf("%v", notificationEn),
								"requestInfo":       req,
								"requestCost": requestCost,
								"requestInfoState":  true,
								"requests": 		 requests,
								"ru": ru,
								"en": en,
								"cost7Day": app.Config.Cost7Day,
								"cost14Day": app.Config.Cost14Day,
								"cost30Day": app.Config.Cost30Day,
								"bankDetails": app.Config.BankDetails,
								"accessFind": accessFind,
								"admin": admin,
							},
						)
					}
				} else {
					c.HTML(
						http.StatusOK,
						"history(logged_in).html",
						gin.H{
							"user": u,
							"errorRuState": errorRu != nil,
							"errorRu": fmt.Sprintf("%v", errorRu),
							"errorEnState": errorEn != nil,
							"errorEn": fmt.Sprintf("%v", errorEn),
							"notificationRuState": notificationRu != nil,
							"notificationRu": fmt.Sprintf("%v", notificationRu),
							"notificationEnState": notificationEn != nil,
							"notificationEn": fmt.Sprintf("%v", notificationEn),
							"requests": requests,
							"ru": ru,
							"en": en,
							"accessFind": accessFind,
							"admin": admin,
							"cost7Day": app.Config.Cost7Day,
							"cost14Day": app.Config.Cost14Day,
							"cost30Day": app.Config.Cost30Day,
							"bankDetails": app.Config.BankDetails,
						},
					)
				}
			} else {
				session.Set("errorRu", "Ошибка загрузки истории")
				session.Save()
			}
		} else {
			c.HTML(
				http.StatusOK,
				"history(logged_in).html",
				gin.H{
					"errorRuState": errorRu != nil,
					"errorRu": fmt.Sprintf("%v", errorRu),
					"errorEnState": errorEn != nil,
					"errorEn": fmt.Sprintf("%v", errorEn),
					"notificationRuState": notificationRu != nil,
					"notificationRu": fmt.Sprintf("%v", notificationRu),
					"notificationEnState": notificationEn != nil,
					"notificationEn": fmt.Sprintf("%v", notificationEn),
					"ru": ru,
					"en": en,
					"cost7Day": app.Config.Cost7Day,
					"cost14Day": app.Config.Cost14Day,
					"cost30Day": app.Config.Cost30Day,
					"bankDetails": app.Config.BankDetails,
				},
			)
		}
	})

	router.GET("/user/pass_recovery/:user", func(c *gin.Context) {
		session := sessions.Default(c)

		user := c.Param("user")
		session.Set("hashPassRecovery", user)
		session.Set("userHashPassRecovery", app.HashPassRecovery[user])
		session.Save()


		c.Redirect(http.StatusMovedPermanently, "/")
	})

	router.POST("/user/registration/", func(c *gin.Context) {
		session := sessions.Default(c)

		email := c.PostForm("email")
		password := c.PostForm("password")
		confirmPassword := c.PostForm("confirm_password")

		ok := app.Registration(email, password, confirmPassword)
		if ok {
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			session.Set("errorRu", "Ошибка авторизации.")
			session.Set("errorEn", "(EN) Ошибка авторизации.")
			session.Save()
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	})

	router.POST("/user/authorisation/", func(c *gin.Context) {
		session := sessions.Default(c)

		email := c.PostForm("email")
		password := c.PostForm("password")

		ok := app.Authorisation(email, password)
		if ok {
			session.Set("isLogin", true)
			session.Set("email", email)
			session.Save()

			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			session.Set("errorRu", "Ошибка авторизации.")
			session.Set("errorEn", "(EN) Ошибка авторизации.")
			session.Save()
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	})

	router.POST("/user/pass_recovery1/", func(c *gin.Context) {
		session := sessions.Default(c)
		email := c.PostForm("email")

		ok := app.ResetPassword1(email)

		if ok {
			session.Set("notificationRu", "Вам на почту было отправлено письмо с дальнейшими инструкциями.")
			session.Set("notificationEn", "(EN) Вам на почту было отправлено письмо с дальнейшими инструкциями.")
			session.Save()
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			session.Set("errorRu", "Ошибка восстановления учётной записи.")
			session.Set("errorEn", "(EN) Ошибка восстановления учётной записи.")
			session.Save()
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	})

	router.POST("/user/pass_recovery2/", func(c *gin.Context) {
		session := sessions.Default(c)
		email := c.PostForm("email")
		password := c.PostForm("password")
		confirmPassword := c.PostForm("confirm_password")

		ok := app.ResetPassword2(email, password, confirmPassword)

		if ok {
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			session.Set("errorRu", "Ошибка восстановления учётной записи.")
			session.Set("errorEn", "(EN) Ошибка восстановления учётной записи.")
			session.Save()
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	})

	router.POST("/user/quit/", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()

		c.Redirect(http.StatusMovedPermanently, "/")
	})

	router.POST("/request/add/", func(c *gin.Context) {
		session := sessions.Default(c)

		email := session.Get("email")
		typeReq := c.PostForm("type")

		ok, id, cost := app.InsertRequest(fmt.Sprintf("%v", email), typeReq)
		if ok {
			session.Set("requestInfo", id)
			session.Set("requestCost", cost)
			session.Save()
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			session.Set("errorRu", "Ошибка добавления заявки.")
			session.Set("errorEn", "(EN) Ошибка добавления заявки.")
			session.Save()
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	})

	router.POST("/castle/find/", func(c *gin.Context) {
		session := sessions.Default(c)
		email := session.Get("email")
		kor := c.PostForm("kor")
		nick := c.PostForm("nick")

		kor_, _ := strconv.Atoi(kor)
		ok := app.InsertFound(fmt.Sprintf("%v", email), nick, uint32(kor_))

		if ok {
			session.Set("notificationRu", "Как только до вас дойдет очередь и цель будет найдена вам будет выслано письмо на почту с информацией о местоположении цели. \n\nПриносим свои извинения за ожидание.")
			session.Set("notificationEn", "(EN) Как только до вас дойдет очередь и цель будет найдена вам будет выслано письмо на почту с информацией о местоположении цели. \n\nПриносим свои извинения за ожидание.")
		} else {
			session.Set("errorRu", "Ошибка поиска замка.")
			session.Set("errorEn", "(EN) Ошибка поиска замка.")
		}
		session.Save()
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	router.GET("/castle/found/", func(c *gin.Context) {
		fnd, status := app.GetFound()

		c.JSON(200, gin.H{
			"status": status,
			"email": fnd.Email,
			"kor": fnd.Kor,
			"purpose": fnd.Name,
		})
	})

	router.POST("/req/accept/", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))

		app.AcceptRequest(int64(id))
		c.Redirect(http.StatusMovedPermanently, "/requests")
	})

	router.POST("/req/cancel/", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))

		app.CancelRequest(int64(id))
		c.Redirect(http.StatusMovedPermanently, "/requests")
	})

	router.POST("/user/language/en", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("en", true)
		session.Set("ru", false)
		session.Save()

		c.Redirect(http.StatusMovedPermanently, "/")
	})

	router.POST("/user/language/ru", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("ru", true)
		session.Set("en", false)
		session.Save()

		c.Redirect(http.StatusMovedPermanently, "/")
	})

	router.GET("/requests", func(c *gin.Context) {
		session := sessions.Default(c)
		isLogin := session.Get("isLogin")
		errorRu := session.Get("errorRu")
		errorEn := session.Get("errorEn")
		ru := session.Get("ru")
		en := session.Get("en")
		if ru != true && en != true {
			ru = true
		}
		notificationRu := session.Get("notificationRu")
		notificationEn := session.Get("notificationEn")
		session.Delete("errorRu")
		session.Delete("errorEn")
		session.Delete("notificationRu")
		session.Delete("notificationEn")
		session.Save()

		if isLogin == true {
			email := session.Get("email")
			requestInfo := session.Get("requestInfo")
			requestCost := session.Get("requestCost")
			u, ok := app.Storage.GetUser(fmt.Sprintf("%v", email))
			admin := u.Role == 1

			requests := app.GetAllRequests()

			if ok {
				accessFind := false
				if u.DurationOfThePrivilege != "00.00.0000" {
					date, _ := time.Parse("2006/1/2", u.DurationOfThePrivilege)
					accessFind = date.Before(time.Now())
				}

				if requestInfo != nil {
					if requestInfo.(int64) > 0 {
						session.Delete("requestInfo")
						session.Delete("requestCost")
						session.Save()
						req, _ := app.GetRequest(requestInfo.(int64))
						c.HTML(
							http.StatusOK,
							"history(admin).html",
							gin.H{
								"user":              u,
								"errorRuState":        errorRu != nil,
								"errorRu":             fmt.Sprintf("%v", errorRu),
								"errorEnState":        errorEn != nil,
								"errorEn":             fmt.Sprintf("%v", errorEn),
								"notificationRuState": notificationRu != nil,
								"notificationRu":      fmt.Sprintf("%v", notificationRu),
								"notificationEnState": notificationEn != nil,
								"notificationEn":      fmt.Sprintf("%v", notificationEn),
								"requestInfo":       req,
								"requestCost": requestCost,
								"requestInfoState":  true,
								"requests": 		 requests,
								"ru": ru,
								"en": en,
								"accessFind": accessFind,
								"admin": admin,
								"cost7Day": app.Config.Cost7Day,
								"cost14Day": app.Config.Cost14Day,
								"cost30Day": app.Config.Cost30Day,
								"bankDetails": app.Config.BankDetails,
							},
						)
					}
				} else {
					c.HTML(
						http.StatusOK,
						"history(admin).html",
						gin.H{
							"user": u,
							"errorRuState": errorRu != nil,
							"errorRu": fmt.Sprintf("%v", errorRu),
							"errorEnState": errorEn != nil,
							"errorEn": fmt.Sprintf("%v", errorEn),
							"notificationRuState": notificationRu != nil,
							"notificationRu": fmt.Sprintf("%v", notificationRu),
							"notificationEnState": notificationEn != nil,
							"notificationEn": fmt.Sprintf("%v", notificationEn),
							"requests": requests,
							"ru": ru,
							"en": en,
							"accessFind": accessFind,
							"admin": admin,
							"cost7Day": app.Config.Cost7Day,
							"cost14Day": app.Config.Cost14Day,
							"cost30Day": app.Config.Cost30Day,
							"bankDetails": app.Config.BankDetails,
						},
					)
				}
			} else {
				session.Set("errorRu", "Ошибка загрузки истории")
				session.Set("errorEn", "(EN) Ошибка загрузки истории")
				session.Save()
			}
		} else {
			c.HTML(
				http.StatusOK,
				"history(admin).html",
				gin.H{
					"errorRuState": errorRu != nil,
					"errorRu": fmt.Sprintf("%v", errorRu),
					"errorEnState": errorEn != nil,
					"errorEn": fmt.Sprintf("%v", errorEn),
					"notificationRuState": notificationRu != nil,
					"notificationRu": fmt.Sprintf("%v", notificationRu),
					"notificationEnState": notificationEn != nil,
					"notificationEn": fmt.Sprintf("%v", notificationEn),
					"ru": ru,
					"en": en,
					"cost7Day": app.Config.Cost7Day,
					"cost14Day": app.Config.Cost14Day,
					"cost30Day": app.Config.Cost30Day,
					"bankDetails": app.Config.BankDetails,
				},
			)
		}
	})

	router.Run(app.Config.Address)
}

func (app *Application) Authorisation(email, password string) bool {
	if valid.IsEmail(email) && len(password) >= 8 {
		return app.Storage.AuthorUser(email, password)
 	} else {
 		return false
	}
}

func (app *Application) Registration(email, password, confirmPassword string) bool {
	if password == confirmPassword && valid.IsEmail(email) && len(password) >= 8 {
		date := time.Now()
		u := storage.User{
			Email: email,
			Password: password,
			DateOfRegistory: date.Format("2006/1/2"),
			DurationOfThePrivilege: "0000/0/0",
			Role: 0,
		}

		return app.Storage.InsertUser(u)
	} else {
		return false
	}
}

func (app *Application) ResetPassword1(email string) bool {
	rndString := RandStringRunes(25)
	app.HashPassRecovery[rndString] = email

	message := "Subject: Сброс пароля \r\n\r\n Для восстановления доступа к аккаунта пройдите по ссылке: "+ app.Config.Address +"/user/pass_recovery/" + rndString
	from := app.Config.Email
	to := email
	host := app.Config.Host
	auth := smtp.PlainAuth("", from, app.Config.Password, host)

	smtp.SendMail(host + ":" + app.Config.Port, auth, from, []string{to}, []byte(message))

	return true
}

func (app *Application) ResetPassword2(email, password, confirmPassword string) bool {
	if password == confirmPassword {
		return app.Storage.UpdateUser(email, password)
	} else {
		return false
	}
}


func (app *Application) InsertRequest(email, typeReq string) (bool, int64, string) {
	date := time.Now()
	valueDay := 0
	cost := "0"

	if typeReq[0:2] == "7 " {
		valueDay = 7
		cost = app.Config.Cost7Day
	} else if typeReq[0:2] == "14" {
		valueDay = 14
		cost = app.Config.Cost14Day
	}else if typeReq[0:2] == "30" {
		valueDay = 30
		cost = app.Config.Cost30Day
	}

	req := storage.Request{
		Email: email,
		DateOfRegistory: date.Format("2006/1/2"),
		ValueDay: uint8(valueDay),
	}
	res1, res2 := app.Storage.InsertRequest(req)

	return res1, res2, cost
}

func (app *Application) GetRequest(id int64) (storage.Request, bool) {
	return app.Storage.GetRequest(id)
}

func (app *Application) GetRequests(email string) []storage.Request {
	return app.Storage.GetRequests(email)
}

func (app *Application) AcceptRequest(id int64) bool {
	return app.Storage.AcceptRequest(id)
}

func (app *Application) CancelRequest(id int64) bool {
	return app.Storage.CancelRequest(id)
}

func (app *Application) GetAllRequests() []storage.Request {
	return app.Storage.GetAllRequests()
}

func (app *Application) InsertFound(email, name string, kor uint32) bool {
	status := "В очереди"

	fnd := storage.Found{
		Email: email,
		Kor: kor,
		Name: name,
		Status: status,
	}

	return app.Storage.InsertFound(fnd)
}

func (app *Application) GetFound() (storage.Found, bool) {
	return app.Storage.GetFound()
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}