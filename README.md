# SWAPI serverless

> [swapi.dev](https://www.swapi.dev) clone written in Go and deployed as lambda function with Netlify.

## WORK IN PROGRESS

Code for this initial draft is in the `develop` branch.

Launch locally with command:
```
$ go generate src/api.go && go run src/api.go -port=8001 -port=8001
```

Open your browser at [http://localhost:8001/api/films](http://localhost:8001/api/films)

Build for production with command:
```
$ make build
```

