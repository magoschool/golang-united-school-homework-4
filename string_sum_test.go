package string_sum

import "testing"

func getErrorMessage(aError error) string {
	if aError != nil {
		return aError.Error()
	}
	return ""
}

func testStringSum(t *testing.T, aInput string, aResult string, aError error) {
	testValue, testError := StringSum(aInput)

	if testError != aError {
		t.Errorf("Invalid ReverseString output: '%v' expected but '%v' found", getErrorMessage(aError), getErrorMessage(testError))
	}
	if testValue != aResult {
		t.Errorf("Invalid ReverseString output: '%v' expected but '%v' found", aResult, testValue)
	}
}

func TestAdd(t *testing.T) {
	testStringSum(t, "3+5", "8", nil)
}

func TestSub(t *testing.T) {
	testStringSum(t, "5-3", "2", nil)
}

func TestAddNeg(t *testing.T) {
	testStringSum(t, "5+-3", "2", nil)
}

func TestAddPos(t *testing.T) {
	testStringSum(t, "5++3", "8", nil)
}

func TestSubNeg(t *testing.T) {
	testStringSum(t, "5--3", "8", nil)
}

func TestSubPos(t *testing.T) {
	testStringSum(t, "5-+3", "2", nil)
}

func TestPosAdd(t *testing.T) {
	testStringSum(t, "+5+3", "8", nil)
}

func TestNegAdd(t *testing.T) {
	testStringSum(t, "-5+3", "-2", nil)
}

func TestArg1Error(t *testing.T) {
	testStringSum(t, "+-5+3", "", errorNotTwoOperands)
}

func TestOneArgError(t *testing.T) {
	testStringSum(t, "-5", "", errorNotTwoOperands)
}

func TestThreeArgError(t *testing.T) {
	testStringSum(t, "5+5+6", "", errorNotTwoOperands)
}

func TestThreeSignError(t *testing.T) {
	testStringSum(t, "5++-5", "", errorNotTwoOperands)
}

func TestIgnoreSpace(t *testing.T) {
	testStringSum(t, "   5   +  -  2  ", "3", nil)
}
