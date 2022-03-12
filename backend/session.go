package main
import(
	//"log"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	//"io/ioutil"
	//"gorm.io/gorm"
	"encoding/json"
	//"time"
	//"fmt"	
)
func createUser(w http.ResponseWriter, r *http.Request) User {
	var du DecodedUser
	err := json.NewDecoder(r.Body).Decode(&du)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return User{}
    }
	np, err := bcrypt.GenerateFromPassword([]byte(du.Password), bcrypt.MinCost)
	if err != nil{
        http.Error(w, err.Error(), http.StatusBadRequest)
        return User{}
	}
	newUser := User{
		Username : du.Username,
		Password : np,
		Firstname : du.Firstname, 
		Lastname : du.Lastname,
		Email : du.Email,
	}
	Db.Create(&newUser)
	return newUser
	/*body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        panic(err)
    }
    log.Println(string(body))
	var jsonUserInfo decodeduser
	err = json.Unmarshal(body, &jsonUserInfo)
	if err != nil{
		panic(err)
	}
	np, err := bcrypt.GenerateFromPassword([]byte(jsonUserInfo.Password), bcrypt.MinCost)
	if err != nil{
		log.Fatal(err)
	}
	newUser := user{jsonUserInfo.Username, np, jsonUserInfo.Firstname, jsonUserInfo.Lastname, jsonUserInfo.Email}
	dbUsers[newUser.Username] = newUser
	log.Println(newUser.Password)
	log.Println(newUser.Firstname)
	*/
	
}
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
// func checkUser(w http.ResponseWriter, r *http.Request) User {
// 	var du DecodedUser
// 	err := json.NewDecoder(r.Body).Decode(&du)
// 	if err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return User{}
//     }
// 	np, err := bcrypt.GenerateFromPassword([]byte(du.Password), bcrypt.MinCost)
// 	if err != nil{
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return User{}
// 	}
// 	newUser := User{
// 		Username : du.Username,
// 		Password : np,
// 	}
// 	return newUser
// }
/*  

	err = r.ParseForm()
	if err != nil{
		fmt.Println("ERROR")
	}username := r.FormValue("username")
	password := r.FormValue("password")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	fmt.Println(username + "   "+ firstname + "   " + lastname+ "   "+ "    " + email+" "  + password)
	np, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil{
		log.Fatal(err)
	}
	newUser := user{username, np, firstname, lastname, email}
	dbUsers[newUser.Username] = newUser
	return newUser
	*/