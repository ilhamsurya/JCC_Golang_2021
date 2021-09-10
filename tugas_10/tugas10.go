package main
import "tugas10/library"
import "fmt"
import "strconv"
func main() {
	fmt.Println("=========== Soal 1 =============")
	var bangunDatar library.HitungBangunDatar
	var bangunRuang library.HitungBangunRuang
	bangunDatar = library.PersegiPanjang{10, 5}
	fmt.Println("===== persegiPanjang")
	fmt.Println("luas      :", bangunDatar.Luas())
	fmt.Println("keliling  :", bangunDatar.Keliling())

	bangunDatar = library.SegitigaSamaSisi{14, 7}
	fmt.Println("===== segitigaSamaSisi")
	fmt.Println("luas      :", bangunDatar.Luas())
	fmt.Println("keliling  :", bangunDatar.Keliling())

	bangunRuang = library.Tabung{14.7, 7}
	fmt.Println("===== tabung")
	fmt.Println("volume      :", bangunRuang.Volume())
	fmt.Println("luasPermukaan  :", bangunRuang.LuasPermukaan())

	bangunRuang = library.Balok{9, 10, 11}
	fmt.Println("===== balok")
	fmt.Println("volume      :", bangunRuang.Volume())
	fmt.Println("luasPermukaan  :", bangunRuang.LuasPermukaan())

  fmt.Println("=========== Soal 2 =============")
	var color []string
	var color_list = []string{"merah", "hijau", "kuning"}
	xiaomi := library.Phone{"Redmi10", "Xiaomi", 2021, color}
	color = library.Phone.AddColors(xiaomi, color_list...)
	xiaomi.Colors = color
	xiaomi.Display(xiaomi)

  fmt.Println("=========== Soal 3 =============")
	fmt.Println(library.LuasPersegi(4, true))
	fmt.Println(library.LuasPersegi(8, false))
	fmt.Println(library.LuasPersegi(0, true))
	fmt.Println(library.LuasPersegi(0, false))

  fmt.Println("=========== Soal 4 =============")
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	kalimat := "" + prefix.(string)

	kumpulanAngka := []int{}

	kumpulanAngka = append(kumpulanAngka, kumpulanAngkaPertama.([]int)[0], kumpulanAngkaPertama.([]int)[1])
	kumpulanAngka = append(kumpulanAngka, kumpulanAngkaKedua.([]int)[0], kumpulanAngkaKedua.([]int)[1])

	jumlah := 0

	for index, item := range kumpulanAngka {
		if index == 0 {
			kalimat += strconv.Itoa(item)
		} else {
			kalimat += "+" + strconv.Itoa(item)
		}
		jumlah += item
	}

	kalimat += "=" + strconv.Itoa(jumlah)

	fmt.Println(kalimat)
}
