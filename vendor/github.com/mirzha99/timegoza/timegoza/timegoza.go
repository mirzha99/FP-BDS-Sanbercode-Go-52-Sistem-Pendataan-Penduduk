package timegoza

import (
	"fmt"
	"strconv"
	"time"
)

type ZaTimes struct {
	Epoch    int64
	Types    string
	TimeZone string
}
type ZaSetTimes struct {
	Detik, Menit, Jam, Tanggal, Bulan, Tahun int
	TimeZone                                 string
}

func EpochTime() int64 {
	timenow := time.Now().Unix()
	return timenow
}
func getIndonesianMonthName(month time.Month) string {
	// Pemetaan nama bulan dalam bahasa Indonesia
	bulan := map[time.Month]string{
		time.January:   "Januari",
		time.February:  "Februari",
		time.March:     "Maret",
		time.April:     "April",
		time.May:       "Mei",
		time.June:      "Juni",
		time.July:      "Juli",
		time.August:    "Agustus",
		time.September: "September",
		time.October:   "Oktober",
		time.November:  "November",
		time.December:  "Desember",
	}

	return bulan[month]
}
func getIndonesianDayName(day time.Weekday) string {
	hari := map[time.Weekday]string{
		time.Sunday:    "Minggu",
		time.Monday:    "Senin",
		time.Tuesday:   "Selasa",
		time.Wednesday: "Rabu",
		time.Thursday:  "Kamis",
		time.Friday:    "Jum'at",
		time.Saturday:  "Sabtu",
	}
	return hari[day]
}
func (z *ZaTimes) HumanTime() string {
	location, err := time.LoadLocation(z.TimeZone)
	if err != nil {
		return "Error : " + err.Error()
	}

	timeInLocation := time.Unix(z.Epoch, 0).In(location)
	Hour, Minute, Second := timeInLocation.Hour(), timeInLocation.Minute(), timeInLocation.Second()
	Day := timeInLocation.Day()
	Month := getIndonesianMonthName(timeInLocation.Month())
	DayWeek := getIndonesianDayName(timeInLocation.Weekday())
	Year := timeInLocation.Year()
	switch z.Types {
	case "Time":
		return strconv.Itoa(Hour) + ":" + strconv.Itoa(Minute) + ":" + strconv.Itoa(Second)
	case "Date":
		return strconv.Itoa(Day) + "-" + Month + "-" + strconv.Itoa(Year)
	case "Day":
		return DayWeek
	case "Kop":
		return DayWeek + "," + strconv.Itoa(Day) + "-" + Month + "-" + strconv.Itoa(Year)
	default:
		return strconv.Itoa(Day) + "-" + Month + "-" + strconv.Itoa(Year) + " (" + strconv.Itoa(Hour) + ":" + strconv.Itoa(Minute) + ":" + strconv.Itoa(Second) + ")"
	}

}
func (z *ZaSetTimes) HumanEpoch() int64 {
	//location
	location, err := time.LoadLocation(z.TimeZone)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	// Menggunakan nilai-nilai dari struktur ZaSetTimes
	epochs := time.Date(z.Tahun, time.Month(z.Bulan), z.Tanggal, z.Jam, z.Menit, z.Detik, 0, location).Unix()

	return epochs
}
