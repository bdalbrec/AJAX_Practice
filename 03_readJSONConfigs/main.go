package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Conn struct {
	Config []Config `json:"connection"`
}

type Config struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// open JSON file
	jf, err := os.Open("dbConfigs.json")
	if err != nil {
		fmt.Printf("Failed to open config file. %v/n", err)
	}
	fmt.Println("Successfully opened config file")

	// defer closing of JSON file so that we can read it
	defer jf.Close()

	// initialize a variable to unmarshall into
	var con Conn

	// read the JSON file into a slice of bytes
	bs, err := ioutil.ReadAll(jf)
	if err != nil {
		fmt.Printf("Could not read config file. %v/n", err)
	}

	//unmarshal the byte slice into the con variable of type Conn
	err = json.Unmarshal(bs, &con)
	if err != nil {
		fmt.Printf("Could not unmarshal JSON. %v", err)
	}

	// Boom! We've unmarshalled the JSON into our program. Print Conn.Config struc for sanity check.
	fmt.Println(con.Config)

}
