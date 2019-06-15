package home

import (
	"log"
	face "o2clock/algorithm-face"
	cFace "o2clock/utils/cface"
	"path/filepath"
)

const (
	VERSION  = "v1.0"
	DATA_DIR = "algorithm-face/dlib-model"
)

func VerifyUser(image []byte, accessToken string) error {
	rec, err := face.NewRecognizer(DATA_DIR)
	if err != nil {
		log.Fatalln(err)
	}
	defer rec.Close()

	dataImage := filepath.Join(DATA_DIR, "mix.jpg")
	faces, err := rec.RecognizeFile(dataImage)
	if err != nil {
		log.Fatalln(err)
	}

	var samples []face.Descriptor
	var totalF []int32
	for i, f := range faces {
		samples = append(samples, f.Descriptor)
		totalF = append(totalF, int32(i))
	}

	//testData := filepath.Join(DATA_DIR, "sample.jpg")
	testf, err := rec.RecognizeSingle(image)
	if err != nil {
		log.Fatalln(err)
	}

	id := cFace.CompareFaces(samples, testf.Descriptor, 0.6)
	if id < 0 {
		log.Println("didn't find known face")
		return nil
	}

	log.Println("id", id)
	log.Println("Image reorganised")
	return nil
}
