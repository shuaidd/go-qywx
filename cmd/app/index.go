package app

import (
	"fmt"
	"github.com/ddshuai/go-qywx/auth"
	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:          "app",
		Short:        "Start App",
		Example:      "go-qywx app -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}

	cfg *auth.CorpInfo
)

func init() {
	auth.CorpConfig = &auth.CorpInfo{
		CorpId:    "ww36e0a51aab349a7d",
		Url:       "https://qyapi.weixin.qq.com/cg-bin",
		DebugMode: false,
		Apps: []auth.Application{
			{
				Name:   "address-book",
				Secret: "AfjvAed_ulqhK0OqTprDQ6xOSnqaT34ll2LsRe0D2NA",
			},
		},
		Callbacks: []auth.Callback{},
	}

	cfg = auth.CorpConfig
}

func run() error {
	fmt.Println("app start")
	token, err := auth.CreateToken(cfg.Url, cfg.CorpId, "hsM5arpdRLJ5r73KTOAnOFxx0F_zgH_prujoa9PwEsM")
	if err != nil {
		return err
	}

	fmt.Println(token)
	return nil
}
