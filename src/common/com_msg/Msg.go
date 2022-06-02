package com_msg

import "fmt"

const (
	ADD_SUCCESS           string = `Add Successed.`
	ADD_FAIL              string = `Add Failed.`
	UPD_SUCCESS           string = `Update Successed.`
	UPD_FAIL              string = `Update Failed.`
	DEL_SUCCESS           string = `Delete Successed.`
	DEL_FAIl              string = `Delete Failed.`
	HAS_BEEN_USED_NOT_DEL string = `The record has been used and cannot be deleted.`
)

func Required(word string) string {
	return fmt.Sprintf(`%s is required.`, word)
}

func PositiveInteger(word string) string {
	return fmt.Sprintf(`%s must be a positive integer.`, word)
}

func ExistsNotUpdate(word string) string {
	return fmt.Sprintf(`%s already exists and cannot be updated.`, word)
}

func HasBeenUsedNotUpdate(word string) string {
	return fmt.Sprintf(`%s has been used and cannot be updated.`, word)
}
