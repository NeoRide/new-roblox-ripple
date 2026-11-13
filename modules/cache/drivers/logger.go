package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.1" )

func OQ0BKYyXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sIrXWdrAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YecbrrZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J2hfPdi9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTYwlYloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IPS4pOrhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OOrgkTUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func glQphaiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thpoPRMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UDKcaaMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m3RkLnGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CX1l6uEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WT4JmaCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Chp9uB4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ykJt4GDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dIA5tLc9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7F7asNPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func epFZAz3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func st9KzE8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93i3EAu6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DO5G9NqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QSLXXjQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DREczubwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B91YGuVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWkP7p2ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YyRq9GIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1bUcQuzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lPIQLMlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQZb3xjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func egByIDXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kydUbSrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UEXWb0vIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dYhrXobWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1j8HcfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pCcgkGgfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tcYNKyj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ldudWWbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3mKQ4vuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 72pubgCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QIy56x2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kI7k3uiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WZtE2wX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eXeddTXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fSZvVux8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zsilusjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7wToePSUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PsXQYxcRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O5Ia6nU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKw5w2GQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZksdPSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbbZwMkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKVsZ3E8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ZXf4AYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHfWJJOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kxji2Y8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRhfB0PGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fZETid0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Nv9lyVLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CwNtogVSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h0LaZTsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mLG8kLqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9haJdVr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qNQQTCpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F3lMeBKqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMOQs7D9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gMw2lWD5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZI9bbAJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n1CGQrTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CMiFP06gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rHWP0FQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JIA6BkrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l04MXs69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IieJfUXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kGaDcA9oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3TSRVSq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tEmacMyrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PidNbsZJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ec0pWaxiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k07sTOSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AU7jStadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8uMobdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LI1mQibXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNx3imKrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J217DmMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oI6QtP14Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ljw0fAMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ypWFWP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJABkH4xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTLDAnOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a8bmmeaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LfIRzNqcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVuuW8hqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cgczAPoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iOHe4CKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cmGCVY3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRwCugBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9JDyqXc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBZuAJFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fODEP20PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvnd3NIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsB53LyEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ncz64YSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mH3oAwVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3W5flKoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GobAFOvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xZ9cTZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X6vPcsmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ComZqPBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FvfHj1rdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWlL8Et0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jO4RaEC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQifIMjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n8ZdvvkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5yDDNcwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bt8f2C82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUpVZlcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u21V13pdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DbP9vPaqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uJQGZF4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DhthkZQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2nGXQBTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LEiZ6PxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l9JK7rAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVDs8gojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZxzOoIJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pOTqm8lwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 69v4BKkyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nn0HJowtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uM82k7AiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QpcB4P4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DyPR8B3fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5qfBQZj1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xXbqJpaQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DcfHNBqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EMFeAnWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZnOP59dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8t6M4LEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHlBsEPFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03DhDDAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jTul8u6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0LycDZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmKE6lVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71q9tifwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sRqx0qoXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GaM0egiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BxplnOeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vvZ0OlJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cg4Wfj9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvPJnvXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oGMQ2KOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hVSX9XYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HPtXbm5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pfNCSVxgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ahk67XmDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XlFEDY5JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vT2deTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TY2h7BpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2DVWq8U2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QAqRTGXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func alBdh7IWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j5DqXgQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f38B7PCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lgYb2h6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zk9IJmNIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rf4vObfsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dz6jAPHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func inDNN92WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMzPb2KWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vx2BWUgQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v379ZLMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e1pKYl4sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bBLo1M2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pJuVati9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mifiVpn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8L2FhpjZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fnC8UV8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ciX3VXCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fhE6hQvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWGvhYc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b1waKJKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YqDfPWdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bXRJfB0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vWcMa8H8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jbnlt8TpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZJNDoIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fwI0AQ2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k2XPil9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mXVeJFIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q7uR8umvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gmKWA5GCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSypeuMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwK9EqPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IoTpcqTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxX7RfeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uy29uA3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rik3djQFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9mhC7FXbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PExs3SMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZTztLzCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r0KPoEOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2fiKCVbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTrlM6s6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E07Js2MAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LilfnmHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s7hOz5nGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 20IaLaRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oGaRkuvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9CfxyliEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hB0JJ4EuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Xx37F8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJHGhIUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Yj14R9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NcHZqqOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bakemzI7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t6gTsvkgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Tl6C9AOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func folONcyyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3tn5DguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uVmnoqcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RyRpaRobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qybuPKsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q57lNyzWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GXweasqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C2MAk2y6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yz9anCLjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 85qDT5sgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HXfh8ZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rVsSxWigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKEdJIbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xahyCBxtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zhAUPZYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EwHoB3zXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RGZLtAGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S08bqrRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2BY0Id0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fqnsc3ngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9MMuR53cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q7Z8oRlUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0moxD6ABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KRkf7MnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f4PzcGasWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func obRBwcoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2hv3km3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rTehQ96MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o6PbcWiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YhlpZyzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qwNyAkHGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ttjDQMsBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 067IywhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6bthC7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKxO7V7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nnbv90VZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vBDhcCEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFVvakYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NoaF9eldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmoW8YUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZDchfpvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 54VTZjlqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7DYnn2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yZydPR2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYNQST07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAOAgPXVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func co76MLBoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9liq56LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RZTA8zqhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qFJB0CxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zkco5q1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GvvxQsLTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08AjDToSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ufjSO9PUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bRD9vIrcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MPeCtu60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YdDiYujgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func htvWR9BpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQq0MvlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVDik4GiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Un1ZF9D3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oL5Xed3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ckwvSgoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3oqEWqftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zC8e5HTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M0xpQ1qFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l4C0K3RgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59cdTAB3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vjgb8Tc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdUqrPa8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hMOQ8HHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyvawtELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTg7u0NjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hFhLxWYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ydpvKXQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLSJffj3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nU5xNhD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9vfKRBBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qfCcLGBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GGTqVt0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6iamZxQhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RbIpZ82qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXUc9lXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func McDwBl5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeiSunhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X66zpJ9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func waIQOgxLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Azb1zgQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFrbz56mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YkQeD7DjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XHpCYdYJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T38xRWuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ymdpVxEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ceJdh4bAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSdLYsYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TlPT3izUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThuaZmylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ECNWNvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQQqOP9KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ma24GJ7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OIXsnH8LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MV9PgDiaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VmKRGmnLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fNxABCzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Y3HsxfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lOOAo9oGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DhTXQ7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0z8orx2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJbmyRVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n6KWSw4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KygQ0sPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b1IVGDTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6vZQV5CLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uPcJZCkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xU5SUqpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A7j3IvjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func baRPV90JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v5weO8cjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hnJIjKxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJxSBuxXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFOF9DUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PIlYBs7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0dPwezAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVVTeoL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9VB0rZ2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w2M25Ku6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lD2CnqRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rgsRwwTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4unZhreWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLKEghy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QcCbsZbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FXXJQz1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbAIOei7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZFttnxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SxKXEuSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d7j2sO3YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HU2baFYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQcVu99mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JvEquK4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P8EueP9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vaMpJl96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMgThetUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func riVT8RsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WNeSMsMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OJimxbJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FxMm8FkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HySeLoUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZgJCTNWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wigy1IpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KVT534nvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0aj0DFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j7hKOvwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wjQgWMbFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4yWgwcJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TC6vUQGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VhzZOkyQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Acp3XwiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6nGltlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Cqco1ZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjjD76H6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BtShMxUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R60XDRIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKPRD1CVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R0u6iTcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WM2mQJDLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1r1RA7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEu5YUlrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJQ0aXb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTnUJk8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJilW7RvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BvQ0H2ELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwsGaYX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d16iRFZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXmorex2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bnI1GmaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sX1bcfaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGnqCbZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aUgIRynkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZdTnmf2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIa8rou0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bITp73r7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7zuqAMSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SjawGia3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6YkUr54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RWKFPIq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NgaMzUPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2nBz0irLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ya4RJycUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2DpX4D6EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7e2TgYXkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLKvqXczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1y7dxYY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TwJtwj5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZBWdVhvsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CTLB7iTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u9P9sEEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kI52xxULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZjblaxaAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7fnewtR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZaLEqZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XnLVsqPBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4QUNzoqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MkKf9LekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 18GTnLvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HmFTkSGxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HroK09zHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9aFa9sPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CkH3EZ5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8MmZsdPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qK7ut8Z9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfIcDGrdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NDUtlJivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wrx9ag8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L1M1BpPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmaR3dHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u4fnqAw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUCn4lvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JXQCt5PkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zin8SajFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func khrH63HAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4lBnrN1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JtaYbKieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func myPhnGEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CYGwytBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FxdzB0KiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qe5HuTeSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cdnSwiDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eN9ZhA7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgQOzZkfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZgEXeifTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rjHi7K16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qFDYkFh8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hQ3mqMTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tafLz80lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bi9BNnsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f2rR2EjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5reUGroOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lj430OsoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o34brn4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fxbOGEALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MSpZmzyQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bj0CnAJhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func StPrSxdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q2biMbHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fkVNb53cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0zOVd5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e0UOBqbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func adjidX7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I3uNn0qJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPKPdpIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkCQwOnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IXWDZyhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PsDk664kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mmgk5mriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0PRHFuWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gHvHqoROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UoOhCzpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TRsCocLXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vMCsvGJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3EBKVSb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKkilAfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9W9Mc4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tjzm6bnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AIyRv7aVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zCt1862YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q8htMoJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vRGiDOsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ofBDpB2AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xRaEcm3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BILAhdEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BzaAMfTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQ8qQXaPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Njh8x2XQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nScsvyi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iuWStNQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B6lMWZRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rbI9CG25Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fd86YttlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4OId3HRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4cGQmHltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jlh1AYnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dNbIA3MxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uDDwcfaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UEZhHvYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q6zPM2XrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KuTWj5XWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yjuedkBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qtgPG7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WylqnZASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hKnpMAysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UCKC38tvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9CUE0K6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pQvcPQAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPY7Rz0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VF3rRADKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hm6eeeEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rKg5o16pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fHpYHkDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func poOSVumvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wOH76qOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8vWvt6qGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jdv71j0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uVAHyCPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rKswyg0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5UOrZLs4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RpuEB9GUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8N9pT4GEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HzjxvQDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KqZcS1mNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mPvghMyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vENtSRjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3oVjyVPCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LWDQwf6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0zrGdpAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hdfQJE99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSkhwcUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXCIbUz5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oku4UcpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RhR9jFyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I5OK4tNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func woM03xScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFeI3B0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 42gWnh6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ULFwTH3bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GTPk6r11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EoLevncvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDy0nFWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func grLS3yFQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWWvbEaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func antTEmlDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HGwM01f9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yGnMOn5aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A7mmmoURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M9ng0pPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wT5vkzHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pf42XAgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dSeoqndBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ariWMfmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QOajbUZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OJTWyWAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0NJ5BsBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Ieuu66ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ffiqgvh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gmZbM9PgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m21iFE60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NoUMsbbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func maBxnXx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JrBSuPE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zZpWqPtjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JOWRYopQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8OqnYZwCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZetznkOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gOoNU89KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BKOw8AANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQ8hkJSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hU2HmNYfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgwIJt4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7yQATmufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j2VvY7EoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 77NRhSMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQtXxdOSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7vnE0lIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThZiDZ0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func htwc7N13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkq8Oiz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mecyFtscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5G0yYaeTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5UK2l0QHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3UethyZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nNRN2teKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ES0aoKeNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OMLVb8DyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3sI7ZCAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUylp4dcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1N5uomvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CYTeEP7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D3fRWUasWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7cijyicSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DAe2r0BkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wi0fmPYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJ39ASCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WUaQkbCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUgQX17AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K8bgxKfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ib17g4mfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TUpcFEjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AsyNEMzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0K4mFvbcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wyXlRVf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func InDYt2yjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gWzqEy63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XGZV9YMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Zr5yn3CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9jGG0tbSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PADaNCXUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8FhfXcLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCYmUHs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzVTAobBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2vFHFtecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ga3XVOn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eI0wa4D5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w10tPqIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSlvO5j3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XppDEiAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMVl1oIOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6xOZYpsrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eyjmGtt9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwhbRimhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4c38LFRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0tWw2r9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JqbAl4BSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jwucnpE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IxeldqkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFkWDRqkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func islr2XK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3J6SckWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EySk0KesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXIzfIDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywdvCPyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sfUxRPTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btIPDfhhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAvWLRZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMhtlzwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lcsllAwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aB5oajegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 87DQSv68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQ7Wg3U9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsdhafaAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUNPamizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m6XDu4YBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MumQ8YWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iHv2VRYdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XoitwDIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yr0NXluAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YtAcdjFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXLtMFRZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ZQ6UM98Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9VVCsA9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I1AwhL7eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qeqd7IGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gl0XtZr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zVqqUBacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3LNa4uPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rL0gEMgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TaSJIBIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xfH02uUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BaW30EoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8bFJYHq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9gitFMkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yFKP6JdkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vek56INkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iYQsk5StWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g3g0WeXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDdrg28jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3j1hpZfmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IrNZqdMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l8iWZIh1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNFVrcX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nktsPotxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W8X2zBBEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TGsJSUIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iYnxt96dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YutFeeUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LSuhTwvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BPCKAsUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vGWfamT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DWSCYoFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mlzal1lPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3QzcbnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4VKvl8mUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9yWhAExvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PySWSf7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jOq5iCVYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ogYqwF9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RnD34Ys1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SsBgdFjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTI9j05jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o69G0jg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AX4LdPL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HqtKmNflWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5mIRT1CKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1RXOpA7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F1W5C6y3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6EDkzcTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RtRikLMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YNq9iB8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9e5zNIdpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozDnxktdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PoMdmqXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OvEFsRLmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CzEPPz2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7D2zCRpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oTWid2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TyCdKh8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func blFiFqD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhMDG0tQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aACcWPwvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGSICQWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eowjzUE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W59UHGakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B2HkducpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VL7qGWhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1vnnTWzhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 678p0oABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5BTLrCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pkSj51LnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func csCKULlcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y8Jg7gNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A4i6Rn0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nqqQu3MAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8kuvCvdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hxd2E6bqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76kQ7sOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AKfYC6lGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9aYU4vYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hk5BlfVwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HvMYnd1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LYEttUbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTC9sWIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func swTU2rkrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2kVdXy7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BlWmWQv6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y5pBRNZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HyFAo0yXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouFYxFjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bWALoxugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Q71oD5cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v2nTTMkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QC4jg5eeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cOFPQKDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z9oDFtXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cpv7C090Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4QF79QzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XpzZgHICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EexuvtCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YANdl6OqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 473lF5xYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oigqvo1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hp6hlip6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func llhquajmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 56UclINaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8aNChbDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y7XGTq1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Y4FGd3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7KlXdW1WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qbC73jmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yA1E5u5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmGFQhHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7zlpvu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9R4e25tUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

