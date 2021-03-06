package sound

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/audio"
)

const (
	sampleRate = 44100
	assetPath  = "asset/sound"
)

var (
	audioContext, _ = audio.NewContext(sampleRate)
)

func InitSE() {
	setSE(assetPath+"/se/Select.wav", selectSE)
	setSE(assetPath+"/se/Collision.wav", collisionSE)
	setSE(assetPath+"/se/GoInside.wav", GoInsideSE)
	setSE(assetPath+"/se/GoOutside.wav", GoOutsideSE)
	setSE(assetPath+"/se/Ledge.wav", LedgeSE)
	setSE(assetPath+"/se/Menu.wav", MenuSE)
	setSE(assetPath+"/se/Save.wav", SaveSE)
}

func InitBGM(bgmname string, fade bool) {
	openBGM(fmt.Sprintf(assetPath+"/bgm/%s", bgmname))
	go playBGM(fade)
}

func ExitBGM() {
	closeBGM()
}

func Exit() {
	closeBGM()
	closeSE(selectSE)
}
