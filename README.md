## Fresh

Fresh could you help with hot reloading at developing
```
go get -u github.com/pilu/fresh
```

## Environment variables
Please create a `.env` file with your credentials
```
WYKOP_APP_KEY=
APP_URL=
```

## Docker
```
docker build -t wykop-rss .
```

```
docker run -d -p 9001:9001 -e "WYKOP_APP_KEY=xxx" "APP_URL=yyy" wykop-rss
```


## Deployment
```
npm -i g now
```

```
now --public -e WYKOP_APP_KEY=xxx APP_URL=yyy
```