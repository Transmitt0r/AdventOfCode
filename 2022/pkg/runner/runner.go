package runner

import (
	"bytes"
	"io"
	"log"
)

type Solution struct {
	Message string
}

type Runner struct {
	Runnables []Runnable
}

type Runnable func([]byte) (Solution, error)

func (d Runner) Run(input io.Reader, partSelect int) error {
	var data = bytes.NewBuffer(nil)
	_, err := io.Copy(data, input)
	if err != nil {
		return err
	}
	solution, err := d.Runnables[partSelect](data.Bytes())
	if err != nil {
		return err
	}
	log.Println(solution.Message)
	return nil
}
