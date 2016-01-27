# dropbox4go
dropbox4go is a Dropbox API binding library for the Go language.

# Usage
```
token = "your access token of Dropbox"
client := http.DefaultClient
svc := dropbox4go.New(client, token)

url := "http://example.com/files/sample.jpg"
ext := path.Ext(url)
switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
	default:
		panic("incorrect extension.")
}

// if you want to upload by URL.
resp, err := client.Get(url)
defer resp.Body.Close()
file := resp.Body

req := dropbox4go.Request{
	File: file,
	Parameters: dropbox4go.Parameters{
		Path: "/home/test" + ext,
		Mode: "overwrite",
		AutoRename: false,
		ClientModified: time.Now().UTC().Format(time.RFC3339),
		Mute: true,
	},
}

fmt.Println("Now uploading...")
result, err := svc.Upload(req)

if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result.Name)
}
```
