package data

// import (
// 	"context"
// 	"errors"
// 	"strings"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"

// 	"github.com/dgrijalva/jwt-go"
// 	"golang.org/x/crypto/bcrypt"

// 	"github.com/ETjoel/task_managment_api/database"
// 	model "github.com/ETjoel/task_managment_api/models/user_model"
// )

// var JwtSecret = []byte("oda_eichro_is_overrated!")

// func RegisterUser(user model.User) error {

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

// 	if err != nil {
// 		return err
// 	}

// 	newUser := model.User{Email: strings.ToLower(user.Email), Password: string(hashedPassword)}

// 	collection := database.GetAuthCollection(database.UsersCollection)
// 	_, err = collection.InsertOne(context.TODO(), newUser)

// 	if err != nil && mongo.IsDuplicateKeyError(err) {
// 		return errors.New("email already in use")
// 	} else if err != nil {
// 		return errors.New("internal server error")
// 	} else {
// 		return nil
// 	}
// }

// func LoginUser(user model.User) (string, error) {
// 	var exitingUser model.User
// 	collection := database.GetCollection(database.UsersCollection)

// 	user.Email = strings.ToLower(user.Email)

// 	err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&exitingUser)

// 	if err != nil {
// 		return "", errors.New("user not found or invalid credentials")
// 	}

// 	if bcrypt.CompareHashAndPassword([]byte(exitingUser.Password), []byte(user.Password)) != nil {
// 		return "", errors.New("invalid email or password")
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"_id":   exitingUser.ID,
// 		"email": exitingUser.Email,
// 	})

// 	jwtToken, err := token.SignedString(JwtSecret)
// 	if err != nil {
// 		return "", errors.New("internal server error: " + err.Error())
// 	}

// 	return jwtToken, nil
// }
