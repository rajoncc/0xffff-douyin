package controller

const (
    STATUS_SUCCESS int32 = 0
    STATUS_FAIL int32 = -1
    USERNAME_MIN_LENGTH int = 6
    USERNAME_MAX_LENGTH int = 15
    PASSWORD_MIN_LENGTH int = 6
    PASSWORD_MAX_LENGTH int = 15
    ADDCOMMENTTYPE int = 1
    DELCOMMENTTYPE int = 2

    //UPLOAD_PATH      string = "/home/ycc/project/go/src/douyin/static/feedfile"
	UPLOAD_PATH      string = "/usr/local/webserver/nginx/html/static/feedfile"
    GET_VIDEO_NUMBER int    = 20
)
