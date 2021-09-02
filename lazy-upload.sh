go build my-app.go
scp -r $PWD/* foundry:burgh-app
rm my-app
