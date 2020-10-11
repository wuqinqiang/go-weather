package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wuqinqiang/go-weather/server"
)

func init() {
	codeCmd.PersistentFlags().IntP("code", "c", 0, "城市码必须是6位的整数")
	weatherCmd.AddCommand(codeCmd)
}

var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "check city weather by city code",
	Args: func(cmd *cobra.Command, args []string) error {
		code, err := cmd.Flags().GetInt("code")
		if err != nil {
			return errors.New("请输入城市码")
		}
		if code == 0 {
			return errors.New("请携带参数-c 或者 --code")
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
