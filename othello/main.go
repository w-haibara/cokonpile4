package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type square struct {
	w, b bool
}

type table []square

var gameSet = fmt.Errorf("game set")

const size = 10

func main() {
	tmp := make([]square, size)
	t := table(tmp)
	r := bufio.NewReader(os.Stdin)

	for {
		for {
			i, err := read(os.Stdout, r, "white")
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			if err := t.putW(i); err != nil {
				fmt.Println("put white error:", err.Error())
				continue
			}

			t.print(os.Stdout)

			if err := t.check(os.Stdout, size); err != nil {
				if err == gameSet {
					os.Exit(0)
				}
				fmt.Println("check error:", err.Error())
				os.Exit(1)
			}

			break
		}

		for {
			i, err := read(os.Stdout, r, "black")
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			if err := t.putB(i); err != nil {
				fmt.Println("put black error:", err.Error())
				continue
			}

			t.print(os.Stdout)
			if err := t.check(os.Stdout, size); err != nil {
				if err == gameSet {
					os.Exit(0)
				}
				fmt.Println("check error:", err.Error())
				os.Exit(1)
			}

			break
		}
	}
}

func (t *table) check(w io.Writer, size int) error {
	wn := 0
	bn := 0
	for _, v := range *t {
		if v.w {
			wn++
		}
		if v.b {
			bn++
		}
	}
	if wn+bn == size {
		fmt.Fprintln(w, "=== game set ===")
		fmt.Fprintln(w, "white:", wn, ", black:", bn)
		switch {
		case wn == bn:
			fmt.Fprintln(w, "draw game")
		case wn > bn:
			fmt.Fprintln(w, "white player win!")
		case wn < bn:
			fmt.Fprintln(w, "black player win!")
		}
		return gameSet
	}

	return nil
}

func read(w io.Writer, r *bufio.Reader, player string) (int, error) {
	fmt.Fprint(w, player+" > ")
	cmd, err := r.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("read command error: %s", err.Error())
	}

	i, err := strconv.Atoi(strings.TrimSuffix(cmd, "\n"))
	if err != nil {
		return 0, fmt.Errorf("atoi error: %s", err.Error())
	}

	return i, nil
}

func (t *table) putW(i int) error {
	if i < 0 || i >= len(*t) {
		return fmt.Errorf("index %d is out of table", i)
	}

	if (*t)[i].w || (*t)[i].b {
		return fmt.Errorf("index %d is already putted", i)
	}

	(*t)[i].w = true

	t.turnOver()

	return nil
}

func (t *table) putB(i int) error {
	if i < 0 || i >= len(*t) {
		return fmt.Errorf("index %d is out of table", i)
	}

	if (*t)[i].w || (*t)[i].b {
		return fmt.Errorf("index %d is already putted", i)
	}

	(*t)[i].b = true

	t.turnOver()

	return nil
}

func (t *table) turnOver() {
	for i := 0; i < len(*t)-2; i++ {
		if !(*t)[i].w && !(*t)[i].b {
			continue
		}

		if (*t)[i].w {
			j := i
			n := 0
			for {
				if j+1 >= len(*t) {
					break
				}

				j++

				if !(*t)[j].b {
					break
				}

				n++
			}
			if n != 0 && (*t)[j].w {
				for k, _ := range (*t)[i+1 : j] {
					(*t)[i+1+k].w = true
					(*t)[i+1+k].b = false
				}
			}
		}

		if (*t)[i].b {
			j := i
			n := 0
			for {
				if j+1 >= len(*t) {
					break
				}

				j++

				if !(*t)[j].w {
					break
				}

				n++
			}
			if n != 0 && (*t)[j].b {
				for k, _ := range (*t)[i+1 : j] {
					(*t)[i+1+k].b = true
					(*t)[i+1+k].w = false
				}
			}
		}
	}
}

func (t *table) print(w io.Writer) {
	s := ""
	for _, v := range *t {
		s += "|"
		switch {
		case v.w:
			s += "o"
		case v.b:
			s += "x"
		default:
			s += " "
		}
	}
	s += "|"

	fmt.Println(s)

}
