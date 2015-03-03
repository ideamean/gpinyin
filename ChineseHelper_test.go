package gpinyin

import (
	"testing"
)

func equal(t *testing.T, got interface{}, want interface{}) {
	if got != want {
		t.Errorf("equal failed,%q == %q，期待值：%q", got, want, want)
	}
}

func TestT2SAndS2TConvert(t *testing.T) {
	const s1 = "台我要1234!#$翻译成繁体的汉字堡垒asdf"
	const s2 = "臺我要1234!#$翻譯成繁體的漢字堡壘asdf"

	r1 := ConvertToTraditionalChinese(s1)
	r2 := ConvertToSimplifiedChinese(s2)

	equal(t, r1, s2)
	equal(t, r2, s1)
}

func TestIsChinese(t *testing.T) {
	equal(t, isChinese("我"), true)
	equal(t, isChinese("臺"), true)
	equal(t, isChinese("A"), false)
}

func TestConvertToPinyinStringPINYIN_WITHOUT_TONE(t *testing.T) {
	const s1 = "台我要1234!#$翻译成繁体的汉字堡垒asdf"
	r1 := ConvertToPinyinString(s1, "-", PINYIN_WITHOUT_TONE)
	r2 := ConvertToPinyinString(s1, "-", PINYIN_WITH_TONE_MARK)
	equal(t, r2, "tái-wǒ-yào-1234!#$-fān-yì-chéng-fán-tǐ-de-hàn-zì-bǎo-lěi-asdf")
	equal(t, r1, "tai-wo-yao-1234!#$-fan-yi-cheng-fan-ti-de-han-zi-bao-lei-asdf")
}
