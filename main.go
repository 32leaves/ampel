package main;

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/stianeikeland/go-rpio"
)

type Color int
const (
    Red    Color = 10
    Yellow Color = 20
    Green  Color = 30
)

type PinMap map[Color]rpio.Pin
type HttpRequest func(http.ResponseWriter, *http.Request)

func main() {
    pinMap := initPins()
    router := initRouter(pinMap)

    log.Fatal(http.ListenAndServe(":8080", router))
    fmt.Println("Running on :8080")
}

func initPins() PinMap {
    r := make(PinMap)
    
    r[Red] = rpio.Pin(1)
    r[Red].Output()

    r[Yellow] = rpio.Pin(2)
    r[Yellow].Output()
    
    r[Green] = rpio.Pin(3)
    r[Green].Output()
    
    return r
}

func initRouter(pinMap PinMap) *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/red", GetLight(pinMap[Red])).Methods("GET")
    router.HandleFunc("/red", SetLight(pinMap[Red], true)).Methods("PUT")
    router.HandleFunc("/red", SetLight(pinMap[Red], false)).Methods("DELETE")
    router.HandleFunc("/yellow", GetLight(pinMap[Yellow])).Methods("GET")
    router.HandleFunc("/yellow", SetLight(pinMap[Yellow], true)).Methods("PUT")
    router.HandleFunc("/yellow", SetLight(pinMap[Yellow], false)).Methods("DELETE")
    router.HandleFunc("/green", GetLight(pinMap[Green])).Methods("GET")
    router.HandleFunc("/green", SetLight(pinMap[Green], true)).Methods("PUT")
    router.HandleFunc("/green", SetLight(pinMap[Green], false)).Methods("DELETE")
    return router
}

func GetLight(p rpio.Pin) HttpRequest {
    return func(w http.ResponseWriter, r *http.Request) {
        if p.Read() == rpio.High {
            w.Write([]byte("on"))
        } else {
            w.Write([]byte("off"))
        }
    }
}

func SetLight(p rpio.Pin, on bool) HttpRequest {
    return func(w http.ResponseWriter, r *http.Request) {
        if on {
            p.High()
            w.Write([]byte("on"))
        } else {
            p.Low()
            w.Write([]byte("off"))
        }
    }
}
