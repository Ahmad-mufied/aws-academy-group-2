package utils

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
)

type UUIDCodec struct{}

func (uc *UUIDCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if val.Type() != reflect.TypeOf(uuid.UUID{}) {
		return bsoncodec.ValueEncoderError{
			Name:     "UUIDEncodeValue",
			Types:    []reflect.Type{reflect.TypeOf(uuid.UUID{})},
			Received: val,
		}
	}

	// Use the MarshalBinary method to get the byte representation
	uuid := val.Interface().(uuid.UUID)
	bytes, _ := uuid.MarshalBinary()

	// Write the bytes with the UUID subtype (0x04)
	return vw.WriteBinaryWithSubtype(bytes, 0x04)
}

func (uc *UUIDCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != reflect.TypeOf(uuid.UUID{}) {
		return bsoncodec.ValueDecoderError{
			Name:     "UUIDDecodeValue",
			Types:    []reflect.Type{reflect.TypeOf(uuid.UUID{})},
			Received: val,
		}
	}

	data, subtype, err := vr.ReadBinary()
	if err != nil {
		return err
	}

	if subtype != 0x04 {
		return fmt.Errorf("unsupported UUID subtype: %v", subtype)
	}

	uuidVal, err := uuid.FromBytes(data)
	if err != nil {
		return err
	}

	val.Set(reflect.ValueOf(uuidVal))
	return nil
}
