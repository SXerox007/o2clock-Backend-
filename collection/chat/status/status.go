package status

import (
	"o2clock/constants/collections"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
)

type AllPersonStatus struct {
	ID          objectid.ObjectID `bson:"_id,omitempty"`
	PersonID    string            `bson:"person_id"`
	Status      bool              `bson:"status"`
	Version     string            `bson:"version"`
	CaptureTime time.Time         `bson:"capture_time"`
}

/**
*
* Set the  person to online/offline
*
**/
func SetUsersStatus(isPersonOnline bool, userId string) error {

}

/**
*
* validate and update status
*
**/
func ValidateAndUpdateUsersStatus(personId string, isPersonOnline string) error {
	data := &AllPersonStatus{}
	filter := bson.M{collections.PARAM_PERSON_ID: personId}

}
