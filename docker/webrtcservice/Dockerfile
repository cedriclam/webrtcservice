FROM            alpine
MAINTAINER      Cedric Lamoriniere <cedric.lamoriniere@gmail.com>
# Add the binary. As it is statically linked, no need to add libc or anything else.
ADD             webrtcservice /
# Open server port
EXPOSE 8080
# Set the entrypoint as the binary, so `docker run <image>` will behave as the binary
ENTRYPOINT      ["/webrtcservice"]
