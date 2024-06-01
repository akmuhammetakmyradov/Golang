package main

import (
	"fanOut-fanIn/imageProcessing"
	"image"
	"strings"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func LoadImage(paths []string) []Job {
	var jobs []Job
	for _, p := range paths {
		job := Job{
			InputPath: p,
			OutPath:   strings.Replace(p, "images/", "images/output/", 1),
		}
		// fmt.Println(job)
		job.Image = imageProcessing.ReadImage(p)
		jobs = append(jobs, job)
	}

	return jobs
}

func Resize(jobs *[]Job) <-chan Job {
	out := make(chan Job, len(*jobs))

	for _, job := range *jobs {
		go func(job Job) {
			job.Image = imageProcessing.ResizeImage(job.Image)
			out <- job
		}(job)
	}
	return out
}

func CollectJobs(input <-chan Job, imageCnt int) []Job {
	var resizedJobs []Job
	for i := 0; i < imageCnt; i++ {
		job := <-input
		resizedJobs = append(resizedJobs, job)
	}
	return resizedJobs
}

func SaveImages(jobs *[]Job) {
	for _, image := range *jobs {
		imageProcessing.WriteImage(image.OutPath, image.Image)
	}
}

func main() {
	images := []string{
		"images/jony.jpeg",
		"images/jony1.jpeg",
		"images/jony2.jpeg",
		"images/jony3.jpeg",
	}

	jobs := LoadImage(images)
	// fan out this function to multiply goroutines
	job := Resize(&jobs)
	// Collect fan in
	resizedJobs := CollectJobs(job, len(images))
	SaveImages(&resizedJobs)
}
