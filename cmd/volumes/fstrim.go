package volumes

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

var lVolumes string

// fstrimCmd represents the fstrim command
var fstrimCmd = &cobra.Command{
	Use:   "fstrim",
	Short: "This triggers fstrim for Longhorn volumes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var responseData map[string]interface{}
		var volumes []string

		client := &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout: 2 * time.Second,
				}).DialContext,
			},
		}

		apiUrl, err := cmd.Flags().GetString("apiUrl")
		if err != nil {
			fmt.Println("Can't find flag `apiUrl`", err)
		}

		if lVolumes == "all" {
			req, err := http.NewRequest("GET", fmt.Sprintf("%s/volumes", apiUrl), nil)
			if err != nil {
				log.Fatal("Error creating HTTP request:", err)
			}

			fmt.Println("Sending volumes: ", req.URL)

			resp, err := client.Do(req)
			if err != nil {
				log.Fatal("Error sending HTTP request:", err)
			}

			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			if err = json.Unmarshal(body, &responseData); err != nil {
				log.Fatal(err)
			}

			for _, id := range responseData["data"].([]interface{}) {
				volumes = append(volumes, id.(map[string]interface{})["id"].(string))
			}
		} else {
			volumes = strings.Split(lVolumes, " ")
		}

		for _, volume := range volumes {
			req, err := http.NewRequest("POST", fmt.Sprintf("%s/volumes/%s?action=trimFilesystem", apiUrl, volume), nil)
			if err != nil {
				log.Fatal("Error creating HTTP request:", err)
			}

			fmt.Println("Sending fstrim: ", req.URL)

			req.Header = http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			}

			resp, err := client.Do(req)
			if err != nil {
				log.Fatal("Error sending HTTP request:", err)
			}

			if resp.StatusCode != 200 {
				log.Fatal("Didn't receive HTTP 200 status code. Stopped.")
			}

			resp.Body.Close()

			time.Sleep(time.Second * 2)
		}

		log.Println("Done.")
	},
}

func init() {
	VolumesCmd.AddCommand(fstrimCmd)
	VolumesCmd.PersistentFlags().StringVarP(&lVolumes, "volumes", "v", "all", "longhorn volumes list")
}
