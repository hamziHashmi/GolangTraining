package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

// exitOnError prints any errors and exits.
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// decimal converts a d/m/s coordinate to decimal degrees.
// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

type location struct {
	lat, long float64
}

type world struct {
	radius float64
}

// rad converts degrees to radians.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// distance calculation using the Spherical Law of Cosines.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

type report struct {
	sol         int
	temperature temperature
	location    location
}
type temperature struct {
	high, low celsius
}
type location2 struct {
	lat, long float64
}
type celsius float64

type report2 struct {
	sol int
	temperature
	location
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func stardate(t time.Time) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

// location with a latitude, longitude in decimal degrees.
type locationtwo struct {
	lat, long float64
}

type honeyBee struct {
	name string
}

func (hb honeyBee) String() string {
	return hb.name
}
func (hb honeyBee) move() string {
	switch rand.Intn(2) {
	case 0:
		return "buzzes about"
	default:
		return "flies to infinity and beyond"
	}
}
func (hb honeyBee) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "pollen"
	default:
		return "nectar"
	}
}

type gopher struct {
	name string
}

func (g gopher) String() string {
	return g.name
}
func (g gopher) move() string {
	switch rand.Intn(2) {
	case 0:
		return "scurries along the ground"
	default:
		return "burrows in the sand"
	}
}
func (g gopher) eat() string {
	switch rand.Intn(5) {
	case 0:
		return "carrot"
	case 1:
		return "lettuce"
	case 2:
		return "radish"
	case 3:
		return "corn"
	default:
		return "root"
	}
}

type animal interface {
	move() string
	eat() string
}

func step(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v %v.\n", a, a.move())
	default:
		fmt.Printf("%v eats the %v.\n", a, a.eat())
	}
}

const sunrise, sunset = 8, 18

// String formats a location with latitude, longitude.
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

type lasertwo int

func (l *laser) talktwo() string {
	return strings.Repeat("pew ", int(*l))
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type persontwo struct {
	age int
}

func (p *persontwo) birthday() {
	p.age++
}
func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return s[i] < s[j] }
	}
	sort.Slice(s, less)
}

type number struct {
	value int
	valid bool
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}
func (n number) String() string {
	if !n.valid {
		return "not set"
	}
	return fmt.Sprintf("%d", n.value)
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		f.Close()
		return err
	}
	_, err = fmt.Fprintln(f, "Don’t just check errors, handle them gracefully.")
	f.Close()
	return err
}

const (
	rows, columns = 9, 9
	empty         = 0
)

// Cell is a square on the Sudoku grid.
type Cell struct {
	digit int8
	fixed bool
}

// Grid is a Sudoku grid.
type Grid [rows][columns]Cell

// Errors that could occur.
var (
	ErrBounds     = errors.New("out of bounds")
	ErrDigit      = errors.New("invalid digit")
	ErrInRow      = errors.New("digit already present in this row")
	ErrInColumn   = errors.New("digit already present in this column")
	ErrInRegion   = errors.New("digit already present in this region")
	ErrFixedDigit = errors.New("initial digits cannot be overwritten")
)

// NewSudoku makes a new Sudoku grid.
func NewSudoku(digits [rows][columns]int8) *Grid {
	var grid Grid
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			d := digits[r][c]
			if d != empty {
				grid[r][c].digit = d
				grid[r][c].fixed = true
			}
		}
	}
	return &grid
}

// Set a digit on a Sudoku grid.
func (g *Grid) Set(row, column int, digit int8) error {
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case !validDigit(digit):
		return ErrDigit
	case g.isFixed(row, column):
		return ErrFixedDigit
	case g.inRow(row, digit):
		return ErrInRow
	case g.inColumn(column, digit):
		return ErrInColumn
	case g.inRegion(row, column, digit):
		return ErrInRegion
	}
	g[row][column].digit = digit
	return nil
}

// Clear a cell from the Sudoku grid.
func (g *Grid) Clear(row, column int) error {
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case g.isFixed(row, column):
		return ErrFixedDigit
	}
	g[row][column].digit = empty
	return nil
}
func inBounds(row, column int) bool {
	if row < 0 || row >= rows || column < 0 || column >= columns {
		return false
	}
	return true
}
func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}
func (g *Grid) inRow(row int, digit int8) bool {
	for c := 0; c < columns; c++ {
		if g[row][c].digit == digit {
			return true
		}
	}
	return false
}
func (g *Grid) inColumn(column int, digit int8) bool {
	for r := 0; r < rows; r++ {
		if g[r][column].digit == digit {
			return true
		}
	}
	return false
}
func (g *Grid) inRegion(row, column int, digit int8) bool {
	startRow, startColumn := row/3*3, column/3*3
	for r := startRow; r < startRow+3; r++ {
		for c := startColumn; c < startColumn+3; c++ {
			if g[r][c].digit == digit {
				return true
			}
		}
	}
	return false
}
func (g *Grid) isFixed(row, column int) bool {
	return g[row][column].fixed
}
func main() {
	fmt.Println("Welcome to Week 1 of Golang Training")
	fmt.Println("===============================================")
	fmt.Println("Enter the Number of Question you want to implement!")
	fmt.Println("You have the following Chocies")
	fmt.Println("21.1 21.2 21.3 21.4 21.5 21.6 21.7 22.1 22.3")
	fmt.Println("= = = = = = = = = = = = = = = = = = = = = = = = = ")
	fmt.Println("23.1 23.2 24.1 24.2 24.3 25.0")
	fmt.Println("= = = = = = = = = = = = = = = = = = = = = = = = = ")
	fmt.Println("26.1 26.2 26.3 26.4 26.5 27.1 27.2 27.3 27.4")
	fmt.Println("= = = = = = = = = = = = = = = = = = = = = = = = = ")
	fmt.Println("27.5 27.6 27.7 28.1 28.2 29.0")
	var questionNo float64
	fmt.Scan(&questionNo)
	switch {
	case questionNo == 21.1:
		{
			var curiosity struct {
				lat  float64
				long float64
			}
			curiosity.lat = -4.5895
			curiosity.long = 137.4417
			fmt.Println(curiosity.lat, curiosity.long)
			fmt.Println(curiosity)
		}
	case questionNo == 21.2:
		{
			type location struct {
				lat  float64
				long float64
			}
			var spirit location
			spirit.lat = -14.5684
			spirit.long = 175.472636
			var opportunity location
			opportunity.lat = -1.9462
			opportunity.long = 354.4734
			fmt.Println(spirit, opportunity)
		}
	case questionNo == 21.3:
		{
			type location struct {
				lat, long float64
			}
			opportunity := location{lat: -1.9462, long: 354.4734}
			fmt.Println(opportunity)
			insight := location{lat: 4.5, long: 135.9}
			fmt.Println(insight)
		}
	case questionNo == 21.4:
		{
			type location struct {
				lat, long float64
			}
			bradbury := location{lat: -4.5895, long: 137.4417}
			curiosity := bradbury
			curiosity.long += 0.0106
			fmt.Println(bradbury, curiosity)
		}
	case questionNo == 21.5:
		{
			type location struct {
				name string
				lat  float64
				long float64
			}
			locations := []location{
				{name: "Bradbury Landing", lat: -4.5895, long: 137.4417},
				{name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
				{name: "Challenger Memorial Station", lat: -1.9462, long: 354.4734},
			}
			fmt.Println(locations)
		}
	case questionNo == 21.6:
		{
			type location struct {
				Lat, Long float64
			}
			curiosity := location{-4.5895, 137.4417}
			bytes, err := json.Marshal(curiosity)
			exitOnError(err)
			fmt.Println(string(bytes))
		}
	case questionNo == 21.7:
		{
			type location struct {
				Lat  float64 `json:"latitude"`
				Long float64 `json:"longitude"`
			}
			curiosity := location{-4.5895, 137.4417}
			bytes, err := json.Marshal(curiosity)
			exitOnError(err)
			fmt.Println(string(bytes))
		}
	case questionNo == 22.1:
		{
			// Bradbury Landing: 4º35'22.2" S, 137º26'30.1" E
			lat := coordinate{4, 35, 22.2, 'S'}
			long := coordinate{137, 26, 30.12, 'E'}
			fmt.Println(lat.decimal(), long.decimal())
		}
	case questionNo == 22.3:
		{
			var mars = world{radius: 3389.5}
			spirit := location{-14.5684, 175.472636}
			opportunity := location{-1.9462, 354.4734}
			dist := mars.distance(spirit, opportunity)
			fmt.Printf("%.2f km\n", dist)
		}
	case questionNo == 23.1:
		{

			bradbury := location2{-4.5895, 137.4417}
			t := temperature{high: -1.0, low: -78.0}
			report := report{sol: 15, temperature: t, location: location(bradbury)}
			fmt.Printf("%+v\n", report)
			fmt.Printf("a balmy %vº C\n", report.temperature.high)

		}

	case questionNo == 23.2:
		{
			report2 := report{
				sol:         15,
				location:    location{-4.5895, 137.4417},
				temperature: temperature{high: -1.0, low: -78.0},
			}
			fmt.Printf("average %vº C\n", report2.temperature)
		}
	case questionNo == 24.1:
		{
			var t interface {
				talk() string
			}
			t = martian{}
			fmt.Println(t.talk())
			t = laser(3)
			fmt.Println(t.talk())
		}
	case questionNo == 24.2:
		{
			day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
			fmt.Printf("%.1f Curiosity has landed\n", stardate(day))
		}
	case questionNo == 24.3:
		{
			curiosity := locationtwo{-4.5895, 137.4417}
			fmt.Println(curiosity)
		}
	case questionNo == 25.0:
		{
			rand.Seed(time.Now().UnixNano())
			animals := []animal{
				honeyBee{name: "Bzzz Lightyear"},
				gopher{name: "Go gopher"},
			}
			var sol, hour int
			for {
				fmt.Printf("%2d:00 ", hour)
				if hour < sunrise || hour >= sunset {
					fmt.Println("The animals are sleeping.")
				} else {
					i := rand.Intn(len(animals))
					step(animals[i])
				}
				time.Sleep(500 * time.Millisecond)
				hour++
				if hour >= 24 {
					hour = 0
					sol++
					if sol >= 3 {
						break
					}
				}
			}
		}
	case questionNo == 26.1:
		{
			answer := 42
			fmt.Println(&answer)
			address := &answer
			fmt.Println(*address)
		}
	case questionNo == 26.2:
		{
			var administrator *string
			scolese := "Christopher J. Scolese"
			administrator = &scolese
			fmt.Println(*administrator)
			bolden := "Charles F. Bolden"
			administrator = &bolden
			fmt.Println(*administrator)
		}
	case questionNo == 26.3:
		{
			rebecca := person{
				name:       "Rebecca",
				superpower: "imagination",
				age:        14,
			}
			birthday(&rebecca)
			fmt.Printf("%+v\n", rebecca)
		}
	case questionNo == 26.4:
		{
			planets := []string{
				"Mercury", "Venus", "Earth", "Mars",
				"Jupiter", "Saturn", "Uranus", "Neptune",
				"Pluto",
			}
			reclassify(&planets)
			fmt.Println(planets)
		}
	case questionNo == 26.5:
		{
			pew := laser(2)
			shout(&pew)
		}
	case questionNo == 27.1:
		{
			var nowhere *int
			if nowhere != nil {
				fmt.Println(*nowhere)
			}
		}
	case questionNo == 27.2:
		{
			var nobody *persontwo
			fmt.Println(nobody)
			nobody.birthday()
		}
	case questionNo == 27.3:
		{
			food := []string{"onion", "carrot", "celery"}
			sortStrings(food, nil)
			fmt.Println(food)
		}
	case questionNo == 27.4:
		{
			var soup []string
			fmt.Println(soup == nil)
			for _, ingredient := range soup {
				fmt.Println(ingredient)
			}
			fmt.Println(len(soup))
			soup = append(soup, "onion", "carrot", "celery")
			fmt.Println(soup)
		}
	case questionNo == 27.5:
		{
			var soup map[string]int
			fmt.Println(soup == nil)
			measurement, ok := soup["onion"]
			if ok {
				fmt.Println(measurement)
			}
		}
	case questionNo == 27.6:
		{
			var p *int
			v := p
			fmt.Printf("%T %v %v\n", v, v, v == nil)
		}
	case questionNo == 27.7:
		{
			n := newNumber(42)
			fmt.Println(n)
			e := number{}
			fmt.Println(e)
		}
	case questionNo == 28.1:
		{
			files, err := ioutil.ReadDir(".")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			for _, file := range files {
				fmt.Println(file.Name())
			}
		}
	case questionNo == 28.2:
		{
			err := proverbs("proverbs.txt")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case questionNo == 29.0:
		{
			s := NewSudoku([rows][columns]int8{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
			})
			err := s.Set(1, 1, 4)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			for _, row := range s {
				fmt.Println(row)
			}

		}

	default:
		fmt.Println("Error ! Enter the question Number from above given Value.")
	}

}
