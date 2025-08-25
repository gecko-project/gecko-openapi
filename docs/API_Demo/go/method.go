package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

type ContentType int //2:text，3:image，4:video

const (
	TextContentType  ContentType = 2
	ImageContentType ContentType = 3
	VideoContentType ContentType = 4
)

type UploadGroup string

const (
	TribeUploadGroup    UploadGroup = "tribe"    //Tribe Community Post
	TribeHubUploadGroup UploadGroup = "tribeHub" //Tribe Community Avatar
	PostUploadGroup     UploadGroup = "post"     //Plaza Post
	TagUploadGroup      UploadGroup = "tag"      //Topic Avatar
	ReportUploadGroup   UploadGroup = "report"   //Report Material
	MediaUploadGroup    UploadGroup = "media"    //Media Resource
)

type ForumUploadAttachmentResp struct {
	UserId    int64       `json:"user_id"`
	FileSize  int64       `json:"file_size"`
	ImgWidth  int64       `json:"img_width"`
	ImgHeight int64       `json:"img_height"`
	Type      ContentType `json:"type"`
	Duration  int64       `json:"duration"` //second
	Content   string      `json:"content"`
	ThumbUrl  string      `json:"thumb_url"`
	FileUrl   string      `json:"file_url"`
}

func ForumUploadAttachment(filePath, thumbPath string, duration int, groupName UploadGroup, isVideo bool) (*ForumUploadAttachmentResp, error) {
	payload := new(bytes.Buffer)
	writer := multipart.NewWriter(payload)

	fh, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	part, err := WriterFormFile(writer, "file", filepath.Base(filePath))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, fh)
	if err != nil {
		return nil, err
	}

	err = writer.WriteField("group_name", string(groupName))
	if err != nil {
		return nil, err
	}

	if isVideo {
		err = writer.WriteField("type", "public/video")
		if err != nil {
			return nil, err
		}
		if duration <= 0 {
			return nil, fmt.Errorf("duration is zero")
		}
		err = writer.WriteField("duration", strconv.Itoa(duration))
		if err != nil {
			return nil, err
		}

		fh, err = os.Open(thumbPath)
		if err != nil {
			return nil, err
		}
		defer fh.Close()

		part, err = WriterFormFile(writer, "thumb", filepath.Base(thumbPath))
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, fh)
		if err != nil {
			return nil, err
		}
	} else {
		err = writer.WriteField("type", "public/image")
		if err != nil {
			return nil, err
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	apiPath := "/openapi/forum/upload/attachment"
	method := "POST"
	resp := new(ForumUploadAttachmentResp)
	err = SendRequest(apiPath, method, writer.FormDataContentType(), payload, resp)
	return resp, err
}
