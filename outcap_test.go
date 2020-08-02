package outcap

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestInput(t *testing.T) {

	c, err := NewContainer('\n')
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	
	// should always use a goroutine when passing redirectd input
	go func() {

		var testInt int
		fmt.Println("Type an int")
		fmt.Scanln(&testInt)

		fmt.Println("The int you entered was:", testInt)


		var testString string
		fmt.Println("Type a single-word string")
		fmt.Scanln(&testString)

		fmt.Println("The single-word string you entered was:", testString)

		cancel()
	}()

	c.WriteToStdin("42\n")
	c.WriteToStdin("Coolness\n")

	select {
	case <-ctx.Done():
		c.Stop()
	}

	if ctx.Err() == context.DeadlineExceeded {
		t.Fatal("Timed Out")
	} else {

		if len(c.OutData) != 4 {
			t.Fatal("Unexpected number of outputs")
		}

		if c.OutData[1] != "The int you entered was: 42" {
			t.Errorf("%q does not match \"The int you entered was: 42\"", c.OutData[1])
		}

		if c.OutData[3] != "The single-word string you entered was: Coolness" {
			t.Errorf("%q does not match \"The single-word string you entered was: 42\"", c.OutData[3])
		}

	}
}

func TestOutput(t *testing.T) {
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
