package main

import (
    "fmt"
    "github.com/golfz/golify"
)

func main() {
    fmt.Println("Start testing")

    //i := int32(2)
    //ip := &i
    //ip = nil

    //a := "abc"
    //var ap *string
    //ap = &a

    b := true
    bp := &b
    //bp = nil


    g1 := golify.Init(bp).
        NotNil(400, "not null").
        AsBoolean(400, "not int")

    if g1.Err != nil {
        fmt.Printf("err code: %d, err msg: %s\n", g1.Err.ErrCode, g1.Err.ErrMsg)
        return
    }

    fmt.Printf("value: %v\n", g1.Value)

}
