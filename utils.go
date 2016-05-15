package main

func MinInt(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func SanitizeText(text string) string {
	end := MinInt(len(text), 256)
	return text[0:end]

}
