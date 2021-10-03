package model

type MessageBase interface {
	ToBytes() ([]byte, error)
	ToJSON() (string, error)
}
