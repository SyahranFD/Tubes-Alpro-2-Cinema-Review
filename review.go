package main
import "fmt"

const NMAX = 999
type Movie struct {
	judul string
	tahunRilis int
	genre string
	deskripsi string
	skor float64
}
type Movies [NMAX]Movie

func main() {
	var tabMovies Movies
	var n, pilihan, i int
	var selesai = false
	var judul, genre string
	n = 0

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
			inputMovie(&tabMovies, &n)
			getAllMovies(tabMovies, n)

		} else if pilihan == 2 {
			getAllMovies(tabMovies, n)
			fmt.Printf("Masukkan nomor film yang ingin diubah (1-%d): ", n)
			fmt.Scan(&i)
			updateMovie(&tabMovies, i-1)
			getAllMovies(tabMovies, n)

		} else if pilihan == 3 {
			getAllMovies(tabMovies, n)
			fmt.Printf("Masukkan nomor film yang ingin dihapus (1-%d): ", n)
			fmt.Scan(&i)
			deleteMovie(&tabMovies, &n, i-1)
			getAllMovies(tabMovies, n)

		} else if pilihan == 4 {
			fmt.Print("Masukkan judul film yang ingin dicari: ")
			fmt.Scan(&judul)
			findMoviesByTitle(tabMovies, n, judul)

		} else if pilihan == 5 {
			fmt.Print("Masukkan genre film yang ingin dicari: ")
			fmt.Scan(&genre)
			findMoviesByGenre(tabMovies, n, genre)

		} else if pilihan == 6 {
			sortMoviesByRating(&tabMovies, n)
			getAllMovies(tabMovies, n)

		} else if pilihan == 7 {
			sortMoviesByReleaseYear(&tabMovies, n)
			getAllMovies(tabMovies, n)

		} else if pilihan == 8 {
			statisticMovie(tabMovies, n)

		} else if pilihan == 9 {
			getAllMovies(tabMovies, n)

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
}

func deleteMovie(tabMovies *Movies, n *int, i int) {
	var j int
	j = i
	for j < *n-1 {
		tabMovies[j] = tabMovies[j+1]
		j++
	}
	*n--
}

func getAllMovies(tabMovies Movies, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("\n%d. %s (%d) - Skor: %.1f\n   Deskripsi: %s\n", i+1, tabMovies[i].judul, tabMovies[i].tahunRilis, tabMovies[i].skor, tabMovies[i].deskripsi)
	}
}

func findMoviesByTitle(tabMovies Movies, n int, judul string) {
	var left, right, mid int
	var found bool
	left = 0
	right = n - 1
	mid = left + right/2
	found = false

	for i := 1; i < n; i++ {
		temp := tabMovies[i]
		j := i - 1
		for j >= 0 && tabMovies[j].judul > temp.judul {
			tabMovies[j+1] = tabMovies[j]
			j--
		}
		tabMovies[j+1] = temp
	}

	for left <= right && !found {
		mid = (left + right) / 2
		if tabMovies[mid].judul == judul {
			found = true
			fmt.Printf("\nFilm Ditemukan!\n")
			fmt.Printf("%s (%d) - Skor: %.1f\n   Deskripsi: %s\n", tabMovies[mid].judul, tabMovies[mid].tahunRilis, tabMovies[mid].skor, tabMovies[mid].deskripsi)
		} else if tabMovies[mid].judul < judul {
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
			fmt.Printf("%d. %s (%d) - Skor: %.1f\n   Deskripsi: %s\n", count, tabMovies[i].judul, tabMovies[i].tahunRilis, tabMovies[i].skor, tabMovies[i].deskripsi)
		}
	}
	if count == 0 {
		fmt.Println("Tidak ada film yang ditemukan dengan genre tersebut.")
	}
}
func sortMoviesByRating(tabMovies *Movies, n int) {
	var i, j, maxIdx int
	var temp Movie
	i = 0
	for i = 0; i < n-1; i++ {
		maxIdx = i
		j = i + 1
		for j < n {
			if tabMovies[j].skor > tabMovies[maxIdx].skor {
				maxIdx = j
			}
			j++
		}
		temp = tabMovies[i]
		tabMovies[i] = tabMovies[maxIdx]
		tabMovies[maxIdx] = temp
	}
}

func sortMoviesByReleaseYear(tabMovies *Movies, n int) {
	var i, j int
	var temp Movie
	for i = 1; i < n; i++ {
		temp = tabMovies[i]
		j = i - 1
		for j >= 0 && tabMovies[j].tahunRilis < temp.tahunRilis {
			tabMovies[j+1] = tabMovies[j]
			j--
		}
		tabMovies[j+1] = temp
	}
}

func statisticMovie(tabMovies Movies, n int) {
	type Genre struct {
		name  string
		count int
	}
	var tabGenre [NMAX]Genre

	var totalRating float64
	var i, j int
	var found bool

	for i = 0; i < n; i++ {
		totalRating += tabMovies[i].skor
		found = false
		j = 0
		for j < i && !found {
			if tabGenre[j].name == tabMovies[i].genre {
				tabGenre[j].count++
				found = true
			}
			j++
		}

		if !found {
			tabGenre[i].name = tabMovies[i].genre
			tabGenre[i].count = 1
		}
	}

	fmt.Printf("\nRata-Rata Rating Seluruh Film: %.2f\n", totalRating/float64(n))
	fmt.Println("Jumlah Film per Genre:")
	for i = 0; i < n; i++ {
		if tabGenre[i].name != "" {
			fmt.Printf("- %s: %d film\n", tabGenre[i].name, tabGenre[i].count)
		}
	}
}
