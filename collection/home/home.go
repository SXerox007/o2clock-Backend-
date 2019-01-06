package home

import (
	"fmt"
	"log"
	face "o2clock/algorithm-face"
	"o2clock/constants/errormsg"
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
		return err
	}
	defer rec.Close()
	//TODO ---> Get the image url from db
	dataImage := filepath.Join(DATA_DIR, "sumit.jpg")

	faces, err := rec.RecognizeFile(dataImage)
	if err != nil {
		return err
	}

	var samples []face.Descriptor
	var totalF []int32
	for i, f := range faces {
		samples = append(samples, f.Descriptor)
		// Each face is unique on that image so goes to its own category.
		totalF = append(totalF, int32(i))
	}

	// Pass samples to the recognizer.
	rec.SetSamples(samples, totalF)

	// Now let's try to classify some not yet known image.

	//testSumit := filepath.Join(DATA_DIR, "sumit.jpg")
	sumit, err := rec.RecognizeSingle(image)
	if err != nil {
		log.Println("Face not recorganise not the same person")
		return status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_FACE_NOT_REC))
	}
	if sumit == nil {
		log.Println("Not a sigle image")
		return status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_NOT_A_SINGLE_FACE))
	}
	log.Println("Image recorganise")
	return nil
}
