package main

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)

func CASPathTranformFunc(key string) string {
	hash := sha1.Sum([]byte(key))
	//fmt.Printf("Print has %v\n", hash)
	hashStr := hex.EncodeToString(hash[:])

	//fmt.Printf("hashstra %v\n", hashStr)

	blocksize := 5
	sliceLen := len(hashStr) / blocksize

	//fmt.Printf("Slice len %v \n", sliceLen)

	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blocksize, (i*blocksize)+blocksize
		paths[i] = hashStr[from:to]
	}

	//fmt.Printf("Final paths %v\n", paths)
	return strings.Join(paths, "/")
}

type PathTransformFunc func(string) string

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

var DefaultPathTransformFunc = func(key string) string {
	return key
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathName := s.PathTransformFunc(key)

	if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
		return err
	}

	filename := "somefilename"

	pathAndFilename := pathName + "/" + filename

	f, err := os.Create(pathAndFilename)
	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}

	log.Printf("written (%d) bytes to disk: %s", n, pathAndFilename)

	return nil
}
