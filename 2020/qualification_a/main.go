package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

// INPUT TEMPLATE START

type MyInput struct {
	rdr         io.Reader
	lineChan    chan string
	initialized bool
}

func (mi *MyInput) start(done chan struct{}) {
	r := bufio.NewReader(mi.rdr)
	defer func() { close(mi.lineChan) }()
	for {
		line, err := r.ReadString('\n')
		if !mi.initialized {
			mi.initialized = true
			done <- struct{}{}
		}
		mi.lineChan <- strings.TrimSpace(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}
}

func (mi *MyInput) readLine() string {
	// if this is the first call, initialize
	if !mi.initialized {
		mi.lineChan = make(chan string)
		done := make(chan struct{})
		go mi.start(done)
		<-done
	}

	res, ok := <-mi.lineChan
	if !ok {
		panic("trying to read from a closed channel")
	}
	return res
}

func (mi *MyInput) readInt() int {
	line := mi.readLine()
	i, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return i
}

func (mi *MyInput) readInt64() int64 {
	line := mi.readLine()
	i, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (mi *MyInput) readInts() []int {
	line := mi.readLine()
	parts := strings.Split(line, " ")
	res := []int{}
	for _, s := range parts {
		tmp, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res = append(res, tmp)
	}
	return res
}

func (mi *MyInput) readInt64s() []int64 {
	line := mi.readLine()
	parts := strings.Split(line, " ")
	res := []int64{}
	for _, s := range parts {
		tmp, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		res = append(res, tmp)
	}
	return res
}

func (mi *MyInput) readWords() []string {
	line := mi.readLine()
	return strings.Split(line, " ")
}

// INPUT TEMPLATE END

////////////////////////////////////////////////////////////////////////////////

func main() {
	f, _ := os.Open("sample.in")
	in := MyInput{rdr: f}
	t := in.readInt()
	for caseNo := 1; caseNo <= t; caseNo++ {
		fmt.Printf("Case #%d: ", caseNo)
		testCase(in)
	}
}

// M is a dict of key int
type M map[int]bool

func testCase(in MyInput) {
	n := in.readInt()
	grid := [][]int{}

	for i := 0; i < n; i++ {
		r := in.readInts()
		grid = append(grid, r)
	}

	rows := make([]M, n)
	cols := make([]M, n)
	for i := 0; i < n; i++ {
		rows[i] = make(M)
		cols[i] = make(M)
	}

	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rows[i][grid[i][j]] = true
			cols[j][grid[i][j]] = true

			if i == j {
				sum += grid[i][j]
			}
		}
	}

	badRows := 0
	badCols := 0
	for i := 0; i < n; i++ {
		if len(rows[i]) != n {
			badRows++
		}
		if len(cols[i]) != n {
			badCols++
		}
	}

	fmt.Printf("%d %d %d\n", sum, badRows, badCols)
}
