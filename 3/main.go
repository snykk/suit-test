package main

import (
	"fmt"
	"sort"
	"strings"
)

type Klasemen struct {
	klubPoin map[string]int
}

func NewKlasemen(klubList []string) *Klasemen {
	klubPoin := make(map[string]int)
	for _, klub := range klubList {
		klubPoin[klub] = 0
	}
	return &Klasemen{klubPoin: klubPoin}
}

func (k *Klasemen) catatPermainan(klubKandang, klubTandang, skor string) {
	skorSplit := strings.Split(skor, ":")
	if len(skorSplit) != 2 {
		fmt.Println("Format skor tidak valid")
		return
	}

	var golKandang, golTandang int
	fmt.Sscanf(skorSplit[0], "%d", &golKandang)
	fmt.Sscanf(skorSplit[1], "%d", &golTandang)

	if golKandang > golTandang {
		k.klubPoin[klubKandang] += 3
	} else if golKandang < golTandang {
		k.klubPoin[klubTandang] += 3
	} else {
		k.klubPoin[klubKandang] += 1
		k.klubPoin[klubTandang] += 1
	}
}

func (k *Klasemen) cetakKlasemen() []string {
	type Klub struct {
		nama string
		poin int
	}

	var klubList []Klub
	for nama, poin := range k.klubPoin {
		klubList = append(klubList, Klub{nama, poin})
	}

	// Urutkan berdasarkan poin tertinggi
	sort.Slice(klubList, func(i, j int) bool {
		return klubList[i].poin > klubList[j].poin
	})

	// Format output
	var hasil []string
	for _, klub := range klubList {
		hasil = append(hasil, fmt.Sprintf("%s: %d", klub.nama, klub.poin))
	}
	return hasil
}

func (k *Klasemen) ambilPeringkat(nomorPeringkat int) string {
	klasemen := k.cetakKlasemen()
	if nomorPeringkat-1 >= 0 && nomorPeringkat-1 < len(klasemen) {
		return strings.Split(klasemen[nomorPeringkat-1], ":")[0]
	}
	return "Peringkat tidak ditemukan"
}

func main() {
	klasemen := NewKlasemen([]string{"Liverpool", "Chelsea", "Arsenal"})

	klasemen.catatPermainan("Arsenal", "Liverpool", "2:1")
	klasemen.catatPermainan("Arsenal", "Chelsea", "1:1")
	klasemen.catatPermainan("Chelsea", "Arsenal", "0:3")
	klasemen.catatPermainan("Chelsea", "Liverpool", "3:2")
	klasemen.catatPermainan("Liverpool", "Arsenal", "2:2")
	klasemen.catatPermainan("Liverpool", "Chelsea", "0:0")

	fmt.Println("Klasemen:")
	for _, hasil := range klasemen.cetakKlasemen() {
		fmt.Println(hasil)
	}

	fmt.Println("Klub peringkat ke-2:", klasemen.ambilPeringkat(2))
}
