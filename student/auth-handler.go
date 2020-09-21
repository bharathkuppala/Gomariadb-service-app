package student

import (
	"net/http"
	"time"

	"github.com/lithammer/shortuuid"
	kafkaservice "github.com/mariaDB/module/kafka-service"
	utility "github.com/mariaDB/module/utilities"
	"golang.org/x/crypto/bcrypt"
)

// Student ...
type Student struct {
	CreatedAt  time.Time  `gorm:"not null"`
	DeletedAt  *time.Time `gorm:"index"`
	ID         string     `gorm:"primary_key"`
	UpdatedAt  time.Time  `gorm:"not null"`
	FirstName  string
	LastName   string
	Password   string
	Email      string
	BloodGroup string
	UserName   string
}

// Register ...
func Register(w http.ResponseWriter, r *http.Request) {
	var student Student
	requestData, err := utility.ToMap(w, r)
	if err != nil {
		utility.SendResponse(w, false, "something wrong with request data")
		return
	}

	r.Body.Close()

	var studentx Student
	// Returns first record which matches my query
	utility.GormDatabase.Where(&Student{
		Email: requestData["email"].(string),
	}).First(&studentx)

	if studentx.Email == requestData["email"].(string) {
		utility.SendResponse(w, false, "this email is already registered with us.")
		return
	}

	student.FirstName = requestData["firstName"].(string) //type assertion
	student.LastName = requestData["lastName"].(string)
	student.Email = requestData["email"].(string)
	student.BloodGroup = requestData["bloodGroup"].(string)
	student.Password = requestData["password"].(string)
	student.UserName = requestData["userName"].(string)

	// Create hash of the password to store it in db
	b, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		utility.SendResponse(w, false, err.Error())
		return
	}
	student.Password = string(b)
	student.ID = shortuuid.New()
	// Insert a values inside table students
	utility.GormDatabase.Model(&Student{}).Create(&student)

	kafkaservice.Publish(student.UserName+"student successfully registered"+time.Now().Format("2006-01-02 15:04:05"), utility.Producer)

	utility.SendResponse(w, true, "registration successful")

}

// Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	var student Student

	requestData, err := utility.ToMap(w, r)
	if err != nil {
		utility.SendResponse(w, false, "something wrong with request data")
		return
	}

	r.Body.Close()

	if requestData["email"].(string) == "" && requestData["password"].(string) == "" {
		utility.SendResponse(w, false, "userName and password cannot be empty")
		return
	}

	student.Email = requestData["email"].(string)
	student.Password = requestData["password"].(string)

	var studentx Student

	// Returns first record which matches my query
	utility.GormDatabase.Where(&Student{
		Email: requestData["email"].(string),
	}).First(&studentx)

	if studentx.ID == "" {
		utility.SendResponse(w, false, "This userName is not registered with us.")
		return
	}

	// Compare password with exisiting user password with password from request data
	if err := bcrypt.CompareHashAndPassword([]byte(studentx.Password), []byte(student.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			utility.Logger.Println(err.Error())
		}

		utility.SendResponse(w, false, "password mismatch occur")
		return
	}

	kafkaservice.Publish(studentx.UserName+" successfully logged in at: "+time.Now().Format("2006-01-02 15:04:05"), utility.Producer)

	utility.SendResponse(w, true, "login in successful")
}
