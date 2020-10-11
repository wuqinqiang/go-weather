package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wuqinqiang/go-weather/server"
	"github.com/wuqinqiang/go-weather/tools"
)

func init() {
	NameCmd.PersistentFlags().StringP("name", "n", "", "input city name")
	weatherCmd.AddCommand(NameCmd)
}

var NameCmd = &cobra.Command{
	Use:   "name",
	Short: "check city weather by city name",
	Args: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return errors.New("please input city name")
		}
		if len(name) == 0 {
			return errors.New("请携带参数-n 或者 --name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		code := tools.CityMap[name]
		if code == 0 {
			fmt.Println("这个城市我不想查")
			os.Exit(1)
		}
		info, err := server.GetWeather(code)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("查询的天气是:", info)
	},
}
