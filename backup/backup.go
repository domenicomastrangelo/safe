package backup

import (
	"archive/zip"
	"io"
	"log"
	"os"

	"github.com/domenicomastrangelo/safe/database"
)

// Everything executes the backup of the
// app, its data and configuration
// to the location previously
// specified inside the configuration
func Everything(backup *bool) {
	if *backup {
		createBackup()
		saveBackup()
	}
}

func createBackup() {
	var (
		input    string = database.DBFileName
		output   string = "backup.zip"
		zipFile  *os.File
		fileInfo os.FileInfo
		header   *zip.FileHeader
		writer   io.Writer
		err      error
	)

	if _, err := os.Stat(input); err != nil {
		log.Println("Backup is not possible. Database file does not exist")
		return
	}

	if zipFile, err = os.Create(output); err != nil {
		log.Fatalln("Could not create the zip file")
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	if fileInfo, err = os.Stat(input); err != nil {
		log.Fatalln("Could not read database file info")
	}

	if header, err = zip.FileInfoHeader(fileInfo); err != nil {
		log.Fatalln("Could not get file info header")
	}

	header.Name = input
	header.Method = zip.Deflate

	if writer, err = zipWriter.CreateHeader(header); err != nil {
		log.Fatalln("Could not create the zip header")
	}

	if _, err := io.Copy(writer, zipFile); err != nil {
		log.Fatalln("Could not write the zip file")
	}
}

func saveBackup() {

}
