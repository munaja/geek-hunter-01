package main

import "strconv"

// Konsep logika yang efisian (non bruteforce) rekursif:
// - Penelusuran kesamaan dilakukan dari ujung ke tengah string atau sebaliknya, dibandingkan lawan posisi string (kanan vs kiri)
//   - Jika ditemukan ketidaksamaan periksa apakah k masih tersedia (k = max berapa kali diubah)
//     - Jika masih boleh berubah
//        - Ubah salah satu mengikuti yang lebih besar
//        - Nilai K berkurang satu karena sudah dipakai
//     - Jika tidak boleh berubah berakhir premature, -1
//   - Jika penelusuran berakhir
//     - Periksa apakah masih ada sisa k
//       - Jika ada gunakan untuk memaksimalkan angka2 menjadi 9
//       - posisi yang tidak pernah berubah konsumsi k 2, selebihnya 1

// Note:
// Kelemahan rekursif adalah ada kemungkinan terdapat bagian sama yg diulang dalam
// perulangan fungsi yang bisa dikeluarkan dari perulangan jika pake loop biasa

// Menggunakan variadic agar tidak perlu menyertakan index pencarian di awal serta
// untuk menampung info index palindrom dan non palindrom untuk memaksimalkan hasil akhir
// index diset ke n[0]

func palindromSearch(s string, k int, n ...int) int {
	// prep
	sLength := len(s) // len(s) simpan variable karena sering diapake

	// check index pada optional paramemeter
	if len(n) == 0 {
		// cek panjang string
		if sLength%2 == 0 { // genap index dimulai dari: hasil bagi - 1
			n = append(n, sLength/2-1)
		} else {
			n = append(n, sLength/2)
		}
	}

	idx := n[0]
	counterPos := sLength - n[0] - 1
	mark := 0
	// jika ketemu non palindrom
	if s[idx] != s[counterPos] {
		// jika k masih ada, boleh ganti kareakter
		if k > 0 {
			// pilih angka yang kecil diganti angka yangg lebih besar
			if s[idx] > s[counterPos] {
				s = s[:counterPos] + string(s[idx]) + s[counterPos+1:]
			} else {
				s = s[:idx] + string(s[counterPos]) + s[idx:]
			}
			k-- // k terus berkurang sampai 0
		} else {
			return -1 // fail karena ada non palindrom sementara k sudah habis
		}
		mark = 1
	}
	x := []int{n[0], mark}
	n = append(x, n[1:]...)

	// check index
	if n[0] > 0 { // belum selesai
		n[0]--
		return palindromSearch(s, k, n...)
	}

	// sudah selesai maksimalkan sisa K yang ada
	idx = 1 // index palindrom mark
	for k > 0 && idx < len(n) {
		// jika angka belum 9 berarti belum maksimal
		if s[idx-1] != 57 {
			counterPos := sLength - idx
			if k >= 2 { // cek k masih diatas
				// palindrom mark = tidak berubah, maksimal kan 9 semua
				if n[idx] == 0 {
					s = s[:idx-1] + "9" + s[idx:counterPos] + "9" + s[counterPos+1:]
					k -= 2
				} else {
					s = s[:idx-1] + "9" + s[idx:counterPos] + "9" + s[counterPos+1:]
					k--
				}
			} else {
				// pernah berubah berarti boleh dimaksimalkan karena perubahan awal hanya consume k = 1
				if n[idx] == 1 {
					s = s[:idx-1] + "9" + s[idx:counterPos] + "9" + s[counterPos+1:]
					break // selesai k pasti habis
				}
			}
		}
		idx++
	}
	result, _ := strconv.Atoi(s)
	return result
}
