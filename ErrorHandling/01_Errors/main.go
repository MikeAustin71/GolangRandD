package main

import (
	"fmt"
	"math"

	ePref "github.com/MikeAustin71/errpref/v2"
)

func main() {

	err := test004()

	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	fmt.Println("******************************")
	fmt.Println("Status Message from 'main()'")
	fmt.Printf("No Errors Detected\n")
	fmt.Printf("******************************\n")

}

func test001() error {

	// NewErrorGenericConfig()
	// Add identification information for this function
	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"main.go",
		"",
		"test001()")

	var value1 int64
	value1 = math.MaxInt64 // 'value1' guaranteed to fail

	var value2 int32

	value2 = 42

	err := (&testFuncs{}).inputParameterOutOfRange(value1, value2)

	if err != nil {

		err2 := errGenericCfg.GetErrorFromFunctionCall(
			"testFuncs",
			"inputParameterOutOfRange()",
			"",
			"",
			err)

		return err2
	}

	return nil
	/*  Output func test001()
	Source File: main.go
	This Function: test001()
	Called Function Type: testFuncs
	Called Function: inputParameterOutOfRange()
	:
	Source File: main.go
	This Function Type: testFuncs
	This Function: inputParameterOutOfRange
	Parameter: value1=9223372036854775807
	Error Context: 'value1' exceeds maximum value of type 'int32' (2147483647)
	Error Message: Error: 'value1' is out of range
	*/
}

func test002() error {

	// NewErrorGenericConfig()
	// Add identification information for this function
	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"main.go",
		"",
		"test002()")

	var value1 int64
	value1 = math.MinInt32

	var value2 int32

	value2 = -1 // 'value1' + 'value2' addition result guaranteed to fail

	err := (&testFuncs{}).inputParameterOutOfRange(value1, value2)

	if err != nil {

		err2 := errGenericCfg.GetErrorFromFunctionCall(
			"testFuncs",
			"inputParameterOutOfRange()",
			"",
			"",
			err)

		return err2
	}

	return nil

	/* Output: test002()
	Source File: main.go
	This Function: test002()
	Called Function Type: testFuncs
	Called Function: inputParameterOutOfRange()
	:
	Source File: main.go
	This Function Type: testFuncs
	This Function: inputParameterOutOfRange
	Called Function: addTwoInts()
	:
	Source File: main.go
	This Function Type: testFuncs
	This Function: addTwoInt32s()
	Error Context: Addition result ('-2147483649') less than minimum value of 'int32' ( -2147483648 )
	Error Message: Addition result out of range
	*/
}

func test003() error {

	// NewErrorGenericConfig()
	// Add identification information for this function
	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"main.go",
		"",
		"test003()")

	var value2 int32

	value2 = 42 // 'value1' + 'value2' addition result guaranteed to fail

	err := new(testFuncs).inputParameterNilPointer(nil, &value2)

	if err != nil {

		err2 := errGenericCfg.GetErrorFromFunctionCall(
			"testFuncs",
			"inputParameterNilPointer()",
			"",
			"",
			err)

		return err2
	}

	return nil
}

/* Output test003()
Source File: main.go
This Function: test003()
Called Function Type: testFuncs
Called Function: inputParameterNilPointer()
:
Source File: main.go
This Function Type: testFuncs
This Function: inputParameterNilPointer
Parameter: value1
Error Context: Invalid method input parameter
Error Message: Error: Input parameter 'value1' is a nil pointer
*/

func test004() error {

	// NewErrorGenericConfig()
	// Add identification information for this function
	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"main.go",
		"",
		"test004()")

	var value1 int64
	value1 = math.MinInt32

	var value2 int32

	value2 = -1 // 'value1' + 'value2' addition result guaranteed to fail

	err := (&testFuncs{}).inputParameterNilPointer(&value1, &value2)

	if err != nil {

		err2 := errGenericCfg.GetErrorFromFunctionCall(
			"testFuncs",
			"inputParameterNilPointer()",
			"",
			"",
			err)

		return err2
	}

	return nil

	/* Output: test004()
	Source File: main.go
	This Function: test004()
	Called Function Type: testFuncs
	Called Function: inputParameterNilPointer()
	:
	Source File: main.go
	This Function Type: testFuncs
	This Function: inputParameterNilPointer
	Called Function Type: testFuncs
	Called Function: inputParameterOutOfRange
	:
	Source File: main.go
	This Function Type: testFuncs
	This Function: inputParameterOutOfRange
	Called Function: addTwoInts()
	:
	Source File: main.go
	This Function Type: testFuncs
	This Function: addTwoInt32s()
	Error Context: Addition result ('-2147483649') less than minimum value of 'int32' ( -2147483648 )
	Error Message: Addition result out of range
	*/
}

type testFuncs struct{}

func (tf testFuncs) inputParameterOutOfRange(value1 int64, value2 int32) error {

	// NewErrorGenericConfig()
	// Add identification information for this function
	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"main.go",
		"testFuncs",
		"inputParameterOutOfRange")

	if value1 > math.MaxInt32 {
		//func (egc *ErrorGenericFactory) GetErrorInputParameterOutOfRange(
		//  parameter string,
		//  path string,
		//  fileName string,
		//  url string,
		//  errorContext string,
		//  errorMessage string) error
		err := errGenericCfg.GetErrorInputParameterOutOfRange(
			fmt.Sprintf("value1=%d", value1),
			"",
			"",
			"",
			fmt.Sprintf("'value1' exceeds maximum value of type 'int32' (%d)", math.MaxInt32),
			"Error: 'value1' is out of range")

		return err

	}

	convertedIntValue1 := int32(value1)

	intResult, err := addTwoInt32s(convertedIntValue1, value2)

	if err != nil {

		err2 := errGenericCfg.GetErrorFromFunctionCall(
			"",
			"addTwoInts()",
			"",
			"",
			err)

		return err2
	}

	fmt.Printf("\nintResult=%d\n", intResult)

	return nil
}

func (tf testFuncs) inputParameterNilPointer(value1 *int64, value2 *int32) error {

	// NewErrorGenericConfig()
	// Add identification information for this function
	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"main.go",
		"testFuncs",
		"inputParameterNilPointer")

	if value1 == nil {

		return errGenericCfg.GetErrorInputParameterNilPointer(
			"value1",
			"",
			"",
			"",
			"Invalid method input parameter",
			"Error: Input parameter 'value1' is a nil pointer")
	}

	if value2 == nil {

		return errGenericCfg.GetErrorInputParameterNilPointer(
			"value2",
			"",
			"",
			"",
			"Invalid method input parameter",
			"Error: Input parameter 'value2' is nil")

	}

	err := tf.inputParameterOutOfRange(*value1, *value2)

	if err != nil {

		return errGenericCfg.GetErrorFromFunctionCall(
			"testFuncs",
			"inputParameterOutOfRange",
			"",
			"",
			err)
	}

	return nil
}

func addTwoInt32s(value1 int32, value2 int32) (int32, error) {

	// NewErrorGenericConfig()
	// Add identification information for this function
	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"main.go",
		"testFuncs",
		"addTwoInt32s()")

	int64Value1 := int64(value1)

	int64Value2 := int64(value2)

	int64Result := int64Value1 + int64Value2

	if int64Result > math.MaxInt32 {

		//func (egc *ErrorGenericFactory) GetErrorNativeToFunction(
		//	parameter string,
		//	path string,
		//	fileName string,
		//	url string,
		//	errorContext string,
		//	errorMessage string)
		return 0, errGenericCfg.GetErrorNativeToFunction(
			"",
			"",
			"",
			"",
			fmt.Sprintf("Addition result ('%d') exceeds maximum value of 'int32' ( %d )",
				int64Result, math.MaxInt32),
			"Addition result out of range")

	}

	if int64Result < math.MinInt32 {

		//func (egc *ErrorGenericFactory) GetErrorNativeToFunction(
		//	parameter string,
		//	path string,
		//	fileName string,
		//	url string,
		//	errorContext string,
		//	errorMessage string)
		return 0, errGenericCfg.GetErrorNativeToFunction(
			"",
			"",
			"",
			"",
			fmt.Sprintf("Addition result ('%d') less than minimum value of 'int32' ( %d )",
				int64Result, math.MinInt32),
			"Addition result out of range")

	}

	return int32(int64Result), nil
}
