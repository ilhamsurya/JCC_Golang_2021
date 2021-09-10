package library
import "math"
import "strconv"
import "fmt"

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi float64
}

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

type TampilDetailPhone interface {
	Display() string
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return 3 * s.Alas
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2*p.Panjang + 2*p.Lebar
}

func (t Tabung) Volume() float64 {
	return math.Pi * t.JariJari * 2 * t.Tinggi
}
func (t Tabung) LuasPermukaan() float64 {
	return math.Pi * 2 * t.JariJari * t.Tinggi
}

func (b Balok) Volume() float64 {
	return b.Panjang * b.Lebar * b.Tinggi
}

func (b Balok) LuasPermukaan() float64 {
	return 2*b.Panjang*b.Lebar + b.Panjang*b.Tinggi + b.Lebar*b.Tinggi
}

func (p Phone) AddColors(warna ...string) []string {
	var color_palette []string
	for _, item := range warna {
		color_palette = append(color_palette, item)
	}
	return color_palette
}

func (p Phone) Display(merk Phone) {
	s := strconv.Itoa(merk.Year)
	fmt.Println("name: ", merk.Name)
	fmt.Println("brand: ", merk.Brand)
	fmt.Println("year: ", s)
	fmt.Println("color: ", merk.Colors)
}

func LuasPersegi(sisi int, stats bool) interface{} {
	var infoLuas interface{}
	s := strconv.Itoa(sisi)
	luasPersegi := sisi * sisi
	l := strconv.Itoa(luasPersegi)
	if sisi > 0 && stats == true {
		infoLuas = "luas persegi dengan sisi " + s + " adalah " + l
	} else if sisi > 0 && stats == false {
		infoLuas = l
	} else if sisi == 0 && stats == true {
		infoLuas = "Maaf anda belum menginput sisi dari persegi"
	} else if sisi == 0 && stats == false {
		infoLuas = nil
	} else {
		infoLuas = "tidak tersedia"
	}
	return infoLuas
}