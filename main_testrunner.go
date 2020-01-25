package main

/*
Quick test runner for go playgrounds
*/

import (
	"fmt"
	"os"
	"testing"
)

func main() {
	fmt.Println("Launching Testsuite")
	testSuite := []testing.InternalTest{
		{Name: "TestCaseA", F: TestCaseA},
	}
	matchString := func(a, b string) (bool, error) { return a == b, nil }
	os.Args = append(os.Args, "-test.v") // Show the logs even if the test is successfull
	testing.Main(matchString, testSuite, nil, nil)
}

func TestCaseA(t *testing.T) {
	t.Log("Hello TestCaseA")
	t.Error("Oh no")
}
