package main

import (
	"fmt"
	"image"
	"log"
	"os"
)

func main() {

	// For debugging, turn fontmap into individual font images
	// No need to keep running, but doesn't hurt to leave it in.
	{
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

		// Write individual char images
		for chr := 0; chr < numBaseChars; chr++ {
			rect := image.Rect(chr*charWidth, 0, (chr+1)*charWidth, charHeight)
			clip := m.(subimager).SubImage(rect)
			writeImage(clip, fmt.Sprintf("char_%v.png", chr))
			charmap[chr] = clip
		}
	}

	// Example usage:

	// Write a single stamp
	stamp := genStamp("foobar", 1)
	writeImage(stamp, "foobar_stamp_1.png")

	// Write 10 pages
	for p := 0; p < 10; p++ {
		page := genPage("foobar", p, 10, 10)
		writeImage(page, fmt.Sprintf("foobar_page_%v.png", p))
		fmt.Println("printed page ", p)
	}
}
