package main

import (
	"flag"
	"fmt"
	"index_data_zinc/config"
	"index_data_zinc/mails"
	"index_data_zinc/zinc"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"

	_ "net/http/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
var goroutineprofile = flag.String("goroutineprofile", "", "write goroutine profile to `file`")

func main() {
	// Profilling cpu
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// Profilling goroutines
	if *goroutineprofile != "" {
		f, err := os.Create(*goroutineprofile)
		if err != nil {
			log.Fatal("could not create goroutine profile: ", err)
		}
		defer f.Close()
		runtime.GC() // Forcing garbage collection to get up-to-date statistics
		if err := pprof.Lookup("goroutine").WriteTo(f, 0); err != nil {
			log.Fatal("could not write goroutine profile: ", err)
		}
	}

	// Profilling memory
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}

	// Verificar que se haya pasado un argumento
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <data_directory>\n", os.Args[0])
	}

	// Obtener el directorio de datos desde los argumentos
	dataDir := os.Args[4]
	mailDir := filepath.Join(dataDir, "/maildir")
	fmt.Printf("mailDir: %s\n", mailDir)

	// Crear índice si no existe
	err := zinc.CreateZincIndex(config.ZincIndex)
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	// Verificar si el directorio existe
	if _, err := os.Stat(mailDir); os.IsNotExist(err) {
		log.Fatalf("Directory does not exist: %v", err)
	}

	// Llamar a la función para procesar correos electrónicos
	if err := mails.ProcessEmails(mailDir); err != nil {
		log.Fatalf("Error processing emails: %v", err)
	}
}
