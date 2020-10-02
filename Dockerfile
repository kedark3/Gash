FROM registry.svc.ci.openshift.org/openshift/release:golang-1.10 as builder
RUN go get github.com/kedark3/gash
WORKDIR /go/src/github.com/kedark3/gash/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/gash gash.go

FROM busybox
LABEL author="Kedar Kulkarni"

WORKDIR /bin
COPY --from=builder /go/bin/gash /bin/gash
COPY --from=builder /go/bin/gash /bin/cp
COPY --from=builder /go/bin/gash /bin/cd
COPY --from=builder /go/bin/gash /bin/mkdir
COPY --from=builder /go/bin/gash /bin/rmdir
COPY --from=builder /go/bin/gash /bin/pwd
COPY --from=builder /go/bin/gash /bin/mv
COPY --from=builder /go/bin/gash /bin/rm
COPY --from=builder /go/bin/gash /bin/touch
COPY --from=builder /go/bin/gash /bin/ls
COPY --from=builder /go/bin/gash /bin/cat

CMD ls