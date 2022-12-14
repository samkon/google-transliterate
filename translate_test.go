package translator

import (
	"testing"
)

// TestTranslator_Translate calls translate.translate.
func TestTranslator_Transliterate(t *testing.T) {
	origin := "كامل الحارونى متفرع من احمد فخرى امام حديقه الطفل مدينه نصر القاهره الدور السابع شقه ١٦"
	dest := "kamil alharunaa mutafarie min ahmad fukharaa amam hadiqih altifl madinuh nasr alqahirah aldawr alsaabie shaqah 16"
	c := Config{
		Proxy:       "http://127.0.0.1:7890",
		UserAgent:   []string{"Custom Agent"},
		ServiceUrls: []string{"translate.google.com.hk"},
	}
	trans := New(c)
	result, err := trans.Transliterate(origin, "auto", "en")

	if result.Text != dest || err != nil {
		t.Fatalf(`%q, %v, Want match for %q, nil`, result.Text, err, dest)
	}
}
