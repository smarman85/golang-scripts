package main

import (
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "log"
)

var GroupMap = map[string]string{
  "techops": "0001",
  "rails":   "0002",
}

type SdmGroups struct {
  Groups map[string][]map[string]string `yaml:"sdm_groups"`
}

func (g *SdmGroups) getGroups() *SdmGroups {
  yamlFile, err := ioutil.ReadFile("users.yaml")
  if err != nil {
    log.Printf("yamlFile.Get err #%v ", err)
  }
  err = yaml.Unmarshal(yamlFile, g)
  if err != nil {
    log.Fatalf("Unmarshal: %v", err)
  }
  return g
}



func main() {
  var g SdmGroups
  g.getGroups()

  fmt.Println(g)
  for key, val := range g.Groups {
    fmt.Println("Key: ", key, " Value: ", val)
    fmt.Println(GroupMap[key])
  }
}
