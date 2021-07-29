//Игра "Угадай число"
package main

/*
V1. Сгенерировать 1..100 и сохранить
V2. Предложить угадать и сохранить
V3. Меньше - сообщение, больше - сообщение
V4. 10 попыток на угадывание, с напоминанием
V5. если X = Y, вывести Успех и перестать перепроходить
V6. если попытки кончились - грустное сообщение
task: -; 321; 23; 12; 13; nothing;
*/
import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//1. генерация чисел
	scs := time.Now().Unix()
	rand.Seed(scs)
	fmt.Println("Выбрал номер от 1 до 100, угадаешь?")
	seed := rand.Intn(100) + 1
	////////////////////////////////////

	//2. чтение с клавиатуры & 4. ограничение в 10 попыток
	reader := bufio.NewReader(os.Stdin)
	win := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("Осталось", 10-guesses, "попыток...")
		fmt.Print("Угадывай давай:")
		input, err := reader.ReadString('\n') //считываем до Enter
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  //удаление \n
		guess, err := strconv.Atoi(input) //строка в число
		if err != nil {
			log.Fatal(err)
		}
		////////////////////////////////////

		//3. предположение
		if guess < seed {
			fmt.Println("Ой, твой вариант МЕНЬШЕ.")
		} else if guess > seed {
			fmt.Println("Ой, твой вариант БОЛЬШЕ")
		} else { //5. угадывание
			win = true
			fmt.Println("\n\tХорошая работа! Ты угадал! B-)")
			break
		}
		////////////////////////////////////
	}
	//6. грустное сообщение если попыток 0
	if !win {
		fmt.Println("Не угадал, или попытки закончились. Я загадал", seed)
	}
	////////////////////////////////////
}
