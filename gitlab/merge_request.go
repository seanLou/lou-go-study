package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	WorkPath     = "/tmp/go_git_name_id.txt"
	PrivateToken = "bcdr3spZM56yUCPxs4Jk"
	GitHost      = "http://gitlab.com"
	PROJECTS     = "/api/v4/projects"
	//SKIP_PROJECT_NAME =
)

type GitProject struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func initGitProject() []GitProject {
	fmt.Println("开始初始化Git项目...")
	// 创建文件
	file, err := os.OpenFile(WorkPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
	handlerError(err)
	defer file.Close()

	var allGitProjects []GitProject

	for currentPage := 1; ; currentPage++ {
		body := getGitProject(currentPage)
		var gitProjects []GitProject
		err := json.Unmarshal(body, &gitProjects)
		handlerError(err)
		if len(gitProjects) < 1 {
			break
		}
		allGitProjects = append(allGitProjects, gitProjects...)
	}
	bytes, err := json.Marshal(allGitProjects)
	handlerError(err)
	n, err := file.Write(bytes)
	handlerError(err)
	fmt.Printf("文件字节数:%d, Git项目总数:%d\n", n, len(allGitProjects))

	return allGitProjects
}

func handlerError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getGitProject(currentPage int) []byte {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr}

	request, err := http.NewRequest(http.MethodGet, GitHost+PROJECTS, nil)

	handlerError(err)

	query := request.URL.Query()

	query.Add("simple", "true")
	query.Add("membership", "true")
	query.Add("per_page", "100")
	query.Add("page", strconv.Itoa(currentPage))
	query.Add("order_by", "id")
	query.Add("sort", "asc")

	request.URL.RawQuery = query.Encode()

	request.Header.Add("PRIVATE-TOKEN", PrivateToken)

	response, err := client.Do(request)

	handlerError(err)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	handlerError(err)

	fmt.Printf("%s\n", body)

	return body
}

func MergeRequest() {
	_, err := os.Lstat(WorkPath)
	var allGitProjects []GitProject
	if err != nil {
		allGitProjects = initGitProject()
	} else {
		file, err := os.OpenFile(WorkPath, os.O_RDONLY, 0755)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(bytes, &allGitProjects)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, gitProject := range allGitProjects {
		var id = gitProject.Id
		var name = gitProject.Name
		fmt.Printf("id:%d, name:%s\n", id, name)

	}
}
