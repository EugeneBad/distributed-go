FROM golang
ARG name
ARG freq
ARG min
ARG max
ARG step
WORKDIR /go/src/app
COPY sensor .
ENTRYPOINT ["./sensor", "--name", "$name", "--freq", "$freq", "--min", "$min", "--max", "$max", "--step", "$step"]