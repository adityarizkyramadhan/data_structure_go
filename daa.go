package data_structure_go

//Aditya Rizky R 215150201111007 DAA
func PerpangkatanRekursif(angkaBulat, angkaPangkat int) int64 {
	if angkaPangkat == 1 {
		return int64(angkaBulat)
	}
	return int64(angkaBulat) * PerpangkatanRekursif(angkaBulat, angkaPangkat-1)
}
