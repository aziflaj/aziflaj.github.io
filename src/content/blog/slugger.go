package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	// for all files in this folder, grab the title, remove the first yyyy-mm-dd- part from the filename, and print
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if name == "slugger.go" {
			continue
		}

		slug := name[11:]
		slug = slug[:len(slug)-3]
		log.Println(slug)

		// read file as a byte array, from the first --- to the second ---

		m, err := os.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}

		pos := 0 // current position in data
		metadataDelim := []byte("---")

		i := bytes.Index(m, metadataDelim) // find meta start
		if i == -1 {
			log.Printf("no meta start %s", name)
			continue
		}
		i += len(metadataDelim) - 1                       // move over "metadata"
		metaStart := pos + i                              // metadata start
		size := bytes.Index(m[metaStart:], metadataDelim) // find closing
		if size == -1 {
			log.Printf("no meta end %s", name)
			continue
		}

		metaEnd := metaStart + size        // metadata end
		pos = metaEnd + len(metadataDelim) // move over "metadata"

		// add slug to metadata right before metaEnd
		m = append(m[:metaEnd], append([]byte("\nslug: "+slug+"\n"), m[metaEnd:]...)...)

		// write back to file
		err = os.WriteFile(name, m, 0644)
		if err != nil {
			log.Printf("error writing file %s: %v", name, err)
		}
	}
}
