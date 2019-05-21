package main

import (
	"encoding/csv"
	"goahrs"
	"flag"
	"log"
	"os"
	"strconv"
)

var newTab [][]string
var q goahrs.Quaternion

var (
	accelxoffset = -0.67185
	accelyoffset = 0.37175
	accelzoffset = -0.00228
	gyroxoffset  = 0.00102
	gyroyoffset  = 0.00561
	gyrozoffset  = -0.00477
)

func main() {
	inName := flag.String("i", "", "Required: telemetry file from DashWare to read")
	flag.Parse()
	if *inName == "" {
		flag.Usage()
		return
	}
	q.Begin(20)
	rows := readSample(*inName)
	appendData(rows)
	writeChanges(*inName, newTab)
}

func readSample(name string) [][]string {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := csv.NewReader(f).ReadAll()
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func appendData(rows [][]string) {

	for i := range rows {
		if i == 0 {
			header := []string{"TIME", "PITCH", "ROLL", "YAW", "WORLDX", "WORLDY", "WORLDZ"}
			newTab = append(newTab, header)
		} else {
			q.UpdateIMU(strToFloat(rows[i][18])-gyroxoffset, strToFloat(rows[i][19])-gyroyoffset, strToFloat(rows[i][17])-gyrozoffset, strToFloat(rows[i][14])-accelxoffset, strToFloat(rows[i][15])-accelyoffset, strToFloat(rows[i][13])-accelzoffset)
			w := q.GetWorldAccel()
			line := []string{
				rows[i][0],
				floatToStr(q.GetPitch()),
				floatToStr(q.GetRoll()),
				floatToStr(q.GetYaw()),
				floatToStr(w[0]),
				floatToStr(w[1]),
				floatToStr(w[2]),
			}
			newTab = append(newTab, line)
		}
	}
}

func writeChanges(name string, data [][]string) {
	f, err := os.Create(name[:len(name)-4] + "-ahrs.csv")
	if err != nil {
		log.Fatal(err)
	}
	err = csv.NewWriter(f).WriteAll(data)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func floatToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func strToFloat(s string) float64 {
	str, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return str
}
