package goffmpeg

import (
	"os"
	"os/exec"
)

func AddAudio(inputVideo string, inputAudio string, outputFile string) error {
	cmd := exec.Command("ffmpeg", "-i", inputVideo, "-i", inputAudio, "-filter_complex", "[0:a:0]volume=0.2[a1];[1:a:0]volume=1.0[a2];[a1][a2]amix=inputs=2[aout]", "-map", "[aout]", "-map", "0:v", "-c:v", "copy", "-c:a", "aac", outputFile, "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
