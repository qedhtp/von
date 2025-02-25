package pkg

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func Pronounce(pronounce_url *string) {
	// mate the GET pronounce request
	response_pronounce, err := http.Get(*pronounce_url)
	if err != nil {
		log.Fatal(err)
	}
	// will be closed once the main function exits
	defer response_pronounce.Body.Close()

	// Create a file to save the voice binary file
	voice_file, err := os.Create("/tmp/voice_tmp.mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer voice_file.Close()

	// Copy the response body to the file and also a variable
	body_pronounce, err := io.ReadAll(response_pronounce.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Write the response body to the file
	_, err = voice_file.Write(body_pronounce)
	if err != nil {
		log.Fatal(err)
	}
	
	err = exec.Command("mpg123","-q","/tmp/voice_tmp.mp3").Run()
	if err != nil {
		log.Fatal(err)
	}
	
	err = os.Remove("/tmp/voice_tmp.mp3")
	if err != nil {
		log.Fatal(err)
	}



}