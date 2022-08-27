package env_test

import (
	"fmt"
	"github.com/fritzkeyzer/go-utils/env"
	"log"
	"os"
)

func ExampleLoadCfgFromEnv() {
	if err := os.Setenv("SOME_SECRET", "asd"); err != nil {
		log.Fatalln(err)
	}

	type Config struct {
		EnvName    string   `env:"ENV_NAME" default:"dev"`
		SomeSecret string   `env:"SOME_SECRET" secret:"true"`
		SomeSlice  []string `env:"SOME_SLICE" default:"'hello', 'world'"`
	}

	var cfg Config
	if err := env.Load(&cfg); err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	fmt.Print(env.Print(&cfg))
	//Output:Config:
	//	EnvName: dev
	//	SomeSecret: '***' (SECRET)
	//	SomeSlice: ['hello' 'world']
}

func ExampleReplaceVars() {
	if err := os.Setenv("SOME_SECRET", "spooky"); err != nil {
		log.Fatalln(err)
	}

	type Config struct {
		EnvName    string   `env:"ENV_NAME" default:"dev"`
		SomeSecret string   `env:"SOME_SECRET" secret:"true"`
		SomeSlice  []string `env:"SOME_SLICE" default:"'hello', 'world'"`
	}

	var cfg Config
	if err := env.Load(&cfg); err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	someExampleString := "environment={ENV_NAME}, secret={SOME_SECRET}"
	replacedString := env.ReplaceVars(someExampleString, &cfg)

	fmt.Print(replacedString)
	//Output:environment=dev, secret=spooky
}

func ExamplePrint() {
	if err := os.Setenv("SOME_SECRET", "spooky"); err != nil {
		log.Fatalln(err)
	}

	type Config struct {
		EnvName    string   `env:"ENV_NAME" default:"dev"`
		SomeSecret string   `env:"SOME_SECRET" secret:"true"`
		SomeSlice  []string `env:"SOME_SLICE" default:"'hello', 'world'"`
	}

	var cfg Config
	if err := env.Load(&cfg); err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	fmt.Print(env.Print(&cfg))
	//Output:Config:
	//	EnvName: dev
	//	SomeSecret: '******' (SECRET)
	//	SomeSlice: ['hello' 'world']
}
