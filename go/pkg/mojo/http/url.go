package http

import "github.com/mojo-lang/core/go/pkg/mojo/core"

func UnmarshalPathParam(params map[string]string, value interface{}, keys ...string) error {
    val := ""
    for _, key := range keys {
        if val, _ = params[key]; len(val) > 0 {
            break
        }
    }

    if len(val) == 0 {
        return nil
    }

    return core.UnmarshalParam(val, value)
}

func UnmarshalQueryParam(params *core.Url_Query, value interface{}, keys ...string) error {
    key := ""
    for _, k := range keys {
        if params.Has(k) {
            key = k
        }
    }
    if len(key) == 0 {
        return nil
    }
    return params.Unmarshal(key, value)
}
