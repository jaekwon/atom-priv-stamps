package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// _ "image/gif"
	// _ "image/jpeg"
	"image/png"

	"github.com/tendermint/go-crypto"
)

const numBaseChars = 16 // 16 for hex
const numCols = 6       // print 6 hex char columns
const numRows = 4       // print 4 hex char rows
const padStampX = 6     // padding around whole stamp
const padStampY = 5     // padding around whole stamp
const padCharX = 0      // padding around each character
const padCharY = 10     // padding around each character
var charmap = map[int]image.Image{}
var charWidth, charHeight int

// Sets a solid color on an image
func clearImage(img image.Image, c color.Color) {
	m := img.(setter)
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			m.Set(x, y, c)
		}
	}
}

// Writes the image to a file
func writeImage(img image.Image, file string) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

// Convenience interface
type subimager interface {
	SubImage(r image.Rectangle) image.Image
}

// Convenience interface
type setter interface {
	Set(x, y int, c color.Color)
}

// genCode will generate 12 bytes given a seed and offset
// seed: a seed phrase
// offset: 0, 1, 2, ...
// returns: 12 bytes
func genCode(seed string, offset int) []byte {
	hash := crypto.Ripemd160([]byte(fmt.Sprintf("%v\n%v", seed, offset)))
	return hash[:12]
}

// genStamp writes an image (stamp) of 6x4 hex letters for genCode result bytes
func genStamp(seed string, offset int) image.Image {
	code := genCode(seed, offset)
	rect := image.Rect(0, 0,
		(numCols)*charWidth+padStampX*2+(numCols-1)*padCharX,
		(numRows)*charHeight+padStampY*2+(numRows-1)*padCharY)
	stamp := image.NewNRGBA(rect)
	clearImage(stamp, color.White)
	for y := 0; y < numRows; y++ {
		for x := 0; x < numCols; x++ {
			// b is the byte we want to write
			b := code[(numCols/2)*y+x/2]
			// chr is the first 4 bits or the last 4
			var chr int
			if x%2 == 0 {
				chr = int((b & 0xF0) >> 4)
			} else {
				chr = int((b & 0x0F))
			}
			// fmt.Printf(">> %v %v %v\n", y, x, chr)
			// destination bounds
			db := image.Rect(
				padStampX+x*(charWidth+padCharX),
				padStampY+y*(charHeight+padCharY),
				padStampX+x*(charWidth+padCharX)+charWidth,
				padStampY+y*(charHeight+padCharY)+charHeight)
			// fmt.Println("stamp", stamp.Bounds(), "db", db, "charmap", charmap[chr].Bounds())
			draw.FloydSteinberg.Draw(stamp, db, charmap[chr], charmap[chr].Bounds().Min)
		}
	}
	return stamp
}

// genPage creates an image with many stamps on a page
func genPage(seed string, pageNum int, numCols int, numRows int) image.Image {
	// First, get sample
	stamp := genStamp(seed, 0)
	stampWidth := stamp.Bounds().Dx()
	stampHeight := stamp.Bounds().Dy()

	rect := image.Rect(0, 0,
		(numCols)*(stampWidth+1)+1,
		(numRows)*(stampHeight+1)+1)
	page := image.NewNRGBA(rect)
	clearImage(page, color.Black)
	for y := 0; y < numRows; y++ {
		for x := 0; x < numCols; x++ {
			stamp := genStamp(seed, numCols*numRows*pageNum+(numCols*y)+x)
			// destination bounds
			db := image.Rect(
				1+x*(stampWidth+1),
				1+y*(stampHeight+1),
				1+x*(stampWidth+1)+stampWidth,
				1+y*(stampHeight+1)+stampHeight)
			draw.FloydSteinberg.Draw(page, db, stamp, stamp.Bounds().Min)
		}
	}
	return page
}
