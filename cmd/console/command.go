package main

import (
	"fmt"
	"hermes/database/migrations"
	"hermes/database/seeders"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "console",
	Short: "CLI para ejecutar comandos en user-service",
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Ejecuta las migraciones de la base de datos",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ejecutando migración...")
		migrations.RunMigration()
		fmt.Println("Migración completada.")
	},
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Ejecuta el seeder para cargar datos iniciales",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ejecutando seeder...")
		seeders.RunSeed()
		fmt.Println("Seeder complete.")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(seedCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
