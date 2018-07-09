package twelve

import "fmt"

func Verse(day int) string {
	days := []string{
		"first", "second", "third", "fourth",
		"fifth", "sixth", "seventh", "eighth",
		"ninth", "tenth", "eleventh", "twelfth"}
	verses := []string{
		"a Partridge in a Pear Tree",
		"two Turtle Doves",
		"three French Hens",
		"four Calling Birds",
		"five Gold Rings",
		"six Geese-a-Laying",
		"seven Swans-a-Swimming",
		"eight Maids-a-Milking",
		"nine Ladies Dancing",
		"ten Lords-a-Leaping",
		"eleven Pipers Piping",
		"twelve Drummers Drumming"}

	str := fmt.Sprintf("On the %s day of Christmas my true love gave to me, ", days[day-1])
	for i := day; i > 0; i-- {
		if i == 1 {
			if day != 1 {
				str += "and "
			}
			str += verses[0] + "."
			continue
		}
		str += verses[i-1] + ", "
	}

	return str
}

func Song() string {
	s := ""
	for i := 1; i < 13; i++ {
		s += Verse(i) + "\n"
	}
	return s
}
