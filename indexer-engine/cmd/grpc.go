package cmd

import (
	"app.io/internal/transport/grpc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var grpcAppCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Run gRPC server",
	Long:  `Delivery api using gRPC transport service`,
	Run: func(cmd *cobra.Command, args []string) {
		grpc.RunServer()
	},
}

func init() {

	grpcAppCmd.PersistentFlags().String("port", "3000", "port of exposing service")
	grpcAppCmd.PersistentFlags().String("host", "0.0.0.0", "Address of server")
	grpcAppCmd.PersistentFlags().String("env", "development", "Envirnoment of application")
	grpcAppCmd.PersistentFlags().String("app", "", "Name of application")

	viper.BindPFlag("host", grpcAppCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("port", grpcAppCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("env", grpcAppCmd.PersistentFlags().Lookup("env"))

	RootCmd.AddCommand(grpcAppCmd)

}
