package main

import (
	"context"
	"e-commerce/cmd/app/container"
	"e-commerce/cmd/app/router"
	_ "e-commerce/cmd/docs"
	"e-commerce/cmd/migrations"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang/glog"
)

func host() {

}
func main() {
	migrations.Migrate()

	container := container.BuildContainer()
	engine := router.InitGinEngine(container)
	server := &http.Server{
		Addr:    ":3000",
		Handler: engine,
	}
	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	glog.Info("Shutdown Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		glog.Fatal("Server Shutdown: ", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		glog.Info("Timeout of 5 seconds.")
	}
	glog.Info("Server exiting")
	fmt.Print(container)
}

// package main

// import (
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/sqlite"
// )

// // Base contains common columns for all tables.
// type Base struct {
// 	ID        string     `json:"uuid" gorm:"unique;not null;index;primary_key"`
// 	CreatedAt time.Time  `json:"created_at"`
// 	UpdatedAt time.Time  `json:"update_at"`
// 	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
// }

// // BeforeCreate will set a UUID rather than numeric ID.
// func (base *Base) BeforeCreate(scope *gorm.Scope) error {
// 	base.ID = uuid.New().String()

// 	return nil
// }

// // User is the model for the user table.
// type User struct {
// 	Base
// 	SomeFlag bool `gorm:"column:some_flag;not null;default:true" json:"some_flag"`
// 	//Profile  Profile `json:"profile"`
// }

// // Profile is the model for the profile table.
// type Profile struct {
// 	Base
// 	Name   string `gorm:"column:name;size:128;not null;" json:"name"`
// 	User   User   `gorm:"association_foreignkey:UserID:"`
// 	UserID uint
// }

// func main() {
// 	db, err := gorm.Open("sqlite3", "test.db")
// 	if err != nil {
// 		panic(err)
// 	}

// 	db.LogMode(true)
// 	db.AutoMigrate(&User{}, &Profile{})

// 	// user := &User{SomeFlag: false}
// 	// if db.Create(&user).Error != nil {
// 	// 	log.Panic("Unable to create user.")
// 	// }

// 	// profile := &Profile{Name: "New User", UserID: user.Base.ID}
// 	// if db.Create(&profile).Error != nil {
// 	// 	log.Panic("Unable to create profile.")
// 	// }

// 	// fetchedUser := &User{}
// 	// if db.Where("id = ?", profile.UserID).Preload("Profile").First(&fetchedUser).RecordNotFound() {
// 	// 	log.Panic("Unable to find created user.")
// 	// }

// 	//fmt.Printf("User: %+v\n", fetchedUser)
// }
