package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

const execMode = 0755
const fileFmt = "ntwrk-%s-%s"
const updateUrl = "https://api.github.com/repos/waits/ntwrk/releases/latest"

type Release struct {
	Name   string  `json:"name"`
	Assets []Asset `json:"assets"`
}

type Asset struct {
	Name string `json:"name"`
	Url  string `json:"browser_download_url"`
}

func update(host string) {
	asset, err := getLatestRelease()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	path := filepath.Join("/", "tmp", asset.Name)
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer file.Close()
	defer os.Remove(path)

	err = file.Chmod(execMode)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Downloading", asset.Url)
	resp, err := http.Get(asset.Url)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	io.Copy(file, resp.Body)

	exec, err := os.Executable()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	os.Rename(exec, exec+".old")
	os.Rename(path, exec)
	os.Remove(exec + ".old")
}

func getLatestRelease() (asset Asset, err error) {
	resp, err := http.Get(updateUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	rel := Release{}
	json.Unmarshal(dat, &rel)

	assetName := fmt.Sprintf(fileFmt, runtime.GOOS, runtime.GOARCH)
	for _, a := range rel.Assets {
		if a.Name == assetName {
			return a, nil
		}
	}
	err = fmt.Errorf("no release for %s/%s", runtime.GOOS, runtime.GOARCH)
	return
}
