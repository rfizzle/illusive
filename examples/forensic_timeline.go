package examples

import (
	"encoding/json"
	"fmt"
	"github.com/rfizzle/illusive"
	"log"
)

func main() {
	client, err := illusive.NewClient(
		"https://illusive.example.com",
		"Basic <Token>",
		illusive.ClientDisableTLSValidation,
		illusive.ClientTimeout(60),
	)

	if err != nil {
		log.Fatal(err)
	}

	results, err := client.ForensicsTimeline("1")
	if err != nil {
		log.Fatal(err)
	}

	_, err = json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range results {
		switch fmt.Sprintf("%s_%s", item.Source, item.Type) {
		case "MANAGEMENT_EVENT":
			var managementEvent illusive.TimelineDetailsManagementEvent
			err = json.Unmarshal(item.Details, &managementEvent)
			if err != nil {
				fmt.Printf("%s\n", string(item.Details))
				log.Fatal(err)
			}
			// Do something with this event type
		case "MACHINE_PROCESS":
			var machineProcess illusive.TimelineDetailsMachineProcess
			err = json.Unmarshal(item.Details, &machineProcess)
			if err != nil {
				fmt.Printf("%s\n", string(item.Details))
				log.Fatal(err)
			}
			// Do something with this event type
		case "MACHINE_NETSTAT":
			var machineNetstat illusive.TimelineDetailsMachineNetstat
			err = json.Unmarshal(item.Details, &machineNetstat)
			if err != nil {
				fmt.Printf("%s\n", string(item.Details))
				log.Fatal(err)
			}
			// Do something with this event type
		case "MACHINE_DNS":
			var machineDns illusive.TimelineDetailsMachineDns
			err = json.Unmarshal(item.Details, &machineDns)
			if err != nil {
				fmt.Printf("%s\n", string(item.Details))
				log.Fatal(err)
			}
			// Do something with this event type
		case "MACHINE_KLIST":
			var machineKlist illusive.TimelineDetailsMachineKlist
			err = json.Unmarshal(item.Details, &machineKlist)
			if err != nil {
				fmt.Printf("%s\n", string(item.Details))
				log.Fatal(err)
			}
			// Do something with this event type
		case "MACHINE_SESSION":
			var machineSession illusive.TimelineDetailsMachineSession
			err = json.Unmarshal(item.Details, &machineSession)
			if err != nil {
				fmt.Printf("%s\n", string(item.Details))
				log.Fatal(err)
			}
			// Do something with this event type
		case "MACHINE_USER_ASSIST":
			var machineUserAssist illusive.TimelineDetailsMachineUserAssist
			err = json.Unmarshal(item.Details, &machineUserAssist)
			if err != nil {
				fmt.Printf("%s\n", string(item.Details))
				log.Fatal(err)
			}
			// Do something with this event type
		case "MACHINE_SHIM":
			var machineShim illusive.TimelineDetailsMachineShim
			err = json.Unmarshal(item.Details, &machineShim)
			if err != nil {
				fmt.Printf("%s\n", string(item.Details))
				log.Fatal(err)
			}
			// Do something with this event type
		case "MACHINE_INSTALLED":
			var machineInstalled illusive.TimelineDetailsMachineInstalled
			err = json.Unmarshal(item.Details, &machineInstalled)
			if err != nil {
				fmt.Printf("%s\n", string(item.Details))
				log.Fatal(err)
			}
			// Do something with this event type
		case "MANAGEMENT_LAST_DEPLOYMENT":
			var managementLastDeployment illusive.TimelineDetailsManagementLastDeployment
			err = json.Unmarshal(item.Details, &managementLastDeployment)
			if err != nil {
				fmt.Printf("%s\n", string(item.Details))
				log.Fatal(err)
			}
			// Do something with this event type
		case "MACHINE_MFT":
			var machineMfg illusive.TimelineDetailsMachineMft
			err = json.Unmarshal(item.Details, &machineMfg)
			if err != nil {
				fmt.Printf("%s\n", string(item.Details))
				log.Fatal(err)
			}
			// Do something with this event type
		case "MACHINE_EVENT_LOG":
			var machineEventLog illusive.TimelineDetailsMachineEventLog
			err = json.Unmarshal(item.Details, &machineEventLog)
			if err != nil {
				fmt.Printf("%s\n", string(item.Details))
				log.Fatal(err)
			}
			// Do something with this event type
		default:
			fmt.Printf("%s\n", string(item.Details))
			log.Fatalf("unknown type: %s_%s\n", item.Source, item.Type)
		}
	}
}
