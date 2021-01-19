package utils

// 快速幂算法
func QuickPow(a uint64, b uint64) uint64 {
	var ans uint64 = 1
	for {
		if b == 0 {
			break
		}
		if b&1 == 1 {
			ans *= a
		}
		a *= a
		b = b >> 1
	}
	return ans
}

// 快速幂取模算法
func QuickPowMod(a uint64, b uint64, m uint64) uint64 {
	var ans uint64 = 1
	mod := a % m
	for {
		if b == 0 {
			break
		}
		if b&1 == 1 {
			ans = (ans * mod) % m
		}
		mod = (mod * mod) % m
		b = b >> 1
	}
	return ans
}

// BKDRHash 算法
func BKDRHash(str string) uint64 {
	seed := uint64(131) // 31 131 1313 13131 131313 etc..
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + uint64(str[i])
	}

	// long int 的上限值，这里可以替换为得到的数据与散列表的大小NHASH进行mod运算 hash % HLEN
	// 0x7fffffff = 2147483647 = (1 << 31) - 1
	return hash & 0x7FFFFFFF
}
