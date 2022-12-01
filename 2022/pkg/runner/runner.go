package runner

import (
	"bytes"
	"errors"
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
	if partSelect >= len(d.Runnables) {
		return errors.New("cannot select this part, check if it is registered yet")
	}
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
