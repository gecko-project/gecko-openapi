package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	ServerHost = "https://tribeopen.alicdn4.com"
	ApiKey     = "8WWG4ge8FTnE29QK9RHomiD4LeThGeMI"
	ApiSecret  = "AI22I2cm4m4Si3I9z8uByUzVJmvZs2l9"
)

func Md5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func Hmac256(key, s string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(s))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func JsonRequest(apiPath string, method string, param interface{}, out interface{}) error {
	if param == nil {
		return SendRequest(apiPath, method, "application/json", nil, out)
	}
	data, err := json.Marshal(param)
	if err != nil {
		return err
	}
	return SendRequest(apiPath, method, "application/json", bytes.NewReader(data), out)
}

func SendRequest(apiPath string, method string, contentType string, body io.Reader, out interface{}) error {
	uri := ServerHost + apiPath

	payload := ""
	if contentType == "application/json" && body != nil {
		data, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		payload = string(data)
		body = io.NopCloser(bytes.NewBuffer(data))
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return err
	}
	nowTime := strconv.FormatInt(time.Now().Unix(), 10)
	req.Header.Add("Accept-Language", "en-us") //en-us,zh-cn
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Api-Key", ApiKey)
	req.Header.Add("Timestamp", nowTime)

	message := Md5(nowTime + ":" + method + ":" + apiPath + ":" + payload)
	signature := Hmac256(ApiSecret, message)
	req.Header.Add("Signature", signature)

	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request fail,status=%v", resp.Status)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	result := new(BaseResp)
	if err = json.Unmarshal(respBody, result); err != nil {
		return err
	}
	if err = result.Err(); err != nil {
		return err
	}
	if len(result.Data) == 0 || out == nil {
		return nil
	}
	if err = json.Unmarshal(result.Data, out); err != nil {
		return err
	}
	return nil
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

var extToMimeType = map[string]string{
	".png":  "image/png",
	".jpg":  "image/jpg",
	".jpeg": "image/jpeg",
	".gif":  "image/gif",
	".mp4":  "video/mp4",
	".mov":  "video/quicktime",
	".webp": "image/webp",
	".zip":  "application/zip",
}

func getFileMimeByExt(filename string) (string, error) {
	ext := strings.ToLower(filepath.Ext(filename))
	mimeType, ok := extToMimeType[ext]
	if ok == false {
		return "", fmt.Errorf("unsupported %s ext", ext)
	}
	return mimeType, nil
}

func WriterFormFile(writer *multipart.Writer, fieldname, filename string) (io.Writer, error) {
	mimeType, err := getFileMimeByExt(filename)
	if err != nil {
		return nil, err
	}
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, escapeQuotes(fieldname), escapeQuotes(filename)))
	h.Set("Content-Type", mimeType)
	return writer.CreatePart(h)
}

type BaseResp struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func (this *BaseResp) Err() error {
	if this.Code != 0 {
		return fmt.Errorf("http error :code=%d，msg=%s", this.Code, this.Message)
	}
	return nil
}
