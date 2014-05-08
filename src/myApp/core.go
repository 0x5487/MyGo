package main

type appError struct {
	Error   error
	Message string
	Code    int
}
