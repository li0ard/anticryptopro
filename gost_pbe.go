package main

import(
	"fmt"
	"encoding/hex"
	"unicode/utf16"
	"github.com/thefish/gogost/gost341194"
	"github.com/thefish/gogost/gost28147"
)

func utf16le(val string) []byte {
	var v []byte
	for _, r := range val {
		if utf16.IsSurrogate(r) {
			r1, r2 := utf16.EncodeRune(r)
			v = append(v, byte(r1), byte(r1>>8))
			v = append(v, byte(r2), byte(r2>>8))
		} else {
			v = append(v, byte(r), byte(r>>8))
		}
	}
	return v
}

func decodeHex(val string) []byte {
	j, _ := hex.DecodeString(val)
	return j
}

func main() {
	fmt.Println("CryptoPro PBE\nby li0ard")
	PASS := "" //"123"
	SALT := "" //"C16E378ABE17ADBC7C29E4F5EA4EEED9"
	fmt.Print("Enter pass: ")
	fmt.Scanln(&PASS)
	fmt.Print("\n")
	fmt.Print("Enter SALT: ")
	fmt.Scanln(&SALT)
	fmt.Println(" ")
	KEY := utf16le(PASS)
	for i := 1; i < 0x7D0 + 1; i++ {
		hasher := gost341194.New(&gost28147.SboxIdGostR341194CryptoProParamSet)
		a := decodeHex(hex.EncodeToString(KEY) + SALT + fmt.Sprintf("%04s", fmt.Sprintf("%x", i)))
		hasher.Write(a)
		KEY = hasher.Sum(nil)
	}
	fmt.Println("KEY  = " + hex.EncodeToString(KEY))
	fmt.Println("SALT = " + SALT[:16])
}
