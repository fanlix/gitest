FROM alpine
ENV VVV gitest-docker
RUN mkdir -p /app
COPY dist/hello /app/
WORKDIR /app
RUN ls -lh . && md5sum ./hello

ENTRYPOINT ["./hello"]
CMD ["abc"]
