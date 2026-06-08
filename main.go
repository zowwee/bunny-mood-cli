package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {

	var mood string

	rootCmd := &cobra.Command{
		Use:   "initiate",
		Short: "Introduction to my cobra project",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello! try: go run . setMood -m happy")
		},
	}

	moodCmd := &cobra.Command{
		Use:   "setMood",
		Short: "set a mood",
		Run: func(cmd *cobra.Command, args []string) {
			output := bunnyMood(mood)
			fmt.Println(output)
		},
	}

	moodCmd.Flags().StringVarP(&mood, "mood", "m", "normal", "set mood")

	rootCmd.AddCommand(moodCmd)

	rootCmd.Execute()
}

func bunnyMood(mood string) string {
	mood = strings.ToLower(mood)
	if mood == "happy" {
		return `
			Happy!

			(\_/)
			( ^_^)
			/ >🌷<
			`
	} else if mood == "sad" {
		return `
			Sad :(

			(\_/)
			( ._.)
			/ >💧<
			`
	} else if mood == "tired" {
		return `
			Tired...

			(\_/)
			( -.-) zZ
			/ >☕<
			`
	} else if mood == "love" || mood == "in love" {
		return `
			In Love!

			(\_/)
			( ♥‿♥)
			/ >💕<
			`
	} else if mood == "rain" || mood == "in the rain" {
		return `
			In The Rain...

			  ☔
			(\_/)
			( ;_;)
			/ >🌧<
			`
	} else if mood == "hungry" {
		return `
			Hungry!

			(\_/)
			( •﹃•)
			/ >🥕<
			`
	} else if mood == "angry" {
		return `
			Angry!

			(\_/)
			( >_<)
			/ >💢<
			`
	} else if mood == "idk" || mood == "confused" {
		return `
			IDK...

			(\_/)
			( •_•)?
			/ >🤷<
			`
	} else if mood == "sleepy" {
		return `
			Sleepy...

			(\_/)
			( -_-) zzZ
			/ >🛏<
			`
	} else if mood == "cool" {
		return `
			Cool!

			(\_/)
			( •_•)
			/ >😎<
			`
	} else if mood == "shy" {
		return `
			Shy...

			(\_/)
			( .///.)
			/ >🌸<
			`
	} else if mood == "excited" {
		return `
			Excited!

			(\_/)
			( ^o^)
			/ >✨<
			`
	} else if mood == "sick" {
		return `
			Sick...

			(\_/)
			( x_x)
			/ >🤒<
			`
	} else {
		return `
			Bunny is unsure of what you typed :( Try Again!

			(\_/)
			( •_•)?
			/ > <\
		`
	}
}
