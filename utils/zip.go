package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/dustin/go-humanize"
)

type WriteCounter struct {
	Total uint64
}

func CloneTemplate(repositoryURL string) {
	download(repositoryURL)

	extract()

	cleanup()
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.printProgress()
	return n, nil
}

func (wc WriteCounter) printProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	fmt.Printf("\r \U0001F9D9 Downloading Scrolls... %s complete", humanize.Bytes(wc.Total))
}

func download(url string) {
	fileUrl := url
	err := downloadFile("./template.zip", fileUrl)
	if err != nil {
		panic(err)
	}
}

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()

	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	fmt.Print("\n")

	out.Close()

	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}

func extract() {

	err := os.Mkdir("Horizon.Modules.Boilerplate", 0755)
	if err != nil {
		fmt.Println(err)
	}

	files, err := unzip("./template.zip", "./Horizon.Modules.Boilerplate")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		moveFile(file)
	}
}

func unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

func cleanup() {
	err := os.RemoveAll("./Horizon.Modules.Boilerplate")
	if err != nil {
		log.Fatal(err)
	}

	err2 := os.Remove("./template.zip")
	if err2 != nil {
		log.Fatal(err)
	}

}
