* statik -src=./html
* go run main.go

# windows
* if blank screen run `CheckNetIsolation.exe LoopbackExempt -a -n="Microsoft.Win32WebViewHost_cw5n1h2txyewy"`
* build on windows `go build -ldflags="-H windowsgui" -o webview-example.exe`
* dll's must be rebuilt for `https://github.com/webview/webview` on local machine and dlls in same folder as exe
* call go func from web `window.goCode("ping").then( (val) => console.log("asynchronous logging has val:",val) );`