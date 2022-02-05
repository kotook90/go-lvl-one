package config

import (
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"os"
)

//В тэгах структуры Арр указано из каких полей json читать переменные окружения
//и какой способ валидации к ним применен( пакет "github.com/asaskevich/govalidator")
type App struct {
	Port      int    `json:"port" valid:"port"`
	DbURL     string `json:"db_url" valid:"url"`
	JaegerURL string `json:"jaeger_url" valid:"url"`
	SentryURL string `json:"sentry_url" valid:"url"`
	Email     string `json:"email" valid:"email"`
}

// Декодирую Json в экземпляр структуры Арр(Арр1)
func Encoding() (App, error) {
	App1 := App{}
	err1 := fmt.Errorf("Не могу открыть JSON файл, проверьте правильность пути ")
	err2 := fmt.Errorf("Ошибка при закрытии файла ")
	err3 := fmt.Errorf("Не могу декодировать JSON файл в структуру ")
	file, err := os.Open("/home/anton/Projects/golangLvlOne/hw9/config/config.json")

	if err != nil {
		fmt.Println(err1, err)
		os.Exit(1)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println(err2, err)
			os.Exit(1)
		}
	}()
	err = json.NewDecoder(file).Decode(&App1)
	if err != nil {
		fmt.Println(err3, err)
		os.Exit(1)
	}

	return App1, nil

}

//С помощью пакета "github.com/asaskevich/govalidator"
//валидирую динные из экземпляра структуры Арр(параметры валидации указаны в тэгах структуры
func Validaate(a App) {
	result, err := govalidator.ValidateStruct(a)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	//Если все ок, то true, если нет - false
	fmt.Println(result)

}
