package runner

import "github.com/projectdiscovery/gologger"

const banner = `
                        __  
  ______________ ______/ /__
 / ___/ ___/ __  / ___/ //_/
/ /__/ /  / /_/ / /__/ / \
\___/_/   \__,_/\___/_/|_| 

	  v0.0.1 by zp857
`

const Version = `v0.0.1`

func showBanner() {
	gologger.Print().Msgf("%v\n", banner)
}
