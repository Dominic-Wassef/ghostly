package s3filesystem

import "github.com/dominic-wassef/ghostly/filesystems"

type S3 struct {
	Key      string
	Secret   string
	Region   string
	Endpoint string
	Bucket   string
}

func (s *S3) Put(fileName, folder string) error {
	return nil
}

func (s *S3) List(prefix string) ([]filesystems.Listing, error) {
	var listing []filesystems.Listing
	return listing, nil
}

func (s *S3) Delete(itemsToDelete []string) bool {
	return true
}

func (s *S3) Get(destination string, items ...string) error {
	return nil 
}