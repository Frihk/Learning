package piscine

// func IsPrime(nb int) bool{
// 	n
// }

func FindNextPrime(nb int) int {
	if IsPrime(nb) == true {
		return nb
	}
	return FindNextPrime(nb + 1)
}
