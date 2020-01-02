package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
)

type resultController struct {
	inputIDPattern *regexp.Regexp
}

type responsec struct {
	Input    int
	Response bool
}

// RegisterControllers ...
func RegisterControllers() {
	rc := newResultController()
	http.Handle("/double_sided_prime/", *rc)
}

func newResultController() *resultController {
	return &resultController{
		inputIDPattern: regexp.MustCompile(`^/double_sided_prime/(\d+)/?`),
	}
}

func (rc resultController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		matches := rc.inputIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		num, err := strconv.Atoi(matches[1])
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		boo, err := rc.isDoubleSidedPrime(num)
		if err != nil {
			fmt.Println(err)
			panic("Server error occured")
		}
		u := responsec{Input: num, Response: boo}
		fmt.Println(u)
		b, err := json.Marshal(u)
		if err != nil {
			fmt.Println(err)
			panic("Server error occured")
		}
		w.Write(b)
	}
}

func (rc resultController) isDoubleSidedPrime(num int) (bool, error) {
	u, err := rc.isPrime(num)
	if err != nil {
		panic("Server error occured ")
	}
	if u == false {
		return false, nil
	}

	// checking for left sided condition
	temp := num
	for temp != 0 {
		temp = temp / 10
		boo, err := rc.isPrime(temp)
		if err != nil {
			panic("Server error occured ")
		}
		if boo == false {
			break
		}
	}
	if temp == 0 {
		return true, nil
	}

	// checking for right sided condition
	temp = 10
	for num%temp != num {
		boo, err := rc.isPrime(num % temp)
		if err != nil {
			panic("Server error occured ")
		}
		if boo == false {
			return false, nil
		}
		temp = temp * 10
	}
	return true, nil
}

func (rc resultController) isPrime(num int) (bool, error) {
	if num == 0 || num == 1 {
		return false, nil
	}
	n := int(math.Floor(math.Sqrt(float64(num)))) + 1
	for i := 2; i <= n; i++ {
		if num%i == 0 {
			return false, nil
		}
	}
	return true, nil
}
