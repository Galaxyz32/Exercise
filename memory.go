package main

func main() {
	const (
		_          = iota
		KB float64 = 1 << (iota * 10)
		MB
		GB
		TB
		PB
		EB
		ZB
	)
	print(KB)
	print(MB)
	print(GB)
	print(TB)
	print(PB)
	print(EB)
	print(ZB)

}
