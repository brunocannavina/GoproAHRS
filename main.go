package main

import (
	"encoding/csv"
	"flag"
	"goahrs"
	"log"
	"os"
	"strconv"
)

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
	inName := flag.String("i", "", "Required: tele metry file from DashWare to read")
	flag.Parse()
	if *inName == "" {
		flag.Usage()
		return
	}
	outName := *inName
	outName = outName[:len(outName)-4] + "-ahrs.csv"
	
	// 20 is sensors sample frequency
	q.Begin(20)
	sensors := readSample(*inName)
	table := appendData(sensors)
	writeChanges(outName, table)
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

func appendData(rows [][]string) [][]string {
	var newTab [][]string
	for i := range rows {
		if i == 0 {
			header := []string{"TIME", "PITCH", "ROLL", "YAW"}
			newTab = append(newTab, header)
		} else {
			gx := strToFloat(rows[i][18])-gyroxoffset
			gy := strToFloat(rows[i][19])-gyroyoffset
			gz := strToFloat(rows[i][17])-gyrozoffset
			ax := strToFloat(rows[i][14])-accelxoffset
			ay := strToFloat(rows[i][15])-accelyoffset
			az := strToFloat(rows[i][13])-accelzoffset

			q.UpdateIMU(gx, gy, gz, ax, ay, az)

			line := []string{
				rows[i][0],
				floatToStr(q.GetPitch()),
				floatToStr(q.GetRoll()),
				floatToStr(q.GetYaw()),
			}
			newTab = append(newTab, line)
		}
	}
	return newTab
}

func writeChanges(name string, data [][]string) {
	f, err := os.Create(name)
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
