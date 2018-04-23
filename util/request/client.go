package request

import (
    "bytes"
    "encoding/json"
    "io"
    "io/ioutil"
    "mime/multipart"
    "net/http"
    "os"
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

func Upload(uri string, filename string) ([]byte, error) {

    bodyBuf := &bytes.Buffer{}
    bodyWrite := multipart.NewWriter(bodyBuf)

    fileWrite, err := bodyWrite.CreateFormFile("media", filename)

    if err != nil {
        return nil, err
    }

    fh, err := os.Open(filename)
    defer fh.Close()

    if err != nil {
        return nil, err
    }

    _, err = io.Copy(fileWrite, fh)

    if err != nil {
        return nil, err
    }

    contentType := bodyWrite.FormDataContentType()
    bodyWrite.Close()

    response, err := http.Post(uri, contentType, bodyBuf)

    defer response.Body.Close()

    if err != nil {
        return nil, err
    }

    return ioutil.ReadAll(response.Body)
}