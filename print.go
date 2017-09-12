package main

import (
	"fmt"
	"image"
	"log"
	"os"
)

func main() {

	reader, err := os.Open("fontmap.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	bounds := m.Bounds()
	charWidth = bounds.Dx() / numBaseChars
	charHeight = bounds.Dy()
	// fmt.Println("charWidth", charWidth, "charHeight", charHeight)

	for chr := 0; chr < numBaseChars; chr++ {
		rect := image.Rect(chr*charWidth, 0, (chr+1)*charWidth, charHeight)
		clip := m.(subimager).SubImage(rect)
		writeImage(clip, fmt.Sprintf("char_%v.png", chr))
		charmap[chr] = clip
	}

	// fmt.Println(genCode("foobar", 1))
	// fmt.Println(genCode("foobar", 1))
	// fmt.Println(genCode("foobar", 2))
	// fmt.Println(genCode("foobar", 3))
	// fmt.Println(genCode("foobaz", 1))
	// fmt.Println(genCode("foobaz", 2))
	// fmt.Println(genCode("foobaz", 3))

	stamp := genStamp("foobar", 1)
	// fmt.Println("stampWidth", stamp.Bounds().Dx(), "stampHeight", stamp.Bounds().Dy())
	writeImage(stamp, "foobar_stamp_1.png")

	for p := 0; p < 10; p++ {
		page := genPage("foobar", p, 10, 10)
		// fmt.Println("pageWidth", page.Bounds().Dx(), "pageHeight", page.Bounds().Dy())
		writeImage(page, fmt.Sprintf("foobar_page_%v.png", p))
		fmt.Println("printed page ", p)
	}
}
