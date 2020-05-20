#!/bin/sh

docker build -t antlr .
docker run -it --mount src="$(pwd)",target=/workspace,type=bind antlr
