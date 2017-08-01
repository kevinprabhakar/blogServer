package user

import (
	"regexp"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"strings"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"blog_server/src/blogServer/mongo"
	"gopkg.in/mgo.v2"
)

type UserController struct {
	session 	*mgo.Session
}

func NewUserController(session *mgo.Session) (*UserController){
	return &UserController{session}
}

func IsValidEmail(email string) (bool) {
	if (len(email) <= 0) {
		return false
	}
	return regexp.MustCompile(`^([a-zA-Z0-9_\.\-\+])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$`).MatchString(email)
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if (err != nil){
		return "", err
	}

	return string(hash), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (self *UserController) SignUp (params string) (User, error){
	dec := json.NewDecoder(strings.NewReader(params))
	var signUpParams SignUpParams
	decodeErr := dec.Decode(&signUpParams)

	if (decodeErr != nil){
		return User{}, decodeErr
	}

	if (!IsValidEmail(signUpParams.Email)){
		return User{}, errors.New("InvalidEmailAddress")
	}

	if(len(signUpParams.Password) < 6){
		return User{}, errors.New("Password is too short")
	}

	var findUser User

	userCollection := mongo.GetUserCollection(mongo.GetDataBase(self.session))

	userFindQuery := bson.M{"email" : signUpParams.Email}
	findUserErr := userCollection.Find(userFindQuery).One(&findUser)

	if (findUserErr != mgo.ErrNotFound){
		return findUser, errors.New("User already exists with this email address")
	}

	hashedPassword, hashingErr := HashPassword(signUpParams.Password)

	if (hashingErr != nil){
		return User{}, hashingErr
	}

	newUser := User{
		Id	: 	bson.NewObjectId(),
		Email 	: 	signUpParams.Email,
		PHash   : 	hashedPassword,
		Name 	: 	signUpParams.Name,
		PostIDList: make([]bson.ObjectId, 0),
		FriendsList: make([]bson.ObjectId, 0),
	}

	insertErr := userCollection.Insert(newUser)

	if (insertErr != nil){
		return User{}, insertErr
	}

	return newUser, nil
}

func (self *UserController) SignIn (params string) (User, error){
	dec := json.NewDecoder(strings.NewReader(params))
	var signInParams SignInParams
	decodeErr := dec.Decode(&signInParams)

	if (decodeErr != nil){
		return User{}, decodeErr
	}

	userFindQuery := bson.M{
		"email" : signInParams.Email,
	}

	var findUser User

	userCollection := mongo.GetUserCollection(mongo.GetDataBase(self.session))
	userFindErr := userCollection.Find(userFindQuery).One(&findUser)

	if (userFindErr != nil){
		return findUser, errors.New("User Not Found with these credentials")
	}

	hashedPassword := CheckPasswordHash(signInParams.Password, findUser.PHash)

	if (!hashedPassword){
		return User{},  errors.New("Invalid Password for this User ID")
	}

	return findUser, nil
}

func (self *UserController) GetUser(accessToken string) (User, error){
	uid, uidError := VerifyAccessToken(accessToken)

	if (uidError != nil){
		return User{}, uidError
	}

	if(!bson.IsObjectIdHex(uid)){
		return User{}, errors.New("Invalid BSON User Id")
	}

	userCollection := mongo.GetUserCollection(mongo.GetDataBase(self.session))

	var findUser User

	findErr := userCollection.Find(bson.M{"_id":bson.ObjectIdHex(uid)}).One(&findUser)

	if (findErr != nil){
		return User{}, findErr
	}

	return findUser, nil
}