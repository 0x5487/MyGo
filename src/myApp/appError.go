package main

type appError struct {
	Ex      error
	Message string
	Code    int
}

func (e *appError) Error() string { return e.Message }
