##### **Short Description**

* None

##### **Request URL**

* {{base\_url}}/openapi/forum/post/createPost

##### **Request Method**

* POST

##### **Headers**

| Field | Example Value | Required | Type | Description |
| :---- | :---- | :---- | :---- | :---- |
| Api-Key | xxxx | Yes | string | Api-Key for authentication |
| Timestamp | 1111111111 | Yes | string | Timestamp in seconds |
| Signature | xxxxx | Yes | string | Calculated request signature |

##### **Request Parameter Example**

{  
  "contents": \[  
    {  
      "content": "Post content @10044 \#mytag \#posttag",  
      "type": 2,  
      "sort": 100  
    },  
    {  
      "content": "{\\"file\_url\\":\\"https://oss.xxx.link/bbs/public/image/30/9d/ac/19/d211-4d4d-b7fb-eb453f03bf0e.png\\",\\"thumb\_url\\":\\"https://oss.xxxx.link/bbs/public/image/30/9d/ac/19/d211-4d4d-b7fb-eb453f03bf0e\_thumb.png\\",\\"file\_size\\":307062,\\"img\_width\\":320,\\"img\_height\\":320}",  
      "type": 3,  
      "sort": 101  
    },  
    {  
      "content": "{\\"file\_url\\":\\"https://oss.xxxx.link/bbs/public/video/02/2c/f3/ae/37e1-45d6-9f9b-3f734c3fecf4.mp4\\",\\"thumb\_url\\":\\"https://oss.xxxx.link/bbs/public/video/02/2c/f3/ae/37e1-45d6-9f9b-3f734c3fecf4\_thumb.png\\",\\"preview\_url\\":\\"https://oss.happygood.link/bbs/public/video/02/2c/f3/ae/37e1-45d6-9f9b-3f734c3fecf4.png\\",\\"file\_size\\":21137014,\\"img\_width\\":320,\\"img\_height\\":320}",  
      "type": 4,  
      "sort": 102  
    }  
  \],  
  "tags": \[  
    "mytag",  
    "posttag"  
  \],  
  "ip\_loc": "Alice City",  
  "tribe\_id": 1922974687233  
}

##### **Request JSON Field Descriptions**

| Field | Required | Type | Description |
| :---- | :---- | :---- | :---- |
| contents | Yes | array | List of post content blocks |
| contents.content | Yes | string | The content of the block |
| contents.type | Yes | int | Content type: 2 for text paragraph, 3 for image URL, 4 for video URL |
| contents.sort | Yes | int | Sort order |
| tags | Yes | array | List of topics/tags |
| ip\_loc | Yes | string | IP-based location. If empty, it will not be displayed. |
| tribe\_id | Yes | int | ID of the paid community |

##### **Success Response Example**

{  
  "code": 0,  
  "message": "success",  
  "data": {  
    "id": 1080018646  
  }  
}

##### **Success Response Parameter Descriptions**

| Parameter | Type | Description |
| :---- | :---- | :---- |
| code | int | Status code |
| msg | string | Message |
| data | object | Response data |
| id | long | The ID of the created post |

##### **Failure Response Example**

{  
  "code": 30014,  
  "message": "Error code: 30014, Error message: Failed to get post list with unknown style parameter"  
}

##### **Failure Response Parameter Descriptions**

| Parameter | Type | Description |
| :---- | :---- | :---- |
| code | int | Error code |
| msg | string | Error message |

##### **Remarks**

* To mention a user, use @username. Regex for extraction: @(\\d{5,10})\\s  
* To use a hashtag, use \#topic. Regex for extraction: \#(\[^\\s\#@\]{2,30})\\s  
* **Note**: It is recommended to add a space at the end of the content to ensure the last @mention or \#hashtag is matched correctly.  
* If the error code is 10002, a "Post not found" interface should be displayed.  
* **Current limitations**: contents.type 2 (text) is limited to one block. contents.type 3 (image) is limited to a maximum of 9 blocks. contents.type 4 (video) is limited to one block. Images and videos cannot be mixed in the same post.  
* When contents.type is 3 or 4, the contents.content value should be the stringified data.content from the response of the /openapi/forum/upload/attachment endpoint.