package http

import (
    "bytes"
    "context"
    "errors"
    "fmt"
    jsoniter "github.com/json-iterator/go"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "io"
    "io/ioutil"
    "net/http"
    "strings"
    "time"
)

func (x *Request) Do(ctx context.Context) (*Response, error) {
    req, err := x.SyncTo()
    if err != nil {
        return nil, err
    }
    if req == nil {
        return nil, nil
    }

    timeout := 3 * time.Second
    if toi := ctx.Value("timeout"); toi != nil {
        if to, ok := toi.(time.Duration); ok {
            timeout = to
        }
    }

    ctx, cancel := context.WithTimeout(ctx, timeout)
    defer cancel()

    // Inform remote service to close the connection after the transaction is complete
    req = req.WithContext(ctx)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("making the http request failed: %w", err)
    }
    defer res.Body.Close()

    return NewResponseFrom(res)
}

func (x *Request) SyncTo() (*http.Request, error) {
    if x != nil {
        if x.Url == nil {
            return nil, errors.New("url is empty, must be set")
        }
        if x.Method == 0 {
            return nil, errors.New("method is empty, must be set")
        }

        var body io.Reader
        if x.Body != nil {
            switch x.Body.GetKind() {
            case core.ValueKind_VALUE_KIND_BYTES:
                body = bytes.NewReader(x.Body.GetBytes())
            case core.ValueKind_VALUE_KIND_STRING:
                body = strings.NewReader(x.Body.GetString())
            case core.ValueKind_VALUE_KIND_ARRAY, core.ValueKind_VALUE_KIND_OBJECT:
                content, err := jsoniter.ConfigDefault.Marshal(x.Body)
                if err != nil {
                    return nil, err
                }
                body = bytes.NewReader(content)
            }
        }

        req, err := http.NewRequest(x.Method.Format(), x.Url.Format(), body)
        if err != nil {
            return nil, fmt.Errorf("creating the http request failed: %w", err)
        }

        if x.Headers != nil {
            for k, vs := range x.Headers.Vals {
                if vs == nil {
                    continue
                }
                for _, v := range vs.Vals {
                    req.Header.Add(k, v)
                }
            }
        }
        return req, nil
    }

    return nil, nil
}

func (x *Request) SyncFrom(request *http.Request) (err error) {
    if x != nil && request != nil {
        if request.Body != nil {
            body, err := ioutil.ReadAll(request.Body)
            if err != nil {
                return err
            }
            x.Body = core.NewBytesValue(body)
        }

        if request.URL != nil {
            x.Url, err = core.NewUrl(request.URL.String())
            if err != nil {
                return err
            }
        }

        x.Method, err = ParseMethod(request.Method)
        if err != nil {
            return err
        }

        x.Version = &Version{
            Major: int32(request.ProtoMajor),
            Minor: int32(request.ProtoMinor),
        }

        if len(request.Header) > 0 {
            x.Headers = NewHeadersFrom(request.Header)
        }
    }

    return nil
}
