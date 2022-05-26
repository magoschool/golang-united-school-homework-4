package string_sum

import "testing"

const msgTwoOperandsError = "StringSum failed: expecting two operands, but received more or less"
const msgEmptyError = "StringSum failed: input is empty"

func getErrorMessage(aError error) string {
	if aError != nil {
		return aError.Error()
	}
	return ""
}

func testStringSum(t *testing.T, aInput string, aResult string, aError string) {
	testValue, testError := StringSum(aInput)

	testErrorText := getErrorMessage(testError)

	if testErrorText != aError {
		t.Errorf("Invalid ReverseString output: '%v' expected but '%v' found", aError, testErrorText)
	}
	if testValue != aResult {
		t.Errorf("Invalid ReverseString output: '%v' expected but '%v' found", aResult, testValue)
	}
}

func TestAdd(t *testing.T) {
	testStringSum(t, "3+5", "8", "")
}

func TestSub(t *testing.T) {
	testStringSum(t, "5-3", "2", "")
}

func TestAddNeg(t *testing.T) {
	testStringSum(t, "5+-3", "2", "")
}

func TestAddPos(t *testing.T) {
	testStringSum(t, "5++3", "8", "")
}

func TestSubNeg(t *testing.T) {
	testStringSum(t, "5--3", "8", "")
}

func TestSubPos(t *testing.T) {
	testStringSum(t, "5-+3", "2", "")
}

func TestPosAdd(t *testing.T) {
	testStringSum(t, "+5+3", "8", "")
}

func TestNegAdd(t *testing.T) {
	testStringSum(t, "-5+3", "-2", "")
}

func TestArg1Error(t *testing.T) {
	testStringSum(t, "+-5+3", "", msgTwoOperandsError)
}

func TestOneArgError(t *testing.T) {
	testStringSum(t, "-5", "", msgTwoOperandsError)
}

func TestThreeArgError(t *testing.T) {
	testStringSum(t, "5+5+6", "", msgTwoOperandsError)
}

func TestThreeSignError(t *testing.T) {
	testStringSum(t, "5++-5", "", msgTwoOperandsError)
}

func TestStringSum(t *testing.T) {
	testStringSum(t, "   5   +  -  2  ", "3", "")
}

func TestEmptySpace(t *testing.T) {
	testStringSum(t, "   5   +  -  2  ", "3", "")
}

func TestWiteSpace(t *testing.T) {
	testStringSum(t, "   ", "", msgEmptyError)
}

func TestEmpty(t *testing.T) {
	testStringSum(t, "", "", msgEmptyError)
}

func TestLetters(t *testing.T) {
	testStringSum(t, "a+b", "", "StringSum failed: strconv.Atoi: parsing \"a\": invalid syntax")
}
