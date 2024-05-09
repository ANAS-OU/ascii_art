package ascii_art

import (
	"log"
	"os"
	"strings"
	"testing"
)

const expectedFile = "txt/expected.txt"
const standardFile = "txt/standard.txt"
const artCharHeight = 8

type TestCase struct {
	input    string
	expected string
	gotten   string
}

func getExpectedOutput(test_ID int) string {
	content, err := os.ReadFile(expectedFile)
	if err != nil {
		log.Fatalf("Error! problem in reading %s.", expectedFile)
	}
	contentList := strings.Split(string(content), "\n")
	begin := test_ID * (artCharHeight + 1)
	end := begin + artCharHeight
	return strings.Join(contentList[begin:end], "\n") + "\n"
}

func TestLowercase(t *testing.T) {
	test := TestCase{
		input:    "hello",
		expected: getExpectedOutput(0),
	}
	test.gotten = ArgProcessor(test.input, standardFile)
	if test.expected != test.gotten {
		t.Fail()
	}
}

func TestUppercase(t *testing.T) {
	test := TestCase{
		input:    "HELLO",
		expected: getExpectedOutput(1),
	}
	test.gotten = ArgProcessor(test.input, standardFile)
	if test.expected != test.gotten {
		t.Fail()
	}
}

func TestMixcase(t *testing.T) {
	test := TestCase{
		input:    "HeLlo HuMaN",
		expected: getExpectedOutput(2),
	}
	test.gotten = ArgProcessor(test.input, standardFile)
	if test.expected != test.gotten {
		t.Fail()
	}
}

func TestDigits(t *testing.T) {
	test := TestCase{
		input:    "1Hello 2There",
		expected: getExpectedOutput(3),
	}
	test.gotten = ArgProcessor(test.input, standardFile)
	if test.expected != test.gotten {
		t.Fail()
	}
}

func TestSymbols_A(t *testing.T) {
	test := TestCase{
		input:    "{Hello & There #}",
		expected: getExpectedOutput(4),
	}
	test.gotten = ArgProcessor(test.input, standardFile)
	if test.expected != test.gotten {
		t.Fail()
	}
}

func TestSymbols_B(t *testing.T) {
	test := TestCase{
		input:    "1a\"#FdwHywR&/()=",
		expected: getExpectedOutput(5),
	}
	test.gotten = ArgProcessor(test.input, standardFile)
	if test.expected != test.gotten {
		t.Fail()
	}
}

func TestSymbols_C(t *testing.T) {
	test := TestCase{
		input:    "\\!\" #$%&'()*+,-./",
		expected: getExpectedOutput(6),
	}
	test.gotten = ArgProcessor(test.input, standardFile)
	if test.expected != test.gotten {
		t.Fail()
	}
}

func TestSymbols_D(t *testing.T) {
	test := TestCase{
		input:    ":;<=>?@",
		expected: getExpectedOutput(7),
	}
	test.gotten = ArgProcessor(test.input, standardFile)
	if test.expected != test.gotten {
		t.Fail()
	}
}

func TestAlphabetic(t *testing.T) {
	test := TestCase{
		input:    "abcdefghijklmnopqrstuvwxyz",
		expected: getExpectedOutput(8),
	}
	test.gotten = ArgProcessor(test.input, standardFile)
	if test.expected != test.gotten {
		t.Fail()
	}
}
