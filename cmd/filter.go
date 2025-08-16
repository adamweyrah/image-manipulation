package cmd

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/adamweyrah/images/internal/processing"
	"github.com/adamweyrah/images/pkg"
	"github.com/spf13/cobra"
)

var filterType string
var outputFormat string

var validTypes = []string{"sepia", "grayscale", "invert"}

// filterCmd creates a new image with filter options applied (sepia|grayscale|invert)
var filterCmd = &cobra.Command{
	Use:   "filter <image-path>",
	Short: "applies a filter (sepia|grayscale) to input image.",
	Long: `filter allows you to specify a '--type=' of (sepia|grayscal|invert) and applies 
them to your original image, creating a new image with the filter applied.`,
	Args:      cobra.ExactArgs(1),
	ValidArgs: validTypes,
	Run: func(cmd *cobra.Command, args []string) {
		imagePath := args[0]

		var format string

		var processor processing.Processor

		switch filterType {
		case "sepia":
			processor = processing.SepiaFilter{}
		case "grayscale":
			processor = processing.GrayscaleFilter{}
		case "invert":
			processor = processing.InvertedFilter{}
		default:
			log.Fatalf("Error: unknown filter type: '%s'\n", filterType)
		}

		file, err := os.Open(imagePath)
		if err != nil {
			log.Fatalf("error opening file: %v\n", err)
		}
		defer file.Close()

		img, detectedFormat, err := image.Decode(file)
		if err != nil {
			log.Fatalf("error decoding the image file: %v\n", err)
		}

		if outputFormat == "" {
			format = detectedFormat
		} else {
			format = outputFormat
		}

		newFilename := pkg.GenerateNewFilename(imagePath, filterType, format)

		outFile, err := os.Create(newFilename)
		if err != nil {
			log.Fatalf("error creating file: %v\n", err)
		}
		defer outFile.Close()

		newImage := processor.Process(img)

		if err := processing.EncodeImage(outFile, newImage, format); err != nil {
			log.Fatalf("failed to encode image: %v\n", err)
		}

		fmt.Printf("Successfully applied '%s' filter. New image saved as '%s'.\n", filterType, newFilename)
	},
}

func init() {
	rootCmd.AddCommand(filterCmd)
	filterCmd.Flags().StringVarP(&filterType, "type", "t", "", "The type of filter to apply (sepia|grayscale).")
	filterCmd.Flags().StringVarP(&outputFormat, "format", "f", "", "The format in which to convert the out file to.")

	filterCmd.MarkFlagRequired("type")
}
