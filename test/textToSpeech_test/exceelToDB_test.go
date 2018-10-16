package textToSpeech

import (
	"demo-code/logging"
	"fmt"
	"github.com/tealeg/xlsx"
	"strings"
	"testing"
)

var (
	logger = logging.GetLogger()
)

func TestParseExcelToDB(t *testing.T) {
	fp := "./L1.xlsx"
	f, err := xlsx.OpenFile(fp)
	if err != nil {
		logger.Error(err)
		return
	}

	for _, sheet := range f.Sheets {
		switch {
		case strings.HasPrefix(sheet.Name, "LessonExercise"): // 等级测试题型
			err = phioncsToDB(sheet)
			if err != nil {
				logger.Error(err)
				return
			}
		default: // 单元测试和等级测试题型

		}
	}

	logger.Info("end........end")

	fmt.Println(f)
}

func phioncsToDB(sheet *xlsx.Sheet) (err error) {

	var (
		sheetName string = sheet.Name
		testName  string
		typeName  string
	)

	for _, row := range sheet.Rows {

		firstCell := row.Cells[0]
		if strings.HasPrefix(firstCell.String(), "test") { // 找到testName
			nameArray := strings.Split(firstCell.String(), "-")
			testName = nameArray[1]

			typeArrays := strings.Split(firstCell.String(), "_")
			typeName = typeArrays[1]

			continue
		}

		typeName := row.Cells[0].String()

		switch {
		case typeName == "3":

		case typeName == "4":
		default:

		}

	}

	logger.Infof("sheetName=%v,testName=%v,typeName=%v", sheetName, testName, typeName)

	return
}

func getExcelEntity(con string) (f *xlsx.File, err error) {
	f, err = xlsx.OpenFile(con)
	if err != nil {
		logger.Error(err)
		return
	}
	return
}
