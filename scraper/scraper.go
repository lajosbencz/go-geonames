package scraper

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/lajosbencz/go-geonames/utils"
	"gorm.io/gorm"
)

func NewScraper(db *gorm.DB, tempDir string, keepFiles bool) (*Scraper, error) {
	var err error = nil
	if !keepFiles && tempDir == "" {
		tempDir, err = ioutil.TempDir("", "go-geoname-files")
		if err != nil {
			return nil, err
		}
	} else {
		if _, err := os.Stat(tempDir); os.IsNotExist(err) {
			os.MkdirAll(tempDir, os.ModeAppend)
		}
	}
	return &Scraper{
		DB:        db,
		TempDir:   tempDir,
		KeepFiles: keepFiles,
	}, err
}

type Scraper struct {
	DB        *gorm.DB
	TempDir   string
	KeepFiles bool
}

func (r *Scraper) downloadBinary(url string, file string) error {
	log.Printf("Downloading url %s to binary file %s\n", url, file)
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http error: %s", resp.Status)
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func (r *Scraper) download(url string, file string, skipLines int) error {
	log.Printf("Downloading url %s to file %s\n", url, file)
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http error: %s", resp.Status)
	}
	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(bufio.ScanLines)

	skippedLines := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		if skipLines > skippedLines {
			skippedLines++
			continue
		}
		out.WriteString(line + "\n")
	}
	err = scanner.Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Scraper) loadIntoSQL(table string, filePath string, sql string) error {
	log.Printf("Loading file %s into table %s\n", filePath, table)
	mysql.RegisterLocalFile(filePath)
	formatted := fmt.Sprintf(sql, table, filePath)
	r.DB.Exec(fmt.Sprintf("TRUNCATE %[1]s", table))
	r.DB.Exec(formatted)
	return nil
}

func (r *Scraper) ScrapeCountries() error {
	filePath := r.TempDir + "/countries.txt"
	if err := r.download(URLCountries, filePath, 0); err != nil {
		return err
	}
	if !r.KeepFiles {
		defer os.Remove(filePath)
	}
	return r.loadIntoSQL("countries", filePath, SQLCountries)
}

func (r *Scraper) ScrapeLanguages() error {
	filePath := r.TempDir + "/languages.txt"
	if err := r.download(URLLanguages, filePath, 1); err != nil {
		return err
	}
	if !r.KeepFiles {
		defer os.Remove(filePath)
	}
	return r.loadIntoSQL("languages", filePath, SQLLanguages)
}

func (r *Scraper) ScrapeLocations() error {
	filePath := r.TempDir + "/allCountries.zip"
	if err := r.downloadBinary(URLLocations, filePath); err != nil {
		return err
	}
	err := utils.Unzip(filePath, r.TempDir)
	if err != nil {
		fmt.Println(filePath)
		return err
	}
	unzippedFile := r.TempDir + "/allCountries.txt"
	if !r.KeepFiles {
		defer os.Remove(filePath)
		defer os.Remove(unzippedFile)
	}
	if _, err := os.Stat(unzippedFile); os.IsNotExist(err) {
		return err
	}
	return r.loadIntoSQL("locations", unzippedFile, SQLLocations)
}

func (r *Scraper) ScrapeLocationAlternates() error {
	filePath := r.TempDir + "/alternateNamesV2.zip"
	if err := r.downloadBinary(URLAlternates, filePath); err != nil {
		return err
	}
	err := utils.Unzip(filePath, r.TempDir)
	if err != nil {
		fmt.Println(filePath)
		return err
	}
	unzippedFile := r.TempDir + "/alternateNamesV2.txt"
	if !r.KeepFiles {
		defer os.Remove(filePath)
		defer os.Remove(unzippedFile)
	}
	if _, err := os.Stat(unzippedFile); os.IsNotExist(err) {
		return err
	}
	return r.loadIntoSQL("location_alternates", unzippedFile, SQLLocationAlternates)
}
