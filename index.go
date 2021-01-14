// // main.go
// package main

// import (
//   "fmt"
//   "log"
//   "net/http"
// )

// func main() {
//   const port = 8000
//   listenAt := fmt.Sprintf(":%d", port)
//   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprint(w, "Hello, World!")
//   })

//   log.Printf("Open the following URL in the browser: http://localhost:%d\n", port)
//   log.Fatal(http.ListenAndServe(listenAt, nil))
// }



package main

import (  
    "fmt"
	"os"
	"flag"
	"github.com/fatih/color"
	
	"github.com/ilyakaznacheev/cleanenv"
)


type Config struct {
    Ports struct {
        Daemon string `yaml:"daemon"`
        Sftp string `yaml:"sftp"`
    } `yaml:"ports"`
	Panel struct {
        AuthenticationTokenId string `yaml:"AuthenticationTokenId"`
        url string `yaml:"url"`
    } `yaml:"panel"`
}


var configPath = "darmon.yml"
var debugMode = true

type Args struct {
	ConfigPath string
}

func main() {
	logMsg("Freddie Darmon", "startup")
	logMsg("Loading configeration file...", "startup")

	var cfg Config

	args := ProcessArgs(&cfg)

	// read configuration from the file and environment variables
	if err := cleanenv.ReadConfig(args.ConfigPath, &cfg); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	startAPI(cfg)
}






























func logMsg(msg string, debug string ) {
	var toSend = ""
	if debug == "debug" && debugMode {
		toSend = toSend + "[DEBUG] - " + msg
		color.Cyan(toSend)
	} else if debug == "startup"{
		toSend = "[STARTUP] - "+toSend + msg
		color.Yellow(toSend)
	} else if debug == "prosess"{
		toSend = "[PROSESS] - "+toSend + msg
		color.Green(toSend)
	}
	
	// fmt.Println(toSend)
}

func loadConfigData(  ) {
	var cfg Config

	args := ProcessArgs(&cfg)

	// read configuration from the file and environment variables
	if err := cleanenv.ReadConfig(args.ConfigPath, &cfg); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
// ProcessArgs processes and handles CLI arguments
func ProcessArgs(cfg interface{}) Args {
	var a Args

	f := flag.NewFlagSet("Example server", 1)
	f.StringVar(&a.ConfigPath, "c", configPath, "Path to configuration file")

	fu := f.Usage
	f.Usage = func() {
		fu()
		envHelp, _ := cleanenv.GetDescription(cfg, nil)
		fmt.Fprintln(f.Output())
		fmt.Fprintln(f.Output(), envHelp)
	}

	f.Parse(os.Args[1:])
	return a
}