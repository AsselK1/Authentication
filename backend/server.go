package main

import(
	// "github.com/dgrijalva/jwt-go"
	"time"
	"github.com/gorilla/mux"
	// "github.com/gorilla/sessions"
	"log"
	"net/http"
	"github.com/satori/go.uuid"
	"fmt"
	"encoding/json"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	//"io/ioutil"	
)

type DecodedUser struct{
	Username string `json:"username"`
	Password string `json:"password"`
	Firstname string  `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`

}
type User struct{
	gorm.Model
	Username string `json:"username" gorm:unique"`
	Password []byte `json:"password"` 
	Firstname string  `json:"firstname"` 
	Lastname string `json:"lastname"`
	Email string `json:"email"`
}
type Submission struct{
	gorm.Model
	ProblemID string
	UserID string
	Author string
	Language string
	Code string
	TimeUsage int
	MemoryUsage int
}
type Session struct{
	SessionID string
	UserID uint
	//User User `gorm:"foreignKey:UserID"`
	Activates time.Time
	Expires time.Time
}

var Db *gorm.DB
var err error
var Gl int
func main(){
	// Loading environment variables
	Gl = 4
	//dialect := os.Getenv("DIALECT") 
	host := os.Getenv("HOST") 
	dbPort := os.Getenv("DBPORT") 
	user := os.Getenv("USER") 
	dbName := os.Getenv("NAME") 
	password := os.Getenv("PASSWORD")
	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)


	//Opening connection to database
	Db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
;
	//Close connection to the database when the main function finishes
	postgresDb, err := Db.DB()
	if err != nil{
		log.Fatal(err)
	}
	defer postgresDb.Close()


	// Make database migrations to the database if they have not already been created
	Db.AutoMigrate(&User{})
	//Db.AutoMigrate(&DecodedUser{})
	Db.AutoMigrate(&Submission{})
	Db.AutoMigrate(&Session{})

	//create users to check whether database is working
	asselchenok := &User{Username : "aaa", Email : "asd@gmail.com"}
	Db.Find(asselchenok)
	// newUser := User{
	// 	Username : "aaa",
	// 	Password : []byte{},
	// 	Firstname : "aaa", 
	// 	Lastname : "aaa",
	// 	Email : "aaa",
	// }
	// Db.Create(&newUser)
	// fmt.Println("%v+", asselchenok)
	handler := mux.NewRouter()
	handler.HandleFunc("/problems/{problemsID}", jsonn).Methods(http.MethodGet)
	handler.HandleFunc("/problems/{problemsID}", sendOptions).Methods(http.MethodOptions)
	handler.HandleFunc("/problems", jsonn).Methods(http.MethodGet)
	handler.HandleFunc("/problems", sendOptions).Methods(http.MethodOptions)
	handler.HandleFunc("/", sendOptions).Methods(http.MethodOptions)
	handler.HandleFunc("/signup", signup).Methods(http.MethodPost)
	handler.HandleFunc("/signup", sendOptions).Methods(http.MethodOptions)
	handler.HandleFunc("/login", login).Methods(http.MethodPost)
	handler.HandleFunc("/login", sendOptions).Methods(http.MethodOptions)
	handler.HandleFunc("/cookie", cookie).Methods(http.MethodPost)
	handler.HandleFunc("/cookie", sendOptions).Methods(http.MethodOptions)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
func sendOptions(w http.ResponseWriter, r *http.Request){
	allowCORS(w, r)
	w.WriteHeader(http.StatusOK)
}
func allowCORS(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
func signup(w http.ResponseWriter, r *http.Request){
	allowCORS(w, r)

	var newUser = createUser(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}
// func viewProblems(w http.ResponseWriter, r *http.Request){
// 	allowCORS(w, r)
// 	sessionCookie, err := r.Cookie("session_token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	sessionToken := sessionCookie.Value
// }
func jsonn(w http.ResponseWriter, r *http.Request){
	allowCORS(w, r)
	var session Session
	cookie, err := r.Cookie("session_token")
	fmt.Println(cookie)
	if err != nil {
		if err == http.ErrNoCookie{
			w.WriteHeader(http.StatusUnauthorized)
			return 
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := cookie.Value
	err = Db.First(&session, "session_id = ?", &sessionToken).Error
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// userID := session.UserID
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{message: get called}"))
	//body, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(body)
	//w.WriteHeader(http.StatusOK)
}
func login(w http.ResponseWriter, r *http.Request){	
	var du DecodedUser
	var user User
	allowCORS(w,r)
	json.NewDecoder(r.Body).Decode(&du)
	Db.First(&user, "username = ?", &du.Username)
	if CheckPasswordHash(du.Password, string(user.Password))==false{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	sessionCookie, err := r.Cookie("session_token")
	if err != nil {
		sessionToken := uuid.NewV4().String()
		newSession := Session{
			SessionID : sessionToken,
			UserID : user.ID,
			Activates : time.Now(),
			Expires : time.Now().Add(12000 * time.Second),
		}		
		sessionCookie = &http.Cookie{
			Name:    "session_token",
			Value:   sessionToken,
			Expires: time.Now().Add(12000 * time.Second),
		}
		Db.Exec("DELETE FROM sessions WHERE user_id = ?", user.ID)
		http.SetCookie(w, sessionCookie)
		Db.Create(&newSession)
	// } else {
	// 	fmt.Println(sessionCookie.Value)
	}
	fmt.Println(sessionCookie.Value)
	// token, err := createToken(user.ID)
	// if err != nil{
	// 	w.WriteHeader(http.StatusUnprocessableEntity)
	// }
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(token)
	//json.NewEncoder(w).Encode(CheckPasswordHash(du.Password, string(user.Password)))
}
// func createToken(userID uint) (string, error){
// 	var err error
// 	atClaims := jwt.MapClaims{}
// 	atClaims["userId"] = userID
// 	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
// 	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
// 	if err != nil {
// 		return "", err
// 	}
// 	return token, nil
// }
func cookie(w http.ResponseWriter, r *http.Request){
	allowCORS(w, r)
	var session Session	
	sessionCookie, err := r.Cookie("session_token")
	sessionToken := sessionCookie.Value
	err = Db.First(&session, "session_id = ?", &sessionToken).Error
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
	}
	w,err:=http.Redirect(w, r, "http://localhost:3000/problems", 301)
}