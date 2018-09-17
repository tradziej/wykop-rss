## Fresh

Fresh could you help with hot reloading at developing
```
go get -u github.com/pilu/fresh
```

## Environment variables
Please create a `.env` file with your credentials
```
WYKOP_APP_KEY=
```

## Docker
```
docker build -t wykop-rss .
```

```
docker run -d -p 9001:9001 -e "WYKOP_APP_KEY=xxx" wykop-rss
```
