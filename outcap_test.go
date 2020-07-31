package outcap

import (
	"fmt"
	"os"
	"testing"
)

func TestStart(t *testing.T) {
	c, err := NewContainer('\n')

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("test")
	fmt.Println("test2")
	fmt.Fprintln(os.Stderr, "stderr error")

	c.Stop()

	fmt.Println(c.OutData)
	fmt.Println(c.ErrorData)

	if len(c.OutData) != 2 {
		t.Fatal("OutData length should be 2")
	}

	if len(c.ErrorData) != 1 {
		t.Fatal("ErrorData length should be 1")
	}

	if c.OutData[0] != "test" {
		t.Errorf("First line should be 'test', instead of %s", c.OutData[0])
	}
	if c.OutData[1] != "test2" {
		t.Errorf("Second line should be 'test2', instead of %s", c.OutData[1])
	}

	if c.ErrorData[0] != "stderr error" {
		t.Errorf("Error output should be 'stderr error', instead of %s", c.ErrorData[0])
	}
}
