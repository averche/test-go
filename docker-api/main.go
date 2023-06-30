package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/exp/slices"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
	}

	fmt.Println(fetchEntrypointAndCmdFromDocker("test-app"))
}

func fetchEntrypointAndCmdFromDocker(imageName string) ([]string, []string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if !strings.Contains(imageName, ":") {
		imageName = imageName + ":latest"
	}

	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't start docker client: %w", err)
	}
	defer dockerClient.Close()

	images, err := dockerClient.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't list images: %w", err)
	}

	// find image ID
	var imageID string
	for _, image := range images {
		if slices.Contains(image.RepoTags, imageName) {
			imageID = image.ID
		}
	}

	if imageID == "" {
		return nil, nil, fmt.Errorf("couldn't find image ID for %q", imageName)
	}

	imageInspect, _, err := dockerClient.ImageInspectWithRaw(ctx, imageID)
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't inspect image: %w", err)
	}

	return imageInspect.Config.Entrypoint, imageInspect.Config.Cmd, nil
}
