package request

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
)

func Get(uri string) ([]byte, error) {

    res, err := http.Get(uri)

    defer  res.Body.Close()

    if err != nil {
        return nil, err
    }

    return ioutil.ReadAll(res.Body)
}

func Post(uri string, obj interface{}) ([]byte, error) {

    jsonData, err := json.Marshal(obj)

    if err != nil {
        return nil, err
    }

    res, err := http.Post(uri, "application/json", bytes.NewBuffer(jsonData))

    defer res.Body.Close()

    if err != nil {
        return nil, err
    }

    return ioutil.ReadAll(res.Body)
}