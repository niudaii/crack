package runner

import "github.com/projectdiscovery/gologger"

const Banner = `
                        __  
  ______________ ______/ /__
 / ___/ ___/ __  / ___/ //_/
/ /__/ /  / /_/ / /__/ / \
\___/_/   \__,_/\___/_/|_| 

	  ` + Version + ` by zp857
`

const Version = `v2.1`

func showBanner() {
	gologger.Print().Msgf("%v\n", Banner)
}
