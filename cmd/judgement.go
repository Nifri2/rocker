package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Messages []Message

type Message struct {
	Code    string `json:"code"`
	Column  int64  `json:"column"`
	File    string `json:"file"`
	Level   string `json:"level"`
	Line    int64  `json:"line"`
	Message string `json:"message"`
}

type TableOfJudgement struct {
	Entry map[string]int
}

func (t *TableOfJudgement) Populate() {
	t.Entry = make(map[string]int)
	t.Entry["DL1001"] = 0
	t.Entry["DL3000"] = 100
	t.Entry["DL3001"] = 0
	t.Entry["DL3002"] = 2
	t.Entry["DL3003"] = 2
	t.Entry["DL3004"] = 10
	t.Entry["DL3006"] = 2
	t.Entry["DL3007"] = 2
	t.Entry["DL3008"] = 2
	t.Entry["DL3009"] = 0
	t.Entry["DL3010"] = 0
	t.Entry["DL3011"] = 10
	t.Entry["DL3012"] = 10
	t.Entry["DL3013"] = 1
	t.Entry["DL3014"] = 15
	t.Entry["DL3015"] = 0
	t.Entry["DL3016"] = 10
	t.Entry["DL3018"] = 10
	t.Entry["DL3019"] = 20
	t.Entry["DL3020"] = 75
	t.Entry["DL3021"] = 10
	t.Entry["DL3022"] = 35
	t.Entry["DL3023"] = 10
	t.Entry["DL3024"] = 10
	t.Entry["DL3025"] = 15
	t.Entry["DL3026"] = 50
	t.Entry["DL3027"] = 5
	t.Entry["DL3028"] = 20
	t.Entry["DL3029"] = 20
	t.Entry["DL3030"] = 5
	t.Entry["DL3032"] = 20
	t.Entry["DL3033"] = 50
	t.Entry["DL3034"] = 20
	t.Entry["DL3035"] = 2
	t.Entry["DL3036"] = 5
	t.Entry["DL3037"] = 5
	t.Entry["DL3038"] = 20
	t.Entry["DL3040"] = 20
	t.Entry["DL3041"] = 5
	t.Entry["DL3042"] = 15
	t.Entry["DL3043"] = 75
	t.Entry["DL3044"] = 10
	t.Entry["DL3045"] = 5
	t.Entry["DL3046"] = 25
	t.Entry["DL3047"] = 2
	t.Entry["DL3048"] = 1
	t.Entry["DL3049"] = 1
	t.Entry["DL3050"] = 3
	t.Entry["DL3051"] = 3
	t.Entry["DL3052"] = 3
	t.Entry["DL3053"] = 3
	t.Entry["DL3054"] = 3
	t.Entry["DL3055"] = 3
	t.Entry["DL3056"] = 3
	t.Entry["DL3057"] = 1
	t.Entry["DL3058"] = 3
	t.Entry["DL3059"] = 2
	t.Entry["DL3060"] = 15
	t.Entry["DL3061"] = 10
	t.Entry["DL4000"] = 10
	t.Entry["DL4001"] = 10
	t.Entry["DL4003"] = 20
	t.Entry["DL4004"] = 10
	t.Entry["DL4005"] = 25
	t.Entry["DL4006"] = 15
	t.Entry["SC1000"] = 3
	t.Entry["SC1001"] = 3
	t.Entry["SC1007"] = 3
	t.Entry["SC1010"] = 3
	t.Entry["SC1018"] = 3
	t.Entry["SC1035"] = 3
	t.Entry["SC1045"] = 3
	t.Entry["SC1065"] = 3
	t.Entry["SC1066"] = 3
	t.Entry["SC1068"] = 3
	t.Entry["SC1077"] = 3
	t.Entry["SC1078"] = 3
	t.Entry["SC1079"] = 3
	t.Entry["SC1081"] = 3
	t.Entry["SC1083"] = 3
	t.Entry["SC1086"] = 3
	t.Entry["SC1087"] = 3
	t.Entry["SC1095"] = 3
	t.Entry["SC1097"] = 3
	t.Entry["SC1098"] = 3
	t.Entry["SC1099"] = 3
	t.Entry["SC2002"] = 3
	t.Entry["SC2015"] = 3
	t.Entry["SC2026"] = 3
	t.Entry["SC2028"] = 3
	t.Entry["SC2035"] = 3
	t.Entry["SC2039"] = 3
	t.Entry["SC2046"] = 3
	t.Entry["SC2086"] = 3
	t.Entry["SC2140"] = 3
	t.Entry["SC2154"] = 3
	t.Entry["SC2155"] = 3
	t.Entry["SC2164"] = 3

}

type Insults struct {
	Insults [][]string
}

func (i *Insults) Populate() {
	i.Insults = [][]string{
		{
			"This is pure poetry in code form.",
			"Your code shines like a diamond among pebbles.",
			"You’ve cracked the code to coding perfection!",
		},
		{
			"Almost flawless! A few tweaks and you’re golden.",
			"Impressive work! Just a couple of rough edges.",
			"Your attention to detail is commendable.",
		},
		{
			"Solid work, but I see room for improvement.",
			"Good effort, though there’s a bit of untapped potential.",
			"Keep it up! A little refinement and this will be stellar.",
		},
		{
			"This is borderline, like spaghetti that’s almost al dente but not quite.",
			"Looks like you got the job done… but just barely.",
			"It works, but have you considered coffee before coding?",
		},
		{
			"Your code’s performance mirrors that of a sloth on vacation.",
			"Functional? Yes. Impressive? Not even close.",
			"Did you write this with your monitor off?",
		},
		{
			"This is like eating plain toast when you were promised a gourmet meal.",
			"If effort counts, then you pass… barely.",
			"I’ve seen more structure in a house of cards.",
		},
		{
			"I’ve seen fewer bugs in a rainforest.",
			"Your code has the elegance of a car crash.",
			"Is this performance art or just bad coding?",
		},
		{
			"I hope you’re not getting paid by the quality of your work.",
			"This code’s best feature is its ability to crash consistently.",
			"I’ve seen better code written on cocktail napkins.",
		},
		{
			"This is what happens when you copy-paste without reading.",
			"You should read this: https://en.m.wikipedia.org/wiki/Reading",
			"It’s hard to believe this came from a human mind.",
		},
		{
			"My condolences to whoever has to maintain this.",
			"Did you accidentally delete the good code and submit this instead?",
			"The only good thing about this code is that it eventually stops running.",
		},
	}
}

func (i *Insults) GetInsult(score int) string {

	if score > 100 {
		score = 100
	}

	if score < 0 {
		score = 0
	}

	score = ((score / 10) - 10) * -1
	return i.Insults[score][rand.IntN(3)]
}

type Shame struct {
	Weight  int
	Message string
}

type ShameMessage struct {
	Score  int
	Shames []Shame
}

func Judgement(cmd *cobra.Command, args []string) {

	jflag, err := rootCmd.Flags().GetBool("json")
	if err != nil {
		log.Fatalf("Error getting json flag: %v", err)
	}

	if jflag {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
			ForceColors:   true,
		})
	}

	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)

	table := TableOfJudgement{}
	table.Populate()

	insults := Insults{}
	insults.Populate()

	for _, file := range args {
		var fn string = file

		content, err := os.ReadFile(fn)
		if err != nil {
			log.Fatal(err)
		}

		var data Messages
		err = json.Unmarshal(content, &data)
		if err != nil {
			log.Fatal(err)
		}

		shame_messages := ShameMessage{
			Score: 100,
		}

		for _, msg := range data {
			shame_messages.Score -= int(table.Entry[msg.Code])
			fmt.Println(msg.Code, table.Entry[msg.Code], shame_messages.Score, msg.Message)

			// append shame message in ShameMessage.Shames
			shame_messages.Shames = append(shame_messages.Shames, Shame{Weight: int(table.Entry[msg.Code]), Message: msg.Message})

			if shame_messages.Score <= 0 {
				shame_messages.Score = 0
				break
			}
		}
		fmt.Println("=====================================")
		fmt.Println(insults.GetInsult(shame_messages.Score))
		fmt.Println("Score: ", shame_messages.Score)
		// fmt.Print(shame_messages)

	}
}
