package http

import (
    "bytes"
    "errors"
    jsoniter "github.com/json-iterator/go"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "io/ioutil"
    "net/http"
)

func Post(url string, req interface{}, resp interface{}) error {
    var content []byte
    var err error

    if req != nil {
        content, err = jsoniter.Marshal(req)
        if err != nil {
            return err
        }
    }

    r, err := http.Post(url, core.ApplicationJson, bytes.NewReader(content))
    if err != nil {
        return err
    }
    defer func() {
        r.Body.Close()
    }()

    if r.StatusCode >= 400 {
        return errors.New("http return error " + r.Status)
    }

    if resp != nil {
        content, err = ioutil.ReadAll(r.Body)
        if err != nil {
            return err
        }

        if len(content) > 0 {
            if err = jsoniter.Unmarshal(content, resp); err != nil {
                return err
            }
        }
    }

    return nil
}

func Get(url string, resp interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }

    defer func() {
        r.Body.Close()
    }()

    if r.StatusCode >= 400 {
        return errors.New("http return error " + r.Status)
    }

    if resp != nil {
        content, err := ioutil.ReadAll(r.Body)
        if err != nil {
            return err
        }

        if len(content) > 0 {
            if err = jsoniter.Unmarshal(content, resp); err != nil {
                return err
            }
        }
    }

    return nil
}

func Delete(url string) error {
    req, err := http.NewRequest("DELETE", url, nil)
    if err != nil {
        return err
    }
    _, err = http.DefaultClient.Do(req)
    return err
}
