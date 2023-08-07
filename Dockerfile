FROM alpine:latest

# install git
RUN apk add --no-cache git

COPY tfsecurity /usr/bin/tfsecurity

## use a non-privileged user
RUN adduser -D tfsecurity
USER tfsecurity

# set the default entrypoint -- when this container is run, use this command
ENTRYPOINT [ "tfsecurity" ]
# as we specified an entrypoint, this is appended as an argument (i.e., `tfsecurity --help`)
CMD [ "--help" ]
