# 🖼️ Take-Home Test: Image Cropping via Border Detection

## 📌 Objective

This program automatically detects and crops the rectangular area inside a black border in `image.png`, using **pure Go standard library**. The cropped image includes the border and is saved as `output.png`.

---

## ✅ Features

- ✅ **Black border detection** using raw pixel inspection.
- ✅ Crops the content **with black border included**.
- ✅ Uses only **Go standard library** (`image`, `image/png`, `os`, `log`).
- ✅ Optionally logs all border pixel coordinates to `border_area.log`.

---
## 🧾 Pixel Detection Logic

```go
r, g, b, a := img.At(x, y).RGBA()
if a > 0 && !(r>>8 == 255 && g>>8 == 255 && b>>8 == 255) {
    // This is part of the border
}
```

## ▶️ How To Run
```go
 go run main.go
```