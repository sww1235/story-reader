/*
StoryReader is an application to read, search and display stories along with
tags provided in a custom markdown based format

story names, authors and tags are stored in a sqlite database in the same directory to allow
faster tag searching and displaying. This is built whenever story files are
opened or manual directory scans
*/

package main

import "flag"

var helpFlag = flag.Bool("help", FALSE, "Prints help message") 

func main() {
	//list of operations:
	// open one story file from command line
	// scan directory of files
	// search database for tag(s)
	// open webserver for display and navigation purposes
	// show usage (help) info

	// want to limit amount of text displayed at once and allow for
	// automatically breaking up into sections
	//
	// also allow for automatic chapter detection and table of contents
	// generation

	// may want to provide separate method for tagging stories related
	// to each other

}
