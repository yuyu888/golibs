package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

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

func QuickPowModEuler(a uint64, b uint64, m uint64) uint64 {
	var e int = Euler(int(m))
	fmt.Println(e)

	ans := QuickPowMod(a/m, b-uint64(e), m)
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

// 根据素数定义判断是否是素数，最low的方法
func IsPrimeLow(n int) bool {
	if n < 2 {
		return false
	}
	max := int(math.Sqrt(float64(n)))
	//max := n / 2
	for i := 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Primenumber(num uint64, primes []uint64) []uint64 {
	if num < 2 {
		return nil
	}
	maxp := uint64(math.Sqrt(float64(num)))
	max_prime := primes[len(primes)-1]
	if max_prime < maxp {
		primes = Primenumber(maxp, primes)
	}
	primeList := primes
	for i := primes[len(primes)-1] + 1; i <= num; i++ {
		isPrime := true
		for _, n := range primes {
			if i%n == 0 {
				isPrime = false
				break
			}
		}
		if isPrime == true {
			primeList = append(primeList, i)
		}
	}
	return primeList
}

func MonteCarloPi(num int) float64 {
	r := 10000
	m := 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= num; i++ {
		x := rand.Intn(r)
		y := rand.Intn(r)
		if (x*x + y*y) <= r*r {
			m++
		}
	}
	pi := float64(4*m) / float64(num)
	fmt.Println(pi)
	pi = math.Trunc(pi*1e2) * 1e-2
	return pi
}

func Baozi(costPrice int, price int, min int, max int, randCount int) int {
	rand.Seed(time.Now().UnixNano())
	profits := 0
	num := 0
	surplus := max - min
	for i := 0; i <= surplus; i++ {
		profitsN := 0
		for j := 0; j <= randCount; j++ {
			personNum := rand.Intn(surplus)
			if personNum >= i {
				profitsN += i * (price - costPrice)

			} else {
				profitsN += personNum*price - costPrice*i
			}
		}
		if profits < profitsN {
			profits = profitsN
			num = i
		}
	}
	return num + min
}

func GetGCDByDivide(a int, b int) int {
	if a%b == 0 {
		return b
	} else {
		return GetGCDByDivide(b, a%b)
	}
}

func GetGCD(a int, b int) int {
	if a == b {
		return b
	} else if a > b {
		return GetGCD(a-b, b)
	} else {
		return GetGCD(b-a, a)
	}
}

func GetLCM(a int, b int) int {
	gcd := GetGCD(a, b)
	return a * b / gcd
}

func Euler(n int) int {
	res := n
	i := 2
	for {
		if i*i > n {
			break
		}
		if n%i == 0 {
			n = n / i
			res = res - res/i
			for {
				if n%i != 0 {
					break
				}
				n = n / i
			}
		}
		i++
	}
	if n > 1 {
		res = res - res/n
	}
	return res
}
