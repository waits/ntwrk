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
	Tag    string  `json:"tag_name"`
	Assets []Asset `json:"assets"`
}

type Asset struct {
	Name string `json:"name"`
	Url  string `json:"browser_download_url"`
}

func (rel *Release) assetForPlatform(os string, arch string) (a Asset, err error) {
	assetName := fmt.Sprintf(fileFmt, os, arch)
	for _, a = range rel.Assets {
		if a.Name == assetName {
			return
		}
	}
	err = fmt.Errorf("no release for %s/%s", os, arch)
	return
}

func update(host string) {
	rel, err := getLatestRelease()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	if rel.Tag == "v"+version {
		fmt.Println("No update available.")
		return
	}

	asset, err := rel.assetForPlatform(runtime.GOOS, runtime.GOARCH)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = applyPatch(asset)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func getLatestRelease() (rel Release, err error) {
	resp, err := http.Get(updateUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	json.Unmarshal(dat, &rel)
	return
}

func applyPatch(asset Asset) error {
	path := filepath.Join("/", "tmp", asset.Name)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	defer os.Remove(path)

	err = file.Chmod(execMode)
	if err != nil {
		return err
	}

	fmt.Println("Downloading", asset.Url)
	resp, err := http.Get(asset.Url)
	if err != nil {
		return err
	}
	io.Copy(file, resp.Body)

	exec, err := os.Executable()
	if err != nil {
		return err
	}

	os.Rename(exec, exec+".old")
	os.Rename(path, exec)
	os.Remove(exec + ".old")

	return nil
}
