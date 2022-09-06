package date

import "time"

const (
	layoutISO      = "2006-01-02"          //2022-12-31
	layoutUS       = "January 2, 2006"     //Desember 31, 2022
	layoutWithTIme = "2006-01-02 15:04:05" //2022-12-31 13:47:29
)

func DateFormat(datetime time.Time, layout string) string {
	t := datetime.Format(layout)
	return t
}
