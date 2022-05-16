package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v3"
)

type KubeConfig struct {
	CurrentContext string    `yaml:"current-context"`
	Contexts       []Context `yaml:"contexts"`
}

type Context struct {
	Name   string
	Active bool
	Data   ContextData `yaml:"context"`
}

type ContextData struct {
	Cluster   string
	Namespace string
	User      string
}

func normalizeCluster(cluster string) string {
	cluster = strings.TrimPrefix(cluster, "api-")
	cluster = strings.TrimSuffix(cluster, "-p1-openshiftapps-com:6443")
	return cluster
}

func contextInfo(c *Context) string {
	// star := ""
	// if c.Active {
	// 	star = "*"
	// }
	return fmt.Sprintf("%s/%s", normalizeCluster(c.Data.Cluster), c.Data.Namespace)
}

func main() {
	var kc KubeConfig
	kubefile := "/Users/jmelis/.kube/config"

	yamlFile, err := ioutil.ReadFile(kubefile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &kc)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, c := range kc.Contexts {
		if kc.CurrentContext == c.Name {
			c.Active = true
			fmt.Println(contextInfo(&c))
		}
	}
}
