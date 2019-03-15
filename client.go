package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"

	"github.com/go-vgo/robotgo"
)

type Map map[string]interface{}

var keycode = Map{
	"`": 41,
	"1": 2,
	"2": 3,
	"3": 4,
	"4": 5,
	"5": 6,
	"6": 7,
	"7": 8,
	"8": 9,
	"9": 10,
	"0": 11,
	"-": 12,
	"+": 13,
	//
	"q":  16,
	"w":  17,
	"e":  18,
	"r":  19,
	"t":  20,
	"y":  21,
	"u":  22,
	"i":  23,
	"o":  24,
	"p":  25,
	"[":  26,
	"]":  27,
	"\\": 28,
	//
	"a": 30,
	"s": 31,
	"d": 32,
	"f": 33,
	"g": 34,
	"h": 35,
	"j": 36,
	"k": 37,
	"l": 38,
	";": 39,
	"'": 40,
	//
	"z": 44,
	"x": 45,
	"c": 46,
	"v": 47,
	"b": 48,
	"n": 49,
	"m": 50,
	",": 51,
	".": 52,
	"/": 53,
	//
	"f1":  59,
	"f2":  60,
	"f3":  61,
	"f4":  62,
	"f5":  63,
	"f6":  64,
	"f7":  65,
	"f8":  66,
	"f9":  67,
	"f10": 68,
	"f11": 69,
	"f12": 70,
	// more
	"esc":     1,
	"tab":     15,
	"ctrl":    29,
	"control": 29,
	"alt":     56,
	"space":   57,
	"shift":   42,
	"enter":   28,
	"cmd":     3675,
	"command": 3675,
}

func main() {
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%s", os.Args[1], os.Args[2]))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	s := robotgo.Start()

	k := 0
	var cursor []int
	for ev := range s {
		if ev.Kind == 4 && ev.Keycode == uint16(keycode["control"].(int)) {
			k++
		}
		if ev.Kind == 4 && ev.Keycode == uint16(keycode["shift"].(int)) {
			k++
		}
		if ev.Kind == 4 && ev.Keycode == uint16(keycode["a"].(int)) {
			k++
		}
		if k == 3 && ev.Kind == 6 {
			cursor = append(cursor, int(ev.X))
			cursor = append(cursor, int(ev.Y))
		}
		if len(cursor) == 4 {
			bitmap := robotgo.CaptureScreen(cursor[0], cursor[1], cursor[2]-cursor[0], cursor[3]-cursor[1])
			bitmapstr := robotgo.TostringBitmap(bitmap)
			robotgo.FreeBitmap(bitmap)

			var reply string
			err = client.Call("ScreenshotService.SendBitmap", bitmapstr, &reply)
			if err != nil {
				log.Fatal(err)
			}
			cursor = []int{}
			k = 0
		}
	}
}
