<h1>trial</h1>
<br>
<h1>cmd</h1>
<p>1. go env

2. GOOS="windows" go build
3. GOOS="linux" go build
</p>
<br>
<p>Memory Management
new()
allocate memory but no INIT
you will get a memory address zeroed storage

make()
Allocate memory and INIT
you will get a memory address non-zeroed storage
GC happens automatically
out of scope or nil
https://pkg.go.dev/runtime

<p>
