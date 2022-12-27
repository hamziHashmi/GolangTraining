package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
func realSensor() kelvin {
	return 0
}

type kelvintwo float64

func measureTemperature(samples int, sensor func() kelvintwo) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%vº K\n", k)
		time.Sleep(time.Second)
	}
}
func fakeSensortwo() kelvintwo {
	return kelvintwo(rand.Intn(151) + 150)
}

var f = func() {
	fmt.Println("Dress up for the masquerade.")
}

type kelvinthree float64

// sensor function type
type sensor func() kelvinthree

func realSensorthree() kelvinthree {
	return 0
}
func calibrate(s sensor, offset kelvinthree) sensor {
	return func() kelvinthree {
		return s() + offset
	}
}
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}
func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}
func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

const (
	width  = 80
	height = 15
)

// Universe is a two-dimensional field of cells.
type Universe [][]bool

// NewUniverse returns an empty universe.
func NewUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

// Seed random live cells into the universe.
func (u Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		u.Set(rand.Intn(width), rand.Intn(height), true)
	}
}

// Set the state of the specified cell.
func (u Universe) Set(x, y int, b bool) {
	u[y][x] = b
}

// Alive reports whether the specified cell is alive.
// If the coordinates are outside of the universe, they wrap around.
func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

// Neighbors counts the adjacent cells that are alive.
func (u Universe) Neighbors(x, y int) int {
	n := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x+h, y+v) {
				n++
			}
		}
	}
	return n
}

// Next returns the state of the specified cell at the next step.
func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	return n == 3 || n == 2 && u.Alive(x, y)
}

// String returns the universe as a string.
func (u Universe) String() string {
	var b byte
	buf := make([]byte, 0, (width+1)*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b = ' '
			if u[y][x] {
				b = '*'
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}

// Show clears the screen and displays the universe.
func (u Universe) Show() {
	fmt.Print("\x0c", u.String())
}

// Step updates the state of the next universe (b) from
// the current universe (a).
func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(x, y, a.Next(x, y))
		}
	}
}

type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

const (
	line         = "======================="
	rowFormat    = "| %8s | %8s |\n"
	numberFormat = "%.1f"
)

type getRowFn func(row int) (string, string)

// drawTable draws a two column table.
func drawTable(hdr1, hdr2 string, rows int, getRow getRowFn) {
	fmt.Println(line)
	fmt.Printf(rowFormat, hdr1, hdr2)
	fmt.Println(line)
	for row := 0; row < rows; row++ {
		cell1, cell2 := getRow(row)
		fmt.Printf(rowFormat, cell1, cell2)
	}
	fmt.Println(line)
}
func ctof(row int) (string, string) {
	c := celsius(row*5 - 40)
	f := c.fahrenheit()
	cell1 := fmt.Sprintf(numberFormat, c)
	cell2 := fmt.Sprintf(numberFormat, f)
	return cell1, cell2
}
func ftoc(row int) (string, string) {
	f := fahrenheit(row*5 - 40)
	c := f.celsius()
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}
func main() {

	fmt.Println("Welcome to Week 3 of Golang Training")
	fmt.Println("===============================================")
	fmt.Println("Enter the Number of Question you want to implement!")
	fmt.Println("You have the following Chocies")
	fmt.Println("14.1 14.2 14.3 14.4 15.0 16.1 16.4 16.5 16.6")
	fmt.Println("= = = = = = = = = = = = = = = = = = = = = =")
	fmt.Println("17.1 17.3 17.4 18.1 18.2 18.3 18.4 18.6")
	fmt.Println("= = = = = = = = = = = = = = = = = = = = = =")
	fmt.Println("19.1 19.2 19.4 19.6 20.0")
	var questionNo float64
	fmt.Scan(&questionNo)

	switch {
	case questionNo == 14.1:
		{
			sensor := fakeSensor
			fmt.Println(sensor())
			sensor = realSensor
			fmt.Println(sensor())
		}
	case questionNo == 14.2:
		{
			measureTemperature(3, fakeSensortwo)
		}
	case questionNo == 14.3:
		{
			f()
		}
	case questionNo == 14.4:
		{
			sensor := calibrate(realSensorthree, 5)
			fmt.Println(sensor())
		}
	case questionNo == 15.0:
		{
			drawTable("ºC", "ºF", 29, ctof)
			fmt.Println()
			drawTable("ºF", "ºC", 29, ftoc)
		}
	case questionNo == 16.1:
		{
			var planets [8]string
			planets[0] = "Mercury"
			planets[1] = "Venus"
			planets[2] = "Earth"
			earth := planets[2]
			fmt.Println(earth)
		}
	case questionNo == 16.3:
		{

		}
	case questionNo == 16.4:
		{
			dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
			for i := 0; i < len(dwarfs); i++ {
				dwarf := dwarfs[i]
				fmt.Println(i, dwarf)
			}
		}
	case questionNo == 16.5:
		{
			planets := [...]string{
				"Mercury",
				"Venus",
				"Earth",
				"Mars",
				"Jupiter",
				"Saturn",
				"Uranus",
				"Neptune",
			}
			planetsMarkII := planets
			planets[2] = "whoops"
			fmt.Println(planets)
			fmt.Println(planetsMarkII)
		}
	case questionNo == 16.6:
		{
			var board [8][8]string
			board[0][0] = "r"
			board[0][7] = "r"
			for column := range board[1] {
				board[1][column] = "p"
			}
			fmt.Print(board)
		}
	case questionNo == 17.1:
		{
			planets := [...]string{
				"Mercury",
				"Venus",
				"Earth",
				"Mars",
				"Jupiter",
				"Saturn",
				"Uranus",
				"Neptune",
			}
			terrestrial := planets[0:4]
			gasGiants := planets[4:6]
			iceGiants := planets[6:8]
			fmt.Println(terrestrial, gasGiants, iceGiants)
			fmt.Println(gasGiants[0])
		}
	case questionNo == 17.3:
		{
			planets := []string{" Venus ", "Earth ", " Mars"}
			hyperspace(planets)
			fmt.Println(strings.Join(planets, ""))
		}
	case questionNo == 17.4:
		{
			planets := []string{
				"Mercury", "Venus", "Earth", "Mars",
				"Jupiter", "Saturn", "Uranus", "Neptune",
			}
			sort.StringSlice(planets).Sort()
			fmt.Println(planets)
		}
	case questionNo == 18.1:
		{
			dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
			dwarfs = append(dwarfs, "Orcus")
			fmt.Println(dwarfs)
		}
	case questionNo == 18.2:
		{
			dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
			dump("dwarfs", dwarfs)
			dump("dwarfs[1:2]", dwarfs[1:2])
		}
	case questionNo == 18.3:
		{
			dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
			dwarfs2 := append(dwarfs1, "Orcus")
			dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")
			fmt.Println(dwarfs1)
			fmt.Println(dwarfs2)
			fmt.Println(dwarfs3)
		}
	case questionNo == 18.4:
		{
			planets := []string{
				"Mercury", "Venus", "Earth", "Mars",
				"Jupiter", "Saturn", "Uranus", "Neptune",
			}
			terrestrial := planets[0:4:4]
			worlds := append(terrestrial, "Ceres")
			fmt.Println(planets)
			fmt.Println(worlds)
		}
	case questionNo == 18.6:
		{
			twoWorlds := terraform("New", "Venus", "Mars")
			fmt.Println(twoWorlds)
			planets := []string{"Venus", "Mars", "Jupiter"}
			newPlanets := terraform("New", planets...)
			fmt.Println(newPlanets)
		}
	case questionNo == 19.1:
		{
			temperature := map[string]int{
				"Earth": 15,
				"Mars":  -65,
			}
			temp := temperature["Earth"]
			fmt.Printf("On average the Earth is %vº C.\n", temp)
			temperature["Earth"] = 16
			temperature["Venus"] = 464
			fmt.Println(temperature)
		}
	case questionNo == 19.2:
		{
			planets := map[string]string{
				"Earth": "Sector ZZ9",
				"Mars":  "Sector ZZ9",
			}
			planetsMarkII := planets
			planets["Earth"] = "whoops"
			fmt.Println(planets)
			fmt.Println(planetsMarkII)
			delete(planets, "Earth")
			fmt.Println(planetsMarkII)
		}
	case questionNo == 19.4:
		{
			temperatures := []float64{
				-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
			}
			frequency := make(map[float64]int)
			for _, t := range temperatures {
				frequency[t]++
			}
			for t, num := range frequency {
				fmt.Printf("%+.2f occurs %d times\n", t, num)
			}
		}
	case questionNo == 19.6:
		{
			var temperatures = []float64{
				-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
			}
			set := make(map[float64]bool)
			for _, t := range temperatures {
				set[t] = true
			}
			if set[-28.0] {
				fmt.Println("set member")
			}
			fmt.Println(set)
		}
	case questionNo == 20.0:
		{
			a, b := NewUniverse(), NewUniverse()
			a.Seed()
			for i := 0; i < 300; i++ {
				Step(a, b)
				a.Show()
				time.Sleep(time.Second / 30)
				a, b = b, a // Swap universes
			}
		}

	default:
		{
			fmt.Println("Error ! Enter the question Number from above given Value.")
		}
	}
}
