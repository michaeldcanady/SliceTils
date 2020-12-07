# SliceTils

## Table Of Contents
[What is SliceTils](#WhatisSliceTils)
[Commands](#Commands)

## What is SliceTils
SliceTils (Slice Utilities) is a library adding in more functionalities and control for slices. Some of the commands have been engineered to accept all types of slices and handle them accordingly; the goal being to standardized and expedite the process.

## Commands
Currently SliceTils only supports int and string slices.

### Universal
 * [Min([]slice)](https://github.com/michaeldcanady/SliceTils/blob/main/SliceTils/SliceTils.go) - This function gets the minimum of any slice provided.
 * [Max([]slice)](https://github.com/michaeldcanady/SliceTils/blob/main/SliceTils/SliceTils.go) - This function gets the maximum of any slice provided.
 * [Equal([]slice,[]slice)](https://github.com/michaeldcanady/SliceTils/blob/main/SliceTils/SliceTils.go) - This function checks whether the two provided slices are equal.
 * [NthRand(int,type,int)](https://github.com/michaeldcanady/SliceTils/blob/main/SliceTils/SliceTils.go) - This function returns a randomly generated slice.
  * The first parameter is the length of the slice.
  * The second is the type of the slice, as a string.
  * The third is only for integer slices, the range for the values.
 * [TwoToOne([][]slice)](https://github.com/michaeldcanady/SliceTils/blob/main/SliceTils/SliceTils.go) - This function turns the provided 2D slice into a 1D slice.
 * [IndexOf(element,[]slice)](https://github.com/michaeldcanady/SliceTils/blob/main/SliceTils/SliceTils.go) - This function returns the index, within the slice, of the element.
 * [RemoveDup([]slice)](https://github.com/michaeldcanady/SliceTils/blob/main/SliceTils/SliceTils.go) - This function returns a modified slice without any repeating values
 * [Convert([]slice,type)](https://github.com/michaeldcanady/SliceTils/blob/main/SliceTils/SliceTils.go) - This function converts provided slice into specified type, as a string.
 * [Contains([]slice,element)](https://github.com/michaeldcanady/SliceTils/blob/main/SliceTils/SliceTils.go) - This function checks if a slice contains the element, it returns a bool (if it is contained) and the index of the value.

### Int

### String
