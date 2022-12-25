//This is  “register user” endpoint that can be used to create new users.

package controllers

import (
    "errors"
    "fmt"
    "jwt-authentication-golang/database"
    "jwt-authentication-golang/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

// Whatever is sent by the client as a JSON body will be
// mapped into the user variable.
func RegisterUser(context *gin.Context) {

    var user models.User // local variable of type models.User

    if err := context.ShouldBindJSON(&user); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        context.Abort()

        return

    }
    var err error
    var haserror bool
    haserror = false
    var errstr string
    errstr = "Following infomation is missing"
    if user.Name == "" {
        errstr = fmt.Sprintf("%v %v", errstr, "Name,")
        // err = errors.New("required name parameter")
        haserror = true
    }
    if user.Username == "" {
        errstr = fmt.Sprintf("%v %v", errstr, "UserName,")
        // err = errors.New("required username parameter")
        haserror = true
    }
    if user.Password == "" {
        errstr = fmt.Sprintf("%v %v", errstr, "Password,")
        // err = errors.New("required password parameter")
        haserror = true
    }
    if user.Email == "" {
        errstr = fmt.Sprintf("%v %v", errstr, "Email,")
        // err = errors.New("required email parameter")
        haserror = true
    }
    fmt.Println("errstr")
    fmt.Println(errstr)
    if haserror == true {
        err = errors.New(errstr)
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    //we hash the password using the bcrypt helpers
    err = user.HashPassword(user.Password)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    record := database.Instance.Create(&user) //Once hashed,we store the user data into the database
    // //If there is an error while saving the data, the application would throw an HTTP Internal Server Error Code 500 and abort the request.
    if record.Error != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
        return
    }
    // // if everything goes well, we send back the user id, name, and email to the client along with a 200 SUCCESS status code.
    // context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
    context.JSON(http.StatusCreated, gin.H{"Status": "user added successfully"})
}
