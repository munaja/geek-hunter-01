package main

// Konsep logika:
// - open bracket berarti tambah stack, tutup brecket berarti kurangi stack
// - waktu tutup bracket cek kesesuaian, jika tidak berarti NO
// - di akhir proses, jika jumlah stack tidak 0 berarti NO

func balanceBracket(s string) string {
	const yes, no = "Yes", "No"
	// key-val open bracket
	openBrackets := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	// key-val close bracket
	closeBrackets := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	// stack, harusnya 0 diakhir proses
	bracketStack := []byte{}

	// check semua bagian string
	for i := range s {
		// check bracket only, non bracket diingore
		if _, ok := openBrackets[s[i]]; ok { // untuk open bracket
			bracketStack = append(bracketStack, s[i])

		} else if desiredBracket, ok := closeBrackets[s[i]]; ok { // untuk close brakcet
			// persiapan
			lastStackIdx := len(bracketStack) - 1
			lastStackNracket := bracketStack[lastStackIdx]

			// jika tutup tidak seusai dengan stack terakhir, proses berakhir premature dengan hasil No
			if desiredBracket != lastStackNracket {
				return no
			}

			// stack berkurang 1
			bracketStack = bracketStack[:lastStackIdx]
		}
	}

	// jika tidak habis (sisa), berarti No
	if len(bracketStack) > 0 {
		return no
	}

	// Sukses
	return yes
}
