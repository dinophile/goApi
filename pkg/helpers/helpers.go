package helpers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// HELPER FUNCTIONS

// ReadCSV - this was common to all of my routes to I pulled it out into its own function. Habit, is this a good
// thing in go? Must investigate more!
// Opens a file and reads contents, then defer waits until the funciton this is called in returns and then
// closes the file
// Seems a little like generators? :thinkemoji:
// returns the contents of a read file
func ReadCSV(w http.ResponseWriter, r *http.Request) [][]string {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		os.Exit(1)
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		os.Exit(1)
	}
	return records
}

// InvertSlice transposes the columns and rows of a given slice
// This one was fun!
// TODO: errors! How to handle errors!
func InvertSlice(input [][]string) []string {
	// Loop through our slice and create a new one line slice by inserting the values into our desired
	// indexes
	// this should take a matrix of any size and transpose the columns and rows for us
	var result []string
	for i, row := range input {
		if i == 0 {
			result = append(result, row...)
		} else {
			var counter int
			for j, item := range row {
				multiplier := i + 1
				if j == 0 {
					result = InsertItem(result, i, item)
					counter = i
				} else {
					counter = counter + multiplier
					result = InsertItem(result, counter, item)
				}
			}
		}
	}
	return result
}

// InsertItem - inserts a new item at a chosen spot in a slice
// TODO: error handling
func InsertItem(slice []string, counter int, itemToInsert string) []string {
	// First create a new item at the end of the slice - value doesn't matter it will get overwritten anyway
	// this expands the capacity of our slice to accomodate the new item
	slice = append(slice, "0")
	// then we shift the contents up one index from the index where we want to insert our new item
	copy(slice[counter+1:], slice[counter:])
	// then overwrite the value in the index where we want to place our new item
	// the value here will have been copied and shifted up one index value
	slice[counter] = string(itemToInsert)
	return slice
}

// FormatSlice will take a slice and create a comman deliniated string of the contents
// and insert a new line every nth character
// TODO: error handling, does not handle the case when the number of lines to print n does not divide into our slice
func FormatSlice(slice []string, n int) string {
	linesToPrint := len(slice) / n
	i := 1
	result := ""
	for i <= linesToPrint {
		start := (i - 1) * n
		end := i * n
		if end > len(slice) {
			end = len(slice)
		}

		result = result + strings.Join(slice[start:end], ",") + "\n"
		i = i + 1
	}
	return result
}
