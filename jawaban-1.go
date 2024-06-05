package main

// Konsep logika:
// - deteksi urutan karakter serupa berurutan berarti deteksi perubahan karakter
// - jika karakter berubah index-bobot bertambah untuk menuju item array selanjutnya
// - jumlahkan bobot karakter sesuai index string yang sedang di cek

func matchStringWeight(s string, q []int) []string {
	lastChar := byte(0)

	// persiapan perhitungan kan hasil
	qLength := len(q)
	charWeight := make([]int, qLength)
	result := make([]string, qLength)
	weightIdx := -1

	// loop tiap karakter dalam s
	for i := range s {
		// cek karakter terakhir karakter
		if lastChar != s[i] { // karakter baru
			// tambah index dan break jika diluar bound slice, sisanya tidak di cek
			weightIdx++
			if weightIdx >= qLength {
				break
			}
			// karena a = 97, maka -96 biar a = 1, dan abaikan non alphabet
			lastChar = s[i]
			if lastChar < 96 || lastChar > 121 {
				lastChar = 96
			}
			charWeight[weightIdx] = int(lastChar - byte(96))

		} else { // bukan karakter baru
			// tinggal tambahkan
			charWeight[weightIdx] += int(lastChar - byte(96))
		}
	}

	// cek hasil akhir
	for i := range charWeight {
		if charWeight[i] == q[i] {
			result[i] = "Yes"
		} else {
			result[i] = "No"
		}
	}

	return result
}

// Yth Tim Pembuat Soal:
// Pada contoh diberikan string 'abbcccd' dengan queries [1, 3, 9, 8]
// Serta disediakan list hasil pembobotan sesuai aturan
// a = 1
// b = 2
// bb = 4
// c = 3
// cc = 6
// ccc = 9
// d = 4

// untuk string 'abbcccd' harusnya yang dipakai
// a = 1
// bb = 4
// ccc = 9
// d = 4
// tapi kenapa query [1, 3, 9, 8] menunjukan hasil [Yes, Yes, Yes, No]?
// bukankah bb = 4, dan bukan 3?

// terimakasih
