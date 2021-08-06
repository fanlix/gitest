FROM alpine
ENV VVV=gitest-docker \
	WEB_HTTP_ADDR=:8080 \
	WEB_HTTP_DIR=/www
RUN mkdir -p /app  /www
COPY dist/hello /app/
COPY www/* /www/

WORKDIR /app
RUN ls -lh . && md5sum ./hello

ENTRYPOINT ["./hello"]
#CMD [":8080", "/www"]
