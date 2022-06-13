package util

import ()

func ValidateUsernamePassword(username string, password string) (bool, string) {
    //username
    var errmsg string
    if username == "" {
        errmsg = "empty username"
        return false, errmsg
    }
    if len(username) < USERNAME_MIN_LENGTH {
        errmsg = "username length < 6"
        return false, errmsg
    }
    if len(username) > USERNAME_MAX_LENGTH {
        errmsg = "username length > 15"
        return false, errmsg
    }
    //password
    if password == "" {
        errmsg = "empty password"
        return false, errmsg
    }
    if len(password) < PASSWORD_MIN_LENGTH {
        errmsg = "password length < 6"
        return false, errmsg
    }
    if len(password) > PASSWORD_MAX_LENGTH {
        errmsg = "password length > 15"
        return false, errmsg
    }
    //username and password not contains special word
    //if

    return true, ""
}
