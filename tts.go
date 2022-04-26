package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func tiktoktts(s string, voice string) (mp3 []byte, err error) {
	u := "https://api16-normal-useast5.us.tiktokv.com/media/api/text/speech/invoke/"

	q := url.Values{
		"text_speaker":     []string{voice},
		"req_text":         []string{s},
		"speaker_map_type": []string{"0"},
	}

	resp, err := http.PostForm(u, q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	res := ttVoiceResult{}
	err = dec.Decode(&res)
	if err != nil {
		return nil, err
	}

	return res.Data.VStr, nil
}

type ttVoiceResult struct {
	Data struct {
		SKey     string `json:"s_key"`
		VStr     []byte `json:"v_str"`
		Duration string `json:"duration"`
	} `json:"data"`
	Extra struct {
		LogID string `json:"log_id"`
	} `json:"extra"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

var voices = []string{
	// DISNEY VOICES
	"en_us_ghostface",    // Ghost Face
	"en_us_chewbacca",    // Chewbacca
	"en_us_c3po",         // C3PO
	"en_us_stitch",       // Stitch
	"en_us_stormtrooper", // Stormtrooper
	"en_us_rocket",       // Rocket

	// ENGLISH VOICES
	"en_au_001", // English AU - Female
	"en_au_002", // English AU - Male
	"en_uk_001", // English UK - Male 1
	"en_uk_003", // English UK - Male 2
	"en_us_001", // English US - Female (Int. 1)
	"en_us_002", // English US - Female (Int. 2)
	"en_us_006", // English US - Male 1
	"en_us_007", // English US - Male 2
	"en_us_009", // English US - Male 3
	"en_us_010", // English US - Male 4

	// EUROPE VOICES
	"fr_001", // French - Male 1
	"fr_002", // French - Male 2
	"de_001", // German - Female
	"de_002", // German - Male
	"es_002", // Spanish - Male

	// AMERICA VOICES
	"es_mx_002", // Spanish MX - Male
	"br_001",    // Portuguese BR - Female 1
	"br_003",    // Portuguese BR - Female 2
	"br_004",    // Portuguese BR - Female 3
	"br_005",    // Portuguese BR - Male

	// ASIA VOICES
	"id_001", // Indonesian - Female
	"jp_001", // Japanese - Female 1
	"jp_003", // Japanese - Female 2
	"jp_005", // Japanese - Female 3
	"jp_006", // Japanese - Male
	"kr_002", // Korean - Male 1
	"kr_003", // Korean - Female
	"kr_004", // Korean - Male 2
}
