package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */


   func timeConversion(s string) string {
    parts := strings.Split(s, ":")
    hour, minute, second := parts[0], parts[1], parts[2][:2]
    ampm := parts[2][2:]

    if ampm == "PM" {
        if hour != "12" {
            hourInt, _ := strconv.Atoi(hour)
            hour = strconv.Itoa(hourInt + 12)
        }
    } else if ampm == "AM" {
        if hour == "12" {
            hour = "00"
        }
    }

    return hour + ":" + minute + ":" + second
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
