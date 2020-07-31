package outcap

import (
	"fmt"
	"os"
	"testing"
)

func TestStart(t *testing.T) {
	c, err := NewContainer()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("test")
	fmt.Println("test2")
	fmt.Fprintln(os.Stderr, "stderr error")

	c.Stop()

	fmt.Println(c.Data)

	if len(c.Data) != 3 {
		t.Error("Data length should be 3")
	}
	if c.Data[0] != "test" {
		t.Errorf("First line should be 'test', instead of %s", c.Data[0])
	}
	if c.Data[1] != "test2" {
		t.Errorf("Second line should be 'test2', instead of %s", c.Data[1])
	}
	if c.Data[2] != "stderr error" {
		t.Errorf("Third line should be 'stderr error', instead of %s", c.Data[2])
	}
}
