FROM scratch
COPY app /
EXPOSE 8080
ENTRYPOINT ["/app"]
