package util

const (
    VIDEO_UPLOAD_LOCATION = "/douyin/static/feedfile/"

    USERNAME_MIN_LENGTH int = 6
    USERNAME_MAX_LENGTH int = 15
    PASSWORD_MIN_LENGTH int = 6
    PASSWORD_MAX_LENGTH int = 15
    USERNAMEREG string = `^(1[345789]\d{1})(\d{8})$`
    PASSWORDREG string = `^[A-Za-z][A-Za-z0-9!#~$*+_]{5,14}$`
)
