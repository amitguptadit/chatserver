
//**************************************************************************
// this is code for POC perpose for developing the chat server using go lang.
//
// Copyright Â© 2/04/2022
// Confidential. All rights reserved.
//@author Amit Gupta
//**************************************************************************

package main

// Common structs
type Config struct {
	Addr     string `json:"addr"`
	RestAddr string `json:"restaddr"`
	LogFile  string `json:"logfile"`
}

type Message struct {
	From    string `json:"from"`
	Content string `json:"content"`
}
