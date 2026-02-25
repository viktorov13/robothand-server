package support

import (
	"io"
	"os"
)

type Service struct{}

func (s *Service) ProcessTicket(
	uuid string,
	email string,
	header string,
	text string,
	file io.Reader,
	filename string,
) error {

	if file != nil {
		f, err := os.Create("uploads/" + filename)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}
	}

	// Здесь можно:
	// — отправить email
	// — сохранить тикет в БД
	// — отправить в CRM

	return nil
}
