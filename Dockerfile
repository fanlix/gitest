FROM alpine
RUN mkdir -p /data
WORKDIR /app
ENV VVV gitest-docker
COPY dist/hello .
RUN ls -lh . && md5sum ./hello

ENTRYPOINT ["hello"]
CMD ["abc"]
