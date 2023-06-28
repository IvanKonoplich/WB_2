package ntp_module

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

// NtpModule is exported function
func NtpModule() {
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(2)
	}

	fmt.Printf("Текущее время: %v \n", time.Now())
	fmt.Println("Ntp время:", t.String())
}
