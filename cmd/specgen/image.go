package specgen

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// imageRegex matches Markdown image syntax with base64 data URLs
var imageRegex = regexp.MustCompile(`!\[([^\]]*)\]\(data:image/(png|jpeg|gif|jpg|webp);base64,([^)]+)\)`)

// ExtractAndSaveImages extracts base64 images from markdown,
// saves them to files, and replaces with relative paths
func ExtractAndSaveImages(markdown, pageID, imagesDir string) (string, error) {
	imageIndex := 0

	result := imageRegex.ReplaceAllStringFunc(markdown, func(match string) string {
		submatches := imageRegex.FindStringSubmatch(match)
		if len(submatches) != 4 {
			return match
		}

		alt := submatches[1]
		ext := submatches[2]
		b64data := submatches[3]

		// Normalize extension
		if ext == "jpeg" {
			ext = "jpg"
		}

		// Generate filename
		imageIndex++
		filename := fmt.Sprintf("%s_img%03d.%s", pageID, imageIndex, ext)

		// Decode base64
		data, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			// If decoding fails, keep original
			return match
		}

		// Save to file
		imagePath := filepath.Join(imagesDir, filename)
		if err := os.WriteFile(imagePath, data, 0644); err != nil {
			// If saving fails, keep original
			return match
		}

		// Return markdown with relative path
		return fmt.Sprintf("![%s](images/%s)", alt, filename)
	})

	return result, nil
}

// CleanImagePlaceholders removes image placeholders like *[Image]*
func CleanImagePlaceholders(markdown string) string {
	// Remove standalone image placeholders
	re := regexp.MustCompile(`\*\[Image\]\*\n*`)
	return re.ReplaceAllString(markdown, "")
}

// ExtractImagesInfo extracts information about images without saving
func ExtractImagesInfo(markdown string) []ImageInfo {
	var images []ImageInfo

	matches := imageRegex.FindAllStringSubmatch(markdown, -1)
	for i, match := range matches {
		if len(match) != 4 {
			continue
		}

		ext := match[2]
		if ext == "jpeg" {
			ext = "jpg"
		}

		// Estimate size from base64 length (base64 is ~4/3 of original)
		estimatedSize := len(match[3]) * 3 / 4

		images = append(images, ImageInfo{
			Index:         i + 1,
			Alt:           match[1],
			Extension:     ext,
			EstimatedSize: estimatedSize,
		})
	}

	return images
}

// ImageInfo contains information about an extracted image
type ImageInfo struct {
	Index         int
	Alt           string
	Extension     string
	EstimatedSize int
}

// HasBase64Images checks if markdown contains base64 images
func HasBase64Images(markdown string) bool {
	return strings.Contains(markdown, "data:image/") && strings.Contains(markdown, ";base64,")
}
