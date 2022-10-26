package main //pada awal kode golang kita perlu menambahkan package + nama packagenya
import "fmt" //import fmt untuk menampilkan output

//deklarasi variabel
var nama string = "Mulki" //var nama_variabel tipe_data = value
var number int = 10       // tipe data integer
var boolean bool = true   // tipe data boolean

//constant
const jumlahHari int = 7 //const merupakan data yang value nya tidak bisa diubah
//multi constant declaration
const (
	jumlahBulan      = 12
	gender           = "Laki-laki, Perempuan"
	tidur       bool = true
)

//array
var hewan [5]string
var hari [5]string //deklarasi array berisi max 5 data dengan tipe data string

func main() { //fungsi di go menggunakan keyword func + nama fungsi, disini nama fungsinya main
	fmt.Println("Hello, World!, Nama saya " + nama) //menampilkan output, disini menggunakan fungsi Println dari package fmt

	// menambahkan elemen array
	hari[0] = "Senin" //mengisi array pada index ke 0
	hari[1] = "Selasa"
	hari[2] = "Rabu"
	hari[3] = "Kamis"
	hari[4] = "Jumat" //mengisi array pada index ke 4
	//hari[5] = "Sabtu"                    //error karena array hanya bisa diisi max 5 data
	fmt.Println(hari[0], hari[1], hari[2]) //menampilkan output array pada index ke 0-2

	hewan = [5]string{"kucing", "kambing", "sapi", "kuda", "ayam"}
	fmt.Println(hewan[0], hewan[1], hewan[2])
}
