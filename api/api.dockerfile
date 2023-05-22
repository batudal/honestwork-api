FROM alpine:latest
RUN mkdir /app
COPY api /app
CMD [ "/app/api"]
