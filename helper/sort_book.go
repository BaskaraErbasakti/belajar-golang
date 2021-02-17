package helper

import (
	"sort"
	"strconv"
	"strings"
)

func SortBook(input string) string {
	ubahKeArray := strings.Split(input, " ")
	urutanKategori := []int{6, 7, 0, 9, 4, 8, 1, 2, 5, 3}
	wadah := []string{}
	for i := 0; i < len(urutanKategori); i++ {
		kategori := []string{}
		for j := 0; j < len(ubahKeArray); j++ {
			if string(ubahKeArray[j][0]) == strconv.Itoa(urutanKategori[i]) {
				kategori = append(kategori, ubahKeArray[j])
			}
		}
		if len(kategori) >= 1 {
			sort.SliceStable(kategori, func(i, j int) bool {
				return kategori[i][2:4] > kategori[j][2:4]
			})
			for k := 0; k < len(kategori); k++ {
				wadah = append(wadah, kategori[k])
			}
		}
	}
	hasilAkhir := strings.Join(wadah, " ")
	return hasilAkhir

}
