package types

import "errors"

// some errors definition
var (
	ErrAccountIsFrozen     = errors.New("Account has been frozen")
	ErrAccountIsLowLevel   = errors.New("Account is low level")
	ErrAccountIDExist      = errors.New("The account ID has been registered!")
	ErrAccountIDNotExist   = errors.New("The account ID is not exist")
	ErrAccountIDNotPermiss = errors.New("You don't have permission to do that!")
	ErrAssetBalance        = errors.New("Insufficient balance!")
	ErrNotAdmin            = errors.New("No adiministrator privileges!")
)
