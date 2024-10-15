package main

import (
	"flag"
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	var (
		flagInterval             = flag.Int("interval", 1, "Interval in seconds")
		flagNotifyMemoryMoreThan = flag.Int("notify-on-memory-more-than", 0, "Notify on change, if notify-with is not set it will beep by default")
		flagNoBeep               = flag.Bool("no-beep", false, "Disable beep")
		flagNotifyWith           = flag.String("notify-with", "", "Run command on notification")
		flagVersion              = flag.Bool("version", false, "Show version")
	)

	flag.Parse()

	if *flagVersion {
		fmt.Printf("memmon for %s, v1.0.0 © 2024 Tony Soekirman", runtime.GOARCH)
		return
	}

	pw := progress.NewWriter()
	pw.Style().Visibility.Time = false
	pw.Style().Visibility.Value = false
	pw.Style().Colors.Percent = text.Colors{text.FgGreen}
	pw.Style().Colors.Tracker = text.Colors{text.FgBlue}
	pw.Style().Colors.Message = text.Colors{text.FgCyan}
	go pw.Render()

	if *flagNotifyMemoryMoreThan > 100 {
		fmt.Println("Memory usage threshold cannot be more than 100%")
		return
	}

	if *flagNotifyWith != "" && *flagNotifyMemoryMoreThan == 0 {
		fmt.Println("Notify command cannot be used without a memory Notify threshold")
		return
	}

	fmt.Printf("Monitoring memory usage every %d seconds. Press Ctrl+C to quit.\n", *flagInterval)
	warningMessage := ""
	commandRunMessage := ""
	tracker := progress.Tracker{
		Message: "Memory Usage",
		Total:   10000,
	}

	pw.AppendTracker(&tracker)

	for {
		v, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if *flagNotifyMemoryMoreThan > 0 && v.UsedPercent > float64(*flagNotifyMemoryMoreThan) {
			warningMessage = fmt.Sprintf("WARNING: Memory usage is more than %d%%", *flagNotifyMemoryMoreThan)

			if !*flagNoBeep {
				beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
			}

			if *flagNotifyWith != "" {
				cmd := exec.Command("sh", "-c", *flagNotifyWith)
				err := cmd.Run()
				if err != nil {
					commandRunMessage = fmt.Sprintf("Error running notify command: %s", err.Error())
				} else {
					commandRunMessage = fmt.Sprintf("Command ran successfully: %s", *flagNotifyWith)
				}
			}
		}

		pw.SetPinnedMessages(
			fmt.Sprintf("Total: %.2f GB", float64(v.Total)/(1024*1024*1024)),
			fmt.Sprintf("Used: %.2f GB", float64(v.Used)/(1024*1024*1024)),
			warningMessage,
			commandRunMessage,
		)
		percentUsed := v.UsedPercent
		tracker.SetValue(int64(percentUsed * 100))
		time.Sleep(time.Duration(*flagInterval) * time.Second)
	}
}
