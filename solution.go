package solution

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)


func GetMessage() string {
	new_emoji_st := emoji.Sprint("Hello :world_map:!")
	return new_emoji_st
}
