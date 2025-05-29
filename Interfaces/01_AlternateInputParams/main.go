package main

import (
	"fmt"
)

type NumericSeparatorDto struct {
	DecimalSeparator   rune // Character used to separate integer and fractional digits ('.')
	ThousandsSeparator rune // Character used to separate thousands (1,000,000,000
	CurrencySymbol     rune // Currency Symbol
}

type NumericSeparatorPair struct {
	InputNumSeps NumericSeparatorDto

	OutputNumSeps NumericSeparatorDto
}

//type INumSepInOut interface {
//	*NumericSeparatorDto | *NumericSeparatorPair
//}

type BigIntNum[T any] struct {
	NumValue           int64
	SignValue          int
	Precision          uint
	DecimalSeparator   rune
	ThousandsSeparator rune
	CurrencySymbol     rune
}

func (bi *BigIntNum[T]) SetNum(
	numVal int64,
	numSign int,
	precision uint) error {

	bi.NumValue = numVal
	bi.SignValue = numSign
	bi.Precision = precision

	return nil
}

func (bi *BigIntNum[T]) SetNumWithNumSeps(
	numVal int64,
	numSign int,
	precision uint,
	numSeps T) error {

	ePrefix := "BigIntNum.SetNumWithNumSeps()"

	bi.NumValue = numVal
	bi.SignValue = numSign
	bi.Precision = precision

	switch v := any(numSeps).(type) {
	case *NumericSeparatorPair:
		bi.CurrencySymbol = v.InputNumSeps.CurrencySymbol
		bi.ThousandsSeparator = v.InputNumSeps.ThousandsSeparator
		bi.DecimalSeparator = v.InputNumSeps.DecimalSeparator
	case *NumericSeparatorDto:
		bi.CurrencySymbol = v.CurrencySymbol
		bi.ThousandsSeparator = v.ThousandsSeparator
		bi.DecimalSeparator = v.DecimalSeparator

	default:
		return fmt.Errorf("%v\n"+
			"Error: 'numSeps' DID NOT MATCH Known Types!\n"+
			"Type is: %v\n\n",
			ePrefix, v)
	}

	return nil
}

func main() {

	CallerNumSepsDto()

	CallerNumSepsPair()
}

func CallerNumSepsDto() {

	ePrefix := "CallerNumSepsDto()"

	usaNumSeps := NumericSeparatorDto{
		DecimalSeparator:   '.',
		ThousandsSeparator: ',',
		CurrencySymbol:     '$',
	}

	bINum := new(BigIntNum[*NumericSeparatorDto])

	err := bINum.SetNumWithNumSeps(
		574,
		1,
		1,
		&usaNumSeps)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n\nSuccessful Completion\n"+
			"%v\n"+
			"bigNum.DecimalSeparator= '%v'\n"+
			"bigNum.ThousandsSeparator= '%v'\n"+
			"bigNum.CurrencySymbol= '%v'\n",
			ePrefix,
			string(bINum.DecimalSeparator),
			string(bINum.ThousandsSeparator),
			string(bINum.CurrencySymbol))
	}

	return
}

func CallerNumSepsPair() {

	ePrefix := "CallerNumSepsPair()"

	usaNumSeps := NumericSeparatorDto{
		DecimalSeparator:   '.',
		ThousandsSeparator: ',',
		CurrencySymbol:     '$',
	}

	numSepsPair := new(NumericSeparatorPair)

	numSepsPair.InputNumSeps = usaNumSeps
	numSepsPair.OutputNumSeps = usaNumSeps

	bINum := new(BigIntNum[any])

	err := bINum.SetNumWithNumSeps(
		574,
		1,
		1,
		numSepsPair)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n\nSuccessful Completion\n"+
			"%v\n"+
			"bigNum.DecimalSeparator= '%v'\n"+
			"bigNum.ThousandsSeparator= '%v'\n"+
			"bigNum.CurrencySymbol= '%v'\n",
			ePrefix,
			string(bINum.DecimalSeparator),
			string(bINum.ThousandsSeparator),
			string(bINum.CurrencySymbol))
	}

	return
}
