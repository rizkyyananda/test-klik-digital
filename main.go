package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	img := loadImage("image.png")
	bounds := img.Bounds()

	minX, minY, maxX, maxY, found := findContentBounds(img, bounds)
	if !found {
		log.Fatal("Tidak ditemukan konten pada gambar.")
	}

	// Tambahkan padding
	padding := 1
	minX = max(0, minX-padding)
	minY = max(0, minY-padding)
	maxX = min(bounds.Max.X-10, maxX+padding)
	maxY = min(bounds.Max.Y-158, maxY+padding)

	croppedImg := cropImage(img, minX, minY, maxX, maxY)
	saveImage("output.png", croppedImg)

	log.Println("Gambar berhasil di-crop dan disimpan sebagai output.png")
	log.Printf("Area yang di-crop persegi hitam: X(%d-%d), Y(%d-%d)", minX, maxX, minY, maxY)
}

// loadImage membuka dan mendekode file gambar
func loadImage(filename string) image.Image {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Gagal membuka gambar: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Gagal mendekode gambar: %v", err)
	}
	return img
}

// findContentBounds mencari batas-batas konten non-background pada gambar
func findContentBounds(img image.Image, bounds image.Rectangle) (minX, minY, maxX, maxY int, found bool) {
	width, height := bounds.Max.X, bounds.Max.Y
	minX, maxX = width, 0
	minY, maxY = height, 0
	found = false

	logFile := createLogFile("border_area.log")
	defer func() {
		if logFile != nil {
			logFile.Close()
		}
	}()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			// Periksa jika pixel bukan putih dan tidak transparan
			if a > 0 && !(r>>8 == 255 && g>>8 == 255 && b>>8 == 255) {
				found = true
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
				if logFile != nil {
					logFile.WriteString(fmt.Sprintf("(X: %d, Y: %d)\n", x, y))
				}
			}
		}
	}
	return
}

// cropImage memotong gambar sesuai batas yang diberikan
func cropImage(img image.Image, minX, minY, maxX, maxY int) image.Image {
	cropWidth := maxX - minX + 1
	cropHeight := maxY - minY + 1
	cropped := image.NewRGBA(image.Rect(0, 0, cropWidth, cropHeight))

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			cropped.Set(x-minX, y-minY, img.At(x, y))
		}
	}
	return cropped
}

// saveImage menyimpan gambar dalam format PNG ke file
func saveImage(filename string, img image.Image) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Gagal membuat file output: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatalf("Gagal menyimpan gambar: %v", err)
	}
}

// createLogFile membuat file log untuk menyimpan koordinat konten (opsional)
func createLogFile(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		log.Printf("Peringatan: gagal membuat log file: %v", err)
		return nil
	}
	file.WriteString("Koordinat batas konten:\n")
	return file
}

// max mengembalikan nilai maksimum dari dua angka
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min mengembalikan nilai minimum dari dua angka
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
