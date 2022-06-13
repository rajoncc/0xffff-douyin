package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    //Set a lower memory limit for multipart forms (default is 32 MiB)
    router.MaxMultipartMemory = 8 << 20 //8M

    setRouter(router)

    router.Run(SERVERPORT)  //port:8765
}
