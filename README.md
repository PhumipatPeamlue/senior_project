# Docker-compose usage
## Start container stack
`docker-compose up -d --build`
## Delete container stack
`docker-compose down`

# REST API
## Video Document
### Search video doc (no keyword)
`GET http://localhost:8080/video_doc/search?page=[NUMBER]&page_size=[NUMBER]`

`http://localhost:8080/video_doc/search?page=1&page_size=5`
```
{
    "data": [
        {
            "_index": "video_doc",
            "_type": "",
            "_id": "xGNEiowBt749NU1AfZR5",
            "_score": 1,
            "_source": {
                "title": "title_a",
                "video_url": "video_a",
                "description": "description_a"
            }
        },
        {
            "_index": "video_doc",
            "_type": "",
            "_id": "6QtIiowBEJHG34akJh-R",
            "_score": 1,
            "_source": {
                "title": "title_b",
                "video_url": "video_b",
                "description": "description_b"
            }
        }
    ],
    "total": 2
}
```
### Search video doc (with keyword)
`GET http://localhost:8080/video_doc/search?page=[NUMBER]&page_size=[NUMBER]&keyword=[TEXT]`

`http://localhost:8080/video_doc/search?page=1&page_size=5&keyword=description_b`
```
{
    "data": [
        {
            "_index": "video_doc",
            "_type": "",
            "_id": "6QtIiowBEJHG34akJh-R",
            "_score": 0.9808291,
            "_source": {
                "title": "title_b",
                "video_url": "video_b",
                "description": "description_b"
            }
        }
    ],
    "total": 1
}
```
### Get video doc
`http://localhost:8080/video_doc/[DOCUMENT's ID]`

`http://localhost:8080/video_doc/6QtIiowBEJHG34akJh-R`

```
{
    "id": "6QtIiowBEJHG34akJh-R",
    "title": "title_b",
    "video_url": "video_b",
    "description": "description_b"
}
```
### Insert video doc
`POST http://localhost:8080/video_doc` + `json's body`
```
http://localhost:8080/video_doc

{
    "title": "title_a",
    "video_url": "video_a",
    "description": "description_a"
}
```
```
{
"message": "add the new video document sucessfully"
}
```
### Update video doc
`PUT http://localhost:8080/video_doc` + `json's body`
```
http://localhost:8080/video_doc

{
    "id": "6QtIiowBEJHG34akJh-R",
    "title": "title_b_updated",
    "video_url": "video_b_updated",
    "description": "description_b_updated"
}
```
```
{
    "message": "update video document successfully"
}
```
### Delete video doc
`DELETE http://localhost:8080/video_doc/[VIDEO_DOC_ID]`

`http://localhost:8080/video_doc/6QtIiowBEJHG34akJh-R`
```
{
    "message": "delete video document sucessfully"
}
```

## Drug document
### Search drug doc (no keyword)
`GET http://localhost:8080/drug_doc/search?page=[NUMBER]&page_size=[NUMBER]`

`http://localhost:8080/drug_doc/search?page=1&page_size=5`
```
{
    "data": [
        {
            "_index": "drug_doc",
            "_type": "",
            "_id": "6wtMiowBEJHG34aksR_j",
            "_score": 1,
            "_source": {
                "trade_name": "trade_name_a",
                "drug_name": "drug_name_a",
                "description": "description_a",
                "preparation": "type 1",
                "caution": "caution_a"
            }
        },
        {
            "_index": "drug_doc",
            "_type": "",
            "_id": "7AtMiowBEJHG34ak2B9X",
            "_score": 1,
            "_source": {
                "trade_name": "trade_name_b",
                "drug_name": "drug_name_b",
                "description": "description_b",
                "preparation": "type 2",
                "caution": "caution_b"
            }
        }
    ],
    "total": 2
}
```
### Search drug doc (with keyword)
`GET http://localhost:8080/drug_doc/search?page=[NUMBER]&page_size=[NUMBER]&keyword=[TEXT]`

`http://localhost:8080/drug_doc/search?page=1&page_size=5&keyword=drug_name_a`
```
{
    "data": [
        {
            "_index": "drug_doc",
            "_type": "",
            "_id": "6wtMiowBEJHG34aksR_j",
            "_score": 0.6931471,
            "_source": {
                "trade_name": "trade_name_a",
                "drug_name": "drug_name_a",
                "description": "description_a",
                "preparation": "type 1",
                "caution": "caution_a"
            }
        }
    ],
    "total": 1
}
```
### Get drug doc
`http://localhost:8080/drug_doc/[DOCUMENT's ID]`

`http://localhost:8080/drug_doc/KIsSiowB9N475D44mK39`

```
{
    "id": "6wtMiowBEJHG34aksR_j",
    "trade_name": "trade_name_a",
    "drug_name": "drug_name_a",
    "description": "description_a",
    "preparation": "type 1",
    "caution": "caution_a"
}
```
### Insert drug doc
`POST http://localhost:8080/drug_doc` + `json's body`
```
http://localhost:8080/drug_doc

{
    "trade_name": "trade_name_b",
    "drug_name": "drug_name_b",
    "description": "description_b",
    "preparation": "type 2",
    "caution": "caution_b"
}
```
```
{
"message": "add new drug document sucessfully"
}
```
### Update drug doc
`PUT http://localhost:8080/drug_doc` + `json's body`
```
http://localhost:8080/drug_doc

{
    "id": "6wtMiowBEJHG34aksR_j",
    "trade_name": "trade_name_a_updated",
    "drug_name": "drug_name_a_updated",
    "description": "description_a_updated",
    "preparation": "type 2",
    "caution": "caution_a_updated"
}
```
```
{
    "message": "update drug document successfully"
}
```
### Delete drug doc
`DELETE http://localhost:8080/drug_doc/[DRUG_DOC_ID]`

`http://localhost:8080/drug_doc/6wtMiowBEJHG34aksR_j`
```
{
    "message": "delete drug document sucessfully"
}
```
## Image
### Get Image
`GET http://localhost:8080/image/[FILENAME]`

`http://localhost:8080/image/photo1.png`
### Get all the image'paths by document ID
`GET http://localhost:8080/image/[DOCUMENT_ID]`

`http://localhost:8080/image/paths/6wtMiowBEJHG34aksR_j`
```
{
    "data": [
        {
            "id": 3,
            "doc_id": "6wtMiowBEJHG34aksR_j",
            "filename": "Screenshot 2566-12-12 at 19.26.15.png",
            "file_path": "uploads/Screenshot 2566-12-12 at 19.26.15.png"
        }
    ]
}
```
### Insert image
`POST http://localhost:8080/image/[DOCUMENT_ID]` + `form-data's body`

```
{
    "message": "insert the image successfully"
}
```
### Delete image
`DELETE http://localhost:8080/image/[IMAGE_ID]`

`http://localhost:8080/image/3`
```
{
    "message": "delete image successfully"
}
```