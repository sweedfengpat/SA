package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/patza/app/controllers"
	"github.com/patza/app/ent"
	"github.com/patza/app/ent/user"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}

type User struct {
	Name     string
	Email    string
	Age      int
	Password string
	Birth    time.Time
	Tel      int
}
type Confirmations struct {
	Confirmation []Confirmation
}
type Confirmation struct {
	AddDate      time.Time
	BookingStart time.Time
	BookingEnd   time.Time
	HoursTime    int
	Owner        int
}

// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewConfirmationController(v1, client)
	controllers.NewBankController(v1, client)
	controllers.NewBillController(v1, client)

	// Set Users Data
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	time01 := "Feb 4, 2014 at 6:05pm (PST)"
	t1, _ := time.Parse(layout, time01)

	users := Users{
		User: []User{
			User{"Chanwit Kaewkasi", "chanwit@gmail.com", 13, "b12345", t1, 07534456756},
			User{"Name Surname", "me@example.com", 16, "c123456", t1, 0234567543},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetName(u.Name).
			SetEmail(u.Email).
			SetAge(u.Age).
			SetPassword(u.Password).
			SetBirth(u.Birth).
			SetTel(u.Tel).
			Save(context.Background())
		if err != nil {
			return
		}
	}

	// Set Banks Data
	bank := []string{"kbank", "scb"}
	for _, b := range bank {
		client.Bank.
			Create().
			SetBank(b).
			Save(context.Background())
	}

	// Set Confirmations Data
	time1 := "Apr 4, 2018 at 8:05pm (PST)"
	time2 := "Apr 5, 2018 at 6:00pm (PST)"
	time3 := "Apr 5, 2018 at 7:00pm (PST)"
	time4 := "Apr 5, 2018 at 7:00pm (PST)"
	time5 := "Apr 5, 2018 at 8:00pm (PST)"
	t01, _ := time.Parse(layout, time1)
	t02, _ := time.Parse(layout, time2)
	t03, _ := time.Parse(layout, time3)
	t04, _ := time.Parse(layout, time4)
	t05, _ := time.Parse(layout, time5)
	h1 := int(t03.Hour() - t02.Hour())
	h2 := int(t05.Hour() - t04.Hour())
	confirmations := Confirmations{
		Confirmation: []Confirmation{
			Confirmation{t01, t02, t03, h1, 1},
			Confirmation{t01, t04, t05, h2, 1},
		},
	}

	for _, con := range confirmations.Confirmation {
		u, err := client.User.
			Query().
			Where(user.IDEQ(int(con.Owner))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Confirmation.
			Create().
			SetAdddate(con.AddDate).
			SetBookingstart(con.BookingStart).
			SetBookingend(con.BookingEnd).
			SetHourstime(con.HoursTime).
			SetOwner(u).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
