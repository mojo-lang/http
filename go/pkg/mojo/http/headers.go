package http

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "net/http"
    "net/textproto"
)

const HeaderTypeName = "Headers"
const HeaderTypeFullName = "mojo.http.Headers"

func NewHeaders() *Headers {
    return &Headers{
        Vals: make(map[string]*core.StringValues),
    }
}

func NewHeadersFrom(header http.Header) *Headers {
    return NewHeaders().SyncFrom(header)
}

// Add adds the key, value pair to the header.
// It appends to any existing values associated with key.
// The key is case insensitive; it is canonicalized by
// CanonicalHeaderKey.
func (x *Headers) Add(key, value string) *Headers {
    key = textproto.CanonicalMIMEHeaderKey(key)
    if v, ok := x.Vals[key]; ok {
        v.Append(value)
    } else {
        x.Vals[key] = core.NewStringValues(value)
    }

    return x
}

// Set sets the header entries associated with key to the
// single element value. It replaces any existing values
// associated with key. The key is case insensitive; it is
// canonicalized by textproto.CanonicalMIMEHeaderKey.
// To use non-canonical keys, assign to the map directly.
func (x *Headers) Set(key, value string) *Headers {
    key = textproto.CanonicalMIMEHeaderKey(key)
    x.Vals[key] = core.NewStringValues(value)
    return x
}

// Get gets the first value associated with the given key. If
// there are no values associated with the key, Get returns "".
// It is case insensitive; textproto.CanonicalMIMEHeaderKey is
// used to canonicalize the provided key. To use non-canonical keys,
// access the map directly.
func (x *Headers) Get(key string) string {
    key = textproto.CanonicalMIMEHeaderKey(key)
    if v, ok := x.Vals[key]; ok {
        values := v.GetVals()
        if len(values) > 0 {
            return values[0]
        }
    }
    return ""
}

// Values returns all values associated with the given key.
// It is case insensitive; textproto.CanonicalMIMEHeaderKey is
// used to canonicalize the provided key. To use non-canonical
// keys, access the map directly.
// The returned slice is not a copy.
func (x *Headers) Values(key string) []string {
    key = textproto.CanonicalMIMEHeaderKey(key)
    if v, ok := x.Vals[key]; ok {
        return v.GetVals()
    }
    return nil
}

// Del deletes the values associated with key.
// The key is case insensitive; it is canonicalized by
// CanonicalHeaderKey.
func (x *Headers) Del(key string) *Headers {
    key = textproto.CanonicalMIMEHeaderKey(key)
    delete(x.Vals, key)
    return x
}

func (x *Headers) AddCookies(cookies ...*Cookie) *Headers {
    if x != nil {
        for _, c := range cookies {
            x.Add("Cookie", c.ToString())
        }
    }
    return x
}

func (x *Headers) AddSetCookies(cookies ...*Cookie) *Headers {
    if x != nil {
        for _, c := range cookies {
            x.Add("Set-Cookie", c.ToString())
        }
    }
    return x
}

func (x *Headers) GetCookies() []*Cookie {
    return x.getCookies("Cookie")
}

func (x *Headers) GetSetCookies() []*Cookie {
    return x.getCookies("Set-Cookie")
}

func (x *Headers) getCookies(key string) []*Cookie {
    if x != nil {
        values := x.Values(key)
        if len(values) > 0 {
            response := http.Response{Header: map[string][]string{}}
            response.Header["Set-Cookie"] = values
            cs := response.Cookies()
            var cookies []*Cookie
            for _, c := range cs {
                cookies = append(cookies, FromHttpCookie(c))
            }
            return cookies
        }
    }
    return nil
}

func (x *Headers) SyncTo(headers http.Header) *Headers {
    if x != nil {
        for key, values := range x.Vals {
            if values != nil {
                for _, value := range values.Vals {
                    headers.Set(key, value)
                }
            }
        }
    }
    return x
}

func (x *Headers) SyncFrom(headers http.Header) *Headers {
    if x != nil {
        for key, values := range headers {
            if values != nil {
                for _, value := range values {
                    x.Add(key, value)
                }
            }
        }
    }
    return x
}
