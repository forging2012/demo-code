package textToSpeech_test

import (
	"demo-code/logging"
	"fmt"
	"github.com/liviosoares/go-watson-sdk/watson"
	"github.com/liviosoares/go-watson-sdk/watson/text_to_speech"
	"github.com/tealeg/xlsx"
	"os"
	"strings"
	"testing"
)

var (
	logger = logging.GetLogger()
)

func TestGetVoiceList(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}
	list, err := client.ListVoices()
	if err != nil {
		t.Error(err)
		return
	}

	for _, v := range list.Voices {
		fmt.Println(v.Customization, v.Gender)
	}
}

func Test2(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(client.ListVoices())

}

func getClient() (client text_to_speech.Client, err error) {

	config := watson.Config{
		Credentials: watson.Credentials{
			Username: "bbdf80a1-bb6f-4595-919f-ddc91b3f152c",
			Password: "H0ByuPPEAlH0",
		},
	}

	client, err = text_to_speech.NewClient(config)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

//  <voice-transformation type="Custom" rate="slow">' + item.get('word') + '</voice-transformation>

func TestBasicTextToSpeech(t *testing.T) {
	client, err := getClient()
	if err != nil {
		logger.Error(err)
		return
	}
	text := "Are you ok？"
	slow := fmt.Sprintf("<voice-transformation type=\"Custom\" rate=\"x-slow\">%s</voice-transformation>", text)
	fmt.Println(slow)
	data, err := client.Synthesize(slow, "en-US_AllisonVoice", "audio/mp3", "")
	if err != nil {
		t.Error(err)
		return
	}

	// write audio to file
	out, err := os.Create(fmt.Sprint(text, ".mp3"))
	if err != nil {
		t.Error(err)
		return
	}
	defer out.Close()
	n, err := out.Write(data)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(n)
}

// excel 解析
func TestCreateVoiceFormExcel(t *testing.T) {
	fp := "./L1.xlsx"
	f, err := xlsx.OpenFile(fp)
	if err != nil {
		logger.Error(err)
		return
	}

	for _, sheet := range f.Sheets {

		// 目录已经存在,则忽略
		if Exists(sheet.Name) {
			logger.Info(sheet.Name)
			continue
		}

		// 创建目录
		err := os.Mkdir(sheet.Name, os.ModePerm)
		if err != nil {
			logger.Error(err)
			return
		}

		// 遍历自然拼读法部分的题目
		if strings.HasPrefix(sheet.Name, "LessonExercise") {

			voiceIndx := 0
			var sonPath string
			for rowIndex, row := range sheet.Rows {

				// 表示一个新的测试题型---重新创建目录
				if strings.HasPrefix(row.Cells[0].String(), "test") {
					// 确定子目录
					tmp := strings.Split(row.Cells[0].String(), "-")
					sonPath = fmt.Sprint(sheet.Name, "/", tmp[1])
					err := os.Mkdir(sonPath, os.ModePerm)
					if err != nil {
						t.Error(err)
						logger.Error(err)
						return
					}

					// 确定声音的cell索引
					for cdx, cell := range row.Cells {
						if strings.Contains(cell.String(), "声音") {
							voiceIndx = cdx
							break
						}
					}

					continue
				}
				//// 找到每个单元格上的英文
				content := row.Cells[voiceIndx].String()
				fmt.Printf("rowIndex=%v, cellIndex=%v,content=%v\n", rowIndex, voiceIndx, fmt.Sprint(sonPath, "/", content))

				err = TextToSpeech(content, fmt.Sprint(sonPath, "/", row.Cells[voiceIndx-1].String()))
				if err != nil {
					logger.Error(err)
					return
				}
			}
		} else { // 单元测试和等级测试部分的题目

			voiceIndx := 0
			var sonPath string
			for rowIndex, row := range sheet.Rows {

				// 表示一个新的测试题型---重新创建目录
				if strings.HasPrefix(row.Cells[0].String(), "test") {

					voiceIndx = 0
					sonPath = ""

					// 确定子目录
					tmp := strings.Split(row.Cells[0].String(), "-")
					sonPath = fmt.Sprint(sheet.Name, "/", tmp[1])
					err := os.Mkdir(sonPath, os.ModePerm)
					if err != nil {
						t.Error(err)
						logger.Error(err)
						return
					}

					// 确定声音的cell索引
					for cdx, cell := range row.Cells {
						if strings.Contains(cell.String(), "声音文件编号") {
							voiceIndx = cdx
							break
						}
					}
					continue
				}

				if voiceIndx == 0 {
					break
				}

				//// 找到每个单元格上的英文
				content := row.Cells[voiceIndx+1].String()
				fmt.Printf("rowIndex=%v, cellIndex=%v,content=%v\n", rowIndex, voiceIndx, fmt.Sprint(sonPath, "/", content))

				err = TextToSpeech(content, fmt.Sprint(sonPath, "/", row.Cells[voiceIndx].String()))
				if err != nil {
					fmt.Println(err)
					logger.Error(err)
					return
				}
			}
		}
	}
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func TextToSpeech(content, fileName string) (err error) {
	client, err := getClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	slow := fmt.Sprintf("<voice-transformation type=\"Custom\" rate=\"x-slow\">%s</voice-transformation>", content)
	data, err := client.Synthesize(slow, "en-US_AllisonVoice", "audio/mp3", "")
	if err != nil {
		logger.Error(err)
		return
	}

	out, err := os.Create(fmt.Sprint(fileName, ".mp3"))
	if err != nil {
		logger.Error(err)
		return
	}
	defer out.Close()
	_, err = out.Write(data)
	return
}

func TestVoice(t *testing.T) {
	msg := "音乐hello音乐"
	fmt.Println(strings.HasPrefix(msg, "音乐"))
	fmt.Println(strings.Contains(msg, "音乐"))
}
