package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	// Configuração do Viper para ler o arquivo de configuração
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo de configuração: %s", err)
	}

	// Lendo a string de conexão do banco de dados
	stringConexao := viper.GetString("databaseConnectionString")

	// Abrindo a conexão com o banco de dados
	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verificando a conexão com o banco de dados
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão está aberta!")

	// Executando a consulta
	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Processando os resultados (exemplo)
	for rows.Next() {
		// Manipular cada linha conforme necessário
	}
}
