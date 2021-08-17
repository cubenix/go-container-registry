package main

import (
	"fmt"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func main() {
	src, err := name.ParseReference("quay.io/tinkerbell/tink")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nsrc\n")
	fmt.Println(src)

	// Fetch the manifest using default credentials.
	img, err := remote.Get(src)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nimg\n")
	fmt.Println(img)
	image, err := img.Image()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nimage\n")
	fmt.Println(image)
	dst, err := name.ParseReference("registry.hub.docker.com/quickdevnotes/go-tink")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\ndst\n")
	fmt.Println(dst)

	err = remote.Write(dst, image, remote.WithAuth(&authn.Basic{Username: "username", Password: "password"}))
	if err != nil {
		panic(err)
	}

	// Prints the digest of registry.example.com/private/repo
	fmt.Println(img.Digest)
}
