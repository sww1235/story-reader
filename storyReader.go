/*
StoryReader is an application to read, search and display stories along with
tags provided in a custom markdown based format

Stories must be located in a specified root story directory or one of its
subdirectories. There can be several story directories, which will each have
separate databases but will be searchable at the same time.

story names, authors and tags are stored in a sqlite database in a root story
directory to allow faster tag searching and displaying. This is built whenever
story files are opened or manual directory scans
*/

package main

import "flag"

var helpFlag bool // if help should be printed
var useWebServer bool // if built in web server should be enabled
var scanDirectory string // scan specified directory for story files to add to database
var fileToOpen string // filename passed in from the command line to open
var configFileLocation string // location of configuration file

var serverPort int // port to use for webserver

var storyDirectories []string


func init() {
	const ( helpUsage = "Show usage information"
		webServerUsage = "Host web server to display formatted stories, use --port or -p to specify port"
		scanDirectoryUsage = "Scan all files in specified directory looking for files in story format. If found, then add or update them to database in root story directory. "
		fileToOpenUsage = "Path to a story file. If file is not in story format, prints to stderr."
		configFileLocationUsage = "Path to program config file. Defaults to ~/.config/story-reader.toml on *nix systems."
		serverPortUsage = "port to use for web server. Defaults to 7331"

		defaultConfigLocation = "~/.config/story-reader.toml"
		defaultServerPort = 7331

)
	flag.BoolVar(&helpFlag, "help", False, helpUsage)
	flag.BoolVar(&helpFlag, "h", False, helpUsage+" (Shorthand)")

	flag.BoolVar(&useWebServer, "webserver", False, webServerUsage)
	flag.BoolVar(&useWebServer, "w", False, webServerUsage+" (Shorthand)")

	flag.StringVar(&scanDirectory, "scan", "", scanDirectoryUsage)
	flag.StringVar(&scanDirectory, "s", "", scanDirectoryUsage)

	flag.StringVar(&fileToOpen, "file", "", fileToOpenUsage)
	flag.StringVar(&fileToOpen, "f", "", fileToOpenUsage)

	flag.StringVar(&configFileLocation, "config", defaultConfigLocation, configFileLocationUsage)
	flag.StringVar(&configFileLocation, "c", defaultConfigLocation, configFileLocationUsage)

	flag.IntVar(&serverPort, "port", defaultServerPort, serverPortUsage)
	flag.IntVar(&serverPort, "p", defaultServerPort, serverPortUsage)
}

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


	flag.Parse() // parse flags first thing

	if helpFlag { // print help information and exit
		PrintDefaults()
	} else if scanDirectory != "" {
		readConfig()
		scanDirectory()
	}
}

func readConfig() {


}



func scanDirectory(string pathToScan) {
	openDB()

}

func openDB() {


}


