FROM ubuntu:20.04

COPY goreleaser-live /

ENTRYPOINT [ "/goreleaser-live" ]
