package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"github.com/tedux/timing/pkg/app"
	"github.com/tedux/timing/pkg/web/handler"
)

func main() {
	srv := app.New()

	rootCmd := &cobra.Command{
		Use:     "srv",
		Aliases: []string{"tim"},
		Short:   "Record time of doing something",
	}

	cmdWeb := &cobra.Command{
		Use:   "web",
		Short: "Open web page at http://localhost:8020/",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			gin.SetMode(gin.DebugMode)
			r := gin.Default()
			r.Use()
			//配置index页面入口
			r.StaticFile("/", "./web/index.html")

			handler.RegisterHandler(r, srv)

			err := r.Run(":8020")
			if err != nil {
				log.Fatalf("Fail to run web: %v", err)
			}
		},
	}

	cmdStart := &cobra.Command{
		Use:  "start [action name]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatalf("Args num is not equal to 1")
			}
			//开始记录
			id, err := srv.StartTiming(args[0])
			if err != nil {
				log.Fatalf("Fail to start action '%s': %v", args[0], err)
			}
			log.Printf("Action id: %v", id)
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
			if input.Text() == "q" || input.Text() == "exit" || input.Text() == "stop" || input.Text() == "quit" {
				err = srv.StopTiming(id)
				if err != nil {
					log.Fatalln(err)
				}
			}
		},
	}

	cmdList := &cobra.Command{
		Use:  "list",
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				log.Fatalf("Args is no need")
			}
			actions, err := srv.ListTimings()
			if err != nil {
				log.Fatalf("Fail to list actions: %v", err)
			}
			encoder := json.NewEncoder(os.Stdout)
			err = encoder.Encode(actions)
			if err != nil {
				log.Fatalln(err)
			}
		},
	}

	cmdListByName := &cobra.Command{
		Use:   "name [string]",
		Short: "list all actions by name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatalf("Args num is not equal to 1")
			}
			actions, err := srv.SearchTimingsByActionAndDt(args[0], "")
			if err != nil {
				log.Fatalf("Fail to list actions by name[%s]: %v", args[0], err)
			}
			encoder := json.NewEncoder(os.Stdout)
			err = encoder.Encode(actions)
			if err != nil {
				log.Fatalln(err)
			}
		},
	}

	cmdListByDt := &cobra.Command{
		Use:   "date [string]",
		Short: "list all actions by name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatalf("Args num is not equal to 1")
			}
			actions, err := srv.SearchTimingsByActionAndDt("", args[0])
			if err != nil {
				log.Fatalf("Fail to list actions by dt[%s]: %v", args[0], err)
			}
			encoder := json.NewEncoder(os.Stdout)
			err = encoder.Encode(actions)
			if err != nil {
				log.Fatalln(err)
			}
		},
	}

	rootCmd.AddCommand(cmdStart, cmdList, cmdWeb)
	cmdList.AddCommand(cmdListByName, cmdListByDt)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
