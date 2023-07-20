package integer

var (
	Counter = 0
)

func SetCounter(c int) {
	if c > Counter {
		Counter = c
	}
}

func GetCounter() int {
	Counter = Counter + 1
	return Counter
}
