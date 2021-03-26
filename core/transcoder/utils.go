package transcoder

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

var errorMap = map[string]string{
	"Unrecognized option 'vaapi_device'":        "you are likely trying to utilize a vaapi codec, but your version of ffmpeg or your hardware doesn't support it. change your codec to libx264 and restart your stream",
	"unable to open display":                    "your copy of ffmpeg is likely installed via snap packages. please uninstall and re-install via a non-snap method.  https://owncast.online/docs/troubleshooting/#misc-video-issues",
	"Failed to open file 'http://127.0.0.1":     "error transcoding. make sure your version of ffmpeg is compatible with your selected codec or is recent enough https://owncast.online/docs/troubleshooting/#misc-video-issues",
	"can't configure encoder":                   "error with codec. if your copy of ffmpeg or your hardware does not support your selected codec you may need to select another",
	"Unable to parse option value":              "you are likely trying to utilize a specific codec, but your version of ffmpeg or your hardware doesn't support it. either fix your ffmpeg install or try changing your codec to libx264 and restart your stream",
	"OpenEncodeSessionEx failed: out of memory": "your NVIDIA gpu is limiting the number of concurrent stream qualities you can support. remove a stream output variant and try again.",
	"Cannot use rename on non file protocol, this may lead to races and temporary partial files": "",
	"No VA display found for device": "vaapi not enabled. either your copy of ffmpeg does not support it, your hardware does not support it, or you need to install additional drivers for your hardware.",
	"Unrecognized option":            "error with codec. if your copy of ffmpeg or your hardware does not support your selected codec you may need to select another",
	"Could not find a valid device":  "your codec is either not supported or not configured properly",
}

var ignoredErrors = []string{
	"Duplicated segment filename detected",
	"Error while opening encoder for output stream",
	"Unable to parse option value",
	"Last message repeated",
	"Option not found",
	"use of closed network connection",
}

func handleTranscoderMessage(message string) {
	log.Debugln(message)

	// Ignore certain messages that we don't care about.
	for _, error := range ignoredErrors {
		if strings.Contains(message, error) {
			return
		}
	}

	// Convert specific transcoding messages to human-readable messages.
	for error, displayMessage := range errorMap {
		if strings.Contains(message, error) {
			log.Error(displayMessage)
			return
		}
	}

	if message == "" {
		return
	}

	// Simply print the transcoding message verbatim.
	log.Error(message)
}
