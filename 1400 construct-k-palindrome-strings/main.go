package main

func main() {
	var m = make(map[string]int)
	var word = "annabelle"
	n_odd := 0
	k := 2
	for _, letter := range word {
		m[string(letter)]++
	}
	for x, v := range m {
		println(x, v)
		if v%2 != 0 {
			n_odd++
		}
	}
	if k < n_odd || k > len(word) {
		println("false")
	} else {
		println("true")
	}
}
