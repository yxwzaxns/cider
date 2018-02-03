package core

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/mholt/archiver"

	git "gopkg.in/src-d/go-git.v4"
)

var (
	CodeURL  string
	WorkPath string
	RootDir  string
	MChan    chan M
	m        *M
)

func init() {
	CodeURL = ""
	WorkPath = ""
	RootDir = "/tmp"
	m = new(M)
}

func StartCI(URL string, msgChan chan M) {
	MChan = msgChan
	CodeURL = URL
	m.URL = URL
	m.info = ""
	// time.Sleep(50000 * time.Millisecond)
	initWorkDir()
	println("initWorkDir finished")
	pullCode()
	println("pullCode finished")
	buildImage()
	println("buildImage finished")
	checkImage()
	println("checkImage finished")
	clean()
	println("CI finished")
}

func initWorkDir() {
	basePath := filepath.Join(RootDir, "cider_workspace")
	path := filepath.Join(basePath, convertURLToPath(CodeURL))

	if _, err := os.Stat(path); os.IsNotExist(err) {
		// os.RemoveAll(path)
		if err := os.MkdirAll(path, 0777); err != nil {
			println(err.Error())
		}
	} else {
		// if dir exist that check if exist .git file
		//
		// to do
		// git pull code
	}
	WorkPath = path
}
func pullCode() {
	projectURL := "https://" + CodeURL + ".git"
	// println("pullCode : ", URL)
	// if dir exist that check if exist .git file
	//
	// to do
	// git pull code
	// else
	// git clone
	gitFilePath := filepath.Join(WorkPath, ".git")

	if _, err := os.Stat(gitFilePath); err == nil {
		// git pull
		// We instance a new repository targeting the given path (the .git folder)
		r, err := git.PlainOpen(WorkPath)
		CheckIfError(err)

		// Get the working directory for the repository
		w, err := r.Worktree()
		CheckIfError(err)

		// Pull the latest changes from the origin remote and merge into the current branch
		// Info("git pull origin")
		w.Pull(&git.PullOptions{RemoteName: "origin"})

		// CheckIfError(err)

		// Print the latest commit that was just pulled
		// ref, err := r.Head()
		// CheckIfError(err)
		// commit, err := r.CommitObject(ref.Hash())
		// CheckIfError(err)

		// fmt.Println(commit)
	} else {
		// git clone
		_, err := git.PlainClone(WorkPath, false, &git.CloneOptions{
			URL: projectURL,
			// Progress: os.Stdout,
		})
		if err != nil {
			println("git clone error : ", err.Error())
		}
	}
}
func buildImage() {
	// put dir to #{CodeURL}.tar
	tarFilePath := filepath.Join(WorkPath, "aim.tar")
	createTar(WorkPath)
	defer deleteFile(tarFilePath)
	// send build req to Docker Daemon
	// ctx, _ := context.WithCancel(context.Background())
	ctx := context.Background()
	if dockerBuildContext, err := os.Open(tarFilePath); err == nil {
		defer dockerBuildContext.Close()
		cli, err := client.NewEnvClient()
		if err != nil {
			fmt.Printf("%s", err.Error())
		}
		options := types.ImageBuildOptions{
			Tags:           []string{CodeURL},
			NoCache:        false,
			SuppressOutput: false,
			Remove:         true,
			ForceRemove:    true,
			PullParent:     true,
		}
		buildResponse, err := cli.ImageBuild(ctx, dockerBuildContext, options)
		if err != nil {
			fmt.Printf("%s", err.Error())
		} else {
			// fmt.Printf("********* %s **********", buildResponse.OSType)
			_, err := ioutil.ReadAll(buildResponse.Body)
			if err != nil {
				fmt.Printf("%s", err.Error())
			} else {
				// timeout := 5 * 60 * 1000
				// timeout := 30 * 1000
				// time.Sleep(time.Duration(timeout) * time.Millisecond)
				// cancel()
				_, err := ioutil.ReadAll(buildResponse.Body)
				if err != nil {
					fmt.Printf("%s", err.Error())
				}
				// fmt.Println(string(response))
			}
		}
	}
}
func checkImage() {

}
func sendNotification(s string) {
	// m.info = s
	// mChan <- *m
}
func clean() {

}

func deleteFile(path string) {
	if _, err := os.Stat(path); err == nil {
		// os.RemoveAll(path)
		if err := os.RemoveAll(path); err != nil {
			println(err.Error())
		}
	}
}
func convertURLToPath(URL string) string {
	return strings.Replace(strings.Replace(URL, "/", "_", -1), ".", "_", -1)

}
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
func createTar(path string) {
	// path := "/tmp/cider_workspace/github_com_yxwzaxns_cider-ci-test"
	dirs := getDirList(WorkPath)
	tarPath := filepath.Join(path, "aim.tar")
	archiver.Tar.Make(tarPath, dirs)
}
func getDirList(path string) []string {
	files := []string{}
	fs, _ := ioutil.ReadDir(path)
	for _, f := range fs {
		files = append(files, filepath.Join(path, f.Name()))
	}
	return files
}
