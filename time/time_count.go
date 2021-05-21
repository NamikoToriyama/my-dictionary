package time

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func timeCount() {

	start1 := time.Now()
	if err := readBytes("words.txt"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	end1 := time.Now()

	start2 := time.Now()
	if err := readLine("words.txt"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	end2 := time.Now()

	start3 := time.Now()
	if err := readLineDelim("words.txt", '\n'); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	end3 := time.Now()

	start := time.Now()
	b, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := string(b)
	lines = ""
	fmt.Print(lines)
	end := time.Now()
	fmt.Printf("%f秒\n", (end1.Sub(start1)).Seconds())
	fmt.Printf("%f秒\n", (end2.Sub(start2)).Seconds())
	fmt.Printf("%f秒\n", (end3.Sub(start3)).Seconds())
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

}

func readBytes(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	b := make([]byte, 10)
	for {
		c, err := file.Read(b)
		if c == 0 {
			break
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		line := string(b[:c])
		line = ""
		fmt.Print(line)

	}
	return nil
}

func readLine(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = ""
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func readLineDelim(filename string, delim byte) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString(delim)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		line = ""
		fmt.Print(line)
	}
	return nil
}
