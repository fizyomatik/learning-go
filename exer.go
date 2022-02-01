package main

import (
	"fmt"
	"time"
)

func main() {
	type placeholder [5]string

	zero := placeholder{
		"QQQ",
		"Q Q",
		"Q Q",
		"Q Q",
		"QQQ",
	}

	one := placeholder{
		" QQ ",
		"  Q ",
		"  Q ",
		"  Q ",
		" QQQ",
	}
	two := placeholder{
		"QQQ",
		"  Q",
		"QQQ",
		"Q  ",
		"QQQ",
	}
	three := placeholder{
		"QQQ",
		"  Q",
		"QQQ",
		"  Q",
		"QQQ",
	}
	four := placeholder{
		"Q Q",
		"Q Q",
		"QQQ",
		"  Q",
		"  Q",
	}
	five := placeholder{
		"QQQ",
		"Q  ",
		"QQQ",
		"  Q",
		"QQQ",
	}
	six := placeholder{
		"QQQ",
		"Q  ",
		"QQQ",
		"Q Q",
		"QQQ",
	}
	seven := placeholder{
		"QQQ",
		"  Q",
		"  Q",
		"  Q",
		"  Q",
	}
	eight := placeholder{
		"QQQ",
		"Q Q",
		"QQQ",
		"Q Q",
		"QQQ",
	}
	nine := placeholder{
		"QQQ",
		"Q Q",
		"QQQ",
		"  Q",
		"QQQ",
	}
	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}
	now := time.Now()
	hour, min, sec := now.Hour(), now.Minute(), now.Second()
	fmt.Printf("hour %d, min %d, sec %d\n", hour, min, sec)

	clock := [...]placeholder{
		digits[hour/10], digits[hour%10],
		digits[min/10], digits[min%10],
		digits[sec/10], digits[sec%10],
	}

	for line := range clock {
		for digit := range clock {
			fmt.Print(clock[digit][line], "  ")
		}
		fmt.Println()
	}

}
