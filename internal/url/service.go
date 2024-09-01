package url

import "errors"

type Service interface {
	GetCurrentOrGenerateKey(url string) (string, error)
	WriteToMemory(key, url string) error
	GetUrlIfExist(key string) (string, error)
}

type service struct {
	Storage   Storage
	Shortener Shortener
}

func NewService(storage Storage, shortener Shortener) Service {
	return &service{Storage: storage, Shortener: shortener}
}

func (s *service) GetCurrentOrGenerateKey(url string) (string, error) {
	if ok, key := s.Storage.IsPresentedByURL(url); ok {
		return key, nil
	}
	short := s.Shortener.Shorten()
	err := s.WriteToMemory(short, url)
	if err != nil {
		return "", err
	}
	return short, nil
}

func (s *service) WriteToMemory(key, url string) error {
	return s.Storage.Write(key, url)
}

func (s *service) GetUrlIfExist(key string) (string, error) {
	if ok, value := s.Storage.IsPresentedByKey(key); ok {
		return value, nil
	} else {
		return "", errors.New("key not found")
	}
}
