package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountByLoop(x uint64) int {
	res := pc[byte(x>>0*8)]
	for i := uint(1); i < 8; i++ {
		res += pc[byte(x>>i*8)]
	}
	return int(res)
}

func PopCountByClearing(x uint64) int {
	res := 0
	for x != 0 {
		x = x & (x - 1)
		res++
	}
	return res
}

func PopCountByShifting(x uint64) int {
	res := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			res++
		}
	}
	return res
}
