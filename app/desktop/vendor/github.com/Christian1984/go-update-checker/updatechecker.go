package updatechecker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/go-version"
)

const Filename string = "latestcheck.json"
const DateFormat string = time.UnixDate

type CheckData struct {
	Timestamp   string `json:"timestamp"`
	Version     string `json:"version"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GithubApiResponseData struct {
	Version     string `json:"tag_name"`
	Name        string `json:"name"`
	Description string `json:"body"`
}

type updateChecker struct {
	Owner           string
	Repo            string
	Software        string
	DownloadLink    string
	MinDaysInterval int
	Verbose         bool
	Message         string
	UpdateAvailable bool
}

func New(owner string, repo string, software string, downloadLink string, minDaysInterval int, verbose bool) updateChecker {
	uc := updateChecker{owner, repo, software, downloadLink, minDaysInterval, verbose, "", false}
	return uc
}

func (uc *updateChecker) processError(err error) {
	if !uc.Verbose {
		return
	}

	fmt.Println("ERROR: " + err.Error())
}

func (uc *updateChecker) requestLatest() (GithubApiResponseData, error) {
	// call https://api.github.com/repos/{uc.Owner}/{uc.Repo}/releases/latest
	var apiResponse GithubApiResponseData

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, requestErr := client.Get("https://api.github.com/repos/" + uc.Owner + "/" + uc.Repo + "/releases/latest")

	if requestErr != nil {
		uc.processError(requestErr)
		return apiResponse, requestErr
	}

	body, ioUtilErr := ioutil.ReadAll(resp.Body)
	if ioUtilErr != nil {
		uc.processError(ioUtilErr)
		return apiResponse, ioUtilErr
	}

	jsonErr := json.Unmarshal(body, &apiResponse)
	if jsonErr != nil {
		uc.processError(jsonErr)
		return apiResponse, jsonErr
	}

	if uc.Verbose {
		fmt.Println("GitHub API Response:")
		fmt.Println(apiResponse)
	}

	return apiResponse, nil
}

func (uc *updateChecker) loadFile() (CheckData, error) {
	var latestCheck CheckData

	file, err := os.Open(Filename)
	if err != nil {
		uc.processError(err)
		return latestCheck, err
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)

	jsonErr := json.Unmarshal(byteValue, &latestCheck)
	if jsonErr != nil {
		uc.processError(err)

		epoch := time.Time{}
		latestCheck.Timestamp = epoch.UTC().Format(DateFormat)
		return latestCheck, jsonErr
	}

	return latestCheck, nil
}

func (uc *updateChecker) writeLatestCheckFile(checkData CheckData) error {
	file, jsonErr := json.Marshal(checkData)
	if jsonErr != nil {
		uc.processError(jsonErr)
		return jsonErr
	}

	fileErr := ioutil.WriteFile(Filename, file, 0644)
	if fileErr != nil {
		uc.processError(fileErr)
		return fileErr
	}

	return nil
}

func (uc *updateChecker) canCheck(latestCheckTimestamp string) bool {
	now := time.Now().UTC()
	lastCheck, timeErr := time.Parse(DateFormat, latestCheckTimestamp)

	if timeErr != nil {
		uc.processError(timeErr)
		return true
	}

	interval := now.Sub(lastCheck.UTC())

	if uc.Verbose {
		fmt.Println("now: " + now.UTC().Format(DateFormat))
		fmt.Println("lastCheck: " + lastCheck.UTC().Format(DateFormat))
		fmt.Printf("Duration since last check: %f days\n", interval.Hours()/3)
	}

	if interval.Hours() > float64(uc.MinDaysInterval)*24 {
		if uc.Verbose {
			fmt.Println("Time is up! Sending request to API...")
		}
		return true
	}

	if uc.Verbose {
		fmt.Println("Checker interval is not up yet! Using cached data...")
	}

	return false
}

func (uc *updateChecker) updateAvailableMessage(checkData CheckData) string {
	s := "=== INFO: A new update is available for " + uc.Software + " ==="
	bars := strings.Repeat("=", len(s))

	link := uc.DownloadLink
	if link == "" {
		link = "https://github.com/" + uc.Owner + "/" + uc.Repo + "/releases"
	}

	return bars + "\n" +
		s + "\n" +
		bars + "\n" +
		"\n" +
		"Version: " + checkData.Version + "\n" +
		"\n" +
		"Title: " + checkData.Name + "\n" +
		"\n" +
		"Description:\n" +
		checkData.Description + "\n" +
		"\n" +
		"Download the latest version here:\n" +
		link + "\n" +
		bars
}

func (uc *updateChecker) noUpdateAvailableMessage(checkData CheckData) string {
	s := "=== INFO: You are running the latestest Version of " + uc.Software + " ==="
	bars := strings.Repeat("=", len(s))

	return bars + "\n" + s + "\n" + bars
}

func (uc *updateChecker) PrintMessage() {
	fmt.Println(uc.Message)
}

func (uc *updateChecker) isCurrentVersionOutdated(currentVersion string, availableVersion string) bool {
	if uc.Verbose {
		fmt.Printf("Comparing current Version %s and available version %s\n", currentVersion, availableVersion)
	}

	vCurr, err := version.NewVersion(currentVersion)
	vAvail, err := version.NewVersion(availableVersion)

	if err != nil {
		uc.processError(err)
		return false
	}

	currOutdated := vCurr.LessThan(vAvail)

	if uc.Verbose {
		if currOutdated {
			fmt.Printf("Current Version %s is older than available version %s\n", vCurr, vAvail)
		} else {
			fmt.Printf("Available Version %s is equal or older than current version %s\n", vAvail, vCurr)
		}
	}

	return currOutdated
}

func (uc *updateChecker) CheckForUpdate(currentVersion string) {
	latestCheck, fileErr := uc.loadFile()
	if fileErr != nil {
		//return false
	}

	if uc.Verbose {
		fmt.Println("Returned from loadFile():")
		fmt.Println(latestCheck)
	}

	if uc.canCheck(latestCheck.Timestamp) {
		apiResponse, apiErr := uc.requestLatest()
		if apiErr == nil {
			now := time.Now()
			snow := now.UTC().Format(DateFormat)

			latestCheck.Timestamp = snow
			latestCheck.Version = apiResponse.Version
			latestCheck.Name = apiResponse.Name
			latestCheck.Description = apiResponse.Description

			uc.writeLatestCheckFile(latestCheck)
		} else {
			uc.processError(apiErr)
		}
	} else if uc.Verbose {
		fmt.Println("Didn't request GitHub API because interval isn't over yet...")
	}

	if uc.isCurrentVersionOutdated(currentVersion, latestCheck.Version) {
		uc.Message = uc.updateAvailableMessage(latestCheck)
		uc.UpdateAvailable = true
	} else {
		uc.Message = uc.noUpdateAvailableMessage(latestCheck)
	}
}
