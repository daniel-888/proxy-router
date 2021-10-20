# syntax=docker/dockerfile:1
FROM scratch
ADD lumerin_amd64 /
ADD lumerinconfig.json /
CMD ["/lumerin"]
