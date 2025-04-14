package main

import (
	"fmt"
	"github.com/melina-delgado/nested-gomod-example/single-package/s3"
)

func main () {
	// Use API
	s3API := s3.ListObjectsAPI
	fmt.Println(s3API)
}