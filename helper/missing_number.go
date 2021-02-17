package helper

import (
	"fmt"
	"strconv"
)

func MissingNumber(x string) int {
	angkaAwalDigit := []int{}
	angkaAkhirDigit := []int{}
	angkaHilang := 0
	for i := 1; i <= 7; i++ {
		if i >= len(x) || i+i >= len(x) {
			break
		}
		angkaPertama, err := strconv.Atoi(x[0:i])
		if err != nil {
			fmt.Println(err)
		}

		angkaKedua, err := strconv.Atoi(x[i : i+i])
		if err != nil {
			fmt.Println(err)
		}

		angkaTerakhir, err := strconv.Atoi(x[len(x)-i:])
		if err != nil {
			fmt.Println(err)
		}

		angkaTerakhirKedua, err := strconv.Atoi(x[len(x)-i-i : len(x)-i])
		if err != nil {
			fmt.Println(err)
		}
		if angkaPertama+1 == angkaKedua {
			angkaAwalDigit = append(angkaAwalDigit, i)
		} else if angkaTerakhir-1 == angkaTerakhirKedua {
			angkaAkhirDigit = append(angkaAkhirDigit, i)
		}
	}
	if len(angkaAwalDigit) >= 1 {
		digitAwal := angkaAwalDigit[len(angkaAwalDigit)-1]
		awal := x[0:digitAwal]
		pembanding := ""
		for j := 0; j < 1000; j++ {
			tambah, err := strconv.Atoi(awal)
			if err != nil {
				fmt.Println(err)
			}
			penambahan := tambah + j
			pembanding = pembanding + strconv.Itoa(penambahan)
			if pembanding != x[0:len(pembanding)] {
				angkaHilang = penambahan
				break
			}
		}
	} else {
		digitAkhir := angkaAkhirDigit[len(angkaAkhirDigit)-1]
		akhir := x[len(x)-digitAkhir:]
		pembanding := ""
		for j := 0; j < 1000; j++ {
			kurang, err := strconv.Atoi(akhir)
			if err != nil {
				fmt.Println(err)
			}
			pengurangan := kurang - j
			pembanding = strconv.Itoa(pengurangan) + pembanding
			if pembanding != x[len(x)-len(pembanding):] {
				angkaHilang = pengurangan
				break
			}
		}
	}
	return angkaHilang
}
