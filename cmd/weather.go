package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wuqinqiang/go-weather/server"
)

var (
	cfgFile    string
	weatherCmd = &cobra.Command{
		Use:   "weather",
		Short: "check weather tool",
		Long:  "that is a check weather tool for command line",
		Args: func(cmd *cobra.Command, args []string) error {
			code, err := cmd.Flags().GetInt("code")
			if err != nil {
				return errors.New("请输入城市码")
			}
			fmt.Println("code的值:", code)
			if code < 100000 || code > 999999 {
				return errors.New("请输入正确的城市码")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			code, _ := cmd.Flags().GetInt("code")
			info, err := server.GetWeather(code)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Println("查询的天气是:", info)
		},
	}
)

func Execute() error {
	return weatherCmd.Execute()
}

func init() {
	cobra.OnInitialize()
	weatherCmd.PersistentFlags().IntP("code", "c", 0, "城市码必须是6位的整数")
	_ = viper.BindPFlag("code", weatherCmd.PersistentFlags().Lookup("code"))
	viper.SetDefault("code", 0)
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

