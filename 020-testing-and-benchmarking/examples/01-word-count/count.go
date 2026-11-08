package words

import "strings"

func Count(text string) int { return len(strings.Fields(text)) }
