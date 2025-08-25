##### **Short Description**

* None

##### **Request URL**

* {{base\_url}}/openapi/forum/upload/attachment

##### **Request Method**

* POST

##### **Headers**

| Field | Example Value | Required | Type | Description |
| :---- | :---- | :---- | :---- | :---- |
| Api-Key | xxxx | Yes | string | Api-Key for authentication |
| Timestamp | 1111111111 | Yes | string | Timestamp in seconds |
| Signature | xxxxx | Yes | string | Calculated request signature |

##### **Request Body Parameters (multipart/form-data)**

| Parameter | Example Value | Required | Type | Description |
| :---- | :---- | :---- | :---- | :---- |
| type | public/video | Yes | string | public/image for images, public/video for videos |
| file | \[object Object\] | Yes | file | The file to upload |
| thumb | \[object Object\] | Yes (for video) | file | Video thumbnail preview image |
| duration | 1 | Yes (for video) | long | Video duration in seconds |
| group\_name | tribe | Yes | string | Group name: tribe for community posts, tribeHub for community avatars, post for general posts, tag for topic avatars, report for report materials, media for media resources |

##### **Success Response Example**

{  
  "code": 0,  
  "message": "success",  
  "data": {  
    "user\_id": 10026,  
    "file\_size": 21137014,  
    "img\_width": 320,  
    "img\_height": 320,  
    "type": 2,  
    "duration": 20,  
    "content": "{\\"file\_url\\":\\"https://oss.domain.link/bbs/public/video/02/2c/f3/ae/37e1-45d6-9f9b-3f734c3fecf4.mp4\\",\\"thumb\_url\\":\\"https://oss.domain.link/bbs/public/video/02/2c/f3/ae/37e1-45d6-9f9b-3f734c3fecf4\_thumb.png\\",\\"preview\_url\\":\\"https://oss.domain.link/bbs/public/video/02/2c/f3/ae/37e1-45d6-9f9b-3f734c3fecf4.png\\",\\"file\_size\\":21137014,\\"img\_width\\":320,\\"img\_height\\":320,\\"duration\\": 20}",  
    "thumb\_url": "https://oss.domain.link/bbs/public/video/02/2c/f3/ae/37e1-45d6-9f9b-3f734c3fecf4\_thumb.png",  
    "file\_url": "https://oss.domain.link/bbs/public/video/02/2c/f3/ae/37e1-45d6-9f9b-3f734c3fecf4.mp4"  
  }  
}

##### **Success Response Parameter Descriptions**

| Parameter | Type | Description |
| :---- | :---- | :---- |
| code | int | Status code |
| message | string | Message |
| data | object | Response data |
| data.user\_id | long | User ID |
| data.file\_size | int | File size of the original image/video |
| data.img\_width | int | Width in pixels of the original image/video |
| data.img\_height | int | Height in pixels of the original image/video |
| data.type | int | File type: 1 for image, 2 for video |
| data.content | string | The content string to be used when creating a post/comment |
| data.content.file\_url | string | URL of the uploaded original file (image/video) |
| data.content.thumb\_url | string | URL of the preview thumbnail |
| data.content.preview\_url | string | URL of the original video preview image (video only) |
| data.content.file\_size | int | File size |
| data.content.img\_width | int | Image width |
| data.content.img\_height | int | Image height |
| data.content.is\_lock | int | Indicates if user has permission to view: 1 for no permission (paid), 0 for permission |
| data.thumb\_url | string | URL of the preview thumbnail |
| data.file\_url | string | URL of the uploaded file |

##### **Failure Response Example**

{  
  "code": 10001,  
  "message": "Error code: 10001, Error message: Invalid input parameters"  
}

##### **Failure Response Parameter Descriptions**

| Parameter | Type | Description |
| :---- | :---- | :---- |
| code | int | Error code |
| msg | string | Error message |

##### **Remarks**

* **Note**: The host address for file uploads should be obtained from the /openapi/forum/conf/serverList endpoint. Choose one randomly from the list.  
* **Note**: A video will have three URLs: the original video, an original preview image, and a thumbnail of the preview image.  
* **Note**: An image will have two URLs: the original image file and a thumbnail. The preview URL will be empty.  
* **Note**: The thumbnail URL is intended for use in list views.