package cmd

import (
	"scoopplus/frontend"

	"github.com/spf13/cobra"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func init() {
	rootCmd.AddCommand(rootCmd_gui)
}

var rootCmd_gui = &cobra.Command{
	Use:   "gui",
	Short: "show GUI applicaiton.",
	Run: func(cmd *cobra.Command, args []string) {
		WailsGUI()
	},
}

func WailsGUI() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "scoopplus",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: frontend.WebAssets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
