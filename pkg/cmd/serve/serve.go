package serve

import (
	"github.com/Benbentwo/bastion-wol/pkg/cmd/common"
	"github.com/Benbentwo/utils/log"
	"github.com/spf13/cobra"
	"net/http"
)

// options for the command
type ServeOptions struct {
	*common.CommonOptions
	port int
}

func NewCmdServe(commonOpts *common.CommonOptions) *cobra.Command {
	options := &ServeOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}
	options.AddServeFlags(cmd)
	cmd.PersistentFlags().IntVarP(&options.port, "port", "p", 8080, "port to serve and listen on")
	// Section to add commands to:

	return cmd
}

func (o *ServeOptions) AddServeFlags(cmd *cobra.Command) {
	o.Cmd = cmd
}

// Run implements this command
func (o *ServeOptions) Run() error {

	log.Logger().Infof("Serving on port %d", o.port)
	http.HandleFunc("/status", StatusPage)
	//http.HandleFunc("/wol", WolPage)
	log.Logger().Fatalf("listen and serve failed due to %s", http.ListenAndServe(":"+common.MustStrConvToStr(o.port), nil))
	return nil
}
