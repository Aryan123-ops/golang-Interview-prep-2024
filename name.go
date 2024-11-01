// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"strings"
// 	"time"
// )

// func main() {
// 	str := []string{"lucky", "bojo", "pojo", "sam", "mojo"}

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter a string: ")
// 	input, _ := reader.ReadString('\n')
// 	input = strings.TrimSpace(input)

// 	found := false

// 	for _, value := range str {
// 		if value == input {
// 			found = true
// 			for i := 1; i <= 100; i++ {
// 				// fmt.Printf("%s BABU \U0001F60E\n", strings.ToUpper(value))
// 				output := fmt.Sprintf("%s BABU \U0001F60E%s", strings.ToUpper(value))
// 				say(output)
// 				time.Sleep(2 * time.Second)
// 			}
// 			break
// 		}
// 	}

// 	if !found {
// 		// fmt.Println("not found")
// 		output := fmt.Sprintf("%snot found \U0001F622")
// 		say(output)
// 	}
// }

// func say(text string) {
// 	// Prepare the command to execute
// 	cmd := exec.Command("espeak", text)
// 	// Run the command and wait for it to complete
// 	if err := cmd.Run(); err != nil {
// 		fmt.Println("Error:", err)
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a string: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		output := fmt.Sprintf("%s \U0001F60E", strings.ToUpper(input))
		fmt.Println(output) // Print to terminal
		say(output)         // Speak using espeak
		time.Sleep(2 * time.Second)
	}
}

func say(text string) {
	// Prepare the command to execute espeak
	cmd := exec.Command("espeak", text)
	// Run the command and wait for it to complete
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"strings"
// 	"time"
// )

// func main() {
// 	str := []string{"lucky", "bojo", "pojo", "sam", "mojo"}

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter baby Name: ")
// 	input, _ := reader.ReadString('\n')
// 	input = strings.TrimSpace(input)

// 	found := false

// 	for _, value := range str {
// 		if value == input {
// 			found = true
// 			for i := 1; i <= 100; i++ {
// 				output := fmt.Sprintf("%s BABU \U0001F60E", strings.ToUpper(value))
// 				// fmt.Println(output) // Print to terminal
// 				say(output) // Speak using espeak
// 				time.Sleep(2 * time.Second)
// 			}
// 			break
// 		}
// 	}

// 	if !found {
// 		output := "not found \U0001F622"
// 		fmt.Println(output) // Print to terminal
// 		say(output)         // Speak using espeak
// 	}
// }

// func say(text string) {
// 	// Print the text to the console
// 	fmt.Println(text)

// 	// Prepare the command to execute espeak
// 	cmd := exec.Command("espeak", text)
// 	// Run the command and wait for it to complete
// 	if err := cmd.Run(); err != nil {
// 		fmt.Println("Error:", err)
// 	}
// }

// package main

// import (
// 	"context"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"os/exec"
// 	"strings"
// 	"time"

// 	"google.golang.org/genproto/googleapis/cloud/speech/v1"
// )

// func main() {
// 	for {
// 		fmt.Println("Say something (or type 'exit' to quit):")

// 		// Record audio using sox
// 		recordAudio("input.wav")

// 		// Recognize speech
// 		text, err := recognizeSpeech("input.wav")
// 		if err != nil {
// 			fmt.Println("Error recognizing speech:", err)
// 			continue
// 		}

// 		if strings.TrimSpace(text) == "exit" {
// 			break
// 		}

// 		fmt.Println("You said:", text)
// 		say(text)

// 		time.Sleep(2 * time.Second)
// 	}
// }

// func recordAudio(filename string) {
// 	cmd := exec.Command("sox", "-d", filename, "trim", "0", "5") // Record for 5 seconds
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	cmd.Run()
// }

// func recognizeSpeech(filename string) (string, error) {
// 	ctx := context.Background()
// 	client, err := speech.NewClient(ctx)
// 	if err != nil {
// 		return "", err
// 	}

// 	data, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return "", err
// 	}

// 	req := &speechpb.RecognizeRequest{
// 		Config: &speechpb.RecognitionConfig{
// 			Encoding:        speechpb.RecognitionConfig_LINEAR16,
// 			SampleRateHertz: 16000,
// 			LanguageCode:    "en-US",
// 		},
// 		Audio: &speechpb.RecognitionAudio{
// 			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
// 		},
// 	}

// 	resp, err := client.Recognize(ctx, req)
// 	if err != nil {
// 		return "", err
// 	}

// 	if len(resp.Results) > 0 && len(resp.Results[0].Alternatives) > 0 {
// 		return resp.Results[0].Alternatives[0].Transcript, nil
// 	}

// 	return "", nil
// }

// func say(text string) {
// 	cmd := exec.Command("espeak", text)
// 	if err := cmd.Run(); err != nil {
// 		fmt.Println("Error:", err)
// 	}
// }
