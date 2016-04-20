package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Fle makes "fle" go on the interwebz
type Fle struct {
	Data string `json:"data"`
}

func getfles() []func(string) string {
	return []func(string) string{
		func(d string) string { return color.New(color.FgYellow).SprintFunc()(d) },
		func(d string) string { return color.New(color.FgBlue).SprintFunc()(d) },
		func(d string) string { return color.New(color.FgRed).SprintFunc()(d) },
		func(d string) string { return color.New(color.FgYellow).SprintFunc()(d) },
		func(d string) string { return color.New(color.FgMagenta).SprintFunc()(d) },
		func(d string) string { return color.New(color.FgCyan).SprintFunc()(d) },
		func(d string) string { return color.New(color.FgWhite).SprintFunc()(d) },
	}
}

func splitargs() string {
	args := os.Args[1:]
	s := ""

	for i := 0; i < len(args); i++ {
		s += args[i] + " "
	}

	return s
}

func stringtoslice(s string) []string {
	return strings.Split(s, "")
}

func main() {

	// Concurrency is not parallelism.
	go func(d string) {
		time.Sleep(50 * time.Millisecond)

		data := url.Values{}
		data.Set("data", d)

		resp, err := http.PostForm("http://localhost:8080/fle", data)
		if err != nil {
			log.Fatal(err)
		}

		fle := Fle{}
		err = json.NewDecoder(resp.Body).Decode(&fle)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(fle.Data)

		os.Exit(0)
	}(splitargs())

	http.HandleFunc("/fle", doFle)
	http.ListenAndServe(":8080", nil)
}

func doFle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	flist := getfles()
	work := r.PostForm.Get("data")

	workS := stringtoslice(work)
	result := ""

	for i := 0; i < len(workS); i++ {
		result += flist[rand.Intn(len(flist))](workS[i])
	}

	ret := Fle{
		result,
	}

	json.NewEncoder(w).Encode(ret)

}
