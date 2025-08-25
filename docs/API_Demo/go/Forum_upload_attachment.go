package main

// go run method.go request.go util.go Forum_upload_attachment.go
func main() {
	//todo:optional
	{
		//demo:image
		resp, err := ForumUploadAttachment("./image.demo.png", "", 0, TribeUploadGroup, false)
		PrintResp(resp, err)
	}
	//todo:optional
	{
		//demo:video
		resp, err := ForumUploadAttachment("./video.demo.mp4", "./video.thumb.png", 120, TribeUploadGroup, true)
		PrintResp(resp, err)
	}
}
