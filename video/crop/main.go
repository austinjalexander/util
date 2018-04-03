package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/satori/go.uuid"
)

const (
	i = "/Users/thebigapple/Desktop/social-promo.mp4"
)

func main() {
	var (
		i = flag.String("i", "", "Input filepath")
		w = flag.Uint("w", 0, "Crop width")
		h = flag.Uint("h", 0, "Crop height")
		x = flag.Uint("x", 0, "Crop top-left-most x")
		y = flag.Uint("y", 0, "Crop top-left-most y")
	)
	flag.Parse()

	outFilepath, err := Crop(*i, *w, *h, *x, *y)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\t***New cropped video file at %s!***\n", outFilepath)
}

// Crop makes system calls to ffmpeg to crop a video.
func Crop(inFilepath string, w, h, x, y uint) (string, error) {
	cmd, args, err := setCmdArgs(inFilepath, w, h, x, y)
	if err != nil {
		return "", err
	}

	stdOutErr, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error: %s: %s", stdOutErr, err)
	}
	return args[len(args)-1], nil
}

func outFilepath() (string, error) {
	vidUUID, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.mp4", filepath.Join(os.TempDir(), vidUUID.String())), nil
}

func setCmdArgs(inFilepath string, w, h, x, y uint) (string, []string, error) {
	out, err := outFilepath()
	if err != nil {
		return "", nil, err
	}

	cmd := "/usr/local/bin/ffmpeg"
	args := []string{
		"-i",
		inFilepath,
		"-vf",
		fmt.Sprintf("crop=%d:%d:%d:%d", w, h, x, y),
		"-crf",
		"1",
		"-c:a",
		"copy",
		out,
	}

	return cmd, args, nil
}
