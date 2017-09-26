package FFT

import (
	"bufio"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/mjibson/go-dsp/fft"
	"io"
	"math/big"
	"os"
)

// log is the default package logger
var log = logger.GetLogger("activity-tibco-rest")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// do eval

	//fmt.Println("start....")
	log.Debug("FFT Activity started.")
	inputFilePath := context.GetInput("inputFilePath").(string)

	sampleSize := context.GetInput("sampleSize").(int)

	outputFilePath := context.GetInput("outputFilePath").(string)

	series := readAndParseFile(inputFilePath, sampleSize)

	fft := runFFT(series)

	//fmt.Println( cmplx.Abs(fft[1]))

	writeCSVOutput(outputFilePath, fft)

	context.SetOutput("status", "success")

	//fmt.Println("Complete....")

	log.Debug("FFT Activity complete.")

	return true, nil
}

func readAndParseFile(filepath string, records int) []float64 {

	startPos := 37
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var values []float64

	b1 := make([]byte, 2)
	i := 1

	for i <= records+startPos {

		_, err := f.Read(b1)
		if err != nil {
			if err == io.EOF {
				fmt.Println("....EOF.....")
				break
			}
			panic(err)
		}

		ten := new(big.Int)
		ten.SetBytes(b1)

		res := float64(ten.Int64()) * 76.7 / 1000000

		if i > startPos {
			values = append(values, res)
			// fmt.Println(res)
		}
		i++

	}

	return values

}

func runFFT(series []float64) []complex128 {

	X := fft.FFTReal(series)

	return X

}

func writeCSVOutput(fileName string, values []complex128) {

	f, err := os.Create(fileName)
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)

	for _, value := range values {
		_, err = fmt.Fprintln(w, value)
		check(err)
	}

	w.Flush()

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
