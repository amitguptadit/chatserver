
//**************************************************************************
// this is code for POC perpose for developing the chat server using go lang.
//
// Copyright Â© 2022
// Confidential. All rights reserved.
//@author Amit Gupta
//**************************************************************************
/**************************************************************************
* main - This is the starting point for start the chat server
* @author Amit Gupta
**************************************************************************/
package main

import (
	//"errors"
	"fmt"
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	fmt.Println("[ welcome in chat server... ]")
	// create the parent context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// set up signals to cancel the context
	sig := make(chan os.Signal, 10)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() { <-sig; cancel() }()

	// run Main command
	if err := Main(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		fmt.Println("Terminating server...")
		os.Exit(1)
	}
}

func Main(ctx context.Context) (err error) {
	// Parse flags
	flag.Parse()

	//set the config file for reading configuration
	conffile := "conf.json"

	conf, err := readConfigFile(conffile)
	if err != nil {
		return err
	}
	err = checkParams(conf)
	if err != nil {
		return err
	}

	// Run server
	server := NewServer(conf.Addr, conf.RestAddr, conf.LogFile)
	return server.Run(ctx)
}