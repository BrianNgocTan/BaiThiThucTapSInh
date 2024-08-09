package words

import (
	"strings"
)

var vulgards = []string{"arse", "ass", "arsehole", "asshole", "balls", "bollocks", "cock", "dick", "prick", "tits", "cunt", "fuck", "fucker", "wank", "jerk off", 
"bugger", "bitch", "whore", "bastard", "piss", "shit", "crap", "fart", "damn", "wanker", "cocksucker", "dickhead", "bollocks", "goddam", "blasted", "motherfucker",
 "blast"}

 func HandlingVulgarWord (world string) string {
	//chars gom cac ki tu trong chuoi word, chuyen cac chu in hoa thanh in thuong
 	var chars []rune
	for _, v:=range world {
		if ('A' <= v && v <= 'Z') {
			chars = append(chars, v + 32)
		} else {
			chars = append(chars, v)
		}
	}

	resultString := string(chars)
	
	flag := false
	for _, v := range vulgards { //vulgard la slice cac chuoi la nhung tu dung tuc
		if strings.Contains(resultString, v) {
			//resultString chua tu dung tuc => world chua tu dung tuc => thay cac nguyen am trong world roi gan cho resultString
			flag = true
			runeSlice := []rune(world)
			for i:=0; i < len(runeSlice); i++ {
				switch runeSlice[i] {
				case 'u':
					runeSlice[i] = '*'
				case 'e':
					runeSlice[i] = '*'
				case 'o':
					runeSlice[i] = '*'
				case 'a':
					runeSlice[i] = '*'
				case 'i':
					runeSlice[i] = '*'
				}
			}
			resultString = string(runeSlice)
			break
		}
	}
	if flag {
		return resultString
	} else {
		return world
	}
}