package main

import (
    "fmt"
    "io/ioutil"
    "regexp"
    "strings"
    "flag"
)

type TestStruct struct {
    id   int
}

func isTagExists(tag string, tags []string) bool {
    for _, t := range(tags) {
        if t == tag {
            return true
        }
    }
    return false
}

func main() {
    kind := flag.String("kind", "service", "kind of file to generate")
    baseName := flag.String("name", "test", "name of file to generate")

    flag.Parse()

    serviceName := *baseName + "Service"
    repositoryName := *baseName + "Repository"

    var path string = ""

    if *kind == "service" || *kind == "svc" {
        path = "cmd/generate/service.template"
    } else if *kind == "repository" || *kind == "repo" {
        path = "cmd/generate/repository.template"
    }

    data, err := ioutil.ReadFile(path)

    if err != nil {
        fmt.Println("Error: ", err)
    }

    tagFounds := make([]string, 0)

    re := regexp.MustCompile(`\${\w+}`)
    tags := re.FindAllString(string(data), -1)

    for _, tag := range(tags) {
        tag := tag[2:len(tag)-1]
        if !isTagExists(tag, tagFounds) {
            tagFounds = append(tagFounds, tag)
        }
    }

    // write to file
    generated := re.ReplaceAllStringFunc(string(data), func(tag string) string {
        tagName := tag[2:len(tag)-1]

        if *kind == "service" || *kind == "svc" {
            if tagName == "ServiceName" {
                return serviceName
            } else if tagName == "repositoryName" {
                return repositoryName
            } else if tagName == "RepositoryName" {
                return strings.Title(repositoryName)
            } else if tagName == "repositoryName" {
                return repositoryName
            }
        } else if *kind == "repository" || *kind == "repo" {
            if tagName == "repositoryName" {
                return repositoryName
            }
        }

        return tag
    })

    var outpath string = ""
    if *kind == "service" || *kind == "svc" {
        outpath = "src/app/services/" + strings.ToLower(*baseName) + ".go"
    } else if *kind == "repository" || *kind == "repo" {
        outpath = "src/app/repositories/" + strings.ToLower(*baseName) + ".go"
    }

    ioutil.WriteFile(outpath, []byte(generated), 0644)
}
