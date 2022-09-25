package main

import "fmt"

func main(){
	var nilaiAkhir = 91
	var nilaiAbsensi = 84

	var lulusNilaiAkhir = nilaiAkhir>80
	var lulusNilaiAbsensi = nilaiAbsensi>80

	var checkLulus = lulusNilaiAkhir&&lulusNilaiAbsensi
	fmt.Println(checkLulus)

}