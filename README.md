<h1 align="center" style="display: flex; align-items: center; justify-content: center;"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/800px-Go_Logo_Blue.svg.png" width="80" style="margin-right: 10px"> API-PRACTICE</h1>

[![License](https://img.shields.io/badge/license-MIT-green)](https://github.com/arkivanov/MVIKotlin/blob/master/LICENSE)

### Functional

* ***CRUD news***.
* ***CD comment***.

### Endpoint's

<table>
<tr><td>HTTP</td><td> URL</td><td>BODY</td><td>RESPONSE</td></tr>
<tr><td>GET</td><td>/api/v1/news</td>
<td>
-
</td>
<td>

```json
[
  {
    "id": 3,
    "title": "Title",
    "description": "Desc",
    "comments": null
  },
  {
    "id": 4,
    "title": "kek",
    "description": "Desc",
    "comments": null
  }
]
```

</td>
</tr>
<tr><td>POST</td><td>/api/v1/news</td>
<td>

```json
{
  "title": "Test"
}
```

</td>
<td>

```json
{
  "id": 5,
  "title": "Test",
  "description": "",
  "comments": null
}
```

</td></tr>
<tr><td>GET</td><td>/api/v1/news/{id}</td><td>-</td>
<td>

```json
{
  "id": 4,
  "title": "kek",
  "description": "Desc",
  "comments": [
    {
      "id": 13,
      "text": "Text",
      "news_id": 4
    }
  ]
}
```

</td></tr>
<tr><td>DELETE</td><td>/api/v1/news/{id}</td><td>-</td>

<td>

```json
{
  "id": 5,
  "title": "Test",
  "description": "",
  "comments": null
}
```

</td></tr>
<tr><td>PUT</td><td>/api/v1/news/{id}</td>

<td>

```json
{
  "title": "Test"
}
```

</td>
<td>

```json
{
  "id": 4,
  "title": "Test",
  "description": "Desc",
  "comments": null
}
```

</td></tr>
<tr><td>POST</td><td>/api/v1/comments</td>

<td>

```json
{
  "text": "Text",
  "news_id": 4
}
```

</td>

<td>

```json
{
  "id": 14,
  "text": "Text",
  "news_id": 4
}
```

</td></tr>
<tr><td>DELETE</td><td>/api/v1/comments/{id}</td><td>-</td>
<td>

```json
"Comment Deleted Successfully!"
```

</td></tr>
</table>

### Technologies

* Language: **Go**
* Libraries: **Viper, GORM, Mux**
* Database: **MySQL**

## Installing

***Firstly*** clone the project.

```git
git clone https://github.com/l1ve4code/go-news-api
```

***Secondly*** run project.

```go
go run .
```


## Author

* Telegram: **[@live4code](https://t.me/live4code)**
* Email: **steven.marelly@gmail.com**
