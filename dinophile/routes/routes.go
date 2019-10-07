package routes

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/pkg/helpers"
)

// ROUTES

// Health is a route I always add to every API!
// Takes no params
// Simple route to determine if API is up and running, so no one can say it isn't! :D
func Health(w http.ResponseWriter, r *http.Request) {
	response := "You're good to...GO! :D\n"
	fmt.Fprint(w, response)
}

// Echo let's us read the contents of a csv file
func Echo(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response string
	for _, row := range records {
		response, err = fmt.Sprintf("%s%s\n", response, strings.Join(row, ",")), errors.New("Response is malformed, check input")
	}
	fmt.Fprint(w, response)
}

// Invert a given slice
// TODO: error handling (sense a common theme?)
func Invert(w http.ResponseWriter, r *http.Request) {
	// Open csv file and read contents, assign them to records variable
	records := helpers.ReadCSV(w, r)
	// invert rows and columns
	invertedSlice := helpers.InvertSlice(records)
	// convert inverted matrix to a string separated by 3 characters
	stringToPrint := helpers.FormatSlice(invertedSlice, 3)
	// create response variable to write to w
	response := fmt.Sprintf("%s", stringToPrint)
	fmt.Fprint(w, response)
}

// Flatten a given slice
// TODO: needs error handling to cover the case our csv not containing strings
func Flatten(w http.ResponseWriter, r *http.Request) {
	// Open csv file and read contents, assign them to records variable
	records := helpers.ReadCSV(w, r)
	// loop through records and join them into a flatten string with commas as separators between characters
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s,", response, strings.Join(row, ","))
	}
	// create response variable to write to w
	response = strings.TrimSuffix(response, ",") + "\n"
	fmt.Fprint(w, response)
}

// Sum the contents of a given slice
func Sum(w http.ResponseWriter, r *http.Request) {
	// Open csv file and read contents, assign them to records variable
	records := helpers.ReadCSV(w, r)
	// loop through the values in each row and convert them to an integer, then add them to the result variable
	result := 0
	for _, row := range records {
		for _, item := range row {
			num, err := strconv.Atoi(item)
			if err != nil {
				errors.Errorf("Unable to convert %s to an integer, check your input", item)
			}
			result += num
		}
	}
	// create response variable to write to w
	response := strconv.Itoa(result) + "\n"
	fmt.Fprint(w, response)
}

// Multiply will multiply the contents of a given slice one item at a time (ex 1 x 2 x 3 etc)
func Multiply(w http.ResponseWriter, r *http.Request) {
	// Open csv file and read contents, assign them to records variable
	records := helpers.ReadCSV(w, r)
	// similar to sum we loop through the values in each row and convert them to an integer, then multiply
	// each value by the value in the result variable
	var result int
	for i, row := range records {
		for j, item := range row {
			num, err := strconv.Atoi(item)
			if err != nil {
				errors.Errorf("Unable to convert %s to an integer, check your input", item)
			}
			if i == 0 && j == 0 {
				result = num
			} else {
				result *= num
			}
		}
	}
	// create response variable to write to w
	response := strconv.Itoa(result) + "\n"
	fmt.Fprint(w, response)
}
