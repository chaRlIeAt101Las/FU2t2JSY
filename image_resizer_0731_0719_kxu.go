// 代码生成时间: 2025-07-31 07:19:06
Usage:
	image-resizer --input-dir="./images" --output-dir="./resized" --width=100 --height=100

Options:
	--input-dir string
		Directory containing input images (default "./images")
	--output-dir string
		Directory where resized images will be saved (default "./resized")
	--width int
		The desired width of the resized images (default 100)
	--height int
		The desired height of the resized images (default 100)
*/

package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// Define the root command
var rootCmd = &cobra.Command{
	Use:   "image-resizer",
	Short: "Batch resize images",
	Long:  `A simple CLI tool to resize images in bulk.`,
	Run:   run,
}

var (
	inputDir   string
	outputDir string
	width     int
	height    int
)

// init is called before main
func init() {
	rootCmd.PersistentFlags().StringVar(&inputDir, "input-dir", "./images", "Directory containing input images")
	rootCmd.PersistentFlags().StringVar(&outputDir, "output-dir", "./resized", "Directory where resized images will be saved")
	rootCmd.PersistentFlags().IntVar(&width, "width", 100, "The desired width of the resized images")
	rootCmd.PersistentFlags().IntVar(&height, "height", 100, "The desired height of the resized images")
}

// run is the main logic of the CLI tool
func run(cmd *cobra.Command, args []string) {
	err := resizeImages(inputDir, outputDir, width, height)
	if err != nil {
		log.Fatalf("Error resizing images: %v", err)
	}
}

// resizeImages resizes images in the specified directory
func resizeImages(inputDir, outputDir string, width, height int) error {
	files, err := os.ReadDir(inputDir)
	if err != nil {
		return fmt.Errorf("failed to read input directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(inputDir, file.Name())
		img, err := loadImage(filePath)
		if err != nil {
			log.Printf("Skipping %s: %v", file.Name(), err)
			continue
		}

		resizedImg := resizeImage(img, width, height)
		outputPath := filepath.Join(outputDir, file.Name())
		err = saveImage(resizedImg, outputPath)
		if err != nil {
			log.Printf("Failed to save %s: %v", file.Name(), err)
			continue
		}
	}

	return nil
}

// loadImage loads an image from the file path
func loadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	return img, nil
}

// resizeImage resizes the given image to the specified width and height
func resizeImage(img image.Image, width, height int) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Src = img
	draw.F = draw.Over
	draw.Draw(m, m.Bounds(), img, image.Pt(0, 0), draw.Src)
	return m
}

// saveImage saves the image to the file path
func saveImage(img image.Image, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	err = jpeg.Encode(file, img, nil)
	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	return nil
}

func main() {
	rootCmd.Execute()
}
