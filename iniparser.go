package iniparser

import (
  "io/ioutil"
  "regexp"
  "strings"
)

var removeCommentRegexp = regexp.MustCompile(`(#.*)|(;.*)`)

type IniConfig struct {
  config map[string]map[string]string
}

func (ic IniConfig) Get(session string, key string) (string, bool) {
  m, ok := ic.config[session]

  if !ok {
    return "", false
  }

  v, ok := m[key]

  if !ok {
    return "", false
  }

  return v, true
}

func (ic IniConfig) GetDefault(key string) (string, bool) {
  return ic.Get("___default___", key)
}

func Parse(path string) (IniConfig, error) {
  b, err := ioutil.ReadFile(path)

  if err != nil {
    return IniConfig{}, err
  }

  return IniConfig{config: getItems(getLines(&b))}, nil
}

func getLines(contentByte *[]byte) []string {

  content := string(*contentByte)

  var result []string

  content = removeCommentRegexp.ReplaceAllString(content, "\n")
  arr := strings.Split(content, "\n")

  for _, v := range arr {
    if len(v) > 0 {
      result = append(result, strings.Trim(v, " "))
    }
  }

  return result
}

func getItems(lines []string) map[string]map[string]string {

  result := make(map[string]map[string]string)

  currentSession := "___default___"

  result[currentSession] = make(map[string]string)

  for _, line := range lines {
    if strings.HasPrefix(line, "[") {
      currentSession = line[1 : len(line)-1]
      result[currentSession] = make(map[string]string)
    } else {
      temp := strings.Split(line, "=")
      key := strings.Trim(temp[0], " ")
      value := strings.Trim(temp[1], " ")
      result[currentSession][key] = value
    }
  }

  return result
}
