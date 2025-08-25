package main

type ForumPostCreatePostReq struct {
	Contents []ForumPostContentItem `json:"contents"`
	Tags     []string               `json:"tags"`
	IpLoc    string                 `json:"ip_loc"`
	TribeId  int64                  `json:"tribe_id"`
}

type ForumPostContentItem struct {
	Content string      `json:"content"`
	Type    ContentType `json:"type"`
	Sort    int         `json:"sort"`
}

type ForumPostCreatePostResp struct {
	Id int `json:"id"` //postId
}

// go run method.go request.go util.go Forum_post_createPost.go
func main() {
	req := new(ForumPostCreatePostReq)
	req.Tags = []string{}
	req.IpLoc = "location"
	req.TribeId = 2053949045761

	req.Contents = make([]ForumPostContentItem, 0, 10)
	//todo:optional
	{
		req.Contents = append(req.Contents, ForumPostContentItem{
			Content: "demo text",
			Type:    TextContentType,
			Sort:    10,
		})
	}
	//todo:optional
	//{
	//	resp, err := ForumUploadAttachment("./image.demo.png", "", 0, TribeUploadGroup, false)
	//	if err != nil {
	//		panic(err)
	//	}
	//	req.Contents = append(req.Contents, ForumPostContentItem{
	//		Content: resp.Content,
	//		Type:    ImageContentType,
	//		Sort:    20,
	//	})
	//}
	//todo:optional
	//{
	//	resp, err := ForumUploadAttachment("./video.demo.mp4", "./video.thumb.png", 120, TribeUploadGroup, true)
	//	if err != nil {
	//		panic(err)
	//	}
	//	req.Contents = append(req.Contents, ForumPostContentItem{
	//		Content: resp.Content,
	//		Type:    VideoContentType,
	//		Sort:    30,
	//	})
	//}
	apiPath := "/openapi/forum/post/createPost"
	method := "POST"
	resp := new(ForumPostCreatePostResp)
	err := JsonRequest(apiPath, method, req, resp)
	PrintResp(resp, err)
}
