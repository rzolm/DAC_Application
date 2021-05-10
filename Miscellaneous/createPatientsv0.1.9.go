package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var sex string
var fname string
var lname string
var address string
var patients [100]string
var position int
var cityPosition int

// type patients struct {
//   status int
//   nhs int
//   gender string
//   firsName string
//   lastName string
//   dob date
//   nhs int
//   address string
//   city string
//   county string
//   postcode string
//   tel01 int
//   tel02 int
//   email string
// }

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Create("patients.txt")

	var telCode string = "01"
	var mobileCode string = "07"

	for i := 0; i <= 100; i++ {
		id := strconv.Itoa(i)
		status := strconv.Itoa(1)
		nhs := createNHS()
		fname := createFirstName()
		lname := createLastName()
		//fullName := (fname + " " + lname)
		//patients[i] = (fullName)
		dob := createDOB()
		address := createAddress()
		city := createCity()
		county := createCounty(position)
		postcode := createPostcode()
		homeTel := createTel()
		mobileTel := createMobile()

		homeTelephone := telCode + homeTel
		mobileTelephone := mobileCode + mobileTel

		emailProvider := createEmail()
		email := fname + "_" + lname + emailProvider

		//write the name to the file
		f1, err := f.WriteString(id + "," + status + "," + nhs + "," + "M" + "," + fname + "," + lname + "," + dob + "," + address + "," + city + "," + county + "," + postcode + "," + homeTelephone + "," + mobileTelephone + "," + email + "\n")
		if err != nil {
			log.Fatal(err)
			fmt.Println(f1)
		}
	}
	check(err)
	defer f.Close()
}

func createNHS() string {
	r1 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(r1)

	//generate 10 digit number - faster when stored in a DB
	//converting to a string as varchar would be slower
	x := strconv.Itoa(r2.Intn(10000000000))

	return x
}

func createFirstName() string {
	file, err := os.Open("male_first_names.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println("unable to read file")
	}

	scanner := bufio.NewScanner(file)

	//generate a random number between 0 - 30 which the for loop stops at
	//this will be the name to use
	r1 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(r1)

	x := r2.Intn(30)

	for i := 0; i <= x; i++ {
		scanner.Scan()
		if i == x {
			fname = scanner.Text()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return fname
}

func createLastName() string {
	file, err := os.Open("last_names.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println("unable to read file")
	}
	//defer file.Close()

	scanner := bufio.NewScanner(file)

	//generate a random number between 0 - 30 which the for loop stops at
	//this will be the name to use
	r1 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(r1)

	x := r2.Intn(40)

	for i := 0; i <= x; i++ {
		scanner.Scan()
		if i == x {
			lname = scanner.Text()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lname
}

func createDOB() string {
	var day int
	var month int
	var year int
	t := time.Now()
	currentYear := t.Year()
	//var dob string

	d1 := rand.NewSource(time.Now().UnixNano())
	m1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.NewSource(time.Now().UnixNano())
	d2 := rand.New(d1)
	m2 := rand.New(m1)
	y2 := rand.New(y1)

	day = d2.Intn(27) + 1
	month = m2.Intn(11) + 1
	age := y2.Intn(90) + 1
	year = currentYear - age

	y := strconv.Itoa(year)
	m := strconv.Itoa(month)
	d := strconv.Itoa(day)

	bd := (y + "-" + m + "-" + d)
	return bd
}

func createAddress() string {

	var patientAddress string

	n1 := rand.NewSource(time.Now().UnixNano())
	n2 := rand.New(n1)

	houseNumberInt := n2.Intn(150)
	houseNumberStr := strconv.Itoa(houseNumberInt)

	//open list of addresses
	file, err := os.Open("address.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println("unable to read file")
	}

	scanner := bufio.NewScanner(file)

	a1 := rand.NewSource(time.Now().UnixNano())
	a2 := rand.New(a1)

	x := a2.Intn(30)

	for i := 0; i <= 30; i++ {
		scanner.Scan()
		if i == x {
			address = scanner.Text()
		}
	}

	patientAddress = houseNumberStr + " " + address
	//fmt.Println("address:",patientAddress)

	return patientAddress
}

func createCity() string {
	var city string

	//open list of cities
	file, err := os.Open("cities.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println("unable to read file")
	}

	scanner := bufio.NewScanner(file)

	c1 := rand.NewSource(time.Now().UnixNano())
	c2 := rand.New(c1)

	position = c2.Intn(20)

	for i := 0; i <= 20; i++ {
		scanner.Scan()
		if i == position {
			city = scanner.Text()
		}
	}
	return city
}

func createCounty(position int) string {
	var county string

	//open list of counties
	file, err := os.Open("counties.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println("unable to read file")
	}

	scanner := bufio.NewScanner(file)

	//match line number to county
	for i := 0; i <= 20; i++ {
		scanner.Scan()
		if i == position {
			county = scanner.Text()
		}
	}

	return county
}

func createPostcode() string {
	var postcode string
	var cityCode string
	var letters = ("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var endCode strings.Builder
	length := 2

	//open list of postcodes
	file, err := os.Open("postcodes.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println("unable to read file")
	}

	scanner := bufio.NewScanner(file)
	for i := 0; i <= 20; i++ {
		scanner.Scan()
		if i == position {
			cityCode = scanner.Text()
		}
	}

	p1 := rand.NewSource(time.Now().UnixNano())
	p2 := rand.New(p1)

	//create random number then convert to a string
	p3 := strconv.Itoa(p2.Intn(40))

	p4 := rand.NewSource(time.Now().UnixNano())
	p5 := rand.New(p4)

	//create random number then convert to a string
	p6 := strconv.Itoa(p5.Intn(10))

	for i := 0; i < length; i++ {
		random := rand.Intn(len(letters))
		randomChar := letters[random]
		endCode.WriteString(string(randomChar))
	}

	postcode = cityCode + p3 + " " + p6 + endCode.String()

	return postcode
}

func createMobile() string {
	r1 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(r1)

	//generate 10 digit number - faster when stored in a DB
	//converting to a string as varchar would be slower
	x := strconv.Itoa(r2.Intn(100000000))

	return x
}

func createTel() string {
	r1 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(r1)

	//generate 10 digit number - faster when stored in a DB
	//converting to a string as varchar would be slower
	x := strconv.Itoa(r2.Intn(1000000000))

	return x
}

func createEmail() string {
	var provider string
	file, err := os.Open("email.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println("unable to read file")
	}

	scanner := bufio.NewScanner(file)

	e1 := rand.NewSource(time.Now().UnixNano())
	e2 := rand.New(e1)

	x := e2.Intn(8)
	for i := 0; i <= 8; i++ {
		scanner.Scan()
		if i == x {
			provider = scanner.Text()
		}
	}
	return provider
}
