package request

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func Get(uri string) ([]byte, error) {

    response, err := http.Get(uri)

    defer response.Body.Close()

    if err != nil {
        return nil, err
    }

    return ioutil.ReadAll(response.Body)
}

func Post(uri string, obj interface{}) ([]byte, error) {

    jsonData, err := json.Marshal(obj)

    if err != nil {
        return nil, err
    }

    response, err := http.Post(uri, "application/json", bytes.NewBuffer(jsonData))

    defer response.Body.Close()

    if err != nil {
        return nil, err
    }

    return ioutil.ReadAll(response.Body)
}
