package http

// Cookies parses and returns the cookies set in the Set-Cookie headers.
func (x *Response) Cookies() []*Cookie {
	if x != nil {
		x.Headers.GetSetCookies()
	}
	return nil
}
