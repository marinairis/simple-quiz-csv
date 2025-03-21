package game

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/marinairis/quiz-go/utils"
)

func (g *GameState) Init() {
	fmt.Print("Seja bem vindo ao quiz com Go!")
	fmt.Print("\033[34;1m Escreva seu nome:\033[0m\n")

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	g.Name = strings.TrimSpace(name)

	fmt.Printf("Vamos ao jogo \033[34;1m%s\033[0m \n", g.Name)
	fmt.Println("------------------------------------")
}

func (g *GameState) ChooseTheme() string {
	fmt.Println("Escolha o tema do quiz:")
	fmt.Println("[1] Anime")
	fmt.Println("[2] Inglês")
	fmt.Println("[3] História do Brasil")
	fmt.Print("\033[35;1mDigite o número correspondente ao tema:\033[0m\n")
	fmt.Println("------------------------------------")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		return "csv/quiz-anime.csv"
	case "2":
		return "csv/quiz-english.csv"
	case "3":
		return "csv/quiz-history-brazil.csv"
	default:
		fmt.Println("Escolha inválida. Usando o tema padrão: \033[35;1mAnime\033[0m")
		return "csv/quiz-anime.csv"
	}
}

func (g *GameState) ProccessCSV(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Error reading csv")
	}

	for index, record := range records {
		if index > 0 {
			correctAnswer, _ := utils.ToInt(record[5])
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  correctAnswer,
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func (g *GameState) Run() {
	timerDuration := 10 * time.Second

	for i, q := range g.Questions {
		fmt.Printf("\033[33;1m %d. %s\033[0m\n", i+1, q.Text)

		for j, opt := range q.Options {
			fmt.Printf("[%d] %s\n", j+1, opt)
		}

		fmt.Println("Digite a resposta:")

		answerCh := make(chan int)
		errCh := make(chan error)

		go func() {
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')

			read = strings.TrimSpace(read)
			answer, err := utils.ToInt(read)

			if err != nil {
				errCh <- err
				return
			}

			answerCh <- answer
		}()

		fmt.Printf("Tempo: \033[36m%d\033[0m segundos\n", int(timerDuration.Seconds()))

		select {
		case answer := <-answerCh:
			if answer == q.Answer {
				fmt.Println("Parabéns você acertou!")
				g.Points += 10
			} else {
				fmt.Println("Vixx tu Errou!")
			}
		case err := <-errCh:
			fmt.Println(err)
		case <-time.After(timerDuration):
			fmt.Println("\nTempo esgotado!")
		}

		fmt.Println("------------------------------------")
	}
}

func (g *GameState) CheckApproval() {
	if g.Points >= 20 {
		g.Approved = true
		fmt.Printf("Parabéns você foi \033[32;1mAPROVADO\033[0m com %d pontos.\n", g.Points)
	} else {
		g.Approved = false
		fmt.Printf("Infelizmente você foi \033[31;1mREPROVADO\033[0m com %d pontos.\n", g.Points)
	}
}
