package main
import (
	"fmt"
	"strings"
)

const NMAX = 999

type Movie struct {
	judul string
	tahunRilis int
	genre string
	deskripsi string
	skor float64
	ulasan string
}
type Movies [NMAX]Movie

func main() {
	var tabMovies Movies
	var n, pilihan, i int
	var selesai = false
	var judul, genre string
	n = 0
	tabMovies, n = SeedMovies()

	for !selesai {
		fmt.Println("\n=== Aplikasi Katalog dan Rating Film (CineReview) ===")
		fmt.Println("1. Tambah Film")
		fmt.Println("2. Ubah Data Film")
		fmt.Println("3. Hapus Data Film")
		fmt.Println("4. Cari Film Berdasarkan Judul")
		fmt.Println("5. Cari Film Berdasarkan Genre")
		fmt.Println("6. Urutkan Film Berdasarkan Rating Tertinggi")
		fmt.Println("7. Urutkan Film Berdasarkan Tahun Rilis Terbaru")
		fmt.Println("8. Tampilkan Statistik")
		fmt.Println("9. Tampilkan Semua Film")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			if n < NMAX {
				inputMovie(&tabMovies, &n)
				getAllMovies(tabMovies, n)
			} else {
				fmt.Println("Kapasitas dalam aplikasi kami sudah penuh, mohon untuk hapus salah satu film terlebih dahulu")
			}

		} else if pilihan == 2 {
			if n == 0 {
				fmt.Println("Belum ada data film, mohon untuk menambahkan film.")
			} else {
				getAllMovies(tabMovies, n)
				fmt.Printf("Masukkan nomor film yang ingin diubah (1-%d): ", n)
				fmt.Scan(&i)
				if i >= 1 && i <= n {
					updateMovie(&tabMovies, i-1)
					getAllMovies(tabMovies, n)
				} else {
					fmt.Println("Nomor film yang anda masukan tidak valid!")
				}
			}

		} else if pilihan == 3 {
			if n == 0 {
				fmt.Println("Belum ada data film, mohon untuk menambahkan film.")
			} else {
				getAllMovies(tabMovies, n)
				fmt.Printf("Masukkan nomor film yang ingin dihapus (1-%d): ", n)
				fmt.Scan(&i)

				if i >= 1 && i <= n {
					deleteMovie(&tabMovies, &n, i-1)
					getAllMovies(tabMovies, n)
				} else {
					fmt.Println("Nomor film yang anda masukan tidak valid!")
				}
			}

		} else if pilihan == 4 {
			if n == 0 {
				fmt.Println("Katalog film masih kosong")
			} else {
				fmt.Print("Masukkan judul film yang ingin dicari: ")
				fmt.Scan(&judul)
				findMoviesByTitle(tabMovies, n, judul)
			}

		} else if pilihan == 5 {
			if n == 0 {
				fmt.Println("Katalog film masih kosong")
			} else {
				fmt.Print("Masukkan genre film yang ingin dicari: ")
				fmt.Scan(&genre)
				findMoviesByGenre(tabMovies, n, genre)
			}

		} else if pilihan == 6 {
			if n == 0 {
				fmt.Println("Katalog film masih kosong")
			} else {
				sortMoviesByRating(&tabMovies, n)
				getAllMovies(tabMovies, n)
			}

		} else if pilihan == 7 {
			if n == 0 {
				fmt.Println("Katalog film masih kosong")
			} else {
				sortMoviesByReleaseYear(&tabMovies, n)
				getAllMovies(tabMovies, n)
			}

		} else if pilihan == 8 {
			statisticMovie(tabMovies, n)

		} else if pilihan == 9 {
			if n == 0 {
				fmt.Println("Katalog film masih kosong")
			} else {
				getAllMovies(tabMovies, n)
			}

		} else if pilihan == 0 {
			selesai = true
			fmt.Println("\nTerima kasih telah menggunakan CineReview!")
		} else {
			fmt.Println("\nPilihan tidak valid.")
		}
	}
}

func inputMovie(tabMovies *Movies, n *int) {
	fmt.Printf("\nMasukkan judul film: ")
	fmt.Scan(&tabMovies[*n].judul)

	fmt.Printf("Masukkan tahun rilis film: ")
	fmt.Scan(&tabMovies[*n].tahunRilis)

	fmt.Printf("Masukkan genre film: ")
	fmt.Scan(&tabMovies[*n].genre)

	fmt.Printf("Masukkan deskripsi film: ")
	fmt.Scan(&tabMovies[*n].deskripsi)

	fmt.Printf("Masukkan skor film: ")
	fmt.Scan(&tabMovies[*n].skor)

	fmt.Printf("Masukkan ulasan film: ")
	fmt.Scan(&tabMovies[*n].ulasan)

	*n++
}

func updateMovie(tabMovies *Movies, i int) {
	fmt.Printf("\nMasukkan judul film baru: ")
	fmt.Scan(&tabMovies[i].judul)

	fmt.Printf("Masukkan tahun rilis film baru: ")
	fmt.Scan(&tabMovies[i].tahunRilis)

	fmt.Printf("Masukkan genre film baru: ")
	fmt.Scan(&tabMovies[i].genre)

	fmt.Printf("Masukkan deskripsi film baru: ")
	fmt.Scan(&tabMovies[i].deskripsi)

	fmt.Printf("Masukkan skor film baru: ")
	fmt.Scan(&tabMovies[i].skor)

	fmt.Printf("Masukkan ulasan film baru: ")
	fmt.Scan(&tabMovies[i].ulasan)
}

func deleteMovie(tabMovies *Movies, n *int, i int) {
	var pass int
	pass = i
	for pass < *n-1 {
		tabMovies[pass] = tabMovies[pass+1]
		pass++
	}
	*n--
}

func printMovie(index int, movie Movie) {
	fmt.Printf("\n%d. %s (%d)\n   Genre: %s\n   Skor: %.1f\n   Deskripsi: %s\n   Ulasan: %s\n", index, movie.judul, movie.tahunRilis, movie.genre, movie.skor, movie.deskripsi, movie.ulasan)
}

func getAllMovies(tabMovies Movies, n int) {
	var i int
	for i = 0; i < n; i++ {
		printMovie(i+1, tabMovies[i])
	}
}

func findMoviesByTitle(tabMovies Movies, n int, judul string) {
	var i, left, right, mid int
	var found bool

	var temp Movie
	var pass int

	left = 0
	right = n - 1
	mid = (left + right) / 2
	found = false

	for i = 1; i < n; i++ {
		temp = tabMovies[i]
		pass = i - 1
		for pass >= 0 && strings.ToLower(tabMovies[pass].judul) > strings.ToLower(temp.judul) {
			tabMovies[pass+1] = tabMovies[pass]
			pass--
		}
		tabMovies[pass+1] = temp
	}

	for left <= right && !found {
		mid = (left + right) / 2
		if strings.EqualFold(tabMovies[mid].judul, judul) {
			found = true
			fmt.Printf("\nFilm Ditemukan!\n")
			printMovie(1, tabMovies[mid])
		} else if strings.ToLower(tabMovies[mid].judul) < strings.ToLower(judul) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if !found {
		fmt.Println("\nFilm tidak ditemukan.")
	}
}
func findMoviesByGenre(tabMovies Movies, n int, genre string) {
	var i, count int
	count = 0
	fmt.Printf("\nHasil pencarian untuk genre '%s':\n", genre)
	for i = 0; i < n; i++ {
		if tabMovies[i].genre == genre {
			count++
			printMovie(count, tabMovies[i])
		}
	}
	if count == 0 {
		fmt.Println("Tidak ada film yang ditemukan dengan genre tersebut.")
	}
}

func sortMoviesByRating(tabMovies *Movies, n int) {
	var i, pass, maxIdx int
	var temp Movie
	i = 0
	for i = 0; i < n-1; i++ {
		maxIdx = i
		pass = i + 1
		for pass < n {
			if tabMovies[pass].skor > tabMovies[maxIdx].skor {
				maxIdx = pass
			}
			pass++
		}
		temp = tabMovies[i]
		tabMovies[i] = tabMovies[maxIdx]
		tabMovies[maxIdx] = temp
	}
}

func sortMoviesByReleaseYear(tabMovies *Movies, n int) {
	var i, pass int
	var temp Movie
	for i = 1; i < n; i++ {
		temp = tabMovies[i]
		pass = i - 1
		for pass >= 0 && tabMovies[pass].tahunRilis < temp.tahunRilis {
			tabMovies[pass+1] = tabMovies[pass]
			pass--
		}
		tabMovies[pass+1] = temp
	}
}

func statisticMovie(tabMovies Movies, n int) {
	type Genre struct {
		name  string
		count int
	}
	var tabGenre [NMAX]Genre
	var totalRating float64
	var i, pass, nGenre int
	var found bool

	if n == 0 {
		fmt.Println("Belum ada film yang dimasukkan, mohon untuk memasukkan film terlebih dahulu")
	} else {
		nGenre = 0
		totalRating = 0.0
		for i = 0; i < n; i++ {
			totalRating += tabMovies[i].skor
			found = false
			pass = 0
			for pass < nGenre && !found {
				if tabGenre[pass].name == tabMovies[i].genre {
					tabGenre[pass].count++
					found = true
				}
				pass++
			}

			if !found {
				tabGenre[nGenre].name = tabMovies[i].genre
				tabGenre[nGenre].count = 1
				nGenre++
			}
		}
		fmt.Printf("\nRata-Rata Rating Seluruh Film: %.2f\n", totalRating/float64(n))
		fmt.Println("Jumlah Film per Genre:")
		for i = 0; i < nGenre; i++ {
			fmt.Printf("- %s: %d film\n", tabGenre[i].name, tabGenre[i].count)
		}
	}
}
