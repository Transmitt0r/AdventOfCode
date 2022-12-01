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
	Runnable func([]byte) (Solution, error)
}

func (d Runner) Run(input io.Reader) error {
	var data = bytes.NewBuffer(nil)
	_, err := io.Copy(data, input)
	if err != nil {
		return err
	}
	solution, err := d.Runnable(data.Bytes())
	if err != nil {
		return err
	}
	log.Println(solution.Message)
	return nil
}
