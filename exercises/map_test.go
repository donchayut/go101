package exercises

func TestConvertArabicToThai(t *testing.T) {
	want := "๒๐๑๙-๐๗-๑๘ ๒๓:๕๖:๔๔"
	given := "2019-07-18 23:56:44"
	
	get := arabicTimeToThaiTime(given)

	if want != get {
		t.Errorf("it should convert arabic number of %q to %q but got %q\n",given,want,get)
	}
}