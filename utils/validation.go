package utils

import "assyifa/models"

func ValidationUser(input *models.Users) bool {
	var result bool
	result = true
	switch {
	case input.Username == "":
		result = false
	case input.Password == "":
		result = false
	case input.NamaLengkap == "":
		result = false
	case input.TanggalLahir == "":
		result = false
	case input.Alamat == "":
		result = false
	}
	return result
}
