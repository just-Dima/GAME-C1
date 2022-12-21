package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var (
	// begin crossplatform support

	scanner       = bufio.NewScanner(os.Stdin)
	iointerface   = "std"
	input         = make(chan string)
	input_batch   = []string{}
	output_batch  = []string{}
	output        = func(s string) { fmt.Print(s) }
	rand_override = []float32{}
	exitfunc      = os.Exit

	// end crossplatform support
	n_str                  = ""
	m_arr                  = [15][50]int64{}
	a1, a2, a3, a4, a5, a6 int64
	c, d, d1, i, j, k      int64
	l, r1, s, s1, t        int64
)

func main() {
	bPRINT("                      OBSTACLE\n")
	bPRINT("                    CREATIVE COMPUTING\n")
	bPRINT("                  MORRISTOWN, NEW JERSEY\n")
	bPRINT("\n")
	bPRINT("\n")
	bPRINT("\n")
	for {
		n_str = bINPUTs("DO YOU WANT INSTRUCTIONS? ")
		if n_str == "YES" {
			bPRINT("THE OBJECT OF THIS GAME IS TO MOVE YOUR CAR'*'\n")
			bPRINT(" BEGINNING AT'S' AND NAVIGATE THROUGH THE OBSTACLES\n")
			bPRINT("'!'&'-' WALLS TO THE SPACE MARKED 'F', YOU MUST LAND\n")
			bPRINT("ON THE SPACE MARKED 'F' ON THE EXACT AMOUNT OF SPACES\n")
			bPRINT("THERE ARE NO DIAGONAL MOVES.\n")
			bPRINT("THERE ARE NO RIGHT TO LEFT MOVES.\n")
			bPRINT("DIRECTIONS NO.1 IS UP.\n")
			bPRINT("DIRECTION NO.2 IS LEFT TO RIGHT\n")
			bPRINT("DIRECTION NO.3 IS DOWN.\n")
			bPRINT("\n")
			bPRINT("SPEED IS THE NO. OF SPACES IN A GIVEN DIRECTION\n")
			break
		} else if n_str == "NO" {
			break
		} else {
			bPRINT("YES OR NO\n")
		}
	}
	a1 = bASC("*")
	a2 = bASC("!")
	a3 = bASC(" ")
	a4 = bASC("S")
	a5 = bASC("F")
	a6 = bASC("-")
l670:
	d1 = 0
	s1 = 0
	for i = 1; i <= 10; i++ {
		for j = 1; j <= 42; j++ {
			r1 = int64(bRND(1) * 1.2)
			if r1 == 0 {
				goto l760
			}
			m_arr[i-1][j-1] = a2
			goto l770
		l760:
			m_arr[i-1][j-1] = a3
		l770:
		}
	}
	m_arr[2-1][2-1] = a4
	m_arr[10-1][40-1] = a3
	m_arr[10-1][41-1] = a5
	m_arr[2-1][3-1] = a3
	m_arr[3-1][2-1] = a3
	for i = 1; i <= 10; i++ {
		m_arr[i-1][0] = a2
		m_arr[i-1][42-1] = a2
	}
	for j = 1; j <= 42; j++ {
		m_arr[0][j-1] = a6
		m_arr[11-1][j-1] = a6
	}
l870:
	k = 1
	l = j
	bPRINT("\n")
	for i = 1; i <= 11; i++ {
		for j = 1; j <= 42; j++ {
			bPRINT(string(rune(m_arr[i-1][j-1])))
		}
		bPRINT("\n")
	}
	i = k
	j = l
	if d1 > 0 {
		goto l1080
	}
l1000:
	bPRINT("OPTION :(A=CONTINUE,B=NEW COURSE,C=STOP)? ")
	n_str = bINPUTs("")
	if n_str == "A" {
		goto l1080
	}
	if n_str == "B" {
		goto l670
	}
	if n_str == "C" {
		exitfunc(0)
	}
	bPRINT("INVALID OPTION\n")
	goto l1000
l1080:
l1100:
	bPRINT("DIRECTION? ")
	d = int64(bINPUTi(""))
	d = bABS(d)
	if d < 1 {
		goto l1100
	}
	if d > 3 {
		goto l1100
	}
	bPRINT("SPEED? ")
	s = int64(bINPUTi(""))
	s = bABS(s)
	d1 = d1 + 1
	s1 = s1 + s
	if d1 != 1 {
		goto l1160
	}
	j = 2
	i = 2
l1160:
	if d == 1 {
		goto l1190
	}
	if d == 2 {
		goto l1230
	}
	if d == 3 {
		goto l1270
	}
l1190:
	m_arr[i-1][j-1] = a3
	for c = 1; c <= s; c++ {
		i = bABS(i - 1)
		if m_arr[i-1][j-1] == a2 {
			goto l1340
		}
		if m_arr[i-1][j-1] == a6 {
			goto l1340
		}
	}
	m_arr[i-1][j-1] = a1
	goto l1310
l1230:
	m_arr[i-1][j-1] = a3
	for c = 1; c <= s; c++ {
		j = j + 1
		if m_arr[i-1][j-1] == a2 {
			goto l1340
		}
		if m_arr[i-1][j-1] == a6 {
			goto l1340
		}
	}
	m_arr[i-1][j-1] = a1
	goto l1310
l1270:
	m_arr[i-1][j-1] = a3
	for c = 1; c <= s; c++ {
		i = i + 1
		if m_arr[i-1][j-1] == a2 {
			goto l1340
		}
		if m_arr[i-1][j-1] == a6 {
			goto l1340
		}
	}
	m_arr[i-1][j-1] = a1
l1310:
	if m_arr[10-1][41-1] != a1 {
		goto l870
	}
	goto l1360
l1340:
	bPRINT("ILLEGAL MOVE...... YOU LOSE!!\n")
	goto l1400
l1360:
	t = (d1 / s1) * 100
	bPRINT("YOU WON!! AND YOUR TIME IS " + fmt.Sprint(t) + "\n")
l1400:
	bPRINT("\n")
	bPRINT("DO YOU WISH TO PLAY AGAIN? ")
	n_str = bINPUTs("")
	if n_str == "YES" {
		goto l670
	}
	exitfunc(0)
}

// begin crossplatform support

var input_batch_i = 0

func bINPUTs(q string) string {
	bPRINT(q)
	if iointerface == "std" {
		if !scanner.Scan() {
			if scanner.Err() != nil {
				exitfunc(0)
			}
		}
		return scanner.Text()
	}
	if iointerface == "js" {
		bPRINT("\n")
		return <-input
	}
	if iointerface == "tests" {
		if len(input_batch) > 0 {
			if len(input_batch) > input_batch_i {
				out := input_batch[input_batch_i]
				input_batch_i++
				return out
			} else {
				input_batch_i = 0
				return input_batch[input_batch_i]
			}
		} else {
			bPRINT("No input!!!\n")
		}
	}
	return ""
}

func bINPUTi(q string) int {
	ret, err := strconv.Atoi(bINPUTs(q))
	if err != nil {
		return 0
	}
	return ret
}

func bPRINT(str string) {
	if iointerface != "tests" {
		output(str)
	} else {
		output_batch = append(output_batch, str)
	}
}

var rand_override_i = 0

func bRND(i int) float32 {
	if len(rand_override) > 0 {
		if len(rand_override) > rand_override_i {
			defer func() { rand_override_i++ }()
			return rand_override[rand_override_i]
		} else {
			rand_override_i = 0
			return rand_override[rand_override_i]
		}
	} else {
		return rand.Float32()
	}
}

// end crossplatform support

func bSGNi(value int64) int64 {
	if value == 0 {
		return 0
	} else if value > 0 {
		return 1
	}
	return -1
}

func bASC(v string) int64 {
	if len(v) == 0 {
		return 0
	}
	return int64(v[0])
}

func bABS(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}
