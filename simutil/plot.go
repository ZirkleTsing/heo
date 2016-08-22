package simutil

import (
	"os/exec"
	"fmt"
	"os"
)

const (
	BAR_PLOT = 0
	LINE_PLOT = 1
)

type PlotType int

func GeneratePlot(outputDirectory string, csvFileName string, plotFileName string, x string, hue string, y string, xticklabelsRotation int, plotType PlotType) {
	var cmd *exec.Cmd

	switch hue {
	case "":
		cmd = exec.Command(
			"./plots.sh",
			"--csv_file_name", outputDirectory + "/" + csvFileName,
			"--plot_file_name", outputDirectory + "/" + plotFileName,
			"--x", x,
			"--y", y,
			"--xticklabels_rotation", fmt.Sprintf("%d", xticklabelsRotation),
			"--plot_type", fmt.Sprintf("%d", plotType),
		)
	default:
		cmd = exec.Command(
			"./plots.sh",
			"--csv_file_name", outputDirectory + "/" + csvFileName,
			"--plot_file_name", outputDirectory + "/" + plotFileName,
			"--x", x,
			"--hue", hue,
			"--y", y,
			"--xticklabels_rotation", fmt.Sprintf("%d", xticklabelsRotation),
			"--plot_type", fmt.Sprintf("%d", plotType),
		)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}
