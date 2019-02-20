## Development

Fresh could you help with hot reloading on development.
```
go get -u github.com/pilu/fresh
```

Project uses `go mod` for dependency management (available from `go 1.11`).
```
go get
```

```
fresh
```

### Environment variables
Please create a `.env` file with your credentials.
```
WYKOP_APP_KEY=
APP_URL=
```

## Docker
```
docker build -t wykop-rss .
```

```
docker run -d -p 9001:9001 --env-file ./.env wykop-rss
```


## Deployment
Warning: Docker deployment on [zeit.co/now](zeit.co/now/) is depreciated for new users.
```
npm i -g now
```

```
now --public -e WYKOP_APP_KEY=xxx -e APP_URL=yyy
```

```
now alias https://wykop-rss-xyz.now.sh abc-wykop-rss
```