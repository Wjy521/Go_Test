package main

import "flag"
import "fmt"
import "bufio"
import "io"
import "os"
import "strconv"
import "time"

import "packs/algorithms/bubblesort"
import "packs/algorithms/qsort"

var infile *string = flag.String("i","unsorted.dat","File contains value for sorting")
var outfile *string = flag.String("o","sorted.dat","File to receive sorted values")
var algorithm *string = flag.String("a","qsort","Sort algorithm")


func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file",infile)
		return 
	}
	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for{
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		str := string(line)
		fmt.Println("this is :",str)

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("The infile=",*infile,"The outfile=",*outfile,"The algorithm=",*algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
			case "qsort":
				qsort.QuickSort(values)
			case "bubblesort" :
				bubblesort.BubbleSort(values)
			default:
				fmt.Println("Sorting algorithm",*algorithm,"is either unknowed or unsupported.")
		}
		t2 := time.Now()
		fmt.Println("The sorting process cost",t2.Sub(t1),"to complete.")
		writeValues(values, *outfile)
	}else {
		fmt.Println(err)
	}
}
