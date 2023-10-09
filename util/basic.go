package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func DateValue(str string) time.Time {
	// layout := "2006-01-02T15:04:05.000Z"
	layout := "2006-01-02"
	// str := "2014-11-12T11:45:26.371Z"

	t, err := time.Parse(layout, str)

	if err != nil {
		return time.Time{}
	}
	return t
}

func ToUUID(str string) uuid.UUID {
	u, e := uuid.Parse(str)
	if e != nil {
		return uuid.Nil
	}
	return u
}

func ToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func ToInt32(s string) int32 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return int32(i)
}

func Int64List2String(ls []int64) string {
	var s string = "("
	//fmt.Println(data)
	for index, j := range ls {
		if index == 0 { //If the value is first one
			s = s + fmt.Sprintf("%v", j)
		} else if len(ls) == index+1 { // If the value is the last one
			s = s + fmt.Sprintf(",%v", j)
		} else {
			s = s + fmt.Sprintf(",%v", j)
		}
	}
	s = s + ")"
	return s
}

func Int64List2Comma(list []int64) string {
	// var result string
	result := []byte(strconv.FormatInt(list[0], 10))
	for n := 1; n < len(list); n++ {
		result = append(result, 44)
		result = append(result, []byte(strconv.FormatInt(list[n], 10))...)
	}
	return string(result)
}

func String2SqlList(list []string) string {
	// var result string
	result := "("
	for i, str := range list {
		// Add the current string to the result, surrounded by single quotes
		if i == 0 {
			result += "'" + str + "'"
		} else {
			result += ",'" + str + "'"
		}
	}
	result += ")"
	return result
}

func Concatinate(s1 string, s2 string) string {
	result := []byte(s1)
	result = append(result, []byte(s2)...)
	return string(result)
}

func RoundTo(value float64, digit float64) float64 {
	num := math.Pow(10, digit)
	rounded := math.Round(value*num) / num
	return rounded
}

// imageToBytes converts an image to a byte slice
func ImageToBytes(img image.Image) []byte {
	var byteData []byte

	// Create an in-memory buffer to write the image data
	buf := new(bytes.Buffer)

	// Encode the image to the buffer as PNG or JPEG
	err := jpeg.Encode(buf, img, nil) // Replace with png.Encode for PNG format
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the byte slice from the buffer
	byteData = buf.Bytes()

	return byteData
}

func Struct2Json(s interface{}) string {
	// Convert struct to JSON
	jsonData, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonData)
	// // Print the JSON data
	// fmt.Printf(string(jsonData))
}
