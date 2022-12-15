package demoparser

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"

	dem "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	msg "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/msg"
	"github.com/pkg/errors"
)

// CsvHeader creates a headerstring of the csv file
func CsvHeader(samplesPerAttack int) []string {
	defer RecoverFromPanic()
	header := []string{}
	l := reflect.TypeOf(AttackData{}).NumField()
	val := reflect.Indirect(reflect.ValueOf(AttackData{}))

	for i := 0; i < l; i++ {

		switch v := (val.Type().Field(i).Type); v.Kind() {
		case reflect.String:
			header = append(header, val.Type().Field(i).Name)
		case reflect.Array:
			continue
		case reflect.Slice:
			continue
		case reflect.Bool:
			header = append(header, val.Type().Field(i).Name)
		case reflect.Uint64:
			header = append(header, val.Type().Field(i).Name)
		case reflect.Float32:
			header = append(header, val.Type().Field(i).Name)
		case reflect.Uint32:
			header = append(header, val.Type().Field(i).Name)
		case reflect.Int32:
			header = append(header, val.Type().Field(i).Name)
		case reflect.Int:
			header = append(header, val.Type().Field(i).Name)
		case reflect.Int64:
			header = append(header, val.Type().Field(i).Name)
		case reflect.Uint8:
			header = append(header, val.Type().Field(i).Name)

		default:
			fmt.Printf("unhandled kind %s\n", v.Kind())
		}
	}
	for j := 1; j <= samplesPerAttack; j++ {
		for k := 0; k < l; k++ {
			switch v := (val.Type().Field(k).Type); v.Kind() {
			case reflect.String:
				//header=append(header, val.Type().Field(i).Name)
			case reflect.Array:
				n := fmt.Sprintf("%sSample%d", val.Type().Field(k).Name, j)
				header = append(header, n)
				//fmt.Println(n)
			case reflect.Slice:
				n := fmt.Sprintf("%sSample%d", val.Type().Field(k).Name, j)
				header = append(header, n)
			case reflect.Bool:
				//header=append(header, val.Type().Field(i).Name)
			case reflect.Uint64:
				//header=append(header, val.Type().Field(i).Name)
			case reflect.Float32:
				//header=append(header, val.Type().Field(i).Name)
			case reflect.Uint32:
				//header=append(header, val.Type().Field(i).Name)
			case reflect.Int32:
				//header=append(header, val.Type().Field(i).Name)
			case reflect.Int:
				//header=append(header, val.Type().Field(i).Name)
			case reflect.Int64:
				//header=append(header, val.Type().Field(i).Name)
			case reflect.Uint8:

			default:
				fmt.Printf("unhandled kind %s\n", v.Kind())
			}

		}
	}

	return header
}

// Returns a mod b, keeping the sign of b
func divisorSignMod(a float64, b float64) float64 {
	return math.Mod(math.Mod(a, b)+b, b)
}

// Normalize an angle to be between -180 and 180
func normalizeAngle(a float32) float32 {
	return float32(-180 + divisorSignMod(float64(a)+180, 360))
}

// checkError prints the error if error is not nil
func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// ParseServerInfo parse the demo's server info
func ParseServerInfo(r io.Reader) *msg.CSVCMsg_ServerInfo {
	var (
		res *msg.CSVCMsg_ServerInfo
	)

	err := ParseWith(r, func(p dem.Parser, h common.DemoHeader) {

		p.RegisterNetMessageHandler(func(info *msg.CSVCMsg_ServerInfo) {
			res = info
		})

	})
	if err != nil {
		return &msg.CSVCMsg_ServerInfo{}
	}

	return res
}

func ParseWith(r io.Reader, f func(parser dem.Parser, header common.DemoHeader)) error {
	dem.DefaultParserConfig.IgnoreErrBombsiteIndexNotFound = true
	p := dem.NewParser(r)

	h, err := p.ParseHeader()
	if err != nil {
		return errors.Wrap(err, "failed to parse header")
	}

	f(p, h)

	err = p.ParseToEnd()
	if err == dem.ErrUnexpectedEndOfDemo {
		fmt.Fprintln(os.Stderr, "WARNING: encountered unexpected end of demo, but the parsed data may still be usable")
	} else if err != nil {
		return errors.Wrap(err, "failed to parse demo to end")
	}

	return nil
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func handleGameEventList(gel *msg.CSVCMsg_GameEventList) map[int32]*msg.CSVCMsg_GameEventListDescriptorT {
	gameEventDescs := make(map[int32]*msg.CSVCMsg_GameEventListDescriptorT)
	for _, d := range gel.GetDescriptors() {
		gameEventDescs[d.GetEventid()] = d
	}
	return gameEventDescs
}

func mapGameEventData(d *msg.CSVCMsg_GameEventListDescriptorT, e *msg.CSVCMsg_GameEvent) map[string]*msg.CSVCMsg_GameEventKeyT {
	data := make(map[string]*msg.CSVCMsg_GameEventKeyT, len(d.Keys))
	for i, k := range d.Keys {
		data[k.GetName()] = e.Keys[i]
	}

	return data
}

func debugGameEvent(d *msg.CSVCMsg_GameEventListDescriptorT, ge *msg.CSVCMsg_GameEvent) (string, map[string]any) {
	const (
		typeStr    = 1
		typeFloat  = 2
		typeLong   = 3
		typeShort  = 4
		typeByte   = 5
		typeBool   = 6
		typeUint64 = 7
	)
	data := make(map[string]any)

	for k, v := range mapGameEventData(d, ge) {
		switch v.GetType() {
		case typeStr:
			data[k] = v.ValString
		case typeFloat:
			data[k] = v.ValFloat
		case typeLong:
			data[k] = v.ValLong
		case typeShort:
			data[k] = v.ValShort
		case typeByte:
			data[k] = v.ValByte
		case typeBool:
			data[k] = v.ValBool
		case typeUint64:
			data[k] = v.ValUint64
		}
	}
	return d.Name, data
}

// FloatToString converts float to string
func FloatToString(f any) string {
	var float float64
	switch f.(type) {
	case float32:
		float = float64(f.(float32))
	case float64:
		float = float64(f.(float64))
	}
	return strconv.FormatFloat(float, 'f', -1, 64)
}

// Writes CSV file of a demo files
func CsvExport(filename string, modelData []AttackData, framerate float64, samplesPerAttack int) error {
	defer RecoverFromPanic()
	dataDest := fmt.Sprintf("./data/%.0f_%s.csv", framerate, filename)
	file, err := os.OpenFile(dataDest, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	fmt.Println("writing csv for ", framerate, filename)
	writer.Write(CsvHeader(samplesPerAttack))
	for _, attackData := range modelData {
		er := writer.Write(attackToString(attackData, samplesPerAttack))
		if er != nil {
			return er
		}
	}

	writer.Flush()
	file.Close()
	fmt.Println("done writing csv for ", framerate, filename)
	return nil
}

var weapontypes = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 101, 102, 103, 104, 105, 106, 107, 201, 202, 203, 204, 205, 206, 301, 302, 303, 304, 305, 306, 307, 308, 309, 310, 311, 401, 405}

// Check whether the weapon type is within out desired weapon types
func Contains(wep int) bool {
	for _, v := range weapontypes {
		if v == wep {
			return true
		}
	}

	return false
}

// recoverFromEOF recovers from EOF
func RecoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("panic ", r, "occured while handling")
	}
}
