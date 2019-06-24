package home

import (
	"fmt"
	"log"
	face "o2clock/algorithm-face"
	"o2clock/constants/errormsg"
	cFace "o2clock/utils/cface"
	"path/filepath"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	if testf == nil {
		log.Println("Not a single face on the image")
		return status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln(errormsg.ERR_NOT_A_SINGLE_FACE))
	}

	id := cFace.CompareFaces(samples, testf.Descriptor, 0.6)
	if id < 0 {
		log.Println("didn't find known face")
		return status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln(errormsg.ERR_FACE_NOT_REC))
	}

	log.Println("id", id)
	log.Println("Image reorganised")
	return nil
}
