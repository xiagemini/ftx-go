# ftx-go
ftx go client,use official API,add some more...
# Usage
1. go get github.com/xiagemini/ftx-go
2. import "github.com/xiagemini/ftx-go"
3. client = ftx.New(api_key, secret_key, sub_account)
# Important
Because my network reason,I use proxy when get request,see function ProxyDo in client.go,and you can chang rest.go if you don't want to use proxy.
