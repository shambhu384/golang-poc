package main

import (
    "fmt"
    "reflect"
    "gopkg.in/go-playground/validator.v9"
    "log"
)
import "encoding/json"



type User struct {
    Name   string `json:"fullname,omitempty" validate:"required`
    Period int `json:"age" validate:"gte=0,lte=130"`
}

type Worker interface {
    Work()
    Smile()
}

func (u User) Work() {
    fmt.Println(u.Name, "has worked for", u.Period, "hrs.")
}

func (u User) Smile() {
    fmt.Println(u.Name, "has worked for  sd ", u.Period, "hrs.")
}

var validate *validator.Validate

func main() {
    validate = validator.New()
    uval := &User{"Scott", 140}
    err := validate.Struct(uval)

    fmt.Fatal(err)

    if err != nil {

        // this check is only needed when your code could produce
        // an invalid value for validation such as interface with nil
        // value most including myself do not usually have code like this.
        if _, ok := err.(*validator.InvalidValidationError); ok {
            fmt.Println(err)
            return
        }

        for _, err := range err.(validator.ValidationErrors) {

            fmt.Println(err.Namespace())
            fmt.Println(err.Field())
            fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
            fmt.Println(err.StructField())     // by passing alt name to ReportError like below
            fmt.Println(err.Tag())
            fmt.Println(err.ActualTag())
            fmt.Println(err.Kind())
            fmt.Println(err.Type())
            fmt.Println(err.Value())
            fmt.Println(err.Param())
            fmt.Println()
        }

        // from here you can create your own error messages in whatever language you wish
    }

    json, err := json.Marshal(uval)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(json))


    uval.Work()

    DoWork(uval)

    pval := &User{"UserPtr", 6}
    DoWork(pval)
}

func DoWork(w Worker) {
    fmt.Println(reflect.TypeOf(w))
    w.Work()
}
