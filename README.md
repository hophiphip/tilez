# tilez
Break up image into a set of square map tiles at each zoom level

## Quick start
```sh
go run .
```

## API Routes
| Method | Route | Description |
|--------|:-----:|:-----------:|
| `GET`  | `/{image}/{zoom}/{x}/{y}` | Retrieve image tile with provided zoom |

| Argument | Description |
|:--------:|:-----------:|
| `image`  | image name without extension |
| `x`      | tile x position on a grid |
| `y`      | tile y position on a grid |
| `zoom`   | grid zoom level |
