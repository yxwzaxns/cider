package core

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/mholt/archiver"

	git "gopkg.in/src-d/go-git.v4"
)

var (
	CodeURL          string
	WorkPath         string
	RootDir          string
	MChan            chan M
	m                *M
	DockerAPIVersion float64
)

func init() {
	CodeURL = ""
	WorkPath = ""
	RootDir = "/tmp"
	m = new(M)
	DockerAPIVersion = 1.36
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
	for {
		if buildImage() == 0 {
			break
		} else {
			DockerAPIVersion -= 0.01
		}
	}
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
			os.Exit(1)
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
func buildImage() int {
	// put dir to #{CodeURL}.tar
	tarFilePath := filepath.Join(WorkPath, "aim.tar")
	createTar(WorkPath)
	defer deleteFile(tarFilePath)

	dockerBuildContext, err := os.Open(tarFilePath)
	defer dockerBuildContext.Close()
	if err != nil {
		panic(err)
	}

	// set DOCKER_API_VERSION
	os.Setenv("DOCKER_API_VERSION", strconv.FormatFloat(DockerAPIVersion, 'f', 2, 64))
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	options := types.ImageBuildOptions{
		Tags:           []string{CodeURL},
		NoCache:        false,
		SuppressOutput: false,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     true,
	}

	// ctx, _ := context.WithCancel(context.Background())
	ctx := context.Background()
	// send build req to Docker Daemon
	buildResponse, err := cli.ImageBuild(ctx, dockerBuildContext, options)
	if err != nil {
		if res, _ := regexp.MatchString("too new", err.Error()); res == true {
			return 1
		} else {
			panic(err)
		}
	} else {
		// fmt.Printf("********* %s **********", buildResponse.OSType)
		_, err := ioutil.ReadAll(buildResponse.Body)
		if err != nil {
			panic(err)
		} else {
			// timeout := 5 * 60 * 1000
			// timeout := 30 * 1000
			// time.Sleep(time.Duration(timeout) * time.Millisecond)
			// cancel()
			// response, err := ioutil.ReadAll(buildResponse.Body)
			// if err != nil {
			// fmt.Printf("%s", err.Error())
			// } else {
			// fmt.Println(string(response))
			// }
			// fmt.Println(string(response))
		}
	}
	return 0
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
