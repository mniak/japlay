package player

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseTime(t *testing.T) {
	testCases := []struct {
		number        int
		secondsString string
	}{
		// Santo santo santo
		{1551004, "8,79253968253968"},
		{3452836, "19,5739002267574"},
		{5153892, "29,217074829932"},
		{6853608, "38,8526530612245"},

		{8663428, "49,1124036281179"},
		{10432908, "59,1434693877551"},
		{12045508, "68,2851927437642"},
		{13880780, "78,6892290249433"},

		{15640552, "88,6652607709751"},
		{17345548, "98,3307709750567"},
		{19115884, "108,366689342404"},
		{20750960, "117,635827664399"},
		{23472268, "133,062743764172"},

		// Other hymns
		{40206640, "227,928798185941"},
		{40733000, "230,912698412698"},
		{41199420, "233,556802721088"},
		{46934928, "266,071020408163"},
		{47797876, "270,963015873016"},
		{48725748, "276,22306122449"},
		{50867816, "288,366303854875"},
		{52139296, "295,574240362812"},
		{52911208, "299,950158730159"},
		{27367700, "155,145691609977"},
		{28242200, "160,103174603175"},
		{29146364, "165,228820861678"},
		{42205548, "239,260476190476"},
		{42692408, "242,020453514739"},
		{43220204, "245,012494331066"},
		{44196052, "250,544512471655"},
		{44770416, "253,800544217687"},
		{45301768, "256,812743764172"},
		{55349764, "313,774172335601"},
	}
	for _, tc := range testCases {
		durationFloat, err := strconv.ParseFloat(strings.Replace(tc.secondsString, ",", ".", 1), 64)
		require.NoError(t, err)
		duration := time.Nanosecond * time.Duration(int64(durationFloat*1_000_000_000))

		t.Run(fmt.Sprintf("%d should be parsed to %v", tc.number, duration), func(t *testing.T) {
			result := ParseTime(tc.number)
			assert.Equal(t, duration.Seconds(), result.Seconds())
		})
	}
}
