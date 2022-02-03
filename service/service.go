package service

import (
	"math/rand"
	"time"
)

var data = map[int]string{
	0: "_", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a",
	11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "z", 18: "y", 19: "x", 20: "w",
	21: "v", 22: "u", 23: "t", 24: "s", 25: "r", 26: "q", 27: "p", 28: "o", 29: "n", 30: "m",
	31: "l", 32: "k", 33: "j", 34: "i", 35: "h", 36: "A", 37: "B", 38: "C", 39: "D", 40: "E",
	41: "F", 42: "G", 43: "H", 44: "I", 45: "J", 46: "K", 47: "L", 48: "M", 49: "N", 50: "O",
	51: "P", 52: "Q", 53: "R", 54: "S", 55: "T", 56: "U", 57: "W", 58: "V", 59: "X", 60: "Y", 61: "Z",
}

func Shorten(longurl string) string {
	rand.Seed(time.Now().UnixNano())
	shorturl := ""
	for i := 0; i < 10; i++ {
		shorturl += data[rand.Intn(62)]
	}
	return shorturl
}
