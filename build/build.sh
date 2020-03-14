GOOS=linux GOARCH=mipsle go build -o omicron.arm.exe && rsync -zzarvh omicron.arm.exe root@omega-92e0:/data && ssh root@omega-92e0.local "killall omicron.arm.exe"
