package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "ugit"
	app.Usage = "Utilities for working with Git"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "workdir, w",
			Usage:  "Working Directory",
			EnvVar: "UGIT_WORKDIR",
		},
		cli.StringFlag{
			Name:   "apitype, t",
			Usage:  "Git API Type: gitlab | bitbucket | github",
			EnvVar: "UGIT_APITYPE",
		},
		cli.StringFlag{
			Name:   "apiurl, a",
			Usage:  "Git API Root URL: e.g. https://gitlab.myserver.net, https://bitbucket.org, https://api.github.com",
			EnvVar: "UGIT_APIURL",
		},
		cli.StringFlag{
			Name:   "apikey, k",
			Usage:  "Git API Key",
			EnvVar: "UGIT_APIKEY",
		},
	}

	//
	app.Commands = []cli.Command{
		{
			Name:  "backup",
			Usage: "Backup all your repos",
			Action: func(c *cli.Context) {
				println("Running task with args: ", c.Args().First(), "and flags: ",
					"workdir=", c.GlobalString("workdir"),
					"apitype=", c.GlobalString("apitype"),
					"apiurl=", c.GlobalString("apiurl"),
					"apikey=", c.GlobalString("apikey"),
				)
			},
		},
		{
			Name:  "restore",
			Usage: "Restore repos from a backup",
			Action: func(c *cli.Context) {
				c.String("lang")
			},
		},
	}
	app.EnableBashCompletion = true
	app.Run(os.Args)
}

// https://help.github.com/articles/importing-a-git-repository-using-the-command-line/

// git clone --bare https://githost.org/extuser/repo.git
// Makes a bare clone of the external repository in a local directory

// Push the locally cloned repository to GitHub using the "mirror" option, which ensures that all references, such as branches and tags, are copied to the imported repository.
// create repo on github
// cd *repo.git*
// git push --mirror https://github.com/ghuser/repo.git
// Pushes the mirror to the new GitHub repository
// Remove the temporary local repository.
// cd ..
// rm -rf repo.git
