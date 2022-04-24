| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `method` | `string` |  | 否 |  | method specifies the HTTP method (GET, POST, PUT, etc.).For client requests, an empty string means GET. |
| `url` | `string` | `Url` | 否 |  | url specifies either the URI being requested (for serverrequests) or the URL to access (for client requests). |
| `version` | `mojo.http.Version` |  | 否 |  | The protocol version for incoming server requests. |
| `headers` | `Map<string, Array<string>>` |  | 否 |  | Headers contains the request header fields either receivedby the server or to be sent by the client. |
| `body` | `Value` |  | 否 |  | body is the request's body, which ban be raw bytes or JSON object |
