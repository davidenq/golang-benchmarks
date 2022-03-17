# Benchmark Switch vs Maps in Golang

## Code

```go
package switch_vs_map

func withSwitch(text string) string {
	message := ""
	switch text {
	case " 1":
		message = "one"
	case " 2":
		message = "two"
	case " 3":
		message = "three"
	case " 4":
		message = "four"
	case " 5":
		message = "five"
	case " 6":
		message = "six"
	case " 7":
		message = "seven"
	case " 8":
		message = "eight"
	case " 9":
		message = "nine"
	case "10":
		message = "ten"
	case "11":
		message = "eleven"
	case "12":
		message = "twelve"
	case "13":
		message = "thirteen"
	case "14":
		message = "fourteen"
	case "15":
		message = "fifteen"
	case "16":
		message = "sixteen"
	case "17":
		message = "seventeen"
	case "18":
		message = "eighteen"
	case "19":
		message = "nineteen"
	case "20":
		message = "twenty"
	}
	return message
}

func withMap(text string) string {
	message := map[string]string{
		" 1": "one",
		" 2": "two",
		" 3": "three",
		" 4": "four",
		" 5": "five",
		" 6": "six",
		" 7": "seven",
		" 8": "eight",
		" 9": "nine",
		"10": "ten",
		"11": "eleven",
		"12": "twelve",
		"13": "thirteen",
		"14": "fourteen",
		"15": "fifteen",
		"16": "sixteen",
		"17": "seventeen",
		"18": "eighteen",
		"19": "nineteen",
		"20": "twenty",
	}
	return message[text]
}

```

## How to test

if you want to see a flame graph you need
to download the dependencies:

```bash
 $ go get ./...
```

```bash
 $ make switchmap       #show results in console
 $ make graph_switchmap #show results in console and open a web browser to show outcomes graphically
```

## Tests

- round 1

outcom

```bash
goos: linux
goarch: amd64
pkg: github.com/davidenq/golang-benchmarks
cpu: Intel(R) Core(TM) i5-10210U CPU @ 1.60GHz
BenchmarkSwitch-8               615475560                 1.929 ns/op
BenchmarkSwitchMiss-8            76664785                 14.55 ns/op
BenchmarkMap-8                    1337336                843.2  ns/op
BenchmarkMapMiss-8                1370751                855.1  ns/op
PASS
ok      github.com/davidenq/golang-benchmarks   6.618s
```

- round 2

```bash
goos: linux
goarch: amd64
pkg: github.com/davidenq/golang-benchmarks
cpu: Intel(R) Core(TM) i5-10210U CPU @ 1.60GHz
BenchmarkSwitch-8               604738747                1.925 ns/op
BenchmarkSwitchMiss-8            77224310                14.50 ns/op
BenchmarkMap-8                    1247374               873.1  ns/op
BenchmarkMapMiss-8                1396983               855.9  ns/op
PASS
ok      github.com/davidenq/golang-benchmarks   6.635s
```

- round 3

```bash
goos: linux
goarch: amd64
pkg: github.com/davidenq/golang-benchmarks
cpu: Intel(R) Core(TM) i5-10210U CPU @ 1.60GHz
BenchmarkSwitch-8               624545629                 1.829 ns/op
BenchmarkSwitchMiss-8            88760383                12.34 ns/op
BenchmarkMap-8                    1376142               855.4  ns/op
BenchmarkMapMiss-8                1382082              1169    ns/op
PASS
ok      github.com/davidenq/golang-benchmarks   6.986s
```

### Final result

Although this is a basic use cases, so it looks like the switch wins.
