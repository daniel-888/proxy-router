

echo "Running: env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o lumerin_amd64/lumerin"

cd /home/sean/Titan/src/lumerin/cmd

env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o lumerin_amd64/lumerin

