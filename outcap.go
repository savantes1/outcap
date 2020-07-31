package outcap

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// container is used to keep track of redirected stdout and stderr
// and hold output collected. Once the Stop() method is called,
// stdout and stderr are restored and collected output is available
// via the Data field.
// IMPORTANT: container is not reusable for collecting output after
// Stop() method is called. If you need to collect output after Stop()
// create new container via NewContainer() function.
type container struct {
	delimiters []rune

	backupStdout *os.File
	writerStdout *os.File
	backupStderr *os.File
	writerStderr *os.File

	data    string
	channel chan string

	Data []string
}

func NewContainer(delims ...rune) (*container, error) {

	rStdout, wStdout, err := os.Pipe()

	if err != nil {
		return nil, err
	}

	rStderr, wStderr, err := os.Pipe()

	if err != nil {
		return nil, err
	}

	c := &container{
		delimiters: delims,

		backupStdout: os.Stdout,
		writerStdout: wStdout,

		backupStderr: os.Stderr,
		writerStderr: wStderr,

		channel: make(chan string),
	}
	os.Stdout = c.writerStdout
	os.Stderr = c.writerStderr

	go func(out chan string, readerStdout *os.File, readerStderr *os.File) {
		var bufStdout bytes.Buffer

		// try to copy buffer from stdout to out channel
		// if it fails, print message to the stderr (not great solution, but can't think of better one)
		if _, err := io.Copy(&bufStdout, readerStdout); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if bufStdout.Len() > 0 {
			out <- bufStdout.String()
		}

		var bufStderr bytes.Buffer

		// try to copy buffer from stderr to out channel
		// ironically, if it fails, print message to stderr...
		if _, err := io.Copy(&bufStderr, readerStderr); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if bufStderr.Len() > 0 {
			out <- bufStderr.String()
		}
	}(c.channel, rStdout, rStderr)

	go func(c *container) {
		for {
			select {
			case out := <-c.channel:
				c.data += out
			}
		}
	}(c)

	return c, nil
}

// Stop() closes redirected stdout and stderr and restores them.
// Also formats collected output data in container.
func (c *container) Stop() {

	if c.writerStdout != nil {
		_ = c.writerStdout.Close()
	}

	if c.writerStderr != nil {
		_ = c.writerStderr.Close()
	}

	// Give it a sec to finish collecting data from buffers?
	time.Sleep(10 * time.Millisecond)

	os.Stdout = c.backupStdout
	os.Stderr = c.backupStderr

	// Separate captured output by delimeters
	c.Data = strings.FieldsFunc(c.data,
		func(r rune) bool {

			for _, elem := range c.delimiters {
				if r == elem {
					return true
				}
			}

			return false
		},
	)

	// // Remove empty items
	// for _, elem := range temp {
	// 	if elem != "" {
	// 		c.Data = append(c.Data, elem)
	// 	}
	// }

}