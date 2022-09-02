package runner

import "github.com/projectdiscovery/gologger"

const banner = `
                        __  
  ______________ ______/ /__
 / ___/ ___/ __  / ___/ //_/
/ /__/ /  / /_/ / /__/ / \
\___/_/   \__,_/\___/_/|_| 

	  ` + Version + ` by zp857
`

const Version = `v2.0`

func showBanner() {
	gologger.Print().Msgf("%v\n", banner)
}
