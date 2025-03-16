package main

import (
	"fmt"
	"strings"
)

func punyaHuruf(kataPertama, kataKedua string) bool { // O(m*n)
	kataPertama = strings.ToLower(kataPertama)
	kataKedua = strings.ToLower(kataKedua)

	for _, huruf := range kataPertama {
		if !strings.ContainsRune(kataKedua, huruf) {
			return false
		}
	}
	return true
}

func punyaHuruf2(kataPertama, kataKedua string) bool { // lebih cepat O(m+n) menggunakan map
	kataPertama = strings.ToLower(kataPertama)
	kataKedua = strings.ToLower(kataKedua)

	hurufAda := make(map[rune]bool)
	for _, huruf := range kataKedua {
		hurufAda[huruf] = true
	}

	for _, huruf := range kataPertama {
		if !hurufAda[huruf] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(punyaHuruf("cat", "antarctica"), punyaHuruf2("cat", "antarctica"))
	fmt.Println(punyaHuruf("cat", "australia"), punyaHuruf2("cat", "australia"))
	fmt.Println(punyaHuruf("cat", "ANTARCTICA"), punyaHuruf2("cat", "ANTARCTICA"))
}
