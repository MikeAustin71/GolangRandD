package main

import (
  "fmt"
)

type NumericSeparatorDto struct {
  DecimalSeparator   rune // Character used to separate integer and fractional digits ('.')
  ThousandsSeparator rune // Character used to separate thousands (1,000,000,000
  CurrencySymbol     rune // Currency Symbol
}

func (numSepDto *NumericSeparatorDto) GetInputSeparators() (*NumericSeparatorDto, error) {

  inputNumSeps := &NumericSeparatorDto{}

  inputNumSeps.CurrencySymbol = numSepDto.CurrencySymbol
  inputNumSeps.DecimalSeparator = numSepDto.DecimalSeparator
  inputNumSeps.ThousandsSeparator = numSepDto.ThousandsSeparator

  return inputNumSeps, nil
}

func (numSepDto *NumericSeparatorDto) GetOutputSeparators() (*NumericSeparatorDto, error) {

  outputNumSeps := &NumericSeparatorDto{}

  outputNumSeps.CurrencySymbol = numSepDto.CurrencySymbol
  outputNumSeps.DecimalSeparator = numSepDto.DecimalSeparator
  outputNumSeps.ThousandsSeparator = numSepDto.ThousandsSeparator

  return outputNumSeps, nil
}

type NumericSeparatorPair struct {
  InputNumSeps NumericSeparatorDto

  OutputNumSeps NumericSeparatorDto
}

func (numPair *NumericSeparatorPair) GetInputSeparators() (*NumericSeparatorDto, error) {

  inputNumSeps := &NumericSeparatorDto{}

  inputNumSeps.CurrencySymbol = numPair.InputNumSeps.CurrencySymbol
  inputNumSeps.DecimalSeparator = numPair.InputNumSeps.DecimalSeparator
  inputNumSeps.ThousandsSeparator = numPair.InputNumSeps.ThousandsSeparator

  return inputNumSeps, nil
}

func (numPair *NumericSeparatorPair) GetOutputSeparators() (*NumericSeparatorDto, error) {

  outputNumSeps := &NumericSeparatorDto{}

  outputNumSeps.CurrencySymbol = numPair.OutputNumSeps.CurrencySymbol
  outputNumSeps.DecimalSeparator = numPair.OutputNumSeps.DecimalSeparator
  outputNumSeps.ThousandsSeparator = numPair.OutputNumSeps.ThousandsSeparator

  return outputNumSeps, nil
}

type IGetNumSeparators interface {
  GetOutputSeparators() (*NumericSeparatorDto, error)

  GetInputSeparators() (*NumericSeparatorDto, error)
}

type BigIntNum struct {
  NumValue           int64
  SignValue          int
  Precision          uint
  DecimalSeparator   rune
  ThousandsSeparator rune
  CurrencySymbol     rune
}

func (bi *BigIntNum) SetNum(
  numVal int64,
  numSign int,
  precision uint) error {

  bi.NumValue = numVal
  bi.SignValue = numSign
  bi.Precision = precision

  return nil
}

func (bi *BigIntNum) SetNumWithNumSeps(
  numVal int64,
  numSign int,
  precision uint,
  numSeps IGetNumSeparators) error {

  //ePrefix := "BigIntNum.SetNumWithNumSeps()"

  bi.NumValue = numVal
  bi.SignValue = numSign
  bi.Precision = precision

  v, ok := numSeps.(IGetNumSeparators)

  if !ok {
    return fmt.Errorf("invalid type %T", numSeps)
  }

  xSeps, err := v.GetInputSeparators()

  if err != nil {
    return err
  }

  bi.DecimalSeparator = xSeps.DecimalSeparator
  bi.ThousandsSeparator = xSeps.ThousandsSeparator
  bi.CurrencySymbol = xSeps.CurrencySymbol

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

  bINum := new(BigIntNum)

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

  bINum := new(BigIntNum)

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
