package server

import "errors"

var (
	// ErrConnectionAlreadyExist is returned when a connection already exist with the same id
	ErrConnectionAlreadyExist = errors.New("Connection already exist")

	// ErrConnectionDidntExist is returned when a connection didn't exist with the same id
	ErrConnectionDidntExist = errors.New("Connection didn't exist")

	// ErrUnknowClientID is returned when a Client is not registred with this id
	ErrUnknowClientID = errors.New("Unknow client id")
)
