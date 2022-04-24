| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `version` | `mojo.http.Version` |  | 否 |  | The protocol version for incoming server requests. |
| `status` | `mojo.http.Status` |  | 否 |  |  |
| `headers` | `Map<string, Array<string>>` |  | 否 |  | Headers contains the request header fields either receivedby the server or to be sent by the client. |
| `body` | `Value` |  | 否 |  | body is the request's body, which ban be raw bytes or JSON object |
