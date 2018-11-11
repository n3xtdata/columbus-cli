package status

import (
	"fmt"
	"github.com/n3xtdata/columbus-cli/utils"
	"github.com/spf13/cobra"
)

var _ = fmt.Printf

var (
	getLong = `
		Display one or many resources
		Prints a table of the most important information about the specified resources.
		You can filter the list using a label selector and the --selector flag. If the
		desired resource type is namespaced you will only see results in your current
		namespace unless you pass --all-namespaces.
		Uninitialized objects are not shown unless --include-uninitialized is passed.
		By specifying the output as 'template' and providing a Go template as the value
		of the --template flag, you can filter the attributes of the fetched resources.`

	getExample = `
		# List all pods in ps output format.
		kubectl get pods
		# List all pods in ps output format with more information (such as node name).
		kubectl get pods -o wide
		# List a single replication controller with specified NAME in ps output format.
		kubectl get replicationcontroller web
		# List deployments in JSON output format, in the "v1" version of the "apps" API group:
		kubectl get deployments.v1.apps -o json
		# List a single pod in JSON output format.
		kubectl get -o json pod web-pod-13je7
		# List a pod identified by type and name specified in "pod.yaml" in JSON output format.
		kubectl get -f pod.yaml -o json
		# Return only the phase value of the specified pod.
		kubectl get -o template pod/web-pod-13je7 --template={{.status.phase}}
		# List all replication controllers and services together in ps output format.
		kubectl get rc,services
		# List one or more resources by their type and names.
		kubectl get rc/web service/frontend pods/web-pod-13je7`
)
var Cmd = &cobra.Command{
	Use:     "status",
	Short:   "shows the current status of columbus",
	Long:    getLong,
	Example: getExample,

	Run: func(cmd *cobra.Command, args []string) {

		var s, i = utils.GetPid()

		if i == 0 {
			fmt.Println("Columubs is not running")
		} else {
			fmt.Println("Columbus is running, PID = " + s)
		}

	},
}
