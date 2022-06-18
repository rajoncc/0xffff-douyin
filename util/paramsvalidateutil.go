package util

import (
    "regexp"
)

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
    //username regexp
    re, err := regexp.Compile(USERNAMEREG)
    if err != nil {
        errmsg = "match error"
        return false, errmsg
    }
    ok := re.MatchString(username)
    if !ok {
        errmsg = "username not right format,need telnum"
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
    //password regexp
    re, err = regexp.Compile(PASSWORDREG)
    if err != nil {
        errmsg = "match error"
        return false, errmsg
    }
    ok = re.MatchString(password)
    if !ok {
        errmsg = "password not right format,need mix Aa0 and !#~$*+_"
        return false, errmsg
    }

    return true, ""
}
