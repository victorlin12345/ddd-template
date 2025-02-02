/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/victorlin12345/ddd-template/internal/infrastructure"
	grpcserver "github.com/victorlin12345/ddd-template/internal/infrastructure/grpc_server"
	"github.com/victorlin12345/ddd-template/internal/presentation"
	"go.uber.org/fx"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fx.New(
			infrastructure.Module,
			presentation.Module,
			fx.Invoke(func(lc fx.Lifecycle, server *grpcserver.GrpcServer, controller *presentation.HelloController) {
				// lc.Append(fx.Hook{OnStart: db.Start, OnStop: db.Stop})
				lc.Append(fx.Hook{OnStart: server.Start, OnStop: server.Stop})
			}),
		).Run()

	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
