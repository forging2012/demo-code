package pkg

import (
	"fmt"
	"github.com/liviosoares/go-watson-sdk/watson"
	"github.com/liviosoares/go-watson-sdk/watson/text_to_speech"
	"os"
	"testing"
)

func TestTextToSpeech(t *testing.T) {

	//config := watson.Config{
	//	Credentials: watson.Credentials{
	//		Username: "1fdda004-cddd-44c6-a0de-7308188d604a",
	//		Password: "ps4uYqY4Z8sO",
	//	},
	//}

	client, err := getClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}

	vl, err := client.ListVoices()
	if err != nil {
		t.Error(err)
		return
	}
	head := vl.Voices[0]
	fmt.Println(head)
	fmt.Println(head.URL)
	fmt.Println(head.Language)

	data, err := client.Synthesize("hello", "audio/mp3", "en-US_AllisonVoice", "")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data)

	//msg, err := client.GetPronunciation("helloWorld", "en-US_AllisonVoice", "spr")
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//
	//fmt.Println(msg)

	//data, err := client.Synthesize("helloWorld", "en-US_AllisonVoice", "audio/wav", "")
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//
	//out, err := os.Create("./output.wav")
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//defer out.Close()
	//n, err := out.Write(data)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//fmt.Println(n)
}

func getClient() (client text_to_speech.Client, err error) {

	config := watson.Config{
		Credentials: watson.Credentials{
			Username: "1fdda004-cddd-44c6-a0de-7308188d604a",
			Password: "ps4uYqY4Z8sO",
		},
	}

	client, err = text_to_speech.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

// .ˌhɛ.ˈlo.ˌwɜld
func TestPronunciation(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}
	data, err := client.GetPronunciation("helloWorld", "", "")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data)
}

func TestTwo(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}

	data, err := client.Synthesize("What are these?", "en-US_AllisonVoice", "audio/mp3", "")
	if err != nil {
		t.Error(err)
		return
	}

	out, err := os.Create("./output.mp3")
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

func TextToSpeech(content string) (err error) {
	client, err := getClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := client.Synthesize(content, "en-US_AllisonVoice", "audio/mp3", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := os.Create("./output.mp3")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()
	_, err = out.Write(data)
	return
}
