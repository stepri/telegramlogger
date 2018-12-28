FROM iron/go
WORKDIR /app
ADD app /app/
EXPOSE 5050 
ENTRYPOINT ["./app"]
