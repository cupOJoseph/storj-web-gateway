package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"storj.io/storj/lib/uplink"
	"storj.io/storj/pkg/storj"
)

const (
	myAPIKey        = "E3KVL6HA5CSL117P67COBD4TBAK7CAHNBHK8A08="
	satellite       = "mars.tardigrade.io:7777"
	myBucket        = "upload-bucket"
	myEncryptionKey = "abc123"
)

func handleUploads(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		http.Redirect(w, r, ".", 301)
	case "POST":
		var encryptionKey storj.Key
		copy(encryptionKey[:], []byte(myEncryptionKey))

		apiKey, err := uplink.ParseAPIKey(myAPIKey)

		if err != nil {
			log.Fatalln("Could not parse api key: ", err)
		}

		r.ParseMultipartForm(10 << 20)

		// FormFile returns the first file for the given key `myFile`
		// it also returns the FileHeader so we can get the Filename,
		// the Header and the size of the file
		file, handler, err := r.FormFile("file")

		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}

		defer file.Close()

		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		// Convert image file to buffer so we can read it as byte
		buff := bytes.NewBuffer(nil)
		if _, err := io.Copy(buff, file); err != nil {
			log.Fatalln("Error: ", err)
			return
		}
		fileBuff, _ := ioutil.ReadAll(buff)

		uploadPath := "uploads/" + handler.Filename

		error := WorkWithLibUplink(satellite, &encryptionKey, apiKey, myBucket, uploadPath, []byte(fileBuff))

		if error != nil {
			log.Fatalln("Error: ", error)
		}
		fmt.Println("Success!")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func WorkWithLibUplink(satelliteAddress string, encryptionKey *storj.Key, apiKey uplink.APIKey,
	bucketName, uploadPath string, dataToUpload []byte) error {
	ctx := context.Background()

	// Create an Uplink object with a default config
	upl, err := uplink.NewUplink(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not create new Uplink object: %v", err)
	}
	defer upl.Close()

	// It is temporarily required to set the encryption key in project options.
	// This requirement will be removed in the future.
	opts := uplink.ProjectOptions{}
	opts.Volatile.EncryptionKey = encryptionKey

	// Open up the Project we will be working with
	proj, err := upl.OpenProject(ctx, satelliteAddress, apiKey, &opts)
	if err != nil {
		return fmt.Errorf("could not open project: %v", err)
	}
	defer proj.Close()

	// Create the desired Bucket within the Project
	_, err = proj.CreateBucket(ctx, bucketName, nil)
	if err != nil {
		return fmt.Errorf("could not create bucket: %v", err)
	}

	// Open up the desired Bucket within the Project
	bucket, err := proj.OpenBucket(ctx, bucketName, &uplink.EncryptionAccess{Key: *encryptionKey})
	if err != nil {
		return fmt.Errorf("could not open bucket %q: %v", bucketName, err)
	}
	defer bucket.Close()

	// Upload our Object to the specified path
	buf := bytes.NewBuffer(dataToUpload)
	err = bucket.UploadObject(ctx, uploadPath, buf, nil)
	if err != nil {
		return fmt.Errorf("could not upload: %v", err)
	}

	// Initiate a download of the same object again
	readBack, err := bucket.OpenObject(ctx, uploadPath)
	if err != nil {
		return fmt.Errorf("could not open object at %q: %v", uploadPath, err)
	}
	defer readBack.Close()

	// We want the whole thing, so range from 0 to -1
	strm, err := readBack.DownloadRange(ctx, 0, -1)
	if err != nil {
		return fmt.Errorf("could not initiate download: %v", err)
	}
	defer strm.Close()

	// Read everything from the stream
	receivedContents, err := ioutil.ReadAll(strm)
	if err != nil {
		return fmt.Errorf("could not read object: %v", err)
	}

	if !bytes.Equal(receivedContents, dataToUpload) {
		return fmt.Errorf("got different object back: %q != %q", dataToUpload, receivedContents)
	}
	return nil
}

func main() {
	port := flag.String("p", "3001", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/upload", handleUploads)

	http.HandleFunc("/ipfs", ipfsRoute)

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func ipfsRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "IPFS Route "+r.URL.Path)
}
