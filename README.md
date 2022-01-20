# tilez
Break up image into a set of square map tiles at each zoom level

## Quick start
```sh
go run .
```

## API Routes
| Method | Route | Description |
|--------|:-----:|:-----------:|
| `GET`  | `/img/:image/:x/:y/:zoom` | Retrieve image tile with provided zoom |
